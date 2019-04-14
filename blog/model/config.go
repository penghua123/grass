package model

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/couchbase/gocb"
	yaml "gopkg.in/yaml.v2"
)

//Config store the config information
type Config struct {
	Couchbase []*Chb `yaml:"couchbase"`
	Thread    int    `yaml:"thread"`
	DocNum    int    `yaml:"docNum"`
	Persist   int    `yaml:"persist"`
	Replicate int    `yaml:"replicate"`
}

//Chb is a struct to storage the information of Couchbase
type Chb struct {
	Driver    string `yaml:"driver"`
	Host      string `yaml:"host"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Bucket    string `yaml:"bucket"`
	Operation string `yaml:"operation"`
	Switch    string `yaml:"switch"`
}

//Parse Parse the config file
func Parse(file string) (cfg *Config, err error) {
	cfg = new(Config)
	body, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(body, cfg)
	if err != nil {
		return
	}
	return
}

//GetBucket parse the couchbase information and collect to the cb database
func (cfg *Config) GetBucket() *Chb {
	for _, chb := range cfg.Couchbase {
		if strings.ToLower(chb.Switch) == "on" && strings.ToLower(chb.Operation) == "write" {
			return chb
		}
	}
	return nil
}

//ConnectionBucket collect the CouchBase database through the given information
func (chb *Chb) ConnectionBucket() (*gocb.Bucket, error) {
	cluster, err := gocb.Connect("couchbase://" + chb.Host)
	if err != nil {
		return nil, err
	}

	cluster.SetServerConnectTimeout(time.Minute * 60)
	cluster.SetFtsTimeout(time.Minute * 60)
	cluster.SetConnectTimeout(time.Minute * 60)
	cluster.SetAnalyticsTimeout(time.Minute * 60)
	cluster.SetN1qlTimeout(time.Minute * 60)

	err = cluster.Authenticate(
		gocb.PasswordAuthenticator{
			Username: chb.Username,
			Password: chb.Password,
		})
	if err != nil {
		return nil, err
	}
	bucket, err := cluster.OpenBucket(chb.Bucket, "")
	if err != nil {
		return nil, err
	}
	bucket.SetBulkOperationTimeout(time.Minute * 60)
	bucket.SetN1qlTimeout(time.Minute * 60)
	bucket.SetOperationTimeout(time.Minute * 60)
	bucket.SetViewTimeout(time.Minute * 60)
	bucket.SetDurabilityPollTimeout(time.Minute * 3)
	bucket.SetDurabilityTimeout(time.Minute * 3)
	return bucket, err
}
