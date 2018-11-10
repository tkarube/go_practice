package main

import (
    "log"
    "time"
)

func sub() {
    log.Println("sub() is running")
    time.Sleep(time.Second)
    log.Println("sub() is finished")
}

func main() {
    log.SetFlags(log.Lmicroseconds)
    for i := 0; i < 10; i++ {
        log.Println("start sub()")
        go sub()
        log.Println("hoge")
        time.Sleep(2 * time.Second)
    }
}
