package inlock

import (
	"encoding/json"
)

type responseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type resultStatus struct {
	Status string                     `json:"status"`
	Result map[string]json.RawMessage `json:"-"`
}

type jsonResponse struct {
	Error  responseError `json:"error"`
	Result resultStatus  `json:"result"`
}
