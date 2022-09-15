FROM golang:alpine3.16 as dev-env

WORKDIR /app

FROM dev-env as build-env
COPY go.mod /app/

COPY . /app/

RUN CGO_ENABLED=0 go build -o /webserver

FROM alpine:3.16.2 as runtime

COPY --from=build-env /webserver /usr/local/bin/webserver

ENV PORT 8081

RUN chmod +x /usr/local/bin/webserver

ENTRYPOINT ["webserver"]
