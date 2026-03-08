package main
import "fmt"

func main() {

	var war1, war2, war3, war4 string
	berhasil := "True"

	i := 1

	for i <= 5 {
		fmt.Print("Percobaan -", i, ": ")
		fmt.Scan(&war1, &war2, &war3, &war4)

		if war1 != "merah" || war2 != "kuning" || war3 != "hijau" || war4 != "ungu" {
			berhasil = "False"
		}

		i++
	}

	fmt.Println(berhasil)
}