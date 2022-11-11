package dolarpy

/*
This is a Golang implementation of the
Dolarpy API (https://github.com/melizeche/dolarPy -  melizeche )
*/

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type apiData map[string]map[string]map[string]float64

func getAPIData() apiData {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://dolar.melizeche.com/api/1.0/", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "DolarpyWrapper/Golang")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	data := apiData{}
	json.Unmarshal(body, &data)

	return data
}

func Reference() float64 {
	data := getAPIData()
	return data["dolarpy"]["bcp"]["referencial_diario"]
}

func Providers() []string {
	data := getAPIData()

	providers := []string{}
	for key := range data["dolarpy"] {
		providers = append(providers, key)
	}

	return providers
}

func Purchase(provider string) float64 {
	data := getAPIData()
	if provider == "" {
		return data["dolarpy"]["bcp"]["compra"]
	} else {
		return data["dolarpy"][provider]["compra"]
	}
}

func Sale(provider string) float64 {
	data := getAPIData()
	if provider == "" {
		return data["dolarpy"]["bcp"]["venta"]
	} else {
		return data["dolarpy"][provider]["venta"]
	}
}

func All() map[string]map[string]float64 {
	data := getAPIData()
	return data["dolarpy"]
}
