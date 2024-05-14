package main

import "fmt"

// Задача 4. Написать метод, который в консоль выводит таблицу умножения.
// На вход метод получает число, до которого выводит таблицу умножения.
// В консоли должна появиться таблица. Например, если на вход пришло число 5, то получим:

func main() {
	drawMultiplicationTable(10)
}

func drawMultiplicationTable(n int) {
	width := len(fmt.Sprintf("%d", n*n))

	fmt.Printf("%*s ", width, "")
	for i := 1; i <= n; i++ {
		fmt.Printf("%*d ", width, i)
	}
	fmt.Println()

	for k := 1; k <= n; k++ {
		fmt.Printf("%*d ", width, k)
		for i := 1; i <= n; i++ {
			fmt.Printf("%*d ", width, i*k)
		}
		fmt.Println()
	}
}
