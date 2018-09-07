package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
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

	p, err := c.PropertyCollector().Create(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var version string
	maxWaitSeconds := int32(180)
	maxObjectUpdates := int32(1000)
	req := &types.WaitForUpdatesEx{
		This:    p.Reference(),
		Version: version,
		Options: &types.WaitOptions{
			MaxWaitSeconds:   &maxWaitSeconds,
			MaxObjectUpdates: maxObjectUpdates,
		},
	}

	resp, err := methods.WaitForUpdatesEx(ctx, c.RoundTripper, req)
	if err != nil {
		log.Fatal(err)
	}
	res := resp.Returnval
	version = res.Version
	fmt.Println(version)
	fmt.Println(res.Truncated)
	fmt.Println(res)

	reqWFU := &types.WaitForUpdates{
		This:    p.Reference(),
		Version: version,
	}

	respWFU, err := methods.WaitForUpdates(ctx, c.RoundTripper, reqWFU)
	if err != nil {
		log.Fatal(err)
	}

	resWFU := respWFU.Returnval
	fmt.Println(resWFU)
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
