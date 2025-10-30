#!/bin/bash

MODE=${1:-before}

SUFFIX="_${MODE}"
echo "============================================"
echo "=== Анализ профилей ==="
echo "============================================"

mkdir -p optimization/analysis${SUFFIX}

# ====== АНАЛИЗ ПРОФИЛЕЙ БЕНЧМАРКОВ ======
echo ""
echo "Анализ профилей бенчмарков (чистая бизнес-логика)..."
echo "--------------------------------------------"

# CPU анализ тестов
if [ -f "optimization/profiles/cpu_test${SUFFIX}.prof" ]; then
  echo "Анализ CPU профиля бенчмарков..."
  go tool pprof -text -nodecount=10 optimization/profiles/cpu_test${SUFFIX}.prof > optimization/analysis${SUFFIX}/cpu_test_top.txt
  echo "✓ optimization/analysis${SUFFIX}/cpu_test_top.txt"
else
  echo "⚠ Профиль optimization/profiles/cpu_test${SUFFIX}.prof не найден"
fi

# Memory анализ тестов
if [ -f "optimization/profiles/mem_test${SUFFIX}.prof" ]; then
  echo "Анализ Memory профиля бенчмарков..."
  go tool pprof -text -nodecount=10 optimization/profiles/mem_test${SUFFIX}.prof > optimization/analysis${SUFFIX}/mem_test_top.txt
  echo "✓ optimization/analysis${SUFFIX}/mem_test_top.txt"
else
  echo "⚠ Профиль optimization/profiles/mem_test${SUFFIX}.prof не найден"
fi

# ====== АНАЛИЗ ПРОФИЛЕЙ HTTP СЕРВЕРА ======
echo ""
echo "Анализ профилей HTTP сервера (с накладными расходами)..."
echo "--------------------------------------------"

# CPU анализ сервера
if [ -f "optimization/profiles/cpu_server${SUFFIX}.prof" ]; then
  echo "Анализ CPU профиля сервера..."
  go tool pprof -text -nodecount=10 optimization/profiles/cpu_server${SUFFIX}.prof > optimization/analysis${SUFFIX}/cpu_server_top.txt
  echo "✓ optimization/analysis${SUFFIX}/cpu_server_top.txt"
else
  echo "⚠ Профиль optimization/profiles/cpu_server${SUFFIX}.prof не найден"
fi

# Memory анализ сервера
if [ -f "optimization/profiles/mem_server${SUFFIX}.prof" ]; then
  echo "Анализ Memory профиля сервера..."
  go tool pprof -text -nodecount=10 optimization/profiles/mem_server${SUFFIX}.prof > optimization/analysis${SUFFIX}/mem_server_top.txt
  echo "✓ optimization/analysis${SUFFIX}/mem_server_top.txt"
else
  echo "⚠ Профиль optimization/profiles/mem_server${SUFFIX}.prof не найден"
fi

echo ""
echo "============================================"
echo "Анализ завершен!"
echo "============================================"
echo "Результаты в папке optimization/analysis${SUFFIX}/"
echo ""
echo "Файлы бенчмарков (для анализа бизнес-логики):"
echo "- cpu_test_top.txt - топ функций по CPU в бенчмарках"
echo "- mem_test_top.txt - топ функций по памяти в бенчмарках"
echo ""
echo "Файлы HTTP сервера (для анализа общей картины):"
echo "- cpu_server_top.txt - топ функций по CPU в HTTP сервере"
echo "- mem_server_top.txt - топ функций по памяти в HTTP сервере"