exeargs =
args =
args1 =
args2 =

ifeq ($(OS),Windows_NT)
    EXE_EXT = .exe
else
    EXE_EXT =
endif

.PHONY: run
run: auth-proxy$(EXE_EXT)
	./auth-proxy $(exeargs)

.PHONY: all
all: auth-proxy$(EXE_EXT)

auth-proxy$(EXE_EXT): *.go
	$(args1) go build $(args) . $(args2)

.PHONY: tests
tests:
	go test ./...

.PHONY: clean
clean:
	rm -rf ./auth-proxy$(EXE_EXT)