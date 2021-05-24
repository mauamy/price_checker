package price_checker

import (
	"io"
	"net/http"
	"strings"
)

func GetPageBody(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyContent, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyContent), err
}

func GetPageBodyLines(url string) ([]string, error) {
	body, err := GetPageBody(url)
	if err != nil {
		return nil, err
	}

	return strings.Split(body, "\n"), nil
}
