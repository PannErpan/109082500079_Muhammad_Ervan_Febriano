package main
import "fmt"

const NMAX int = 21

func main() {
	var suara [NMAX]int
	var x, totalMasuk, totalSah int
	var i, ketua, wakil int

	fmt.Scan(&x)

	for x != 0 {
		totalMasuk++

		if x >= 1 && x <= 20 {
			suara[x]++
			totalSah++
		}

		fmt.Scan(&x)
	}

	ketua = 1
	for i = 2; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			ketua = i
		}
	}

	wakil = -1
	for i = 1; i <= 20; i++ {
		if i != ketua {
			if wakil == -1 || suara[i] > suara[wakil] {
				wakil = i
			}
		}
	}

	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", totalSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
