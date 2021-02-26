package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

func TestHandleFunc_POST_Success(t *testing.T) {
	w := httptest.NewRecorder()
	data := url.Values{}
	data.Set("first_name", "John")
	data.Set("last_name", "Doe")
	data.Set("email", "email@example.com")
	data.Set("phone_number", "0819999999")
	req, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	oContent, err := ioutil.ReadFile(dataFile)
	if err != nil {
		t.Fail()
	}

	defer ioutil.WriteFile(dataFile, oContent, os.ModeAppend)

	handleFunc(w, req)
	if w.Code != http.StatusOK {
		t.Fail()
	}

	content, err := ioutil.ReadFile(dataFile)
	if err != nil {
		t.Fail()
	}
	var forms []formInput
	err = json.Unmarshal(content, &forms)
	if err != nil {
		t.Fail()
	}
	form := forms[len(forms)-1]
	if form.FirstName != data.Get("first_name") ||
		form.LastName != data.Get("last_name") ||
		form.Email != data.Get("email") ||
		form.PhoneNumber != data.Get("phone_number") {
		t.Fail()
	}
}

func TestHandleFunc_POST_BadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	data := url.Values{}
	data.Set("first_name", "John")
	data.Set("last_name", "Doe")
	data.Set("email", "email@example.com")
	// data.Set("phone_number", "0819999999")
	req, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	handleFunc(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func TestHandleFunc_GET_Success(t *testing.T) {
	w := httptest.NewRecorder()

	cases := []string{"/", "/index.html", "/form.html"}
	for _, c := range cases {
		req, _ := http.NewRequest(http.MethodGet, c, nil)
		handleFunc(w, req)
		if w.Code != http.StatusOK {
			t.Fail()
		}
	}
}

func TestHandleFunc_GET_NotFound(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/norm.html", nil)
	handleFunc(w, req)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestHandleFunc_PUT_NotFound(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/", nil)
	handleFunc(w, req)
	if w.Code != http.StatusNotFound {
		t.Fail()
	}
}

func TestRun(t *testing.T) {
	oLoadEnv := loadEnv
	loadEnv = func(filename ...string) (err error) {
		os.Setenv("PORT", "8080")
		return
	}
	defer func() {
		loadEnv = oLoadEnv
		r := recover()
		if r != nil {
			t.Fail()
		}
	}()
	srv := run()
	time.Sleep(1 * time.Second)
	srv.Shutdown(context.TODO())
}
