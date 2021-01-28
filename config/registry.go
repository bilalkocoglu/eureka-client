package config

import (
	"github.com/bilalkocoglu/eureka-client/discovery"
	"github.com/bilalkocoglu/eureka-client/model"
	"github.com/bilalkocoglu/eureka-client/store"
)

func SetRegistryConfigForCloud(cloudConfig model.CloudConfig, cfg *store.RegistryConfig) {
	for _, s := range cloudConfig.PropertySources {
		url := s.Source["registry.url"]
		if cfg.URL == "" && url != nil && url != "" {
			cfg.URL = url.(string)
		}

		username := s.Source["registry.username"]
		if cfg.Username == "" && username != nil && username != "" {
			cfg.Username = username.(string)
		}

		password := s.Source["registry.password"]
		if cfg.Password == "" && password != nil && password != "" {
			cfg.Password = password.(string)
		}

		useSSL := s.Source["registry.use-ssl"]
		if useSSL != nil && useSSL != "" {
			cfg.UseSSL = useSSL.(bool)
		}

	}
}

func ServiceRegister(cfg store.RegistryConfig) {
	config := discovery.RegistrationVariables{ServiceRegistryURL: cfg.URL, UserName: cfg.Username, Password: cfg.Password}
	go discovery.ManageDiscovery(config)
}
