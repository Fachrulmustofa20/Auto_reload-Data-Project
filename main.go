package main

import (
	"assignment3/helper"
	"fmt"
	"net/http"
)

func main() {
	go helper.CreateJson()
	http.HandleFunc("/", helper.ReloadWeb)
	fmt.Println("Running in port 8080")
	http.ListenAndServe(":8080", nil)
}
