package main

import (
	"fmt"
	"time"

	"github.com/kenesparta/fullcycle-client-server-api/server/constants"
)

var UrlServer = fmt.Sprintf("http://127.0.0.1%s/cotacao", constants.Port)

const (
	MaxRequestTimeout = 300 * time.Millisecond
	TemplateName      = "bid.go.tmpl"
	FileNameCotacao   = "cotacao.txt"
)
