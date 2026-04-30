package main
import "fmt"

const NMAX int = 1000
type tabBerat [NMAX]float64
func main() {
	var A tabBerat
	var n, i int
	var min, max float64

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	min = A[0]
	max = A[0]

	for i = 1; i < n; i++ {
		if A[i] < min {
			min = A[i]
		}
		if A[i] > max {
			max = A[i]
		}
	}

	fmt.Println(min, max)
}