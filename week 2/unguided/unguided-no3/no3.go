package main
import "fmt"

func main() {
	var beratGram int
	var kg, sisa int
	var biayaKg, biayaSisa, total int

	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&beratGram)

	kg = beratGram / 1000
	sisa = beratGram % 1000
	biayaKg = kg * 10000

	if sisa == 500 {
		biayaSisa = sisa * 5
	} else if sisa < 500 {
		biayaSisa = sisa * 15
	} else {
		biayaSisa = 0
	}

	total = biayaKg + biayaSisa

	fmt.Println("\n===== Detail Perhitungan =====")
	fmt.Println("Detail berat :", kg, "kg +", sisa, "gram")
	fmt.Println("Detail biaya : Rp.", biayaKg, "+ Rp.", biayaSisa)
	fmt.Println("Total biaya  : Rp", total)
}