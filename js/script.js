// Copyright (c) 2022 Janet DoAll rights reserved
//
// Created by: Janet Do
// Created on: Dec 2022
// This file contains the JS functions for index.html

/**
 * Get API info.
 */
// code from: https://thecatapi.com/
"use strict"

const getCat = async (URLAddress) => {
  try {
    const result = await fetch(URLAddress)
    const jsonData = await result.json()
    console.log(jsonData[0].url)
    console.log(jsonData[0])
    document.getElementById("api-image").innerHTML = '<img src="' + jsonData[0].url + '" alt="cat image" width= 25% height 20%>'

  } catch (err) {
    console.log(err)
  }
}

getCat("https://api.thecatapi.com/v1/images/search")