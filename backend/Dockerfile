FROM golang:1.22.4-alpine3.20

WORKDIR /app

COPY . /app

RUN go build -o /application cmd/main.go

EXPOSE 8080

CMD ["/application"]
