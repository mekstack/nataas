package models

type Route struct {
	Name        string `yaml:"name"`
	Prefix      string `yaml:"prefix"`
	ClusterName string `yaml:"cluster"`
}
