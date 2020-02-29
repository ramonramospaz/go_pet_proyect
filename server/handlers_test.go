package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestListUserService(t *testing.T) {

	// Borrar todos los otros casos

	req, err := http.NewRequest("DELETE", "/user", nil)

	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	deleteAllUsers(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("The users cant be deleted")
	}

	// Test listUsers with no data
	expected := []byte("[]\n")

	req, err = http.NewRequest("GET", "/user", nil)

	if err != nil {
		t.Fatal(err)
	}

	res = httptest.NewRecorder()

	listUsers(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("The list of user fail")
	}

	if bytes.Compare(expected, res.Body.Bytes()) != 0 {
		t.Errorf(" Expected %v, Return %v, %v, %v, %v", string(expected),
			res.Body,
			bytes.Compare(expected,
				res.Body.Bytes()), expected, res.Body.Bytes())
	}

	// Test Include Data
	req, err = http.NewRequest("POST", "/user", strings.NewReader(`{
		"id":15281565,          
		"name":"Ramon Ramos",        
		"age":38,        
		"birthCountry":"Venezuela"
	}`))

	res = httptest.NewRecorder()

	createUser(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("The user creation fail")
	}

	// Test listUsers with data
	expected = []byte("[{\"id\":15281565,\"name\":\"Ramon Ramos\",\"age\":38,\"birthCountry\":\"Venezuela\"}]\n")

	req, err = http.NewRequest("GET", "/user", nil)

	if err != nil {
		t.Fatal(err)
	}

	res = httptest.NewRecorder()

	listUsers(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("The list of user fail")
	}

	if bytes.Compare(expected, res.Body.Bytes()) != 0 {
		t.Errorf(" Expected %v, Return %v, %v, %v, %v", string(expected),
			res.Body,
			bytes.Compare(expected,
				res.Body.Bytes()), expected, res.Body.Bytes())
	}

}

func TestFindUserService(t *testing.T) {

	// Borrar todos los otros casos

	req, err := http.NewRequest("DELETE", "/user", nil)

	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()

	deleteAllUsers(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("The users cant be deleted")
	}

	// Test Include Data 1
	req, err = http.NewRequest("POST", "/user", strings.NewReader(`{
		"id":15281565,          
		"name":"Ramon Ramos",        
		"age":38,        
		"birthCountry":"Venezuela"
	}`))

	res = httptest.NewRecorder()

	createUser(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("The user creation fail")
	}

	// Test Include Data 2
	req, err = http.NewRequest("POST", "/user", strings.NewReader(`{
		"id":15281567,          
		"name":"Fabiola Ramos",        
		"age":37,        
		"birthCountry":"Venezuela"
	}`))

	res = httptest.NewRecorder()

	createUser(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("The user creation fail")
	}

	// Test listUsers with no data
	expected := []byte("[{\"id\":15281565,\"name\":\"Ramon Ramos\",\"age\":38,\"birthCountry\":\"Venezuela\"}]\n")

	req, err = http.NewRequest("GET", "/user/15281565", nil)

	if err != nil {
		t.Fatal(err)
	}

	res = httptest.NewRecorder()

	findUserByID(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("The list of user fail")
	}

	if bytes.Compare(expected, res.Body.Bytes()) != 0 {
		t.Errorf(" Expected %v, Return %v, %v, %v, %v", string(expected),
			res.Body,
			bytes.Compare(expected,
				res.Body.Bytes()), expected, res.Body.Bytes())
	}

}
