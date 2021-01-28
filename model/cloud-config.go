package model

type CloudConfig struct {
	Name            string                      `json:"name"`
	Profiles        []string                    `json:"profiles"`
	Label           string                      `json:"label"`
	Version         string                      `json:"version"`
	State           string                      `json:"state"`
	PropertySources []CloudConfigPropertySource `json:"propertySources"`
}

type CloudConfigPropertySource struct {
	Name   string                 `json:"name"`
	Source map[string]interface{} `json:"source"`
}
