package inlock
/*
 * INLOCK API
 *
 * This is the documentation of the INLOCK API.
 *
 * API version: 1.4.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

type OrderBookRespSell struct {
	// Amount of the order
	Amount string `json:"amount,omitempty"`
	// Price of the order
	Price string `json:"price,omitempty"`
	// the order is the user's own order
	Own bool `json:"own,omitempty"`
}
