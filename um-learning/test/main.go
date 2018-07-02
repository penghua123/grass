package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/units"
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
	//vcen := Vc{Host: "10.138.", Username: "user", Password: "password", Port: "443"}
	url := "https://administrator@vsphere.local:Admin!23@10.138.0.218/sdk"
	//"https://" + vcen.Username + ":" + vcen.Password + "@" + vcen.Host + "/sdk"
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

	for _, vm := range vms[0:2] {
		fmt.Printf("%s: %s\n", vm.Summary.Config.Name, vm.Summary.Config.GuestFullName)
	}

	var hss []mo.HostSystem
	err = v.Retrieve(ctx, []string{"HostSystem"}, []string{"summary"}, &hss)
	if err != nil {
		log.Fatal(err)
	}

	// Print summary per host (see also: govc/host/info.go)

	tw := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "Name:\tUsed CPU:\tTotal CPU:\tFree CPU:\tUsed Memory:\tTotal Memory:\tFree Memory:\t\n")

	for _, hs := range hss {
		totalCPU := int64(hs.Summary.Hardware.CpuMhz) * int64(hs.Summary.Hardware.NumCpuCores)
		freeCPU := int64(totalCPU) - int64(hs.Summary.QuickStats.OverallCpuUsage)
		freeMemory := int64(hs.Summary.Hardware.MemorySize) - (int64(hs.Summary.QuickStats.OverallMemoryUsage) * 1024 * 1024)
		fmt.Fprintf(tw, "%s\t", hs.Summary.Config.Name)
		fmt.Fprintf(tw, "%d\t", hs.Summary.QuickStats.OverallCpuUsage)
		fmt.Fprintf(tw, "%d\t", totalCPU)
		fmt.Fprintf(tw, "%d\t", freeCPU)
		fmt.Fprintf(tw, "%s\t", (units.ByteSize(hs.Summary.QuickStats.OverallMemoryUsage))*1024*1024)
		fmt.Fprintf(tw, "%s\t", units.ByteSize(hs.Summary.Hardware.MemorySize))
		fmt.Fprintf(tw, "%d\t", freeMemory)
		fmt.Fprintf(tw, "\n")
	}

	_ = tw.Flush()

}
