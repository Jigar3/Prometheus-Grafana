FROM golang:1.13-alpine

COPY . /app

WORKDIR /app

RUN go mod tidy

CMD [ "go run main.go" ]