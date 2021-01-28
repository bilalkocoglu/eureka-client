package discovery

func ManageDiscovery(configs RegistrationVariables) {
	manager := new(EurekaRegistrationManager)
	manager.RegisterWithServiceRegistry(configs)
	manager.StoreOtherMSInfo(configs)
	manager.SendHeartBeat(configs)
}
