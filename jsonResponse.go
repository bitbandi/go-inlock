package inlock

import (
	"encoding/json"
)

type responseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type resultStatus struct {
	Status string `json:"status"`
	Result map[string]json.RawMessage
}

type jsonResponse struct {
	Error  responseError `json:"error"`
	Result resultStatus  `json:"result"`
}

func (rs *resultStatus) UnmarshalJSON(b []byte) error {
	var r map[string]json.RawMessage
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}
	s, ok := r["status"]
	if ok {
		_ = json.Unmarshal(s, rs.Status)
		delete(r, "status")
	}
	rs.Result = r
	return nil
}
