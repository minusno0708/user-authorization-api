package main

import (
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"testing"
	"fmt"

	"user-register-api/domain"
)

const endpoint = "http://localhost:8080"

var response struct {
	Message string `json:"message"`
	User domain.User `json:"user"`
}

type errorString struct {
	message string
}

func (e *errorString) Error() string {
	return e.message
}

func sendRequest(method string, endpoint string, jsonBody *bytes.Buffer) (*http.Response, error) {
	var req *http.Request
	var err error

	if jsonBody != nil {
		req, err = http.NewRequest(method, endpoint, jsonBody)
	} else {
		req, err = http.NewRequest(method, endpoint, nil)
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

func verifyExpectedResponse(resp *http.Response, expectedStatusCode int, expectedMessage string) error {
	if resp.StatusCode != expectedStatusCode {
		return &errorString{message: fmt.Sprintf("Expected status code %v, got %v", expectedStatusCode, resp.StatusCode)}
	}

	responseData, _ := ioutil.ReadAll(resp.Body)

	err := json.Unmarshal(responseData, &response)
	if err != nil {
		return err
	}

	if response.Message != expectedMessage {
		return &errorString{message: fmt.Sprintf("Expected message %v, got %v", expectedMessage, response.Message)}
	}

	return nil
}

func TestConnectionApi(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedMessage := "Connection Successful"

	resp, err := sendRequest("GET", endpoint, nil)
	if err != nil {
		t.Fatal(err)
	}
	
	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}