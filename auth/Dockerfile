FROM golang:1.19 AS build

ADD . /grpc-auth

WORKDIR /grpc-auth

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./grpc-auth ./cmd/server/main.go

RUN rm -rf .git

RUN rm -rf tmp



FROM alpine:latest

WORKDIR /

COPY --from=build ./grpc-auth/grpc-auth ./grpc-auth

COPY --from=build ./grpc-auth/config/config-local.yml ./config/config-local.yml

EXPOSE 5001 6001 7001

ENTRYPOINT ["/grpc-auth"]
