package main

import (
	"fmt"
)

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
func permutasi(n, r int) int {
	if n < r {
		return 0
	}
	return faktorial(n) / faktorial(n-r)
}
func kombinasi(n, r int) int {
	if n < r {
		return 0
	}
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {

	var a, b, c, d int
	fmt.Print("Masukkan nilai a, b, c, d: ")
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	Pac := permutasi(a, c)
	Cac := kombinasi(a, c)
	Pbd := permutasi(b, d)
	Cbd := kombinasi(b, d)

	fmt.Printf("p(a, c) = %d\n", Pac)
	fmt.Printf("c(a, c) = %d\n", Cac)
	fmt.Printf("p(b, d) = %d\n", Pbd)
	fmt.Printf("c(b, d) = %d\n", Cbd)
}
