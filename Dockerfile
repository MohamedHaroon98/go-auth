FROM golang:1.21.0-alpine

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy 

CMD go run main.go -b 0.0.0.0
