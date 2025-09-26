package main

import (
	"bufio"
	"fmt"
	"os"
)

type CustomString string
type FilterableByStr interface {
	filter(removePredicate func(char rune) bool) (result string)
}

func (str CustomString) filter(removePredicate func(char rune) bool) (result string) {
	for _, char := range []rune(str) {
		if !removePredicate(char) {
			result += string(char)
		}
	}
	return
}

/* Для вирішення завдання використовувати комплексно функції, методи та інтерфейси (для обробки рядків писати свої функції та методи) */
/* Видалити всі голосні літери з рядка */
func main() {
	var str FilterableByStr

	fmt.Print("Enter a text: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str = CustomString(scanner.Text())
	newStr := str.filter(func(char rune) bool {
		return char == 'а' || char == 'у' || char == 'е' || char == 'ї' || char == 'і' || char == 'о' || char == 'є' || char == 'я' || char == 'и' || char == 'ю' ||
			char == 'А' || char == 'У' || char == 'Е' || char == 'Ї' || char == 'І' || char == 'О' || char == 'Є' || char == 'Я' || char == 'И' || char == 'Ю' ||
			char == 'a' || char == 'e' || char == 'u' || char == 'i' || char == 'o' ||
			char == 'A' || char == 'E' || char == 'U' || char == 'I' || char == 'O'
	})

	fmt.Printf("Result string: %s\n", newStr)
}
