package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://www.baidu.com/"
	res, err := http.Head(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(url, ":", res.Status, res.Header)
	fmt.Println("#####################################")
	res, err = http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("body :", string(data))
	fmt.Println("#####################################")
	res, err = http.Post("https://open.ys7.com", "application/x-www-form-urlencoded", strings.NewReader("name=abc"))
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("body :", string(data))

	fmt.Println("#####################################")

	/*res, err = http.PostForm("http://www.01happy.com/demo/accept.php", url.Values{"key": {"Value"}, "id": {"123"}})
		if err != nil {
			fmt.Println(err)
		}
	    fmt.Println("body :", string(data))*/
	httpDo()
}

func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}
