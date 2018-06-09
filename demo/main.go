package main

import (
    "fmt"
    "github.com/couchbase/go-couchbase"
)

func main(){
    bucket,err :=couchbase.GetBucket("http://hua:penghua@10.15.158.52:8091/","default","hua")
    if err !=nil {
        fmt.Println(err)
    }
    defer bucket.Close()
    fmt.Println("Connected to Couchbase Bucket '%s'\n",bucket.Name)
}
