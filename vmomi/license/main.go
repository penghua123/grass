package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/license"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	uri := "https://administrator@vsphere.local:Admin!23@10.160.164.46"
	u, err := soap.ParseURL(uri)
	if err != nil {
		log.Fatal(err)
	}

	c, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		log.Fatal(err)
	}

	sc, err := vim25.NewClient(ctx, c.RoundTripper)
	if err != nil {
		log.Fatal(err)
	}

	//p := property.DefaultCollector(sc)
	m := license.NewManager(sc)
	l, err := m.List(ctx)
	if err != nil {
		log.Fatal(err)
	}
	data2JSON(l, "licenseList.json")

	licenseInfo, err := m.Decode(ctx, "Name")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(licenseInfo)
}

func data2JSON(v interface{}, filename string) error {
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}
