package search

import "log"


type Result structure {
    File string
    content string
}

type Matcher interface {
    Search(feed *Feed, searchTerm string) ([]*result,error)
    func Match(matcher Matcher,feed *Feed, searchTerm string)([]*Result, error)
}

func Match (matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result){
    searchResults, err :=matcher.Search(feed, earchTerm)
    if err !=nil{
        log.Println(err)
        return
    }

    for _, result := range searchResults { 
        cfo
    }

}

func Display(results chan *Result){
    for result := range results {
        log.Printf("%s:\n%s\n\n", result.Field, result.Content)
    }

}
