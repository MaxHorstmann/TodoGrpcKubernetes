FROM golang:1.18 AS build
WORKDIR /src
RUN apt-get update && apt-get install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
COPY . .
RUN protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/todo/services.proto 
RUN mkdir -p bin && go build -o bin ./...

FROM golang:1.18
WORKDIR /app
COPY --from=0 /src/bin/server /app
EXPOSE 1234
CMD /app/server