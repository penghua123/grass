package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"grass/blog/model"
	"log"
	"net/http"
	"os"

	"gopkg.in/gocb"
	//	"gopkg.in/mgo.v2/bson"
)

var bucket *gocb.Bucket
var bucketName string

type Person struct {
	Name string
	Age  int
}

type Blog struct {
	Title   string
	FileDir string
}

//响应消息结构
type Resp struct {
	Errno  string
	Errmsg string
	Data   interface{}
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\n"))
}

//统一响应消息
func RespJson(resp *Resp, w http.ResponseWriter) {
	resp.Errno = "0"
	if resp.Errmsg != "ok" {
		resp.Errno = "100" // 错误编码
	}
	data, _ := json.Marshal(resp)
	w.Write(data)
}

func publish(w http.ResponseWriter, r *http.Request) {
	//0. 响应消息
	var resp Resp
	resp.Errmsg = "ok"
	defer RespJson(&resp, w)
	//1. 获取前端消息
	title := r.FormValue("title")
	content := r.FormValue("content")
	//正文特别多的话，可以生成文件来保存
	filehash := sha256.Sum256([]byte(title))
	filename := fmt.Sprintf("%x", filehash)
	fmt.Println(filename)
	f, err := os.Create("static/blogfile/" + filename)
	if err != nil {
		fmt.Println("failed to create file ", err)
		resp.Errmsg = err.Error()
		return
	}
	defer f.Close()
	f.WriteString(content) //写入文件
	//2. 保存到数据库
	blog := Blog{title, "blogfile/" + filename}
	_, err = bucket.Insert(blog.Title, blog.FileDir, 0)
	if err != nil {
		fmt.Println("failed to insert mongo", blog.Title, err)
	}
}

func lists(w http.ResponseWriter, r *http.Request) {
	//1. 组织响应消息
	var resp Resp
	resp.Errmsg = "ok"
	defer RespJson(&resp, w)
	//2. 查询数据库

	sStmt := fmt.Sprintf("select `%s`.* from `%s`", bucketName)
	query := gocb.NewN1qlQuery(sStmt)
	blogs, err := bucket.ExecuteN1qlQuery(query, nil)
	if err != nil {
		log.Fatal("ERROR EXECUTING N1QL QUERY:", err)
	}

	//3. 响应消息赋值处理
	resp.Data = blogs
}

func main() {
	cfg, err := model.Parse("config.yaml")
	if err != nil {
		log.Fatal("Error:", err)
	}
	//连接到couchbase
	chb := cfg.GetBucket()
	if chb == nil {
		log.Fatal("Couchbase configure error!")
	}
	bucket, err := chb.ConnectionBucket()
	if err != nil {
		log.Fatal(err)
	}

	defer bucket.Close() // 本函数返回时自动执行

	http.HandleFunc("/ping", pong)                        //测试用
	http.Handle("/", http.FileServer(http.Dir("static"))) //提供静态文件服务的根目录
	http.HandleFunc("/publish", publish)                  //发表博客
	http.HandleFunc("/lists", lists)                      //发表博客
	http.ListenAndServe(":8086", nil)                     //启动http服务器
}
