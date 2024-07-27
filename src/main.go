package main

import (
	"log"
	"src/db"
	"strconv"
	"net/http"
	"encoding/json"
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
        w.Write([]byte("note was created with id " + strconv.Itoa(message.Id)))
    } else {
        w.Write([]byte("please, input text message"))
    }
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/message", MessageHandler)

	err := http.ListenAndServe(":6030", mux)
	log.Fatal(err)
	
}