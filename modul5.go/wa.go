<<<<<<< HEAD
package main

import "fmt"

func bintang(n, i int) {
	if i > n {
		return
	}

	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()

	bintang(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	bintang(n, 1)
}
=======
package main

import "fmt"

func bintang(n, i int) {
	if i > n {
		return
	}

	for j := 0; j < i; j++ {
		fmt.Print("*")
	}
	fmt.Println()

	bintang(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	bintang(n, 1)
}
>>>>>>> 6e1b229356f910a2ca82b3a2ea45dcc52276fc6e
