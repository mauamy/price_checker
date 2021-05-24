package price_checker

type PriceChecker struct {
	shopProviders map[string]ShopProvider
}

func NewPriceChecker() *PriceChecker {
	pc := &PriceChecker{}
	pc.initialProviders()

	return pc
}

func (pc *PriceChecker) initialProviders() {
	pc.shopProviders = make(map[string]ShopProvider, 0)

	pc.registerProvider(NewPlazaJapan())
	pc.registerProvider(NewMecchaJapan())
}

func (pc *PriceChecker) registerProvider(sp ShopProvider) {
	pc.shopProviders[sp.GetKey()] = sp
}

func (pc *PriceChecker) GetShopProvider(providerKey string) ShopProvider {
	return pc.shopProviders[providerKey]
}
