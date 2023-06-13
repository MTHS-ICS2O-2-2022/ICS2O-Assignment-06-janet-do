// Created by: Janet Do
// Created on: Dec 2022
//
// This program displays, "Random Cat Image Generator, API"

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := http.Client{}
	response, err := client.Get("https://api.thecatapi.com/v1/images/search")
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

	fmt.Println("")
	fmt.Println(string(body))
	fmt.Println("\nHave a random image of a cat")
	fmt.Println("\n\nDone")
}
