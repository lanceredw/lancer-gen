.PHONY: all run clean help

APP = lancer-gen

## linux: Compile and package Linux
.PHONY: linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(RACE) -o ./bin/${APP}-linux64 ./main.go

## win: Compile and package win
.PHONY: win
win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(RACE) -o ./bin/${APP}-win64.exe ./main.go

## mac: Compile and package Mac
.PHONY: mac
mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(RACE) -o ./bin/${APP}-darwin64 ./main.go

build:
	@go build -o ${APP}

## Compile Windows, Linux, Mac platforms
.PHONY: all
all:win linux mac

run:
	@go run ./

.PHONY: tidy
tidy:
	@go mod tidy

## test: Run unit test.
.PHONY: test
test:
	@$(MAKE) go.test

## Clean up binary files
clean:
	@if [ -f ./bin/${APP}-linux64 ] ; then rm ./bin/${APP}-linux64; fi
	@if [ -f ./bin/${APP}-win64.exe ] ; then rm ./bin/${APP}-win64.exe; fi
	@if [ -f ./bin/${APP}-darwin64 ] ; then rm ./bin/${APP}-darwin64; fi

help:
	@echo "make - Format Go code and compile to generate binary files"
	@echo "make mac - Compile Go code to generate binary files for Mac"
	@echo "make linux - Compile Go code to generate Linux binary files"
	@echo "make win - Compile Go code to generate Windows binary files"
	@echo "make tidy - Execute go mod tidy"
	@echo "make run - Run Go code directly"
	@echo "make clean - Remove compiled binary files"
	@echo "make all - Compile binary files for multiple platforms"