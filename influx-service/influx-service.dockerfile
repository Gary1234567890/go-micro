FROM alpine:latest

RUN mkdir /app

COPY influxServiceApp /app

CMD [ "/app/influxServiceApp"]