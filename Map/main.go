package main

func main() {
	//var colors map[string]string

	//colors := make(map[string]string)

	//colors["white"] = "#ffffff"

	//delete(colors, "white")

	colors := map[string]string{
		"red":    "#ff0000",
		"orange": "#fca103",
		"yellow": "#f8fc03",
	}
	//We are coping the reference
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		println(color + " -> " + hex)
	}
}
