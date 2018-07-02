package main

import (
	"encoding/json"
	"fmt"
	"grass/model"
	"log"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/gocb"
)

type VcData struct {
	BootTime              string `json:"bootTime"`
	ChangeTime            string `json:"changeTime"`
	ChangedByUserId       int    `json:"changedByUserId"`
	DocKey                string `json:"docKey"`
	DocTime               string `json:"docTime"`
	DocType               string `json:"docType"`
	FullName              string `json:"fullName"`
	Hz                    int    `json:"hz"`
	Id                    int    `json:"id"`
	InstanceUuid          string `json:"instanceUuid"`
	LicenseProductName    string `json:"licenseProductName"`
	LicenseProductVersion string `json:"licenseProductVersion"`
	MemorySize            int    `json:"memorySize"`
	Moref                 string `json:"moref"`
	Name                  string `json:"name"`
	NumCpuCores           int    `json:"numCpuCores"`
	NumCpuPackages        int    `json:"numCpuPackages"`
	NumCpuThreads         int    `json:"numCpuThreads"`
	PowerState            string `json:"powerState"`
	ProductLineId         string `json:"productLineId"`
	PropertyVersion       string `json:"propertyVersion"`
	UmaId                 int    `json:"umaId"`
	Uuid                  string `json:"uuid"`
	VcCollectionId        int    `json:"vcCollectionId"`
	VcId                  int    `json:"vcId"`
	Version               string `json:version`
}

func main() {

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
	if err != nil {
		log.Fatal(err)
	}
	user := model.User{
		UserName: "Tom",
		Password: "123456789",
		City:     "Shanghai",
		Age:      "25",
		Point:    "13762",
	}

	var s model.UserList
	s.Users = append(s.Users, user)
	s.Users = append(s.Users, model.User{UserName: "Mary", Password: "374659321", City: "Xian", Age: "23", Point: "8977"})

	var d []map[string]interface{}
	d, err = model.Query(bucket)
	if err != nil {
		log.Fatal(err)
	}

	var items []gocb.BulkOp
	for id, doc := range d {
		var v VcData
		fmt.Println("Get Row:", id, doc)
		mapstructure.Decode(doc, &v)
		str, err := json.Marshal(v)
		fmt.Println(v)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, &gocb.InsertOp{Key: v.DocKey, Value: str})
	}
	err = bucket.Do(items)
	if err != nil {
		log.Fatal(err)
	}
	err = bucket.Close()
	if err != nil {
		log.Fatal(err)
	}
	//for id, u := range s.Users {
	//	model.Insert(bucket, fmt.Sprintf("%d", id+1), u)
	//}

	//model.Insert(string(id), user)
	//time.Sleep(10 * time.Second)
	//model.OpsJSON("user.json", "parse")

	/*docs, err := model.Query(bucket)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(docs)

	model.Data2JSON(docs, "paul.json")
	err = bucket.Close()*/
	//model.OpsJSON("paul.json", "parse", bucket)
}
