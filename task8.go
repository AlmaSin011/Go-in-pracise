package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// строки в go байтовые. если нужно работать со сторойокак с символами - используем руны
func main() {
	phrase := readString()
	words := strings.Fields(phrase)
	abbrRunes := []rune{}

	for _, word := range words {
		runes := []rune(word)
		if len(runes) > 0 && unicode.IsLetter(runes[0]) {
			abbrRunes = append(abbrRunes, runes[0])
		}
	}
	abbr := []rune(strings.ToUpper(string(abbrRunes)))

	// 1. Разбейте фразу на слова, используя `strings.Fields()`
	// 2. Возьмите первую букву каждого слова и приведите
	//    ее к верхнему регистру через `unicode.ToUpper()`
	// 3. Если слово начинается не с буквы, игнорируйте его
	//    проверяйте через `unicode.IsLetter()`
	// 4. Составьте слово из получившихся букв и запишите его
	//    в переменную `abbr`
	// ...

	fmt.Println(string(abbr))
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
