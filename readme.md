## Create topic in kafka
kafka-topics --create --topic messages --partitions 1 --replication-factor 1 --bootstrap-server kafka:9092

## 
kafka-log-dirs --describe --bootstrap-server kafka:9092

kafka-topics --bootstrap-server kafka:9092 --topic messages --describe


kafka-console-producer --bootstrap-server kafka:9092 --topic messages
kafka-console-consumer --bootstrap-server kafka:9092 --topic messages
kafka-topics --list --bootstrap-server kafka:9092