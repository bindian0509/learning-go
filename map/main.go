package main

import "fmt"

func main() {

	colors := map[string]string{
		"RED":    "#FF0000",
		"GREEN":  "#008000",
		"BLUE":   "#0000FF",
		"YELLOW": "#FFFF00",
	}
	fmt.Println(colors)
	fmt.Println(colors["RED"])
	//delete
	delete(colors, "RED")
	fmt.Println(colors)
	//manipulate
	colors["RED"] = "LAAL"
	fmt.Println(colors["RED"])
	fmt.Println(colors)

	// two ways of defining the maps
	var rangs map[string]string
	fmt.Println(rangs)

	rangeela := make(map[string]string)
	fmt.Println(rangeela)

	changeMap(colors)
	printMap(colors)

}

func printMap(c map[string]string) {
	for col, hex := range c {
		fmt.Println("Hex code for", col, "is", hex)
	}
}

func changeMap(c map[string]string) {
	c["LAAL"] = "Ishq"
}
