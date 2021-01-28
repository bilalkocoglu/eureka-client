package store

import (
	"strings"
)

type RegistryConfig struct {
	URL      string `envconfig:"REGISTRY_URL"`
	Username string `envconfig:"REGISTRY_USERNAME"`
	Password string `envconfig:"REGISTRY_PASSWORD"`
	UseSSL   bool   `envconfig:"USE_SSL"`
	AppName  string `envconfig:"APPLICATION_NAME"`
	AppPort  string `envconfig:"APPLICATION_PORT"`
}

type ServiceInfos struct {
	Services []ServiceApp
}

type ServiceApp struct {
	Name             string
	InstanceBaseUrls []string
}

type RegisteredServiceInfos struct {
	Application RegisteredApp `json:"applications"`
}

type RegisteredApp struct {
	Version string                `json:"versions__delta"`
	Hash    string                `json:"apps__hashcode"`
	Apps    []RegistryApplication `json:"application"`
}

type RegistryApplication struct {
	Name      string             `json:"name"`
	Instances []RegistryInstance `json:"instance"`
}

type RegistryInstance struct {
	InstanceId       string       `json:"instanceId"`
	Hostname         string       `json:"hostname"`
	App              string       `json:"app"`
	IpAddress        string       `json:"ipAddr"`
	Status           string       `json:"status"`
	OverriddenStatus string       `json:"overriddenStatus"`
	Port             RegisterInfo `json:"port"`
	SecurePort       RegisterInfo `json:"securePort"`
	CountryId        int          `json:"countryId"`
}

type RegisterInfo struct {
	Value    int    `json:"$"`
	IsActive string `json:"@enabled"`
}

var RegisteredServices *ServiceInfos

func (si ServiceInfos) GetServiceUrl(serviceName string) string {
	for _, s := range si.Services {
		if strings.EqualFold(s.Name, serviceName) {
			if len(s.InstanceBaseUrls) > 0 {
				return s.InstanceBaseUrls[0]
			}
		}
	}
	return ""
}
