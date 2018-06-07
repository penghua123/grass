package main

import (
	"fmt"
	"log"
	"encoding/json"
	"gopkg.in/gocb"
)


func main(){
	cluster, err := gocb.Connect("couchbase://10.158.15.52")
		if err != nil {
		log.Fatal(err)
	}
	err = cluster.Authenticate(
		gocb.PasswordAuthenticator{
			Username: "hua",
			Password: "penghua",
		})
	if err != nil {
		log.Fatal(err)
	}
	bucket, err := cluster.OpenBucket("hua", "")


	myQuery := gocb.NewN1qlQuery("SELECT * FROM `paul`")
	rows, err := bucket.ExecuteN1qlQuery(myQuery, nil)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Select finished")

	var items []gocb.BulkOp
	doc := make(map[string]interface{})
	id :=1
	//var beer map[string]interface{}
	for rows.Next(&doc) {
		id=id+1
		str, err := json.Marshal(doc)
		if err != nil {
			fmt.Println(err)
		}
		if id%2 == 0 {
        		items = append(items, &gocb.ReplaceOp{Key: fmt.Sprintf("%s", id-1), Value: str})
			//cas, _ := bucket.Get(fmt.Sprintf("%s", id+1), &beer)
			//bucket.Replace(fmt.Sprintf("%s", id-1), doc,cas, 0)
		} else {
			items = append(items, &gocb.ReplaceOp{Key: fmt.Sprintf("%s", id+1), Value: str})
			//cas, _ := bucket.Get(fmt.Sprintf("%s", id-1), &beer)
			//bucket.Replace(fmt.Sprintf("%s", id+1), doc,cas, 0)
 		}
	}
	
	
	fmt.Println("start replace")
	err = bucket.Do(items)
	if err != nil {
		fmt.Println(err)
	}
}
