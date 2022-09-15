## Golang Kafka application to emit and consume messages every second
Application is based on the library `github.com/lovoo/goka`.
### Build container image
```
$ podman build . -t lubouski/topicreader:v1
$ podman tag topicreader:v1 lubowsky/kafka-reader:v1
$ podman push lubowsky/kafka-reader:v1
```
