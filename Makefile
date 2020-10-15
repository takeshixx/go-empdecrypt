build:
ifeq ($(OS),Windows_NT)
	go build $(ldflags) -o ./bin/empdecrypt.exe ./cmd/empdecrypt
else
	go build $(ldflags) -o ./bin/empdecrypt ./cmd/empdecrypt
endif

run:
	go run ./cmd/empdecrypt/main.go

clean:
	rm -rf ./bin

test:
	go test -timeout 20s -count=1 -cover ./...