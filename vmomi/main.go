package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
)

type Vc struct {
	Host     string
	Username string
	Password string
	Port     string
}

func main() {
	ctx, _ := context.WithCancel(context.Background())
	vcen := Vc{Host: "IP", Username: "user", Password: "password", Port: "443"}
	url := "https://" + vcen.Username + ":" + vcen.Password + "@" + vcen.Host + "/sdk"
	fmt.Println(url)
	u, err := soap.ParseURL(url)
	c, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		fmt.Println(err)
	}

	defer c.Logout(ctx)

	// Create view of VirtualMachine objects
	m := view.NewManager(c.Client)

	v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		log.Fatal(err)
	}

	defer v.Destroy(ctx)

	// Retrieve summary property for all machines
	// Reference: http://pubs.vmware.com/vsphere-60/topic/com.vmware.wssdk.apiref.doc/vim.VirtualMachine.html
	var vms []mo.VirtualMachine
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary"}, &vms)
	if err != nil {
		log.Fatal(err)
	}

	// Print summary per vm (see also: govc/vm/info.go)

	for _, vm := range vms {
		fmt.Printf("%s: %s\n", vm.Summary.Config.Name, vm.Summary.Config.GuestFullName)
	}

}
