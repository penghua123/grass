package model

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func Parse(file string) (*Config, error) {
	cfg := new(Config)
	body, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(body, cfg)
	if err != nil {
		fmt.Println(err)
	}
	return cfg, err
}
