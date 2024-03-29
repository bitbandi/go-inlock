package inlock
/*
 * INLOCK API
 *
 * This is the documentation of the INLOCK API.
 *
 * API version: 1.4.5
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

import (
	"encoding/json"
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type APIResult struct {
	Status       string `json:"status"`
	ResponseData map[string]json.RawMessage
}

type APIResponse struct {
	Error  APIError  `json:"error,omitempty"`
	Result APIResult `json:"result"`
}

func (rs *APIResult) UnmarshalJSON(b []byte) error {
	var r map[string]json.RawMessage
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}
	s, ok := r["status"]
	if ok {
		_ = json.Unmarshal(s, &rs.Status)
		delete(r, "status")
	}
	rs.ResponseData = r
	return nil
}
