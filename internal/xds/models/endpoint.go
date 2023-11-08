package models

type Endpoint struct {
	Address string `yaml:"address"`
	Port    uint32 `yaml:"port"`
}
