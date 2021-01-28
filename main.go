package eureka_client

import (
	"github.com/bilalkocoglu/eureka-client/discovery"
	"github.com/bilalkocoglu/eureka-client/store"
)

func ServiceRegister(cfg store.RegistryConfig) {
	config := discovery.RegistrationVariables{ServiceRegistryURL: cfg.URL, UserName: cfg.Username, Password: cfg.Password}
	go discovery.ManageDiscovery(config)
}
