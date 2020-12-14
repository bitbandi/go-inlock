package inlock

import "encoding/json"

type OrderBookFoo struct {
	Amount float64 `json:"amount"`
	Price  float64 `json:"price"`
	Own    bool    `json:"own"`
}

type OrderBook struct {
	MarketId int            `json:"marketid"`
	Market   string         `json:"market"`
	Sell     []OrderBookFoo `json:"sell"`
	Buy      []OrderBookFoo `json:"buy"`
}

func (i *Inlock) getTokenMarket(pair int) (orderbook OrderBook, err error) {
	r, err := i.client.do("GET", "private/getTokenMarket", nil, true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result.Result["getTokenMarket"], &orderbook)
	return
}

