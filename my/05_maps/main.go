package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "0xFF0000",
		"green": "0xFF00",
		"blue":  "0xFF",
	}

	// var colors map[string]string
	// colors := make(map[string]string)
	colors["white"] = "0xFFF"

	printMap(&colors)
	printMap(&colors)
}

func printMap(c *map[string]string) {
	(*c)["white"] = "white"
	for color, hex := range *c {
		fmt.Println("Hex code for", color, "is", hex)
	}
	(*c) = make(map[string]string)
}
