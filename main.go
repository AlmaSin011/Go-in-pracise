package main

import (
	"fmt"
	"time"
)

func main() {
	var s string
	fmt.Scan(&s)
	d, err := time.ParseDuration(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s, "=", d.Minutes(), "min")

	var a = 123
	fmt.Println(a)

	var sa = "test string"
	fmt.Println(sa)

	const sdf = "sdf"
	fmt.Println(sdf)
}
