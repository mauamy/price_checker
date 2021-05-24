package price_checker

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

const MecchaJapanProviderKey = "MCJP"

type MecchaJapan struct {
	Shop
}

func NewMecchaJapan() *MecchaJapan {
	mjp := &MecchaJapan{}
	mjp.Name = "Meccha Japan"
	mjp.BaseURL = "https://meccha-japan.com/"
	return mjp
}

func (mjp *MecchaJapan) GetName() string {
	return mjp.Name
}

func (mjp *MecchaJapan) GetBaseUrl() string {
	return mjp.BaseURL
}

func (mjp *MecchaJapan) GetKey() string {
	return MecchaJapanProviderKey
}

func (mjp *MecchaJapan) ParseProduct(productUrl string) (*Product, error) {
	resp, err := http.Get(productUrl + "?SubmitCurrency=1&id_currency=1")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	priceSpan := doc.Find("[itemprop='price']")
	priceValue, ok := priceSpan.Attr("content")
	if !ok {
		return nil, fmt.Errorf("could not get price value")
	}
	price, err := strconv.ParseFloat(priceValue, 64)
	if err != nil {
		return nil, err
	}

	instock := true
	stock := -1
	availabilitySpan := doc.Find("#product-availability-mobile")

	if availabilitySpan != nil && strings.Contains(availabilitySpan.Text(), "Out-of-Stock") {
		instock = false
		stock = 0
	}

	return &Product{
		URL:     productUrl,
		Price:   price,
		Stock:   stock,
		InStock: instock,
	}, nil
}
