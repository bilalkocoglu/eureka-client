package discovery

type RegistrationManager interface {
	Manage(configs RegistrationVariables)
	RegisterWithServiceRegistry()
	SendHeartBeat(configs RegistrationVariables)
	DeRegisterFromServiceRegistry(configs RegistrationVariables)
	StoreOtherMSInfo(config RegistrationVariables)
}

type RegistrationVariables struct {
	ServiceRegistryURL string
	UserName           string
	Password           string
	AppName            string
	AppPort            string
	UseSSL             bool
}
