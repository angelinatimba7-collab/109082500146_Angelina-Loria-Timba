<<<<<<< HEAD
package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah suku: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print(fib(i), " ")
	}
=======
package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah suku: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print(fib(i), " ")
	}
>>>>>>> 6e1b229356f910a2ca82b3a2ea45dcc52276fc6e
}