package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"os/exec"
)

func main() {
	resp, err := http.Get("http://hgdownload.soe.ucsc.edu/goldenPath/hg38/database")
	if err != nil {
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error.")
	}
	src := string(body)
	re, _ := regexp.Compile(">refGene.sql<")
	src = re.FindString(src)
	fmt.Println(src)
	//url := "http://hgdownload.soe.ucsc.edu/goldenPath/hg38/database/refGene.sql"
	//downloadSql(url)
	cmd := "wget http://hgdownload.soe.ucsc.edu/goldenPath/hg38/database/refGene.sql"
	f, err :=exec.LookPath(cmd) 
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
}


func downloadSql (url string) {
	fs := http.FileServer(http.Dir(url))
	http.Handle("./", http.StripPrefix("./", fs))
}
