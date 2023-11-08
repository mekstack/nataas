package models

type Listener struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	Port    uint32 `yaml:"port"`
}
