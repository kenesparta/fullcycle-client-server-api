package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/kenesparta/fullcycle-client-server-api/server/dollar"
)

func RequestDollarPrice(ctx context.Context) (*dollar.Cotacao, error) {
	ctx, cancel := context.WithTimeout(ctx, MaxRequestTimeout)
	defer cancel()
	client := http.Client{Timeout: MaxRequestTimeout}
	request, reqErr := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		UrlServer,
		nil)
	if reqErr != nil {
		return nil, reqErr
	}

	request.Header.Set("Accept", "application/json")
	response, respErr := client.Do(request)
	if respErr != nil {
		log.Printf("timeout error %v", respErr)
		return nil, respErr
	}

	if response.StatusCode < http.StatusOK || response.StatusCode >= 300 {
		return nil, fmt.Errorf("bad response, error: %v", response.StatusCode)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("close stream error %v", err)
			return
		}
	}(response.Body)

	readBody, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		log.Printf("read error %v", readErr)
		return nil, readErr
	}

	var conv dollar.Cotacao
	unmErr := json.Unmarshal(readBody, &conv)
	if unmErr != nil {
		return nil, unmErr
	}

	return &conv, nil
}
