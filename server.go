package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const MAX_REQUEST_LEN = 50 

func toJSON(m interface{}) string {
	js, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(string(js), ",", ", ", -1)
}

func sendResponse(w http.ResponseWriter, iCode int){
  w.WriteHeader(iCode)
  w.Write([]byte(http.StatusText(iCode)))
}

func handler(w http.ResponseWriter, r *http.Request) {
	str := r.URL.RawQuery
	if len(str) > MAX_REQUEST_LEN {
		sendResponse(w, http.StatusRequestEntityTooLarge)
	} else {
		if values, err := url.ParseQuery(str); err != nil {
			//return a bad query response
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 bad request" + err.Error()))
		} else {
			fmt.Fprintf(w, toJSON(values))
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
