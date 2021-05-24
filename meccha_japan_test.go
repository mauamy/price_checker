package price_checker

import (
	"fmt"
	"testing"
)

func TestMecchaJapan_ParseProduct(t *testing.T) {
	pc := NewPriceChecker()
	mjpp := pc.GetShopProvider(MecchaJapanProviderKey)

	productUrls := []string{
		"https://meccha-japan.com/en/trading-card-game/26263-eevee-heroes-booster-box-pokemon.html",
		"https://meccha-japan.com/en/trading-card-game/23236-display-blow-master-ichigeki-pokemon-card.html",
	}

	for _, prodUrl := range productUrls {
		pi, err := mjpp.ParseProduct(prodUrl)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(pi)
	}
}
