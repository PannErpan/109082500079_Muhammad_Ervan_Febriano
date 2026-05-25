package main

import "fmt"

const NMAX = 100000

type ArrInt [NMAX]int

func insertionSort(A *ArrInt, n int) {
	var i, j, temp int

	for i = 1; i < n; i++ {
		temp = A[i]
		j = i - 1

		for j >= 0 && A[j] > temp {
			A[j+1] = A[j]
			j--
		}

		A[j+1] = temp
	}
}

func median(A ArrInt, n int) int {
	if n%2 != 0 {
		return A[n/2]
	}
	return (A[n/2-1] + A[n/2]) / 2
}

func main() {
	var A ArrInt
	var x, n int

	for {
		fmt.Scan(&x)

		if x == -5313 {
			break
		} else if x == 0 {
			insertionSort(&A, n)
			fmt.Println(median(A, n))
		} else {
			A[n] = x
			n++
		}
	}
}
