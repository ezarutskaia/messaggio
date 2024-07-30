package main

import (
	"src/db"
)

func main() {
	DataBase := *db.Engine()
	DataBase.AutoMigrate(&db.Messages{})
}
