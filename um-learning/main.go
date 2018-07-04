package main

import (
	"context"
	"fmt"
	"grass/um-learning/model"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/event"
	"github.com/vmware/govmomi/units"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	cfg, err := model.Parse(path + "/config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	ctx, _ := context.WithCancel(context.Background())

	u, err := soap.ParseURL(cfg.Vc[0])
	c, err := govmomi.NewClient(ctx, u, true)
	if err != nil {
		fmt.Println(err)
	}

	defer c.Logout(ctx)

	// Create view of VirtualMachine objects
	m := view.NewManager(c.Client)

	v1, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve summary property for all machines
	// Reference: http://pubs.vmware.com/vsphere-60/topic/com.vmware.wssdk.apiref.doc/vim.VirtualMachine.html
	var vms []mo.VirtualMachine
	err = v1.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary"}, &vms)
	if err != nil {
		log.Fatal(err)
	}

	// Print summary per vm (see also: govc/vm/info.go)

	for _, vm := range vms[0:2] {
		fmt.Printf("%s: %s\n", vm.Summary.Config.Name, vm.Summary.Config.GuestFullName)
	}

	v2, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"HostSystem"}, true)
	if err != nil {
		log.Fatal(err)
	}

	defer v2.Destroy(ctx)
	var hss []mo.HostSystem
	err = v2.Retrieve(ctx, []string{"HostSystem"}, []string{"summary"}, &hss)
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
	fmt.Println("#######################################")
	m1 := event.NewManager(c.Client)
	/*ref := types.ManagedObjectReference{"type", "char"}
	//ref := m1.Common.Reference()
	historyCollect := event.NewHistoryCollector(c.Client, ref)
	//baseEvent, err := historyCollect.LatestPage(ctx)
	//fmt.Println(baseEvent)
	baseEvent, err := historyCollect.ReadNextEvents(ctx, 10)
	fmt.Println(baseEvent)*/
	end := time.Now()
	start := end.AddDate(0, -1, 0)
	fmt.Println(start)
	fmt.Println(end)
	filter := types.EventFilterSpec{
		Time: &types.EventFilterSpecByTime{
			BeginTime: &start,
			EndTime:   &end,
		},
		UserName: &types.EventFilterSpecByUsername{
			SystemUser: false,
		},
		EventTypeId: []string{
			"com.vmware.license.AddLicenseEvent",
			"com.vmware.license.AssignLicenseEvent",
			"VmCreatedEvent",
			"VmMigratedEvent",
			"VmPoweredOffEvent",
			"VmPoweredOnEvent",
			"VmRelocatedEvent",
			"VmSuspendedEvent",
		},
		MaxCount: 200,
	}

	historyCollect, err := m1.CreateCollectorForEvents(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	event, err := historyCollect.ReadNextEvents(ctx, 20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event)
}
