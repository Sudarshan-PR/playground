# Build Stage
FROM golang:1.21-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o main-bin .

# Deployment Stage
FROM alpine:3.14 as deploy

RUN echo "Hello World"

WORKDIR /app

COPY --from=build /app/main-bin .

EXPOSE 8000

CMD ["/app/main-bin"]
