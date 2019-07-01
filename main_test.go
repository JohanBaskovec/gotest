package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCountWordsHandler(t *testing.T) {
	bodyObject := TranslationQuoteRequest{
		Source:       "hello world hello ",
		PricePerWord: 0.54,
	}
	body, err := json.Marshal(bodyObject)
	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequest("POST", "/countwords", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CountWordsHandler)

	handler.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expectedObject := TranslationQuote{
		Words: map[string]int{
			"hello": 2,
			"world": 1,
		},
		TotalPrice: 0.54 * 2,
	}
	expectedJsonBytes, err := json.Marshal(expectedObject)
	if err != nil {
		t.Fatal(err)
	}
	expected := string(expectedJsonBytes)
	if responseRecorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseRecorder.Body.String(), expected)
	}
}
