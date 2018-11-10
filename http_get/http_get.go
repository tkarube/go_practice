// program_name <IP address of VS>

package main

import (
        "log"
        "flag"
        "net/http"
)

func main() {
    flag.Parse()
    url := flag.Arg(0)
    url = "http://" + url
    resp, err := http.Get(url)
    log.Println("HTTP GET for " + url)
    if err != nil {
        panic(err)
    }
    log.Println(resp)
}
