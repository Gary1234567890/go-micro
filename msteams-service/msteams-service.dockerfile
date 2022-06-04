FROM alpine:latest

RUN mkdir /app

COPY msteamsServiceApp /app

CMD [ "/app/msteamsServiceApp"]