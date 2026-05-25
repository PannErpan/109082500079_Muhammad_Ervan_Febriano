package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NMAX = 1001

type Pemain struct {
	nama   string
	gol    int
	assist int
}

type arrPemain [NMAX]Pemain

func lebihBesar(a, b Pemain) bool {
	if a.gol > b.gol {
		return true
	} else if a.gol == b.gol {
		return a.assist > b.assist
	}
	return false
}

func SelectionSort(A *arrPemain, n int) {
	var i, j, idxMax int
	var temp Pemain

	for i = 0; i < n-1; i++ {
		idxMax = i
		for j = i + 1; j < n; j++ {
			if lebihBesar(A[j], A[idxMax]) {
				idxMax = j
			}
		}
		temp = A[i]
		A[i] = A[idxMax]
		A[idxMax] = temp
	}
}

func main() {
	var n int
	var A arrPemain

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan Data Input :")
	fmt.Scan(&n)
	fmt.Scanln()

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()

		data := strings.Fields(line)
		panjang := len(data)

		A[i].gol, _ = strconv.Atoi(data[panjang-2])
		A[i].assist, _ = strconv.Atoi(data[panjang-1])
		A[i].nama = strings.Join(data[:panjang-2], " ")
	}

	SelectionSort(&A, n)

	fmt.Println()
	fmt.Println("Hasil Sorting :")

	for i := 0; i < n; i++ {
		fmt.Println(A[i].nama, A[i].gol, A[i].assist)
	}
}