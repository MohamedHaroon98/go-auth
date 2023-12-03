FROM golang:1.21.0-alpine as builder
WORKDIR /usr/src/app
COPY . .
RUN go build -o go_auth .

FROM alpine:3.18.3
WORKDIR /bin
COPY --from=builder /usr/src/app .
USER nobody
EXPOSE 8080
CMD ["go_auth"]