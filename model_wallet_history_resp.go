package swagger
/*
 * INLOCK API
 *
 * This is the documentation of the INLOCK API.
 *
 * API version: 1.4.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

type WalletHistoryResp struct {
	// List of event records.
	Events []WalletHistoryRespEvents `json:"events,omitempty"`
}
