package inlock

import "encoding/json"

func (i *Inlock) GetAvailableCollateralCoins() (coins []int, err error) {
	r, err := i.client.do("GET", "public/getAvailableCollateralCoins", nil, false)
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
	err = json.Unmarshal(response.Result.Result["getAvailableCollateralCoins"], &coins)
	return
}
