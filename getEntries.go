package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
  "os"
  "strings"
)
var keyHeader strings.Builder

func main() {

	url := "https://cloud.tenable.com/audit-log/v1/events"

	req, _ := http.NewRequest("GET", url, nil)
  accessKey := os.Getenv("ACCESS_KEY")  
  secretKey := os.Getenv("SECRET_KEY")

  if accessKey == ""{
    panic("AK not defined")
  }

  if secretKey == ""{
    panic("SK not defined")
  }

  fmt.Fprintf(&keyHeader, "accessKey=%s; secretKey=%s", accessKey, secretKey)

	req.Header.Add("x-apikeys", keyHeader.String())

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
