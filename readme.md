## Установка
git clone https://github.com/ezarutskaia/messaggio

cd messaggio

docker compose build

docker compose up -d

docker exec -ti messaggio-kafka-1 kafka-topics --create --topic messages --partitions 1 --replication-factor 1 --bootstrap-server kafka:9092

docker exec -ti service_1 go run automigrate.go

docker exec -ti service_1 go run consumer.go

## Консольные команды для выполнения запросов к Api

curl --location 'f8ba8ca4d3f8.vps.myjino.ru/message' \
--header 'Content-Type: application/json' \
--data '{
    "message": "test"
}'

curl --location 'f8ba8ca4d3f8.vps.myjino.ru/statistic'