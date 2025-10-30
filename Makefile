MODE ?= before

run:
	go run ./cmd/main.go

fmt:
	go fmt ./...

test:
	go test ./...

profiling:
	./optimization/scripts/profiling.sh $(MODE)

analyze:
	./optimization/scripts/analyze.sh $(MODE)