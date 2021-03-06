package main

import (
	"encoding/json"
	"fmt"

	"github.com/couchbase/go-couchbase"
)

type Event struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Likes int    `json:"likes"`
}

func NewEvent(name string) *Event {
	return &Event{"Event", name, 0}
}

func NewEventJson(jsonbytes []byte) (event *Event) {
	err := json.Unmarshal(jsonbytes, &event)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func LikeEvent(id string, bucket *couchbase.Bucket) {
	var event Event
	err := bucket.Get(id, &event)
	if err != nil {
		fmt.Println(err)
	}
	event.Likes++
	bucket.Set(id, 0, event)

}

func main() {
	bucket, err := couchbase.GetBucket("http://hua:penghua@IP:8091/", "default", "bucket")
	if err != nil {
		fmt.Println(err)
	}
	defer bucket.Close()
	fmt.Println("Connected to Couchbase Bucket '%s'\n", bucket.Name)

	event := NewEvent("Couchbase collect")
	err = bucket.Set("cc2014", 0, event)

	event = NewEvent("Couspher India")
	err = bucket.Set("gc2015", 0, event)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Set Successifully")

	var eventSearch Event
	err = bucket.Get("cc2014", &eventSearch)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Get Successifully:", &eventSearch)

	for i := 0; i < 30; i++ {
		go func() {
			LikeEvent("cc2014", bucket)
		}()

	}
}
