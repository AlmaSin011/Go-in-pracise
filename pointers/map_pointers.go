package main

import "fmt"

// Если функция изменяет содержимое карты, передавайте ее как значение.
func changeMap(m map[string]int) {
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
}

// Если функция полностью заменяет саму карту,
// передавайте ее как значение и возвращайте новое значение.
func clearMap(m map[string]int) map[string]int {
	m = map[string]int{}
	return m
}

// Либо передавайте карту как указатель.
func clearMapPtr(m *map[string]int) {
	*m = map[string]int{}
}

func main() {
	m := map[string]int{}
	m["four"] = 4
	changeMap(m)
	fmt.Println(m)
	// map[four:4 one:1 three:3 two:2]

	m = clearMap(m)
	fmt.Println(m)
	// map[]
}
