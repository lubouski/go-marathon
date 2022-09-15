FROM golang:alpine3.16 as dev-env

WORKDIR /app

FROM dev-env as build-env
COPY go.mod /go.sum /app/
RUN go mod download

COPY . /app/

RUN CGO_ENABLED=0 go build -o /topicreader

FROM alpine:3.16.2 as runtime

COPY --from=build-env /topicreader /usr/local/bin/topicreader

ENV TOPIC monitoring
ENV GROUP monitoring-group
ENV BROKERS kafka-1:9092

RUN chmod +x /usr/local/bin/topicreader

ENTRYPOINT ["topicreader"]
