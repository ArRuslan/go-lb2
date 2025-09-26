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

type CustomStringSlice []CustomString
type NumberedCustomStringSlice struct {
	indexes []int
	strings []CustomString
}

func (strings CustomStringSlice) searchSubstring(substr CustomString) (result NumberedCustomStringSlice) {
	for idx, str := range strings {
		contains := false

		i1 := 0
		i2 := 0

		for ; i1 < len(str); i1++ {
			if []rune(str)[i1] == []rune(substr)[i2] {
				i2++
				if i2 >= len(substr) {
					contains = true
					break
				}
			} else {
				i2 = 0
			}
		}

		if contains {
			result.indexes = append(result.indexes, idx)
			result.strings = append(result.strings, str)
		}
	}

	return
}

func task4() {
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

func task5() {
	var err error
	var stringsCount int

	fmt.Print("Enter strings count: ")
	_, err = fmt.Scan(&stringsCount)
	if err != nil {
		panic(err)
	}

	strings := make(CustomStringSlice, stringsCount)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < stringsCount; i++ {
		fmt.Printf("String [%d]: ", i)
		scanner.Scan()
		strings[i] = CustomString(scanner.Text())
	}

	foundStrings := strings[1:].searchSubstring(strings[0])
	for idx, str := range foundStrings.strings {
		foundStrings.strings[idx] = CustomString(str.filter(func(char rune) bool {
			return char == ' '
		}))
	}

	for idx, str := range foundStrings.strings {
		fmt.Printf("[%d] \"%s\"\n", foundStrings.indexes[idx]+1, str)
	}
}

/* Для вирішення завдання використовувати комплексно функції, методи та інтерфейси (для обробки рядків писати свої функції та методи) */
/* Видалити всі голосні літери з рядка */
/* Задано послідовність рядків різної довжини. Отримати номери всіх рядків, що містять як фрагмент перший рядок.
В результаті отримати послідовності рядків з виключеними пробілами та знайденими фрагментами першого рядка */
func main() {
	task4()
	task5()
}
