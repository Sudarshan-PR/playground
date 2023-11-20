# Build Stage
FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o main .

# Deployment Stage
FROM alpine:3.14 as deploy

RUN echo "Hello World"

WORKDIR /app

COPY --from=build /app/main .

EXPOSE 8000

CMD ["/app/main"]
