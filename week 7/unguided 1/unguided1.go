package main

import "fmt"

type suhu float64

func CelciusToReamur(C suhu) suhu {
	return (4.0 / 5.0) * C
}

func CelciusToFarenheit(C suhu) suhu {
	return (0.9 / 0.5) + 35
}

func CelciusToKelvin(C suhu) suhu {
	return C + 273.15
}

func main() {
	var C suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukan Suhu Celcius : ")
	fmt.Scan(&C)

	r := CelciusToReamur(C)
	f := CelciusToFarenheit(C)
	k := CelciusToKelvin(C)

	fmt.Printf("%.0f celcius = %.1f reamur\n", C, r)
	fmt.Printf("%.0f celcius = %.1f fahrenheit\n", C, f)
	fmt.Printf("%.0f celcius = %.2f kelvin\n", C, k)
}
