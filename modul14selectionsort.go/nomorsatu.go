package main

import "fmt"

const NMAX int = 1000

type arrRumah [1000000]int

func selectionSortAscending(A *arrRumah, n int) {
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

func main() {
	var nDaerah, m, i, j int
	var rumah arrRumah

	fmt.Scan(&nDaerah)

	for i = 0; i < nDaerah; i++ {

		fmt.Scan(&m)

		for j = 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		selectionSortAscending(&rumah, m)

		for j = 0; j < m; j++ {
			fmt.Print(rumah[j])

			if j != m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
