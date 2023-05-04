BINARY_NAME = ventilar-api
# BINARY_NAME = $(shell cat go.mod | grep module | cut -d " " -f 2 | rev | cut -d "/" -f 1 | rev)
BUILD_FOLDER = build

all: build run

build:
	go build -o ${BUILD_FOLDER}/${BINARY_NAME}
# GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
# GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
# GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows main.go

run: build
	./${BUILD_FOLDER}/${BINARY_NAME}

dev:
	~/go/bin/CompileDaemon -build="go build -o ${BUILD_FOLDER}/${BINARY_NAME}" -command="./${BUILD_FOLDER}/${BINARY_NAME}"

test:
	go test -v ./... -count=1
# count=1 makes it not to cache

test_cov:
	mkdir coverage
	go test ./... -coverprofile=./coverage/coverage.out

dep:
	go get github.com/githubnemo/CompileDaemon
	go get github.com/joho/godotenv
	go get github.com/gofiber/fiber/v2
	go get gorm.io/gorm
	go get gorm.io/driver/postgres
	go get github.com/withmandala/go-log
	go get github.com/golangci/golangci-lint@v1.52.2
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin 
	go get github.com/automation-co/husky@0.2.16


install:
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all
# If you are using go mod, make sure to add "golangci-lint" to your go.mod file.

lint-fix:
	golangci-lint run --enable-all --fix

ctimage:
	docker build -t auto-api:v0.2 .
	
clean:
	go clean
	rm -Rf ./${BUILD_FOLDER} ./coverage
#  rm ${BINARY_NAME}-darwin
#  rm ${BINARY_NAME}-linux
#  rm ${BINARY_NAME}-windows
