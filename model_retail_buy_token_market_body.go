package inlock
/*
 * INLOCK API
 *
 * This is the documentation of the INLOCK API.
 *
 * API version: 1.4.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

type BuyTokenMarketBody struct {
	// Amount to buy.
	Amount string `json:"amount"`
	// Trading pair of the market e.g.: XBTILK.
	Pair string `json:"pair"`
	// Limit of cost. Spend at most limit amount.
	Limit string `json:"limit,omitempty"`
	// Limit of price. The average price of the buy order is capped at this value. limit and limit_price can not be specified at the same time.
	LimitPrice string `json:"limit_price,omitempty"`
}
