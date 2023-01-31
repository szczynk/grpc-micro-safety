FROM golang:1.19 AS build

ADD . /grpc-safety

WORKDIR /grpc-safety

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o ./grpc-safety ./cmd/server/main.go

RUN rm -rf .git

RUN rm -rf tmp



FROM alpine:latest

WORKDIR /

RUN apk update && apk add bash && apk --no-cache add tzdata

COPY --from=build ./grpc-safety/grpc-safety ./grpc-safety

COPY --from=build ./grpc-safety/config/config-local.yml ./config/config-local.yml

EXPOSE 5004 6004 7004

ENTRYPOINT ["/grpc-safety"]
