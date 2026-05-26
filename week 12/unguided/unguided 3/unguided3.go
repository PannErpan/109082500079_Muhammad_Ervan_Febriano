package main

import "fmt"

func insertionSort(arr []int, n int) {
	var i, j, key int
	for i = 1; i < n; i++ {
		key = arr[i]
		j = i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func cekJarak(arr []int, n int) {
	if n <= 1 {
		fmt.Println("Data berjarak 0")
		return
	}

	selisih := arr[1] - arr[0]

	for i := 2; i < n; i++ {
		if arr[i]-arr[i-1] != selisih {
			fmt.Println("Data berjarak tidak tetap")
			return
		}
	}

	fmt.Printf("Data berjarak %d\n", selisih)
}

func main() {
	var arr [100]int
	var x, n int

	for {
		fmt.Scan(&x)
		if x < 0 {
			break
		}
		arr[n] = x
		n++
	}

	insertionSort(arr[:], n)

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
	cekJarak(arr[:], n)
}