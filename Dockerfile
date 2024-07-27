FROM golang:latest

WORKDIR /go/src

COPY ./src /go/src

EXPOSE 6050

CMD ["go", "run", "main.go"]