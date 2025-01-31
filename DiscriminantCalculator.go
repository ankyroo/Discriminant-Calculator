package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Println("Discriminant Calculator")

	fmt.Println("Enter a:")
	fmt.Scan(&a)

	fmt.Println("Enter b:")
	fmt.Scan(&b)

	fmt.Println("Enter c:")
	fmt.Scan(&c)

	D := (b * b) - 4*(a*c)

	if D > 0 {
		var x1, x2 float64

		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)

		fmt.Println("Your equation has 2 roots\nD = " + fmt.Sprint(D))
		fmt.Println("x₁ = " + fmt.Sprint(x1) + "\nx₂ = " + fmt.Sprint(x2))

	} else if D == 0 {
		var x float64

		x = (-b) / (2 * a)

		fmt.Println("Your equation has 1 root, because D = " + fmt.Sprint(D))
		fmt.Println("x = " + fmt.Sprint(x))
	} else if D < 0 {
		fmt.Println("Your equation has no roots, because D < 0\nD = " + fmt.Sprint(D))
	}
	fmt.Scan(&a)
}
