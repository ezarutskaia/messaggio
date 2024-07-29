package main

import (
	"fmt"
	"log"
	"src/db"
	"strconv"
	"net/http"
	"encoding/json"
	"src/producer"
)

func MessageHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	engine := *db.Engine()
	var requestData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	textMessage := requestData["message"]

	if textMessage != "" {
        message := &db.Messages{Text: textMessage}
		engine.Create(&message)
		producer.MessageProducer(strconv.Itoa(message.Id))
        w.Write([]byte("note was created with id " + strconv.Itoa(message.Id)))
    } else {
        w.Write([]byte("please, input text message"))
    }
}

func GetStatistic(w http.ResponseWriter, r *http.Request) {
	engine := *db.Engine()
	var totalCount int64
    engine.Model(&db.Messages{}).Count(&totalCount)

    var trueCount int64
    engine.Model(&db.Messages{}).Where("processed = ?", true).Count(&trueCount)

    response := map[string]int64{
        "total_count":  totalCount,
        "processed_count_true": trueCount,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("Server started.")

	mux := http.NewServeMux()
	mux.HandleFunc("/message", MessageHandler)
	mux.HandleFunc("/statistic", GetStatistic)

	err := http.ListenAndServe(":6030", mux)
	log.Fatal(err)
	
}