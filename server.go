package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	//	"strings"
	"github.com/DMXMax/tenableProxy/TenableLogs"
)

const MAX_REQUEST_LEN = 50

func toJSON(m interface{}) string {
	js, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return (string(js))
}

func sendResponse(w http.ResponseWriter, iCode int) {
	w.WriteHeader(iCode)
	w.Write([]byte(http.StatusText(iCode)))
}

func validateKeys(t url.Values) bool {
	for key := range t {
		fmt.Println(key)
	}
	return true
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/audit-events" {
		http.NotFound(w, r)
	} else {
		str := r.URL.RawQuery
		if len(str) > MAX_REQUEST_LEN {
			sendResponse(w, http.StatusRequestEntityTooLarge)
		} else {
			if _, err := url.ParseQuery(str); err != nil {
				//return a bad query response
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("400 bad request" + err.Error()))
			} else {
				//fmt.Fprintf(w, toJSON(values))
        //f := TenableLogs.Filter{str}dd
        if resp, err := TenableLogs.GetLogEntries(TenableLogs.Filter{str}); err == nil{

          fmt.Fprintf(w, string(resp))
        }else{
          fmt.Fprintf(w, err.Error())
        }
			}
		}
	}
}

func main() {
	http.HandleFunc("/", http.NotFound)
	http.HandleFunc("/audit-events", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
