FROM golang:alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o apiGateway ./cmd/api

RUN chmod +x /app/apiGateway

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/apiGateway /app

CMD ["/app/apiGateway"]