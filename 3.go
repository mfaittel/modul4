package main

import "fmt"

// Procedure to print the sequence based on the given number n
func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	// Print the final 1
	fmt.Printf("%d\n", n)
}

func main() {
	var n int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&n)

	// Call the procedure to print the sequence
	cetakDeret(n)
}
