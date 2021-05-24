package price_checker

import (
	"testing"
)

func TestPlazaJapan_ParseProduct(t *testing.T) {
	pc := NewPriceChecker()
	pjp := pc.GetShopProvider(PlazaJapanProviderKey)

	productUrl := "https://www.plazajapan.com/4521329322254/"
	pi, err := pjp.ParseProduct(productUrl)
	if err != nil {
		t.Error(err)
	}

	if pi.URL != productUrl {
		t.Errorf("Wrong set product URL, want %s, got %s", productUrl, pi.URL)
	}

	if pi.Price != 39.6 {
		t.Errorf("Wrong Price: want %f, got %f", 39.6, pi.Price)
	}

	if pi.Stock != 0 {
		t.Errorf("Wrong Stock: want %d, got %d", 2, pi.Stock)
	}

	if pi.InStock {
		t.Errorf("Product should not be in stock")
	}
}
