package main

import (
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"testing"

	"user-register-api/domain"
)

func TestDeleteUserParamsNotExist(t *testing.T) {
	expectedStatusCode := http.StatusNotFound
	expectedMessage := "404 page not found"

	resp, err := sendRequest("DELETE", endpoint+"/user", nil)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != expectedStatusCode {
		t.Fatalf("Expected status code %v, got %v", expectedStatusCode, resp.StatusCode)
	}

	responseData, _ := ioutil.ReadAll(resp.Body)
	
	responseMessage := string(responseData)
  	if responseMessage != expectedMessage {
		t.Fatalf("Expected message %v, got %v", expectedMessage, responseMessage)
	}
}

func TestDeleteUserBodyNotExist(t *testing.T) {
	expectedStatusCode := http.StatusBadRequest
	expectedMessage := "Body does not exist"

	resp, err := sendRequest("DELETE", endpoint+"/user/"+userID, nil)
	if err != nil {
		t.Fatal(err)
	}
	
	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteUserPasswordNotExist(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := &domain.User{}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("DELETE", endpoint+"/user/"+userID, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}
	
	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteUserUserNotFound(t *testing.T) {
	expectedStatusCode := http.StatusNotFound
	expectedMessage := "User not found"

	requestBody := &domain.User{
		Password: "testpass",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("DELETE", endpoint+"/user/"+"not_exist_user", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}
	
	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteUserPasswordNotCorrect(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedMessage := "Password is incorrect"

	requestBody := &domain.User{
		Password: "not_correct_pass",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("DELETE", endpoint+"/user/"+userID, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}
	
	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteUserSuccess(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedMessage := "User can be deleted"

	requestBody := &domain.User{
		Password: "testpass",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := sendRequest("DELETE", endpoint+"/user/"+userID, bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}
	
	err = verifyExpectedResponse(resp, expectedStatusCode, expectedMessage)
	if err != nil {
		t.Fatal(err)
	}
}
