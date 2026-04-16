package main
import"fmt"

type angka int 
type kata string

type buku struct {
	judul kata 
	penulis kata 
	penerbit kata 
	tahunterbit angka 
	jumlahhalaman angka
}

func main(){
	var b buku 

	fmt.Println("=== INPUT BIODATA BUKU ===")
	fmt.Print("Masukan Judul Buku : ")
	fmt.Scan(&b.judul)

	fmt.Print("Masukan Nama Penulis : ")
	fmt.Scan(&b.penulis)

	fmt.Print("Masukan Penerbit : ")
	fmt.Scan(&b.penerbit)

	fmt.Print("Masukan Tahun Terbit : ")
	fmt.Scan(&b.tahunterbit)

	fmt.Print("Masukan Jumlah Halaman : ")
	fmt.Scan(&b.jumlahhalaman)

	fmt.Println("\n=== BIODATA BUKU ===")
	fmt.Println("Judul Buku : ", b.judul)
	fmt.Println("Penulis : ", b.penulis)
	fmt.Println("Tahun Terbit : ", b.tahunterbit)
	fmt.Println("Jumlah Halaman : ", b.jumlahhalaman)
}