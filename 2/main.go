// Задача 2. Написать функцию/метод, которая на вход
// получает массив положительных целых чисел произвольной длины.
// Например [42, 12, 18].
// На выходе возвращает массив чисел, которые являются общими
// делителями для всех указанных числе. В примере это будет [2, 3, 6].
package main

import (
	"fmt"
	"math"
)

func main() {
	testArr := []int{42, 12, 18}
	fmt.Println(commonDivisor(testArr))
}

func commonDivisor(input []int) []int {
	result := input[0]
	for i := 1; i < len(input); i++ {
		result = gcd(result, input[i])
	}
	commonDivisors := make([]int, 0)
	for i := 1; i <= int(math.Sqrt(float64(result))); i++ {
		if result%i == 0 {
			commonDivisors = append(commonDivisors, i)
			if result/i != i {
				commonDivisors = append(commonDivisors, result/i)
			}
		}
	}

	return commonDivisors
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
