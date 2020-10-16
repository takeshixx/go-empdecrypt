build:
ifeq ($(OS),Windows_NT)
	go build -o ./bin/empdecrypt.exe ./cmd/empdecrypt
else
	go build -o ./bin/empdecrypt ./cmd/empdecrypt
endif

dll:
	gcc -m32 -shared -o bin/Matrix42.Common.AppVerificator.dll decoy/matrix.c 

all: clean build dll

run:
	go run ./cmd/empdecrypt/main.go

clean:
ifeq ($(OS),Windows_NT)
	rm -Recurse -Force ./bin
else
	rm -rf ./bin || true
endif
	

test:
	go test -timeout 20s -count=1 -cover ./...