package main

import "fmt"

func main() {
	var text string
	var width int
	var res string
	fmt.Scanf("%s %d", &text, &width)

	if len(text) <= width {
		res = text
		fmt.Println(res)
		return
	}
	res = text[:width] + "..."
	fmt.Println(res)
}
