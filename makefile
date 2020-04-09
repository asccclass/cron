PROJECT?=https://github.com/justgps/sherry
RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILDTIME?=$(shell date -u '+%Y-%m-%d %H:%M:%S')
APP?=app
PORT?=10003
ContainerName?=sherry/cron
DBSERVER?=MySQLx

init:
	go get github.com/robfig/cron	
clean:
	sh clean.sh
delapp:
	rm -f ${APP}

build: 
	GOOS=linux GOARCH=amd64 go build -tags netgo \
	-ldflags "-s -w" \
	-o ${APP}

docker: build
	docker build -t ${ContainerName} .
	rm -f ${APP}

run: docker
	docker run --restart=always -d -v /etc/localtime:/etc/localtime:ro --name cron \
	--link ${DBSERVER} -p ${PORT}:80 ${ContainerName}
	docker ps -a
test: docker
	docker run --rm -v /etc/localtime:/etc/localtime:ro --name cron -p ${PORT}:80 ${ContainerName}
	docker ps -a
stop:
	docker stop cron
