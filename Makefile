# 基本的な Go コマンド
GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOVERSION=$(GOCMD) version

GOPHER = "ʕ◔ϖ◔ʔʔ"

.PHONY: dist build-windows build-linux build-mac

gopher:
	@echo ${GOPHER}
version:
	@${GOVERSION}
init:
	@${GOCMD} mod init naoland
	@${GOGET}
build:
	@${GOBUILD} -o ./app app.go
run:
	go run app.go
clean:
	rm ./app
build-windows:
	echo "Build for windows10"
	GOOS=windows GOARCH=amd64 go build -o dist/windows/nemprice-win.exe app.go
	echo "Done!"
build-linux:
	echo "Build for linux"
	GOOS=linux GOARCH=amd64 go build -o dist/linux/nemprice-linux app.go
	echo "Done!"
build-mac:
	echo "Build for macOS(Darwin)"
	GOOS=darwin GOARCH=amd64 go build -o dist/macos/nemprice-mac app.go
