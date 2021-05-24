package price_checker

type Shop struct {
	Name    string
	BaseURL string
}

type Product struct {
	URL     string
	Price   float64
	Stock   int
	InStock bool
}

type ShopProvider interface {
	GetName() string
	GetBaseUrl() string
	GetKey() string
	ParseProduct(productUrl string) (*Product, error)
}
