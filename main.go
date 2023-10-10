package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Country struct {
	Name       Name                
	Cca2       string              `json:"cca2"`
	Cca3       string              `json:"cca3"`
	Region     string              `json:"region"`
	Subregion  string              `json:"subregion"`
	Currencies map[string]Currency `json:"currencies"`
}

type Currency struct {
	Symbol string `json:"symbol"`
}

type Name struct {
	Official string `json:"official"`
}

func fetcher(param string) ([]Country, error) {
	url := fmt.Sprintf("https://restcountries.com/v3.1/region/%s?fields=name,cca2,cca3,region,subregion,currencies", param)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data []Country
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	data, err := fetcher("europe")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
	
	fmt.Println("-----------------")
	fmt.Println("-----------------")
	fmt.Println("-----------------")

	data, err = fetcher("north america")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
}
