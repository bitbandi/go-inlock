package swagger
/*
 * INLOCK API
 *
 * This is the documentation of the INLOCK API.
 *
 * API version: 1.4.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

type SwapHistoryRespEvents struct {
	// Unique ID of event.
	EvId int64 `json:"ev_id,omitempty"`
	// Time of creation. Unix timestamp in UTC.
	Created int64 `json:"created,omitempty"`
	// Type of event.
	Type_ string `json:"type,omitempty"`
	// From amount.
	FromAmt string `json:"from_amt,omitempty"`
	// From ticker.
	FromTicker string `json:"from_ticker,omitempty"`
	// To ticker.
	ToTicker string `json:"to_ticker,omitempty"`
	// Price.
	Price string `json:"price,omitempty"`
	// To amount.
	ToAmt string `json:"to_amt,omitempty"`
	// From coin is the primary (e.g. for price display) or the to coin.
	FromPrimary bool `json:"from_primary,omitempty"`
}
