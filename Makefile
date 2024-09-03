darwin_amd: clean
	mkdir -p bin
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -mod=vendor -o ./bin/bplivetool ./app/*.go

darwin_arm: clean
	mkdir -p bin
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -mod=vendor -o ./bin/bplivetool ./app/*.go

clean:
	rm -rf bin

