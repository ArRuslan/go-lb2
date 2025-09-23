package main

import (
	"fmt"
)

func sign(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}

	return 0
}

/* При вирішенні задачі використовувати методи(??) та функції, а також покажчики при передачі параметрів в методи (???) */
/* Обчислити значення виразу z = (sign(x) + sign(y)) * sign(x + y),
При вирішенні задачі визначити і використати функцію sign:
-1 при x < 0
0 при x = 0
1 при x > 0 */
func main() {
	var err error
	var x, y int

	fmt.Print("Enter X: ")
	_, err = fmt.Scan(&x)
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter Y: ")
	_, err = fmt.Scan(&y)
	if err != nil {
		panic(err)
	}

	z := (sign(x) + sign(y)) * sign(x+y)
	fmt.Printf("Result: %d\n", z)
}
