package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	sum := 0
	for _, num := range nums { // get index and element value. We ignore index by underscore
		sum += num
	}
	fmt.Println(sum)

	m := map[string]string{"a": "apple",
		"b": "banana"}
	for key, val := range m {
		fmt.Printf("%s -> %s\n", key, val)
	}

	maps := map[string]string{"a": "apple", "b": "banana"}
	for key := range maps {
		fmt.Println("key:", key)
	}
}
