package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CatData struct {
	URL string `json:"url"`
}

func getCat(URLAddress string) {
	response, err := http.Get(URLAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var jsonData []CatData
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(jsonData) > 0 {
		fmt.Println(jsonData[0].URL)
		// Update the HTML with the cat image
		// ...
	} else {
		fmt.Println("No cat data found.")
	}
}

func main() {
	URLAddress := "https://api.thecatapi.com/v1/images/search"
	getCat(URLAddress)

	fmt.Println("\nDone.")
}
