package main
import "fmt"

func main(){
	var berat,tinggi,bmi float64

	fmt.Print("Masukan Berat badan (KG) = ")
	fmt.Scan(&berat)
	fmt.Print("Masukan Tinggi badan (M) = ")
	fmt.Scan(&tinggi)

	bmi = berat / (tinggi * tinggi)

	fmt.Println("Nilai BMI anda adalah: ", bmi)
}
