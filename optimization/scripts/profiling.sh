#!/bin/bash

MODE=${1:-before}
SUFFIX="_${MODE}"

echo "=== Запуск профилирования с нагрузочным тестированием ==="

mkdir -p optimization/benchmarks
mkdir -p optimization/profiles

# ====== ПРОФИЛИРОВАНИЕ БЕНЧМАРКОВ ======
echo "============================================"
echo "Запуск бенчмарков с профилированием..."
echo "============================================"

# Запуск с CPU профилированием
go test -bench=. -benchmem \
  -cpuprofile=optimization/profiles/cpu_test${SUFFIX}.prof \
  -memprofile=optimization/profiles/mem_test${SUFFIX}.prof \
  ./internal/services/statssvc/... > optimization/benchmarks/benchmark${SUFFIX}.txt 2>&1

echo "Профили бенчмарков созданы!"
echo "- optimization/profiles/cpu_test${SUFFIX}.prof"
echo "- optimization/profiles/mem_test${SUFFIX}.prof"
echo ""

# ====== ПРОФИЛИРОВАНИЕ HTTP СЕРВЕРА ======
echo "============================================"
echo "Запуск HTTP сервера для нагрузочного теста..."
echo "============================================"

echo "Запуск сервера..."
go run cmd/main.go &
sleep 3

SERVER_PID=$(lsof -ti:8080)
if [ -z "$SERVER_PID" ]; then
  echo "Ошибка: Сервер не найден на порту 8080"
  exit 1
fi

echo "Сервер запущен с PID: $SERVER_PID"

echo "Запуск CPU профилирования в фоне..."
go tool pprof -proto -seconds=20 http://localhost:8080/debug/pprof/profile > optimization/profiles/cpu_server${SUFFIX}.prof &
CPU_PID=$!

echo "Запуск memory профилирования в фоне..."
go tool pprof -proto -seconds=20 http://localhost:8080/debug/pprof/heap > optimization/profiles/mem_server${SUFFIX}.prof &
MEM_PID=$!

echo "Запуск trace в фоне..."
curl -o optimization/profiles/trace${SUFFIX}.out http://localhost:8080/debug/pprof/trace?seconds=20 &
TRACE_PID=$!

sleep 2

# Генерация рандомных payload'ов для тестов: 100, 1000, 10000, 1000000
python3 - <<'PY'
import json, os, random
os.makedirs('optimization/data', exist_ok=True)
def gen(n):
    return {"values":[random.random()*200-100 for _ in range(n)]}
for n in (1000, 10000, 100000):
    with open(f'optimization/data/values_{n}.json','w') as f:
        json.dump(gen(n), f, separators=(',',':'))
print("Random payloads generated: 1000, 10000, 100000")
PY

echo "Запуск нагрузочного тестирования с hey..."

# Тест 1: Маленький массив
echo "Тестирование маленького набора данных (1000 элементов)..."
hey -n 10000 -c 50 -m POST \
  -H "Content-Type: application/json" \
  -D optimization/data/values_1000.json \
  http://localhost:8080/stats | tee optimization/benchmarks/hey_small${SUFFIX}.txt

# Тест 2: Средний массив
echo "Тестирование среднего набора данных (10000 элементов)..."
hey -n 1000 -c 25 -m POST \
  -H "Content-Type: application/json" \
  -D optimization/data/values_10000.json \
  http://localhost:8080/stats | tee optimization/benchmarks/hey_medium${SUFFIX}.txt

# Тест 3: Большой массив
echo "Тестирование большого набора данных (100000 элементов)..."
hey -n 500 -c 10 -m POST -H "Content-Type: application/json" \
  -D optimization/data/values_100000.json \
  http://localhost:8080/stats | tee optimization/benchmarks/hey_large${SUFFIX}.txt

echo "Ожидание завершения профилирования..."
wait $CPU_PID
wait $MEM_PID
wait $TRACE_PID

echo "Остановка сервера..."
kill $SERVER_PID

echo ""
echo "============================================"
echo "Профилирование завершено!"
echo "============================================"
echo "Созданные файлы:"
echo ""
echo "Бенчмарки (чистая бизнес-логика):"
echo "- optimization/profiles/cpu_test${SUFFIX}.prof"
echo "- optimization/profiles/mem_test${SUFFIX}.prof"
echo ""
echo "HTTP сервер (с накладными расходами):"
echo "- optimization/profiles/cpu_server${SUFFIX}.prof"
echo "- optimization/profiles/mem_server${SUFFIX}.prof"
echo "- optimization/profiles/trace${SUFFIX}.out"