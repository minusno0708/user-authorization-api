package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"testing"
	"bytes"

	"user-register-api/domain"
)

const endpoint = "http://localhost:8080"

var response struct {
	Message string `json:"message"`
	User domain.User `json:"user"`
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

func TestConnectionApi(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedMessage := "Connection Successful"

	resp, err := sendRequest("GET", endpoint, nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != expectedStatusCode {
		t.Fatalf("Expected status code %v, got %v", expectedStatusCode, resp.StatusCode)
	}
	responseData, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(responseData, &response)
	if err != nil {
		t.Fatal(err)
	}

  	if response.Message != expectedMessage {
		t.Fatalf("Expected message %v, got %v", expectedMessage, response.Message)
	}
}