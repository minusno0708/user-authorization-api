package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"user-register-api/domain"
)

const endpoint = "http://localhost:8080"

type requestBody struct {
	TokenString string `json:"token"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type responseBody struct {
	User        domain.User `json:"user"`
	TokenString string      `json:"token"`
}

type errorString struct {
	message string
}

func (e *errorString) Error() string {
	return e.message
}

type header struct {
	key   string
	value string
}

func setToken(tokenString string) *header {
	return &header{
		key:   "Token",
		value: tokenString,
	}
}

func (h *header) ToArray() []*header {
	return []*header{h}
}

func sendRequest(method string, endpoint string, header []*header, sendingBody *bytes.Buffer) (*http.Response, error) {
	var req *http.Request
	var err error

	if sendingBody != nil {
		req, err = http.NewRequest(method, endpoint, sendingBody)
	} else {
		req, err = http.NewRequest(method, endpoint, nil)
	}

	for _, h := range header {
		req.Header.Set(h.key, h.value)
	}

	if err != nil {
		return nil, err
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func verifyExpectedResponse(resp *http.Response, expectedStatusCode int, expectedMessage string) (*responseBody, error) {
	var response struct {
		Message     string      `json:"message"`
		User        domain.User `json:"user"`
		TokenString string      `json:"token"`
	}

	if resp.StatusCode != expectedStatusCode {
		return nil, &errorString{message: fmt.Sprintf("Expected status code %v, got %v", expectedStatusCode, resp.StatusCode)}
	}

	responseData, _ := ioutil.ReadAll(resp.Body)

	err := json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, err
	}

	if response.Message != expectedMessage {
		return nil, &errorString{message: fmt.Sprintf("Expected message %v, got %v", expectedMessage, response.Message)}
	}

	responseBody := &responseBody{
		User:        response.User,
		TokenString: response.TokenString,
	}

	return responseBody, nil
}

func TestConnectionApi(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedMessage := "Connection Successful"

	resp, err := sendRequest("GET", endpoint, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}
