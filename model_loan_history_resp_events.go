package swagger
/*
 * INLOCK API
 *
 * This is the documentation of the INLOCK API.
 *
 * API version: 1.4.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

type LoanHistoryRespEvents struct {
	// Unique ID of the event.
	EvId int64 `json:"ev_id,omitempty"`
	// Time of creation. Unix timestamp in UTC.
	Created int64 `json:"created,omitempty"`
	// Type of event.
	Type_ string `json:"type,omitempty"`
	// Amount related to the event. Optional, can be null.
	Amt string `json:"amt,omitempty"`
	// Ticker related to the event. Optional, can be null.
	Ticker string `json:"ticker,omitempty"`
}
