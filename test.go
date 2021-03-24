package main

import (
	"fmt"
)

// Extra desafio tenico ***

// Escreva um método que retorne o N elemento da sequencia abaixo:
// Posições: 0  1  2  3  4  5  6   7   8   9  10
// Valores:  0  2  3  5  6  8  9  11  12  14  15

// func findNext(pos int) int {
// 	// Sua implementação
// }

// Ex.:
// 	func (0) = 0
// 	func (1) = 2
// 	func (2) = 3
// 	func (3) = 5
// 	[…]
// 	func (10) = 15
// 	func (11) = ?

func main() {
	next := findLele(10)

	fmt.Println("next:", next)
}

func findNext(pos int) int {

	arr := make([]int, 0)
	value := 0
	for i := 0; i <= pos; i++ {
		b := i % 2
		multiplier := 2

		if b == 0 {
			multiplier = 1
		}

		if i > 0 {
			value += multiplier
		}

		arr = append(arr, value)
	}

	fmt.Println(arr)

	return arr[pos]
}

func findLele(pos int) int {
	// 	(0 * 2) - 0 = 0
	// (1 * 2) - 0 = 2
	// (2 * 2) - 1 = 3
	// (3 * 2) - 1 = 5
	// (4 * 2) - 2 = 7
	// (5 * 2) - 2 = 8
	// ...
	// findNext = (pos) => {
	// 	const rule = pos % 2 == 0 ? pos / 2 : (pos - 1) / 2;
	// 	const number = pos * 2 - rule;
	// 	return number;
	//   };
	//   console.log(findNext(11));

	rule := pos / 2
	if (pos % 2) != 0 {
		rule = (pos - 1) / 2
	}

	number := pos*2 - rule

	return number
}
