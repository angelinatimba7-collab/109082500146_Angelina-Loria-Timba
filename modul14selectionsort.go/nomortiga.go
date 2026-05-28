package main

import "fmt"

const NMAX int = 1000000

type arrInt [NMAX]int

func insertionSort(A *arrInt, n int) {
	var i, j, temp int

	for i = 1; i < n; i++ {

		temp = A[i]
		j = i

		for j > 0 && temp < A[j-1] {
			A[j] = A[j-1]
			j--
		}

		A[j] = temp
	}
}

func hitungMedian(A arrInt, n int) int {

	if n%2 == 1 {
		return A[n/2]
	}

	return (A[(n/2)-1] + A[n/2]) / 2
}

func main() {

	var data arrInt
	var x, n int

	n = 0

	for {

		fmt.Scan(&x)

		if x == -5313 {
			break
		}

		if x == 0 {

			insertionSort(&data, n)

			fmt.Println(hitungMedian(data, n))

		} else {

			data[n] = x
			n++
		}
	}
}
