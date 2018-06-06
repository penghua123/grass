package model

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	user1 := User{
		UserName: "Jim",
		PassWord: "7868541427",
		City:     "Beijing",
		Age:      "22",
		Point:    "29831",
	}

	for id := 3; id < 10; id++ {
		err := Insert(fmt.Sprintf("%d", id), user1)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("Success id:", id)
	}

}

func TestOpsJSON(t *testing.T) {
	u := User{
		UserName: "Jim",
		PassWord: "7868541427",
		City:     "Beijing",
		Age:      "22",
		Point:    "29831",
	}
	data := new(Users)
	data.Writer = "hua"
	for id := 3; id < 10; id++ {
		ux := u
		data = append(data, ux)
	}

	Data2JSON(data, "hua.json")
	//OpsJSON("hua.json", "parse")

}
