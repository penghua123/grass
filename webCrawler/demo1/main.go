package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"

	_ "github.com/lib/pq"
	"github.com/opesun/goquery"
)

type news struct {
	title string
	url   string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "huap"
	password = ""
	dbname   = "test"
)

func saveNewsDB(newsChan chan news, wSaveOk *sync.WaitGroup) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stmt := "CREATE TABLE IF NOT EXISTS news(_id INTEGER PRIMARY KEY,title TEXT NOT NULL,url TEXT NOT NULL);"
	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	for e := range newsChan {
		stmt = fmt.Sprintf("INSERT INTO news (title,url) VALUES ($1,$2)")
		_, err = db.Exec(stmt, e.title, e.url)
		if err != nil {
			log.Fatal(err)
		}
	}
	wSaveOk.Done()
}

var count int64

func getBaiduNews(index chan int, newsChan chan news, wGetOk *sync.WaitGroup) {
	for {
		offset, ok := <-index
		if !ok {
			break
		}
		var stringBuffer bytes.Buffer
		stringBuffer.WriteString("https://www.baidu.com/home/pcweb/data/mancardwater?id=2&offset=")
		stringBuffer.WriteString(strconv.Itoa(offset))
		stringBuffer.WriteString("&sessionId=15180565112719&crids=&version=&pos=52&newsNum=52&blacklist_timestamp=0&indextype=manht&_req_seqid=0xab0aac7f0000ef5d&asyn=1&t=1518056583617&sid=1428_21082_20719")
		url := stringBuffer.String()
		payload := strings.NewReader("params=lKyJiYt65SeTS%252BaO4ZqdonwEyY%252BSzr7oFr8kEpP8j0H%252FqMMPELQv9UibydyQRz10kcdSQHQmgvca2yvKfGSX5R%252BU8ByLp6rS4CRiH%252B%252FYxok%253D%26&encSecKey=1dd9fd1745dae8af1c0a678baf62803becdabb6685b9cf756ce101accb9daa291f408a848e84d83a344fe98db6d3ea2abf63a278f98191c4234bd201a5bbfba1faadc509bacde313e693ecf0aceace909b5e8a168be20e34b0eef3640a45b075a6b4c1ff581cea91debaa69d125326e218d09bb01cc490ad09fe5c1d24746047")
		req, _ := http.NewRequest("GET", url, payload)
		req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
		req.Header.Add("accept-encoding", "gzip, deflate, br")
		req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8")
		req.Header.Add("cache-control", "no-cache")
		req.Header.Add("connection", "keep-alive")
		req.Header.Add("cookie", "BDUSS=WoxNkZQYXdTfkdjWUd3YUVDNTIwTkV-NlpRRTYtbWhPOUEwZ356OGdyY0h4ak5aSVFBQUFBJCQAAAAAAAAAAAEAAACtM2Q2am92ZXpnAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAc5DFkHOQxZd; __cfduid=dd4f0bca4c50d4c7833592743d9934d741500540178; BAIDUID=43DFC2B31E07D893EC68CC4C31A5C7DC:FG=1; PSTM=1504776284; BIDUPSID=4F2127EE6005F11465294EE835520DE7; BD_UPN=123353; pgv_pvi=5320228864; BDRCVFR[e7VUaW6Ywr3]=aeXf-1x8UdYcs; BD_HOME=1; BD_CK_SAM=1; BDRCVFR[Oi7iajNidCC]=9xWipS8B-FspA7EnHc1QhPEUf; PSINO=5; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; H_PS_PSSID=1428_21082_20719; sug=3; sugstore=1; ORIGIN=2; bdime=21110")
		req.Header.Add("host", "www.baidu.com")
		req.Header.Add("upgrade-insecure-requests", "1")
		req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/64.0.3282.119 Chrome/64.0.3282.119 Safari/537.36")
		req.Header.Add("postman-token", "0b307632-c6ae-bb88-6a78-cb196aa43b4e")
		res, err := http.DefaultClient.Do(req)
		if err == nil && res.StatusCode == 200 {
			bytes, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			info := string(bytes)
			if strings.Contains(info, "\"errNo\": \"0\"") {
				info := strings.Split(strings.Split(info, "\",\"isEnd\": \"0\",")[0], "{\"errNo\": \"0\",\"html\" : \"")[1]
				nodes, parseError := goquery.Parse(strings.NewReader(info))
				if parseError == nil {
					nodes.Find("a.s-title-yahei").Each(func(index int, element *goquery.Node) {
						var new news
						for _, v := range element.Node.Attr {
							if v.Key == "data-title" {
								new.title = v.Val
							}
							if v.Key == "data-link" {
								new.url = v.Val
							}
						}
						newsChan <- new
						fmt.Println(new)
					})
				}
			}
		}
		res.Body.Close()
	}
	wGetOk.Done()
}

func main() {
	fmt.Println(os.Args)
	var count int

	if len(os.Args) > 1 {
		size, err := strconv.ParseInt(os.Args[1], 0, 64)
		if nil != err {
			panic(err)
		}
		count = int(size)
	}
	if count == 0 {
		count = 10
	}

	index := make(chan int)
	newsChan := make(chan news, 100)
	var wGetOk sync.WaitGroup
	wGetOk.Add(runtime.NumCPU())
	for m := 0; m < runtime.NumCPU(); m++ {
		go getBaiduNews(index, newsChan, &wGetOk)
	}
	var wSaveOk sync.WaitGroup
	wSaveOk.Add(1)
	go saveNewsDB(newsChan, &wSaveOk)
	go func() {
		defer close(newsChan)
		wGetOk.Wait()
	}()

	go func() {
		defer close(index)
		for i := 1; i < count; i++ {
			index <- i
		}
	}()
	wSaveOk.Wait()
}
