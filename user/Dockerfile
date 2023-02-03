FROM golang:1.19 AS build

ADD . /grpc-user

WORKDIR /grpc-user

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./grpc-user ./cmd/server/main.go

RUN rm -rf .git

RUN rm -rf tmp



FROM alpine:latest

WORKDIR /

COPY --from=build ./grpc-user/grpc-user ./grpc-user

COPY --from=build ./grpc-user/config/config-local.yml ./config/config-local.yml

EXPOSE 5002 6002 7002

ENTRYPOINT ["/grpc-user"]
