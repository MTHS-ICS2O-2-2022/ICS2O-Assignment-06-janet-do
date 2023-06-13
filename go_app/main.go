// Copyright (c) 2022 Janet DoAll rights reserved
//
// Created by: Janet Do
//This program generates a different cat API everytime it is run 
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

func getCat(urlAddress string) {
	response, err := http.Get(urlAddress)
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
	urlAddress := "https://api.thecatapi.com/v1/images/search"
	getCat(urlAddress)

	fmt.Println("\nDone.")
}
