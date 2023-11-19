#####################
# Build Proto files #
#####################
FROM golang:1.21-alpine as buildProto

# Install protobuff deps
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
WORKDIR /protos

# Get proto files from repo
RUN wget https://github.com/Sudarshan-PR/playground-protos/archive/refs/heads/main.zip && unzip main.zip
WORKDIR /protos/playground-protos-main
RUN mkdir protos

# Generate Protobuff files
RUN protoc --go_out=protos --go_opt=paths=source_relative --go-grpc_out=protos --go-grpc_opt=paths=source_relative ./notification.proto


####################
# Build code Stage #
####################
FROM golang:1.21-alpine AS buildCode

WORKDIR /app

# Install project dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy proto files
COPY --from=buildProto /protos/playground-protos-main/protos .

# Copy source code
COPY . .
RUN go build -v -o main .


####################
# Deployment Stage
####################
FROM alpine:3.14 as deploy

WORKDIR /app

COPY --from=buildCode /app/main .

EXPOSE 8090
EXPOSE 9000

CMD ["/app/main"]
