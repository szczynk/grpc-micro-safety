FROM golang:1.19 AS build

ADD . /grpc-gateway

WORKDIR /grpc-gateway

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./grpc-gateway ./cmd/server/main.go

RUN rm -rf .git

RUN rm -rf tmp



FROM alpine:latest

WORKDIR /

COPY --from=build ./grpc-gateway/grpc-gateway ./grpc-gateway

COPY --from=build ./grpc-gateway/config/config-local.yml ./config/config-local.yml

COPY --from=build ./grpc-gateway/doc ./doc

EXPOSE 5000 7000

ENTRYPOINT ["/grpc-gateway"]
