package inlock

import "encoding/json"

type SavingCoin struct {
	CoinId             int     `json:"coin_id"`
	Currency           string  `json:"currency"`
	Symbol             string  `json:"symbol"`
	Type               string  `json:"type"`
	Apy                float64 `json:"apy_kind,string"`
	ApyIlk             float64 `json:"apy_ilk,string"`
	CollBalance        float64 `json:"coll_balance,string"`
	FiatPrice          float64 `json:"fiat_price,string"`
	PendingBalance     float64 `json:"pending_balance,string"`
	PendingInt         float64 `json:"pending_int,string"`
	PendingIntFiat     float64 `json:"pending_int_fiat,string"`
	SavingsBalance     float64 `json:"savings_balance,string"`
	SavingsBalanceFiat float64 `json:"savings_balance_fiat,string"`
	Balance            float64 `json:"tm_balance,string"`
	Total              float64 `json:"total_int,string"`
	TotalFiat          float64 `json:"total_int_fiat,string"`
}

type SavingOverview struct {
}

type getSavingsDataResult struct {
	Coins []SavingCoin `json:"coins"`
}

func (i *Inlock) GetSavingsData(fiat string) (SavingCoins []SavingCoin, err error) {
	r, err := i.client.do("GET", "private/getSavingsData", map[string]string{"fiat": fiat}, true)
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
	var res getSavingsDataResult
	err = json.Unmarshal(response.Result.Result["getSavingsData"], &res)
	return res.Coins, err
}
