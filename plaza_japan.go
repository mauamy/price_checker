package price_checker

import (
	"fmt"
	"github.com/buger/jsonparser"
	"regexp"
)

const PlazaJapanProviderKey = "PLZJP"

type PlazaJapan struct {
	Shop
}

func NewPlazaJapan() *PlazaJapan {
	pj := &PlazaJapan{}
	pj.Name = "PlazaJapan"
	pj.BaseURL = "https://www.plazajapan.com/"
	return pj
}

func (pj *PlazaJapan) GetName() string {
	return pj.Name
}

func (pj *PlazaJapan) GetBaseUrl() string {
	return pj.BaseURL
}

func (pj *PlazaJapan) GetKey() string {
	return PlazaJapanProviderKey
}

func (pj *PlazaJapan) ParseProduct(productUrl string) (*Product, error) {
	pageBody, err := GetPageBody(productUrl + "?setCurrencyId=3")
	if err != nil {
		return nil, err
	}

	rex := regexp.MustCompile(`var BCData = (.*);`)
	matches := rex.FindStringSubmatch(pageBody)
	if len(matches) < 2 {
		return nil, fmt.Errorf("could not find product informations")
	}

	productData := []byte(matches[1])

	price, err := jsonparser.GetFloat(productData, "product_attributes", "price", "without_tax", "value")
	if err != nil {
		return nil, err
	}

	inStock, err := jsonparser.GetBoolean(productData, "product_attributes", "instock")
	if err != nil {
		return nil, err
	}

	var stock int = 0
	if inStock {
		stockI64, err := jsonparser.GetInt(productData, "product_attributes", "stock")
		stock = int(stockI64)
		if err != nil {
			return nil, err
		}
	}

	return &Product{
		URL:     productUrl,
		Price:   price,
		Stock:   stock,
		InStock: inStock,
	}, nil
}
