package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"temp/internal/xds/models"
)

type Envoy struct {
	Listeners []models.Listener `yaml:"listener_config"`
	Clusters  []models.Cluster  `yaml:"cluster_config"`
	Routes    []models.Route    `yaml:"route_config"`
	TLS       models.TLS        `yaml:"tls_config"`
}

func New(configPath string) (*Envoy, error) {
	cfg := &Envoy{}
	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
