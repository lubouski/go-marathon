## Kafka producer and consumer demo app with Kafka and Zookeeper
Repository holds Dockerfiles to run `Kafka` and `Zookeeper` locally as containers, and `app` folder has Dockerfile and source code to containerize golang application which has producer and consumer capabilities. 
### Build Kafka and Zookeeper images, and run them:
```
$ podman build . -t lubouski/kafka:3.2.1
$ cd zookeeper
$ podman build . -t lubouski/zookeeper:3.2.1
# Create shared Kafka network to have simplier DNS names
$ podman network create kafka
# We need to run zookeeper first, Kafka has dependencie on running Zookeeper
$ podman run -d --rm --name zookeeper-1 --net kafka lubouski/zookeeper:3.2.1
$ podman run -d --rm --name kafka-1 --net kafka lubouski/kafka:3.2.1
```
### Exec to Kafka container and create topics for app:
```
$ podman exec -ti kafka-1 bash
# /kafka/bin/kafka-topics.sh --create --bootstrap-server kafka-1:9092 --topic monitoring
# /kafka/bin/kafka-topics.sh --create --bootstrap-server kafka-1:9092 --topic monitoring-group-table
```
### Run app container, and check logs to ensure that it's working:
Application will emit messages every second.
```
$ podman run -d --rm --name topicreader --net kafka lubouski/topicreader:v1
$ podman logs topicreader                                                                                                                                                  
2022/09/14 04:35:50 [Processor monitoring-group] setup generation 1, claims=map[string][]int32{"monitoring":[]int32{0}}
2022/09/14 04:35:51 key = some-key, counter = 0, msg = some-value
$ podman logs topicreader                                                                                                                                                 
2022/09/14 04:35:50 [Processor monitoring-group] setup generation 1, claims=map[string][]int32{"monitoring":[]int32{0}}
2022/09/14 04:35:51 key = some-key, counter = 0, msg = some-value
2022/09/14 04:35:52 key = some-key, counter = 0, msg = some-value
2022/09/14 04:35:53 key = some-key, counter = 0, msg = some-value
``` 
