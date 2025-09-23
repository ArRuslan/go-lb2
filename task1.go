package main

import (
	"fmt"
	"math/rand"
)

/* Для введення/виведення масиву та його обробки використовувати функцію */
/* Дан масив A(N). Знайти: max(a[2], a[4], ..., a[2k]) + min(a[1], a[3], ..., a[2k+1]) */
func main() {
	var err error
	var n int

	fmt.Print("Enter N (at least 2): ")
	_, err = fmt.Scan(&n)
	if err != nil {
		panic(err)
	}

	if n < 2 {
		panic("Number must be greater than 2!")
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(128)
	}

	fmt.Printf("Array: %v\n", arr)

	max_ := arr[0]
	min_ := arr[1]

	for i, num := range arr {
		if i%2 == 0 {
			max_ = max(max_, num)
		} else {
			min_ = min(min_, num)
		}
	}

	result := max_ + min_
	fmt.Printf("Result: %d\n", result)
}
