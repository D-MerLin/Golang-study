package main

import "fmt"

func mysqrt(x float64) float64 {
	x = float64(x)
	t := 0
	z := 1.0
	for t < 10 {
		z -= ( z * z - x) / (2 * z)
		t ++
	}
	return z
}	

func main(){
	fmt.Println(mysqrt(9))
}