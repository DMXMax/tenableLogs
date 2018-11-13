package TenableLogs 

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Filter struct {
	FilterString string
}

func (f Filter) getFilterString() string {
	if f.FilterString == "" {
		return ""
	} else {

		return "?" + f.FilterString
	}
}


var keyHeader strings.Builder

func GetLogEntries(f Filter) ([]byte, error) {

	url := "https://cloud.tenable.com/audit-log/v1/events" + f.getFilterString()

	fmt.Println(url)

	req, _ := http.NewRequest("GET", url, nil)
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")

	if accessKey == "" {
		return nil, errors.New("AK not defined")
	}

	if secretKey == "" {
		return nil, errors.New("SK not defined")
	}

	fmt.Fprintf(&keyHeader, "accessKey=%s; secretKey=%s", accessKey, secretKey)

	req.Header.Add("x-apikeys", keyHeader.String())

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body, nil
}
