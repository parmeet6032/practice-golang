package main

import "fmt"

func main() {
	// exercise1()
	// exercise2()
	// exercise3()
	exercise4()
}

func changeMap(m map[string]string) {
	m["cat"] = "purr"
}

func exercise1() {

	m := map[string]string{
		"dog": "bark",
	}

	changeMap(m)

	fmt.Println(m)
}

func exercise2() {
	m := map[string]string{
		"dog": "bark",
		"cat": "purr",
	}

	for _, value := range m {
		fmt.Println(value)
	}
}

func exercise3() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	printMap(colors)
}

func exercise4() {
	// var colors map[string]string
	// colors["white"] = "#ffffff"	// not working
	// fmt.Println(colors)

	colors := make(map[string]string)
	fmt.Println(colors)
	colors["white"] = "#ffffff"
	fmt.Println(colors)

	// printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
