package types

func (a *Anigo) GetProviders() ([]ProviderPlugin, bool) {

	return _GetServicesOrProviders[ProviderPlugin](*a, "providers")
}

func (a *Anigo) GetServices() ([]ServicePlugin, bool) {

	return _GetServicesOrProviders[ServicePlugin](*a, "services")
}

// //---------------------------------------------------------// //

func (a *Anigo) GetProviderNames() ([]string, bool) {

	return _GetServicesOrProvidersNames[ProviderPlugin](*a, "provider")
}

func (a *Anigo) GetServiceNames() ([]string, bool) {

	return _GetServicesOrProvidersNames[ServicePlugin](*a, "services")
}

// //---------------------------------------------------------// //

func (a *Anigo) GetDirectUrl(url string) (string, bool) {
	providers, _ := a.GetProviders()
	domain := _GetDomain(url)

	for _, prov := range providers {
		if url, ok := prov.Handler(domain, url); ok {

			return url, true
		}

	}

	return url, false
}
