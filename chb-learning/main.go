package main

import (
	"chb-learning/model"
	"fmt"
)

func main() {
	user := model.User{
		UserName: "Tom",
		PassWord: "123456789",
		City:     "Shanghai",
		Age:      "25",
		Point:    "13762",
	}

	var s model.UserList
	s.Users = append(s.Users, user)
	s.Users = append(s.Users, model.User{UserName: "Mary", PassWord: "374659321", City: "Xian", Age: "23", Point: "8977"})

	id := 1
	for testR, u := range s.Users {
		id = id + 1
		fmt.Println("testR:", testR)
		model.Insert(fmt.Sprintf("%d", id), u)
	}

	//for k := 0; k < s.Users(); k++ {
	//	id = id + 1
	//	model.Insert(string(id), s.Users(k))
	//}

	//model.Insert(string(id), user)
	model.OpsJSON("user.json", "parse")
	model.Select()
	//model.Data2JSON("", "h.json")
	//select {}
}
