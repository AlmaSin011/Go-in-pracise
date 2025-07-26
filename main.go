package main

import (
	"fmt"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

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
