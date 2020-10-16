build:
ifeq ($(OS),Windows_NT)
	go build -o ./bin/empdecrypt.exe ./cmd/empdecrypt
else
	go build -o ./bin/empdecrypt ./cmd/empdecrypt
endif

dll:
	mkdir bin
	gcc -m32 -shared -o bin/Matrix42.Common.AppVerificator.dll decoy/matrix.c 

embedded:
ifeq ($(OS),Windows_NT)
	mkdir resources
	xcopy bin\Matrix42.Common.AppVerificator.dll resources
	xcopy EmpCrypt.exe resources
	go build --tags embedded -o ./bin/empdecrypt.exe ./cmd/empdecrypt
else
	mkdir resources || true
	cp bin/Matrix42.Common.AppVerificator.dll resources
	cp EmpCrypt.exe resources
	go build --tags embedded -o ./bin/empdecrypt ./cmd/empdecrypt
endif


all: clean build dll

run:
	go run ./cmd/empdecrypt/main.go

clean:
ifeq ($(OS),Windows_NT)
	-del ".\bin\" ".\resources\"
else
	rm -rf ./bin ./resources
endif
	

test:
	go test -timeout 20s -count=1 -cover ./...