package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"./slack"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello, goma-bot")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := mux.NewRouter()
	r.HandleFunc("/slack", slack.HandlerSlack)
	addr := ":" + port
	http.ListenAndServe(addr, r)
}
