package main

// Pretty-printing cocktail recipes in JSON format

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)

type Recipe struct {
	Id string `json: "id"`
	Name string `json: "name"`
	Description string `json: "description"`
	Method string `json: "method"`
	Variations string `json: "variations"`
	Ingredients []struct {
		Parts string `json: "parts"`
		Amount string `json: "amount"`
		AmountUnits string `json: "amountUnits"`
		IngredientName string `json: "ingredientName"`
	} `json: "ingredients"`
	Glass string `json: "glass"`
	Garnish string `json: "garnish"`
}

func prettyPrintJSON(r Recipe) {
	fmt.Println(r.Name)
	if r.Description != "" {
		fmt.Println(r.Description)
	}
	fmt.Println(r.Method)
	fmt.Println(r.Glass)
	if r.Garnish != "" {
		fmt.Println(r.Garnish)
	}
	if r.Variations != "" {
		fmt.Println(r.Variations)
	}
	for _, v := range r.Ingredients {
		if v.Parts != "" {
			fmt.Println(v.Parts, " part(s) /", v.Amount, v.AmountUnits, v.IngredientName)
		} else {
			if v.AmountUnits != "" {
				fmt.Println(v.Amount, v.AmountUnits, v.IngredientName)
			} else {
				fmt.Println(v.Amount, v.IngredientName)
			}
		}
	}
}

func main() {
	var recipe Recipe
	f, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	unmarshal_err := json.Unmarshal(f, &recipe)
	if unmarshal_err != nil {
		fmt.Println(unmarshal_err)
	}

	prettyPrintJSON(recipe)
}

