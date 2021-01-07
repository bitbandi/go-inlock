package inlock

import (
	"encoding/json"
	"strconv"
)

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

type OffersRates struct {
	Ticker  string  `json:"ticker"`
	ApyKind float64 `json:"apy_kind,string"`
	ApyIlk  int     `json:"apy_ilk,string"`
}

type Offers struct {
	IlkReq int           `json:"ilk_req,string"`
	Rates  []OffersRates `json:"rates"`
}

type offersResult struct {
	Offers []Offers `json:"offers"`
}

func (i *Inlock) GetAutoLendApr() (Offers []Offers, err error) {
	r, err := i.client.do("GET", "public/getAutoLendApr", nil, false)
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
	var res offersResult
	err = json.Unmarshal(response.Result.Result["getAutoLendApr"], &res)
	return res.Offers, err
}

type LoanOffers struct {
	LoanAmt float64     `json:"loan_amt"`
	IntDebt float64 `json:"int_debt"`
	Apr     float64 `json:"apr"`
	CollId  int     `json:"coll_id"`
	CollAmt float64 `json:"coll_amt"`
	CostIlk float64 `json:"cost_ilk"`
	CostUsd float64 `json:"cost_usd"`
}

func (i *Inlock) GetPubCustIndLoanOffers(Amt float64, OverCollaterization int, Duration int, CollaterizationId int) (loanOffers LoanOffers, err error) {
	params := map[string]string{
		"coin_id": "9",
		"amt":     strconv.FormatFloat(Amt, 'f', 8, 64),
		"oc":      strconv.FormatUint(uint64(OverCollaterization), 10),
		"dur":     strconv.FormatUint(uint64(Duration), 10),
		"coll":    strconv.FormatUint(uint64(CollaterizationId), 10),
	}
	r, err := i.client.do("GET", "public/getPubCustIndLoanOffers", params, false)
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
	var res LoanOffers
	err = json.Unmarshal(response.Result.Result["getPubCustIndLoanOffers"], &res)
	return res, err
}
