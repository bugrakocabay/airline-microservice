FROM alpine:latest

RUN mkdir /app

COPY collectorApp /app

CMD ["/app/collectorApp"]