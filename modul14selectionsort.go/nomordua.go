package main

import "fmt"

const NMAX int = 1000000

type arrInt [NMAX]int

func selectionSortAscending(A *arrInt, n int) {
	var i, j, idxMin, temp int

	for i = 0; i < n-1; i++ {
		idxMin = i

		for j = i + 1; j < n; j++ {
			if A[j] < A[idxMin] {
				idxMin = j
			}
		}

		temp = A[i]
		A[i] = A[idxMin]
		A[idxMin] = temp
	}
}

func selectionSortDescending(A *arrInt, n int) {
	var i, j, idxMax, temp int

	for i = 0; i < n-1; i++ {
		idxMax = i

		for j = i + 1; j < n; j++ {
			if A[j] > A[idxMax] {
				idxMax = j
			}
		}

		temp = A[i]
		A[i] = A[idxMax]
		A[idxMax] = temp
	}
}

func main() {
	var nDaerah, m int
	var i, j int

	var ganjil, genap arrInt
	var gCount, eCount int
	var x int

	fmt.Scan(&nDaerah)

	for i = 0; i < nDaerah; i++ {

		gCount = 0
		eCount = 0

		fmt.Scan(&m)

		for j = 0; j < m; j++ {

			fmt.Scan(&x)

			if x%2 == 1 {
				ganjil[gCount] = x
				gCount++
			} else {
				genap[eCount] = x
				eCount++
			}
		}

		selectionSortAscending(&ganjil, gCount)
		selectionSortDescending(&genap, eCount)

		for j = 0; j < gCount; j++ {
			fmt.Print(ganjil[j])

			if j != gCount-1 || eCount > 0 {
				fmt.Print(" ")
			}
		}

		for j = 0; j < eCount; j++ {
			fmt.Print(genap[j])

			if j != eCount-1 {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}
