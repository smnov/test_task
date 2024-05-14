package main

import (
	"fmt"
	"math"
)

// Задача 3. Написать функцию/метод, которая возвращает массив простых чисел
// в диапазоне (2 числа - минимальное и максимальное) заданных чисел.
// Например, на вход переданы 2 числа: от 11 до 20  (диапазон считается включая граничные значения).
// На выходе программа должна выдать [11, 13, 17, 19]

func main() {
	giveSimpleNums(11, 20)
}

func giveSimpleNums(a, b int) []int {
	var result []int
	for i := a; i <= b; i++ {
		if numIsSimple(i) {
			result = append(result, i)
		}
	}
	fmt.Printf("result: %v\n", result)
	return result
}

func numIsSimple(n int) bool {
	if n < 2 {
		return false
	}
	sq_root := int(math.Sqrt(float64(n)))
	for i := 2; i <= sq_root; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
