package inlock
/*
 * INLOCK API
 *
 * This is the documentation of the INLOCK API.
 *
 * API version: 1.4.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

type RetailStartWithdrawBody struct {
	// Withdrawal address.
	Address string `json:"address"`
	// Ticker of the coin.
	Ticker string `json:"ticker"`
	// network of chain of a multichain asset.
	Network string `json:"network"`
	// Amount to withdraw.
	Amount string `json:"amount"`
}
