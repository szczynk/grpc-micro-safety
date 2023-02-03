FROM golang:1.19 AS build

ADD . /grpc-mail

WORKDIR /grpc-mail

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./grpc-mail ./cmd/server/main.go

RUN rm -rf .git

RUN rm -rf tmp



FROM alpine:latest

WORKDIR /

COPY --from=build ./grpc-mail/grpc-mail ./grpc-mail

COPY --from=build ./grpc-mail/config/config-local.yml ./config/config-local.yml

COPY --from=build ./grpc-mail/internal/templates ./internal/templates

EXPOSE 7003

ENTRYPOINT ["/grpc-mail"]
