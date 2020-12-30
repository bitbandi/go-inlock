package inlock

import "encoding/json"

type getBestAvailableOffersResult struct {
	BestApr float64 `json:"bestApr,string"`
}

func (i *Inlock) GetBestAvailableOffers() (bestAvailableOffers float64, err error) {
	r, err := i.client.do("GET", "public/getBestAvailableOffers", nil, false)
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
	var res getBestAvailableOffersResult
	err = json.Unmarshal(response.Result.Result["getBestAvailableOffers"], &res)
	return res.BestApr, err
}

type Currency struct {
	CoinId   int    `json:"coin_id"`
	Ticker   string `json:"ticker"`
	LongName string `json:"long_name"`
}

type Fiats struct {
	Currency
}

type getFiatsResult struct {
	Fiats []Fiats `json:"fiats"`
}

func (i *Inlock) GetFiats() (Fiats []Fiats, err error) {
	r, err := i.client.do("GET", "public/getFiats", nil, false)
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
	var res getFiatsResult
	err = json.Unmarshal(response.Result.Result["getFiats"], &res)
	return res.Fiats, err
}

type Coins struct {
	Currency
	Type string `json:"type"`
}

type getAvailableCoinsResult struct {
	AvailableCoins []Coins `json:"available-coins"`
}

func (i *Inlock) GetAvailableCoins() (AvailableCoins []Coins, err error) {
	r, err := i.client.do("GET", "public/getavailablecoins", nil, false)
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
	var res getAvailableCoinsResult
	err = json.Unmarshal(response.Result.Result["getavailablecoins"], &res)
	return res.AvailableCoins, err
}

