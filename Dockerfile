FROM golang:alpine

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download

ADD . .

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 80

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main