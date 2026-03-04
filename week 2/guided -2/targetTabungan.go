package main

import "fmt"

func main() {
	var target, tabungan, total, hari int

	fmt.Print("Masukan Target uang yang ingin di capai: ")
	fmt.Scan(&target)

	total = 0
	hari = 0

	for total < target {
		hari++
		fmt.Print("Masukan Minimal tabungan hari ke-: ", hari)
		fmt.Scanln(&tabungan)

		total = total + tabungan
	}
	fmt.Printf("Selamat Target Telah Tercapai dalam %d dari: \n", hari)
	fmt.Printf("Total Tabungan Anda Terkumpul %d : ", total)
}
