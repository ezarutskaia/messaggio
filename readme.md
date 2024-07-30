## Установка
git clone https://github.com/ezarutskaia/messaggio

docker compose build

docker compose up -d

docker exec -ti service_1 go run automigrate.go

docker exec -ti service_1 go run consumer.go