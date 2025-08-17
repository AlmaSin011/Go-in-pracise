package main

import "fmt"

func addval(nval int, delta int) { // модиифицируется копия, ориганал не меняется
	nval += delta
}

func addptr(nptr *int, delta int) { // меняется оригинальное значение потому что передается указатель
	// идет доступ оп указателю и меняется значение
	*nptr += delta
}

func main() {
	var iptr *int
	fmt.Println(iptr)

	i := 42
	iptr = &i         // здесь адрес переменной в памяти.
	fmt.Println(iptr) // выводим адрес

	fmt.Println(*iptr) // выводим то значение на которое ссылается указатель

	n := 42
	addval(n, 3)
	fmt.Println(n)

	addptr(&n, 3)
	fmt.Println(n)
}
