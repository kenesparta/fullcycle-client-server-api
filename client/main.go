package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	price, reqErr := RequestDollarPrice(ctx)
	if reqErr != nil {
		log.Printf("error requesting: %v", reqErr)
		return
	}

	saveErr := saveToFile(price.Bid)
	if saveErr != nil {
		log.Printf("error saving file %v", saveErr)
		return
	}
}
