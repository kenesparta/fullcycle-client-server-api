package main

import (
	"context"
	"log"
	"net/http"

	"github.com/kenesparta/fullcycle-client-server-api/server/constants"
	"github.com/kenesparta/fullcycle-client-server-api/server/dollar"
	"github.com/kenesparta/fullcycle-client-server-api/server/handlers"
)

func main() {
	dollar.CreateTables(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", handlers.HandleDollarPrice)
	mux.HandleFunc("/read", handlers.ReadDollar)
	errServer := http.ListenAndServe(constants.Port, mux)
	if errServer != nil {
		log.Printf("read error %v", errServer)
		return
	}
}
