package main

import "fmt"

func main() {
	var limite int
	fmt.Scanln(&limite)
	var out float64
	if limite == 0 {
		out = 1
	} else {
		for i := 1; i <= limite; i++ {
			out += float64((1.0 / float64(factorial(i))))
		}
	}
	fmt.Println(1.0 + out)
}

func factorial(n int) uint64 {
	var f uint64 = 1

	for i := 1; i <= n; i++ {
		f *= uint64(i)
	}

	return f
}
