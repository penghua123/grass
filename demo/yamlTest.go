package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Frequency struct {
	Value int64  `yaml:"value"`
	Unit  string `yaml:"unit"`
}

type Config struct {
	PgDb      *PgDatabase `yaml:"database"`
	Sample    *Sample     `yaml:"sample"`
	Frequency Frequency   `yaml:"frequency"`
	IsTest    bool        `yaml:"-"`
	GOVC_URL  string      `yaml:"govcurl"`
}

type Sample struct {
	Type       string   `yaml:"type"`
	Percentage float64  `yaml:"percentage"`
	Protected  []string `ymal:"protected"`
}

type PgDatabase struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Switch   string `yaml:"switch"`
	Port     string `yaml:"-"`
	Sslmode  string `yaml:"sslmode"`
}

func (cfg *Config) Parse(file string) {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(body, cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var cfg Config
	cfg.Parse("vcsim-ctl/config.yaml")
	for _, vm := range cfg.Sample.Protected {
		fmt.Println(vm)
	}

}
