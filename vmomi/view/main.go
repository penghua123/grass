package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
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

	/*p := property.DefaultCollector(sc)

	var dst interface{}
	v := view.NewContainerView(sc, p.Reference())
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"config.product.fullName"}, &dst)
	if err != nil {
		fmt.Println("Retrieve has error!!!", err)
	}
	data2JSON(dst, "view.dst.json")
	var vms []mo.VirtualMachine
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary.config.memorySizeMB", "summary.config.guestFullName", "summary.vm", "summary.config.name"}, &vms)
	if err != nil {
		fmt.Println("Retrieve has error!!!", err)
	}
	data2JSON(vms, "view.vms.json")*/

	m := view.NewManager(sc)
	cv, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		fmt.Println(err)
	}
	var vms2 []mo.VirtualMachine
	err = cv.Retrieve(ctx, []string{"VirtualMachine"}, []string{"guest.hostName", "guest.ipAddress", "runtime.bootTime", "runtime.host", "runtime.powerState", "summary.config.cpuReservation", "summary.config.memorySizeMB", "summary.config.guestFullName", "summary.vm", "summary.config.name"}, &vms2)
	if err != nil {
		fmt.Println(err)
	}
	data2JSON(vms2, "view.cvVms.json")
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
