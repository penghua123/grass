package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/gocb"
)

type User struct {
	UserName string
	Password string
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

//Insert  Insert data to bcouchabse
//func Insert(bucket *gocb.Bucket, ID string, user User) error {
func Insert(bucket *gocb.Bucket, ID string, user User) error {

	//cluster, _ := gocb.Connect("couchbase://http://10.158.5.52")
	//bucket, _ := cluster.OpenBucket("hua", "peng")

	str, err := json.Marshal(user)
	if err != nil {
		return err
	}

	var items []gocb.BulkOp
	items = append(items, &gocb.InsertOp{Key: ID, Value: str})
	//items = append(items, gocb.BulkOp{Key: ID, Value: str})

	err = bucket.Do(items)
	if err != nil {
		return err
	}

	return nil
}

//func Select(bucket *gocb.Bucket) interface{} {
func Query(bucket *gocb.Bucket) ([]map[string]interface{}, error) {
	//cluster, _ := gocb.Connect("couchbase://http://10.158.5.52")
	//bucket, _ := cluster.OpenBucket("hua", "peng")

	myQuery := gocb.NewN1qlQuery("SELECT `paul`.* FROM `paul` where docType='VcBaseHost'")
	rows, err := bucket.ExecuteN1qlQuery(myQuery, nil)
	if err != nil {
		log.Println(err)
	}

	//myQuery := gocb.NewN1qlQuery("SELECT * FROM `hua` WHERE city=$1 ")
	//var myParams []gocb.BulkOp
	//myParams = append(myParams, []interface{}{city})
	//rows, err := bucket.ExecuteN1qlQuery(myQuery, myParams)
	var docs []map[string]interface{}
	doc := make(map[string]interface{})
	for rows.Next(&doc) {
		//fmt.Printf("Return row: %#v\n", doc)
		docs = append(docs, doc)
		doc = make(map[string]interface{})
	}
	for id, doc1 := range docs {
		fmt.Println("Get Row:", id, doc1)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	return docs, err
}

//OpsJSON operate json data
func OpsJSON(path string, ops string, bucket *gocb.Bucket) error {
	if ops == "parse" {
		return parseJSON(path, bucket)
	}
	return nil
}
func parseJSON(path string, bucket *gocb.Bucket) error {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	/*data := new(Data)
	json.Unmarshal(body, data)
	for _, u := range data.Users {
		fmt.Printf("%#v\n", u)
	}*/

	var d2 Data
	json.Unmarshal(body, &d2)
	fmt.Println("var d2 data assgin:")

	data := new(Data)
	json.Unmarshal(body, data)
	var items []gocb.BulkOp
	for id, u1 := range data.Users {
		fmt.Println(u1)
		b, err := json.MarshalIndent(u1, "", "    ")
		if err != nil {
			return err
		}
		items = append(items, &gocb.InsertOp{Key: fmt.Sprintf("%d", id), Value: b})
	}
	err = bucket.Do(items)
	if err != nil {
		return err
	}

	//for _, u := range d2.Users {
	//	fmt.Printf("%#v\n", u)
	//}
	//b, err := json.MarshalIndent(v, "", "    ")
	//if err != nil {
	//	return err
	//}

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

func map2json(chbResult map[string]string, filename string) error {
	body, err := json.MarshalIndent(chbResult, "", "    ")
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return err
	}
	err = ioutil.WriteFile(filename, body, 0644)
	if err != nil {
		return err
	}
	return err
}

/*func map2struct(chbResult map[string]string) User {
	var user User
	mapstructure.Decode(chbResult, user)
	return user
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()        //结构体属性值
	structFieldValue := structValue.FieldByName(name) //结构体单个属性值

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type() //结构体的类型
	val := reflect.ValueOf(value)              //map值的反射值

	var err error
	if structFieldType != val.Type() {
		val, err = TypeConversion(fmt.Sprintf("%v", value), structFieldValue.Type().Name()) //类型转换
		if err != nil {
			return err
		}
	}

	structFieldValue.Set(val)
	return nil
}
*/
