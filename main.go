package main

import (
	"encoding/json"
	"fmt"
)

type EventList struct {
	Events []struct {
		Id         string
		Action     string
		Crud       string
		Received   string
		Is_Failure bool
	}
	Pagination struct {
		Total int
		Limit int
	}
}

var el EventList

func main() {
	//var f =Filter{"f=date.gt:2018-11-08"}
	var f = Filter{"limit=98"}

	if str, err := GetLogEntries(f); err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(string(str))
		json.Unmarshal(str, &el)
		if printable, err := json.MarshalIndent(el, "", "\t"); err != nil {
			panic(err)
		} else {
			fmt.Println(string(printable))
		}
		//fmt.Println(el)
		//fmt.Println(len(el.Events))
	}
}
