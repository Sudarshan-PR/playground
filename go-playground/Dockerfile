FROM golang:1.21-alpine

WORKDIR /app

RUN apk add rabbitmq-c-utils

COPY . .

CMD ["./listen.sh"]
