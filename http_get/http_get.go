package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://10.0.11.138")
	fmt.Println(resp)
}
