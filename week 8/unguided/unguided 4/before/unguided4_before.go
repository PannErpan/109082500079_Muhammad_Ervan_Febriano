package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0

	for *n < NMAX {
		fmt.Scanf("%c", &ch)

		if ch == '\n' || ch == '\r' {
			continue
		}

		if ch == '.' {
			break
		}

		t[*n] = ch
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func main() {
	var tab tabel
	var m int
	isiArray(&tab, &m)
	balikanArray(&tab, m)
	cetakArray(tab, m)
}
