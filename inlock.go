// Package Inlock is an implementation of the Inlock API in Golang.
package inlock

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const (
	API_BASE = "https://api.inlock.io/inlock/api/v1.0" // Inlock API endpoint
)

// New returns an instantiated inlock struct
func New(apiKey, apiSecret string) *Inlock {
	client := NewClient(apiKey, apiSecret)
	return &Inlock{client}
}

// NewWithCustomHttpClient returns an instantiated inlock struct with custom http client
func NewWithCustomHttpClient(apiKey, apiSecret string, httpClient *http.Client) *Inlock {
	client := NewClientWithCustomHttpConfig(apiKey, apiSecret, httpClient)
	return &Inlock{client}
}

// NewWithCustomTimeout returns an instantiated inlock struct with custom timeout
func NewWithCustomTimeout(apiKey, apiSecret string, timeout time.Duration) *Inlock {
	client := NewClientWithCustomTimeout(apiKey, apiSecret, timeout)
	return &Inlock{client}
}

// handleErr gets JSON response from inlock API en deal with error
func handleErr(r jsonResponse) error {
	if r.Error.Code != 0 {
		return errors.New(r.Error.Message)
	}
	return nil
}

// inlock represent a inlock client
type Inlock struct {
	client *client
}

// set enable/disable http request/response dump
func (i *Inlock) SetDebug(enable bool) {
	i.client.debug = enable
}

type Tokens struct {
	AccessToken        string `json:"access_token"`
	RefreshToken       string `json:"refresh_token"`
	ExpiredIn          int    `json:"expired_in"`
}

func (i *Inlock) Login(username string, password string) (tokens Tokens, err error) {
	payload := map[string]string{
		"username": username,
		"password": password,
	}
	r, err := i.client.do("POST", "public/login", payload, true)
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
	err = json.Unmarshal(response.Result.Result["login"], &tokens)
	return
}

func (i *Inlock) RefreshToken(refreshToken string) (tokens Tokens, err error) {
	payload := map[string]string{
		"refresh_token": refreshToken,
	}
	r, err := i.client.do("POST", "public/refresh", payload, true)
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
	err = json.Unmarshal(response.Result.Result["refresh"], &tokens)
	return
}
