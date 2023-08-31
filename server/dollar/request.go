package dollar

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/kenesparta/fullcycle-client-server-api/server/constants"
	"github.com/kenesparta/fullcycle-client-server-api/server/utils"
)

func RequestDollarPrice(ctx context.Context) (*Conversion, error) {
	ctx, cancel := context.WithTimeout(ctx, constants.MaxRequestTimeout)
	defer cancel()
	client := http.Client{Timeout: constants.MaxRequestTimeout}
	request, reqErr := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		constants.URLCotacao,
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

	var conv Conversion
	unmErr := json.Unmarshal(readBody, &conv)
	if unmErr != nil {
		return nil, unmErr
	}

	fmt.Println(string(readBody))
	return &conv, nil
}

func SaveDollarPrice(ctx context.Context) ([]byte, error) {
	conv, reqErr := RequestDollarPrice(ctx)
	if reqErr != nil {
		return nil, reqErr
	}

	conv.Id = uuid.New().String()
	cotacao := Cotacao{
		Id:         conv.Id,
		Code:       conv.Data.Codein,
		Codein:     conv.Data.Codein,
		Name:       conv.Data.Name,
		High:       utils.StrToFloat(conv.Data.High),
		Low:        utils.StrToFloat(conv.Data.Low),
		VarBid:     utils.StrToFloat(conv.Data.VarBid),
		PctChange:  utils.StrToFloat(conv.Data.PctChange),
		Bid:        utils.StrToFloat(conv.Data.Bid),
		Ask:        utils.StrToFloat(conv.Data.Ask),
		Timestamp:  conv.Data.Timestamp,
		CreateDate: conv.Data.CreateDate,
	}

	saveErr := Save(ctx, cotacao)
	if saveErr != nil {
		return nil, saveErr
	}

	convBytes, marshErr := json.Marshal(cotacao)
	if marshErr != nil {
		return nil, marshErr
	}

	return convBytes, nil
}
