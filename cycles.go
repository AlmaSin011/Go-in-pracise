package main

import "fmt"

func main() {
	//var x1, y1, x2, y2, d float64
	//fmt.Scan(&x1, &y1, &x2, &y2)
	//d = math.Sqrt((math.Pow((x1 - x2), 2)) + (math.Pow((y1 - y2), 2)))
	//fmt.Println(d)

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	const n = 10
	for i := range n {
		fmt.Print(i, " ")
	}

}
