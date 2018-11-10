package main

import (
    "flag"
    "fmt"
)

func main() {
    flag.Parse()
    args := flag.Args()
    fmt.Println(flag.Arg(0))
    fmt.Println(args[0])
}
