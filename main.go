package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/countwords", CountWordsHandler).Methods(http.MethodPost, http.MethodOptions)
	// sets the Access-Control-Allow-Methods automatically for matched routes
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(corsMiddleware)
	fmt.Printf("Server started on port 8090.")
	log.Fatal(http.ListenAndServe(":8090", router))
}

func CountWordsHandler(writer http.ResponseWriter, request *http.Request) {
	requestJson, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	countWordsForm := TranslationQuoteRequest{}
	err = json.Unmarshal(requestJson, &countWordsForm)
	// TODO: better json validation
	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}
	quote := ComputeQuote(countWordsForm)
	responseJson, err := json.Marshal(quote)

	if err != nil {
		http.Error(writer, err.Error(), 500)
		return
	}

	writer.Header().Set("content-type", "application/json")
	_, _ = writer.Write(responseJson)
	request.Body.Close()
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4201")
		writer.Header().Set("Access-Control-Allow-Headers", "*")
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if request.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(writer, request)
	})
}
