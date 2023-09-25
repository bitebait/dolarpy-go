package dolarpy

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// APIData represents the structure of the API response
type APIData struct {
	Dolarpy map[string]map[string]float64 `json:"dolarpy"`
}

// getAPIData retrieves the API data from the remote server
func getAPIData() (*APIData, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://dolar.melizeche.com/api/1.0/", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "DolarpyWrapper/Golang")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	var data APIData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &data, nil
}

// Reference returns the referential exchange rate from BCP
func Reference() (float64, error) {
	data, err := getAPIData()
	if err != nil {
		return 0, fmt.Errorf("failed to get API data: %w", err)
	}
	return data.Dolarpy["bcp"]["referencial_diario"], nil
}

// Providers returns a list of available providers
func Providers() ([]string, error) {
	data, err := getAPIData()
	if err != nil {
		return nil, fmt.Errorf("failed to get API data: %w", err)
	}

	providers := make([]string, 0, len(data.Dolarpy))
	for key := range data.Dolarpy {
		providers = append(providers, key)
	}

	return providers, nil
}

// Purchase returns the purchase exchange rate from the specified provider
// If provider is empty, returns the rate from BCP
func Purchase(provider string) (float64, error) {
	data, err := getAPIData()
	if err != nil {
		return 0, fmt.Errorf("failed to get API data: %w", err)
	}

	rate := data.Dolarpy["bcp"]["compra"]
	if provider != "" {
		rate = data.Dolarpy[provider]["compra"]
	}

	return rate, nil
}

// Sale returns the sale exchange rate from the specified provider
// If provider is empty, returns the rate from BCP
func Sale(provider string) (float64, error) {
	data, err := getAPIData()
	if err != nil {
		return 0, fmt.Errorf("failed to get API data: %w", err)
	}

	rate := data.Dolarpy["bcp"]["venta"]
	if provider != "" {
		rate = data.Dolarpy[provider]["venta"]
	}

	return rate, nil
}

// All returns all available exchange rates from all providers
func All() (map[string]map[string]float64, error) {
	data, err := getAPIData()
	if err != nil {
		return nil, fmt.Errorf("failed to get API data: %w", err)
	}
	return data.Dolarpy, nil
}
