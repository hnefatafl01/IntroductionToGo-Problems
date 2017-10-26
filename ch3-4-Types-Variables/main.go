package main;

import "fmt"


func main() {
	nums()
	strs()
	bools()
	// mult(2, 8)
	fmt.Println(32132 * 42452)
	// printer()
	// tempConverter()
	lengthConverter()
}

func nums() {
	fmt.Println("1 + 1 =", 1 + 1)
	fmt.Println("1.0 + 1.0 =", 1.0 + 1.0)
}

func strs() {
	var x string = "Hello World"
	fmt.Println(len(x))
	fmt.Println(x[1])
	fmt.Println("Hello " + "World")
}

func bools() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println((true && false) || (false && true) || !(false && false))
}

func printer() {
	var (
		a = 10.0
		b = "turtle"
	)
	fmt.Print("Enter a number " + b + ": ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2 + a
	fmt.Println(output)
}

func tempConverter() {
	var F float64
	fmt.Print("Enter degrees F: ")
	fmt.Scanf("%f", &F)
	output := (F - 32) * 5/9
	fmt.Println(output)
}

func lengthConverter() {
	var meters float32
	fmt.Print("Enter length ft:")
	fmt.Scanf("%f", &meters)
	output := meters * 0.3048
	fmt.Println(output)
}

// func mult(n1, n2) {
// 	fmt.Println(n1 * n2)
// }

// 128| 64| 32|16 | 8 | 4 | 2 | 1 
// 2^7|2^6|2^5|2^4|2^3|2^2|2^1|2^0
//  1 | 1 | 1 | 1 | 1 | 1 | 1 | 1

 