package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/gocb"
)


type User struct {
	UserName string
	PassWord string
	City     string
	Age      string
	Point    string
}

type User2 struct {
	UserName    string `json:"userName"`
	UserIP      string `json:"userIP"`
	UserAddress string `json:"userAddress"`
}

//type Users []User `json:"users"`

type Data struct {
	Users  []User `json:"users"`
	Writer string `json:"writer"`
}
type UserList struct {
	Users []User `json:"users"`
}

//user := User{
//	UserName: "Tom",
//	PassWord: "123456789",
//	City:     "Shanghai",
//	Age:      "25",
//	Point:    "13762",
//}

//Insert  Insert data to bcouchabse
func Insert(ID string, user User) error {
	cluster, _ := gocb.Connect("couchbase://IP")
	bucket, _ := cluster.OpenBucket("user", "password")
	str, err := json.Marshal(user)
	if err != nil {
		return err
	}

	//var s UserList
	//s.Users = append(s.Users, user)
	//s.Users = append(s.Users, User{})

	var items []gocb.BulkOp
	items = append(items, &gocb.InsertOp{Key: ID, Value: str})

	err = bucket.Do(items)
	if err != nil {
		return err
	}

	_ = bucket.Close()
	return err
}

func Select() {
	cluster, _ := gocb.Connect("couchbase://IP")
	bucket, _ := cluster.OpenBucket("user", "password")

	myQuery := gocb.NewN1qlQuery("SELECT * FROM hua")
	rows, err := bucket.ExecuteN1qlQuery(myQuery, nil)

	//myQuery := gocb.NewN1qlQuery("SELECT * FROM `hua` WHERE city=$1 ")
	//var myParams []gocb.BulkOp
	//myParams = append(myParams, []interface{}{city})
	//rows, err := bucket.ExecuteN1qlQuery(myQuery, myParams)

	var row interface{}
	for rows.Next(&row) {
		fmt.Printf("Row: %+v\n", row)
	}

	if err = rows.Close(); err != nil {
		fmt.Printf("Couldn't get all the rows: %s\n", err)
	}

	_ = bucket.Close()

}

//OpsJSON operate json data
func OpsJSON(path string, ops string) error {
	if ops == "parse" {
		return parseJSON(path)
	}
	return nil
}
func parseJSON(path string) error {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	data := new(Data)
	json.Unmarshal(body, data)
	for _, u := range data.Users {
		fmt.Printf("%#v\n", u)
	}
	var d2 Data
	json.Unmarshal(body, &d2)
	fmt.Println("var d2 data assgin:")
	for _, u := range d2.Users {
		fmt.Printf("%#v\n", u)
	}
	//fmt.Println("JSONRead first output\n", string(data))

	/*var u User2
	//json.Unmarshal([]byte(data), &u)
	err = json.Unmarshal(data, &u)
	if err != nil {
		return err
	}
	fmt.Println("JSONRead second output")
	fmt.Printf("%+v", u)
	fmt.Println()

	fmt.Println("JSONRead third output")
	fmt.Printf("%T", u)
	fmt.Println()*/
	return nil
}
