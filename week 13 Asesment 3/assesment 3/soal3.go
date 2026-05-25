package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func posisi(T tabPartai, n int, nama int) int {
	var i int
	for i = 0; i < n; i++ {
		if T[i].nama == nama {
			return i
		}
	}
	return -1
}

func main() {
	var p tabPartai
	var n, x, idx int

	fmt.Println("Masukkan proses input suara :")
	fmt.Scan(&x)

	for x != -1 {
		idx = posisi(p, n, x)

		if idx == -1 {
			p[n].nama = x
			p[n].suara = 1
			n++
		} else {
			p[idx].suara++
		}

		fmt.Scan(&x)
	}

	var i, j int
	var temp partai

	for i = 1; i < n; i++ {
		temp = p[i]
		j = i - 1

		for j >= 0 && p[j].suara < temp.suara {
			p[j+1] = p[j]
			j--
		}

		p[j+1] = temp
	}

	fmt.Println()
	fmt.Println("Hasil Perhitungan suara :")

	for i = 0; i < n; i++ {
		fmt.Printf("%d(%d) ", p[i].nama, p[i].suara)
	}
}