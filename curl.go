package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://www.jayteesf.com"

func main() {
	//client := &http.Client{}
	//resp, err := client.Get(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("doh: ", err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", body)
	}
}
