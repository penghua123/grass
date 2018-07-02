package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
)

var conf = flag.String("conf", "config.json", "proxy config")

func main() {
	flag.Parse()
	config, err := NewConfig(*conf)
	if err != nil {
		log.Fatalf("Config Error: %s", err.Error())
	}

	if len(config.Proxy) > 0 {
		for _, proxy := range config.Proxy {
			http.Handle(proxy.Route(), proxy.Handler())
		}
	}
	if path := flag.Arg(0); path != "" {
		http.ListenAndServe(config.Port, nil)
	}
}

type Config struct {
	Port  string `json:"p"`
	Proxy []Pro
}

type Pro struct {
	Host string
	Path string
}

func (p *Pro) Route() string {
	return filepath.Clean(fmt.Sprintf("/%s/*", p.Path))
}

func (p *Pro) Clean() string {
	return filepath.Clean(fmt.Sprintf("/%s", p.Path))
}

func (p *Pro) Handler() http.Handler {
	u, e := url.Parse(p.Host)
	if e != nil {
		log.Fatal("Bad destination.")
	}
	return http.StripPrefix(p.Clean(), httputil.NewSingleHostReverseProxy(u))
}

func NewConfig(file string) (*Config, error) {
	conf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var c *Config
	err = json.Unmarshal(conf, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
