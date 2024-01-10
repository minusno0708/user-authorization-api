package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"testing"

	"user-register-api/domain"
)

const endpoint = "http://localhost:8080"

var response struct {
	Message string `json:"message"`
	User domain.User `json:"user"`
}

func TestConnectionApi(t *testing.T) {
	expectedStatus := http.StatusOK
	expectedMessage := "Connection Successful"

	resp, err := http.Get(endpoint)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != expectedStatus {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
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