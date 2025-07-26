package main

import "fmt"

func main() {
	var source string
	var times int
	var result string

	fmt.Scan(&source, &times)

	for i := 1; i <= times; i++ {
		result = result + source
	}
	fmt.Println(result)
}
