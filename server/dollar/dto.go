package dollar

type Cotacao struct {
	Id         string  `json:"id"`
	Code       string  `json:"code"`
	Codein     string  `json:"codein"`
	Name       string  `json:"name"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	VarBid     float64 `json:"varBid"`
	PctChange  float64 `json:"pctChange"`
	Bid        float64 `json:"bid"`
	Ask        float64 `json:"ask"`
	Timestamp  string  `json:"timestamp"`
	CreateDate string  `json:"create_date"`
}

type Conversion struct {
	Id   string   `json:"id"`
	Data ConvData `json:"USDBRL"`
}

type ConvData struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}
