package models

type TLS struct {
	Name string `yaml:"name"`
	Cert string `yaml:"cert"`
	Key  string `yaml:"key"`
}
