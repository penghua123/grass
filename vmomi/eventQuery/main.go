package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/vmware/govmomi/event"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vim25"
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
	sc, err := vim25.NewClient(ctx, c.RoundTripper)
	if err != nil {
		log.Fatal(err)
	}
	filter := eventFilterSpec(false)
	m := event.NewManager(sc)
	events, err := m.QueryEvents(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	data2JSON(events, "nonSystemUser.json")
}

func eventFilterSpec(systemUser bool) types.EventFilterSpec {
	end := time.Now()
	start := end.Add(-48 * time.Hour)
	return types.EventFilterSpec{
		Time: &types.EventFilterSpecByTime{
			BeginTime: &start,
			EndTime:   &end,
		},
		UserName: &types.EventFilterSpecByUsername{
			SystemUser: systemUser,
		},
		EventTypeId: []string{
			"com.vmware.license.AddLicenseEvent",
			"com.vmware.license.AssignLicenseEvent",
			"VmCreatedEvent",
			"VmMigratedEvent",
			"VmPoweredOnEvent",
			"VmPoweredOffEvent",
			"VmRelocatedEvent",
			"VmSuspendedEvent",
			"VmCloneEvent",
			"VmDeployedEvent",
			"VmRemovedEvent",
			"VmRenamedEvent",
		},
		MaxCount: 1000,
	}
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
