package main

import(
    "log"
    "os"

    _"GOlenrning/matchers"
    "GOlenrning/search"
)

func init(){
    log.SetOutput(os.Stdout)

}

func main(){
    search.Run("president")
}
