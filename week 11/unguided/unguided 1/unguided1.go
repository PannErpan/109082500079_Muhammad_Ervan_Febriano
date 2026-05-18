package main

import "fmt"

const NMAX int = 21

func main() {
	var suara [NMAX]int
	var x int
	var totalMasuk, totalSah, i int

	fmt.Scan(&x)

	for x != 0 {
		totalMasuk++

		if x >= 1 && x <= 20 {
			suara[x]++
			totalSah++
		}

		fmt.Scan(&x)
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", totalSah)

	for i = 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Println(i, ":", suara[i])
		}
	}
}
