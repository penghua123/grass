package model

type Config struct {
	Vc []string `ymal:"vc"`
}

type Vc struct {
	Host     string
	Username string
	Password string
	Port     string
}
