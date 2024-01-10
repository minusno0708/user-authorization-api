package main

import (
	"bytes"
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

func TestSigninBodyNotExist(t *testing.T) {
	expectedStatus := http.StatusBadRequest
	expectedMessage := "Body does not exist"

	req, err := http.NewRequest("POST", endpoint+"/signin", nil)
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != expectedStatus {
		t.Fatalf("Expected status code 400, got %v", resp.StatusCode)
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

func TestSigninUserIDNotExist(t *testing.T) {
	expectedStatus := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := &domain.User{
		Username: "testuser",
		Password: "testpass",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", endpoint+"/signin", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != expectedStatus {
		t.Fatalf("Expected status code 401, got %v", resp.StatusCode)
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

func TestSigninPasswordNotExist(t *testing.T) {
	expectedStatus := http.StatusUnauthorized
	expectedMessage := "Body is not valid"

	requestBody := &domain.User{
		UserID: "testuser",
		Username: "testuser",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", endpoint+"/signin", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != expectedStatus {
		t.Fatalf("Expected status code 401, got %v", resp.StatusCode)
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

func TestSigninSuccessUsernameExist(t *testing.T) {
	expectedStatus := http.StatusCreated
	expectedMessage := "User created successfully"

	requestBody := &domain.User{
		UserID: "testuser_id1",
		Username: "testusername",
		Password: "testpass",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", endpoint+"/signin", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != expectedStatus {
		t.Fatalf("Expected status code 201, got %v", resp.StatusCode)
	}

	responseData, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Message != expectedMessage {
		t.Fatalf("Expected message %v, got %v", expectedMessage, response.Message)
	}

	if response.User.UserID != requestBody.UserID {
		t.Fatalf("Expected user_id %v, got %v", requestBody.UserID, response.User.UserID)
	}
	if response.User.Username != requestBody.Username {
		t.Fatalf("Expected username %v, got %v", requestBody.Username, response.User.Username)
	}
	if response.User.Password != requestBody.Password {
		t.Fatalf("Expected password %v, got %v", requestBody.Password, response.User.Password)
	}
}

func TestSigninSuccessUsernameNotExist(t *testing.T) {
	expectedStatus := http.StatusCreated
	expectedMessage := "User created successfully"

	requestBody := &domain.User{
		UserID: "testuser_id2",
		Password: "testpass",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", endpoint+"/signin", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != expectedStatus {
		t.Fatalf("Expected status code 201, got %v", resp.StatusCode)
	}

	responseData, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Message != expectedMessage {
		t.Fatalf("Expected message %v, got %v", expectedMessage, response.Message)
	}

	if response.User.UserID != requestBody.UserID {
		t.Fatalf("Expected user_id %v, got %v", requestBody.UserID, response.User.UserID)
	}
	if response.User.Username != requestBody.UserID {
		t.Fatalf("Expected username %v, got %v", requestBody.UserID, response.User.Username)
	}
	if response.User.Password != requestBody.Password {
		t.Fatalf("Expected password %v, got %v", requestBody.Password, response.User.Password)
	}
}

func TestSigninUserConflict(t *testing.T) {
	expectedStatus := http.StatusConflict
	expectedMessage := "User already exists"

	requestBody := &domain.User{
		UserID: "testuser_id1",
		Username: "testusername",
		Password: "testpass",
	}

	jsonString, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", endpoint+"/signin", bytes.NewBuffer(jsonString))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != expectedStatus {
		t.Fatalf("Expected status code 209, got %v", resp.StatusCode)
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