default: build

build:
	go build -v -o ci-illuminate cmd/main.go

build-pi:
	GOOS=linux GOARCH=arm GOARM=7 go build -v -o ci-illuminate cmd/main.go

clean:
	go clean