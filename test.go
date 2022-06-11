package main

import (
	"fmt"
)

func aad(x float64, y float64) float64 {
	return x + y

}

// func frg() {
// 	// var f float64 = math.Sqrt(9 + 16)
// 	fmt.Println("Hello", rand.Intn(100))

// }

const x int = 5

func multiple(a, b string) (string, string) {
	return a, b

}

func main() {
	n1 := 1.9
	n2 := 7.3
	a := &n1
	w1, w2 := "hello", "hi"
	fmt.Println(aad(n1, n2))
	fmt.Println(multiple(w1, w2))
	fmt.Println(a)
	fmt.Println(*a)
}
