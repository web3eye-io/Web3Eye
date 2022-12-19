# install
## docker
```shell
docker run -d --name zookeeper -p 2181:2181 wurstmeister/zookeeper
docker run -d --name kafka -p 9092:9092 -e KAFKA_BROKER_ID=0 -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 --link zookeeper -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://192.168.49.1:9092 -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 -t wurstmeister/kafka
```

## helm install in k8s
```shell
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install -f 01-kafka.yaml kafka bitnami/kafka
kc get svc
```

# useful command
```shell
kafka-topics.sh --create --bootstrap-server kafka-headless:9092 --replication-factor 1 --partitions 1 --topic dog
kafka-topics.sh --list --bootstrap-server kafka-headless:9092
kafka-console-producer.sh --broker-list kafka-headless:9092 --topic coffe
kafka-console-consumer.sh --bootstrap-server kafka-headless:9092 --topic coffe 


kafka-consumer-groups.sh --bootstrap-server 127.0.0.1:9092 --list
kafka-consumer-groups.sh --bootstrap-server 127.0.0.1:9092 --group default --describe
kafka-topics.sh --bootstrap-server 127.0.0.1:9092 --topic coffe --describe

kafka-consumer-groups.sh --group default  --bootstrap-server kafka-headless:9092 --describe
```