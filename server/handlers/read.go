package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/kenesparta/fullcycle-client-server-api/server/dollar"
)

func HandleDollarPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ctx := context.Background()
	conv, saveErr := dollar.SaveDollarPrice(ctx)
	if saveErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, writeErr := w.Write(conv)
	if writeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("sucessfull response")
}

func ReadDollar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ctx := context.Background()
	cotList, err := dollar.Read(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	convBytes, marshErr := json.Marshal(cotList)
	if marshErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, writeErr := w.Write(convBytes)
	if writeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
