package main

import "fmt"

func volume(r, t float64) float64 {
	return 3.14 * r * r * t
}

func massa(r, t, p float64) float64 {
	v := volume(r, t)
	return v * p
}

func display(m1, m2 float64) {
	if m1 == m2 {
		fmt.Println("Balance")
	} else {
		if m1 > m2 {
			fmt.Println(m1 - m2)
		} else {
			fmt.Println(m2 - m1)
		}
	}
}

func main() {
	var r float64
	var tKiri, tKanan float64
	var mjKiri, mjKanan float64
	var massaKiri, massaKanan float64

	fmt.Print("Masukan Jari-Jari alas tabung : ")
	fmt.Scan(&r)

	fmt.Print("Masukan Tinggi Zat Cair Tabung kiri : ")
	fmt.Scan(&tKiri)

	fmt.Print("Masukan massa jenis zat cair tabung kiri : ")
	fmt.Scan(&mjKiri)

	fmt.Print("Masukan Tinggi Zat Cair Tabung kanan : ")
	fmt.Scan(&tKanan)

	fmt.Print("Masukan massa jenis zat cair tabung kanan : ")
	fmt.Scan(&mjKanan)

	massaKiri = massa(r, tKiri, mjKiri)
	massaKanan = massa(r, tKanan, mjKanan)

	display(massaKiri, massaKanan)
}
