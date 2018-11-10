// program_name <IP address of VS>
// send 10 HTTP GET requests in 3 seconds 

package main

import (
    "flag"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
    "time"
)

func sub(url string) {
    resp, err := http.Get(url)
    log.Println("HTTP GET for " + url)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    log.Println(strings.TrimRight(string(body),"\n"))
}

func main() {
    flag.Parse()
    url := flag.Arg(0)
    url = "http://" + url

    log.SetFlags(log.Lmicroseconds)
    for i := 0; i < 10; i++ {
        log.Println("start sub()")
        go sub(url)
        log.Println("hoge")
        time.Sleep(2 * time.Second)
    }
}
