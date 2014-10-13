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

func bold(t string) string {
	return "\033[1m" + t + "\033[0m"
}

func prettyPrintJSON(r Recipe) {
	fmt.Println(bold(r.Name))
	if r.Description != "" {
		fmt.Println(r.Description)
	}
	fmt.Println(r.Method)
	fmt.Println(r.Glass)
	if r.Garnish != "" {
		fmt.Println(r.Garnish)
	}
	if r.Variations != "" {
		fmt.Println("Variations: ")
		fmt.Println("\t", r.Variations)
	}
	fmt.Println("Ingredients: ")
	for _, v := range r.Ingredients {
		if v.Parts != "" {
			fmt.Println("\t", bold(v.Parts), " part(s) /", bold(v.Amount), v.AmountUnits, v.IngredientName)
		} else {
			if v.AmountUnits != "" {
				fmt.Println("\t", bold(v.Amount), v.AmountUnits, v.IngredientName)
			} else {
				fmt.Println("\t", bold(v.Amount), v.IngredientName)
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
