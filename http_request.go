package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	req_type := os.Args[1]
	url := os.Args[2]
	
	if req_type == "GET" {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		
		defer resp.Body.Close()
		
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Printf(string(body))
	}
}