package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ToshihitoKon/goma-bot/slack"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	fmt.Println("Hello, goma-bot")

	r := mux.NewRouter()
	r.HandleFunc("/slack", slack.HandlerSlack)
	addr := ":" + port
	http.ListenAndServe(addr, r)
}
