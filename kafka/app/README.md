## Golang Kafka application to emit and consume messages every second
Application is based on the library `github.com/lovoo/goka`.
### Build container image
```
$ podman build . -t lubouski/topicreader:v1
$ podman tag topicreader:v1 lubowsky/kafka-reader:v1
$ podman push lubowsky/kafka-reader:v1
```
### Deployment to Kubernetes
At the deployment we could configure `topic` name and `group` name which would be used by application to write and read events from. As well as `brokers` addresses with a ports.
```
  containers:
  - env:
    - name: BROKERS
      value: strimzi-kafka-bootstrap:9092
    - name: TOPIC
      value: monitoring
    - name: GROUP
      value: monitoring-group
``` 
So these variable would be propagated to the container and then picked up by the application via `env variables`.

