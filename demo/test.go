package main

import (
    "fmt"
    "strconv"
    "time"
)

func main() {
    t := time.Now()
    fmt.Println(t)

    fmt.Println(t.UTC().Format(time.UnixDate))

    fmt.Println(t.Unix())

    timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)
    fmt.Println(timestamp)
    timestamp = timestamp[:10]
    fmt.Println(timestamp)
}
