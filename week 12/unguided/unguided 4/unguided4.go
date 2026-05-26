package main

import "fmt"

const NMAX = 100

type Buku struct {
	judul     string
	penulis   string
	penerbit  string
	tahun     int
	eksemplar int
	rating    int
}

type DaftarBuku [NMAX]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].penerbit,
			&pustaka[i].tahun,
			&pustaka[i].eksemplar,
			&pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	max := pustaka[0]

	for i := 1; i < n; i++ {
		if pustaka[i].rating > max.rating {
			max = pustaka[i]
		}
	}

	fmt.Println(max.judul, max.penulis, max.penerbit, max.tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku
	var j int

	for i := 1; i < n; i++ {
		temp = pustaka[i]
		j = i - 1

		for j >= 0 && pustaka[j].rating < temp.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}

		pustaka[j+1] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	batas := 5
	if n < 5 {
		batas = n
	}

	for i := 0; i < batas; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n, r int) {
	kiri := 0
	kanan := n - 1
	var tengah int

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2

		if pustaka[tengah].rating == r {
			fmt.Println(
				pustaka[tengah].judul,
				pustaka[tengah].penulis,
				pustaka[tengah].penerbit,
				pustaka[tengah].tahun,
				pustaka[tengah].eksemplar,
				pustaka[tengah].rating)
			return
		} else if r > pustaka[tengah].rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	fmt.Println("Tidak ada buku dengan rating seperti itu")
}

func main() {
	var pustaka DaftarBuku
	var n, r int

	fmt.Scan(&n)

	DaftarkanBuku(&pustaka, n)
	CetakTerfavorit(pustaka, n)

	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)

	fmt.Scan(&r)
	CariBuku(pustaka, n, r)
}
