package main
import "fmt"

const NMAX = 1000
type ArrRumah [NMAX]int

func selectionAsc(A *ArrRumah, n int) {
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

func selectionDesc(A *ArrRumah, n int) {
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
	var n, m, x int
	var ganjil, genap ArrRumah
	var ng, ne int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		ng = 0
		ne = 0

		fmt.Scan(&m)

		for j := 0; j < m; j++ {
			fmt.Scan(&x)

			if x%2 != 0 {
				ganjil[ng] = x
				ng++
			} else {
				genap[ne] = x
				ne++
			}
		}

		selectionAsc(&ganjil, ng)
		selectionDesc(&genap, ne)

		for j := 0; j < ng; j++ {
			fmt.Print(ganjil[j], " ")
		}

		for j := 0; j < ne; j++ {
			fmt.Print(genap[j], " ")
		}

		fmt.Println()
	}
}
