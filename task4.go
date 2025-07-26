package main

import "fmt"

func main() {

	var code string
	fmt.Scan(&code)
	var lang string

	switch code {
	case "ru":

		lang = "Russian"
	case "en":
		lang = "English"
	case "fr ":
		lang = "French"
	default:
		lang = "Unknown"

	}
	fmt.Println(lang)
}
