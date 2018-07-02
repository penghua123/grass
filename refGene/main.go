package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"os/exec"
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
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
	cmd := "wget /usr/local/bin/wget http://hgdownload.soe.ucsc.edu/goldenPath/hg38/database/refGene.sql"
	f, err :=exec.LookPath(cmd) 
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
	CollectError("huap")
}


func downloadSql (url string) {
	fs := http.FileServer(http.Dir(url))
	http.Handle("./", http.StripPrefix("./", fs))
}

func CollectError (user string){
	db, err := sql.Open("mysql", user+":VMw@re.c0m00@tcp(localhost:3306)/test?charset=utf8")
	checkErr(err)
	fmt.Println("MySQL collected successed!")

	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("XXXXX", "Gopher", "2012-12-09")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func checkErr(err error) {
    if err != nil {
        fmt.Println("Error is ", err)
    }
}

