FROM golang:1.15-alpine as build

WORKDIR /build

COPY ./backend ./

RUN go mod download

RUN go build ./main.go

FROM alpine:3.9.6

WORKDIR /app

COPY --from=build /build/main ./

ENTRYPOINT ./main