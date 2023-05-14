package main

import "fmt"

func main() {
	// 생성 방법 1
	//var colors map[string]string

	// 생성 방법 2
	//colors := make(map[string]string)

	// 생성 방법 3 -> 생성 + 값 초기화
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	colors["white"] = "#ffffff"

	// map 원소 삭제
	delete(colors, "white")

	fmt.Printf("%+v", colors)
	fmt.Println()
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%+v: %+v", color, hex)
		fmt.Println()
	}
}
