// program_name <IP address of VS>

package main

import (
        "log"
        "io/ioutil"
        "flag"
        "net/http"
        "strings"
        "time"
)

var client = http.Client {
    Timeout: time.Millisecond * 100,
    Transport: &http.Transport{
        MaxIdleConnsPerHost: 32,
        DisableKeepAlives: true,
    },
}

func get_error (err error) {
    panic(err)
}

func http_request (url string) {
    log.Println("HTTP GET for " + url)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        get_error(err)
    }
    resp, err := client.Do(req)
    if err != nil {
        get_error(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        get_error(err)
    }
    log.Println(strings.TrimRight(string(body),"\n"))
}

func main() {
    flag.Parse()
    url := flag.Arg(0)
    url = "http://" + url
    log.SetFlags(log.Lmicroseconds)

    for i := 0; i < 300; i++ {
        go http_request(url)
        time.Sleep(100 * time.Millisecond)
    }
}
