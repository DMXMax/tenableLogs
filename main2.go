package main

import (
	"fmt"
)

func main() {
	var f = Filter{"limit:30&f=date.gt:2018-11-08"}

	url := "https://cloud.tenable.com/audit-log/v1/events" + f.getFilterString()
	fmt.Println(url + f.getFilterString())
}
