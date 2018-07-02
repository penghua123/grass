package main

import (
	"fmt"
	"log"
	"um-chb/model"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/gocb"
)

type Id struct {
	ID string `json:"id"`
}

func main() {
	cluster, err := gocb.Connect("couchbase://10.158.15.49")
	model.CheckError(err)
	err = cluster.Authenticate(
		gocb.PasswordAuthenticator{
			Username: "hua",
			Password: "penghua",
		})
	model.CheckError(err)

	bucket, err := cluster.OpenBucket("hua", "")
	query := gocb.NewN1qlQuery("select meta().* from `hua`")
	var b interface{}
	rows, _ := bucket.ExecuteN1qlQuery(query, b)

	var row interface{}
	var doc Id
	//var wg sync.WaitGroup
	var items []gocb.BulkOp
	i := 0
	for rows.Next(&row) {
		//fmt.Println(docs)
		//fmt.Println(row)
		mapstructure.Decode(row, &doc)
		//fmt.Println(doc)
		var value interface{}
		c, err := bucket.Get(doc.ID, &value)
		//fmt.Println(value)
		model.CheckError(err)
		id := doc.ID

		items = append(items, &gocb.RemoveOp{Key: id, Cas: c})
		//fmt.Println(id)
		//bucket.Remove(id, cas)
		//wg.Add(1)
		//go func(id string, c gocb.Cas) {
		//	defer wg.Done()
		//	bucket.Remove(id, c)
		//}(id, c)
		//wg.Wait()
		if (i % 100) == 0 {
			fmt.Println(i)
		}
		i += 1
	}
	err = bucket.Do(items)
	if err != nil {
		log.Fatal(err)
	}

}
