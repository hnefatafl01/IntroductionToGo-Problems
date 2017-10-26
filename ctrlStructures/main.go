package main

import "fmt"

func main() {
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for i := 1; i < 10; i++ {
	// 	// fmt.Println(i)
	// 	if (i % 2 == 0) {
	// 		fmt.Println(i, "even")
	// 	} else {
	// 		fmt.Println(i, "odd")
	// 	}
	// }

	// for i := 0; i < 10; i++ {
	// 	switch i {
	// 		case 0: fmt.Println("Zero")
	// 		case 1: fmt.Println("One")
	// 		case 2: fmt.Println("Two")
	// 		case 3: fmt.Println("Three")
	// 		case 4: fmt.Println("Four")
	// 		case 5: fmt.Println("Five")
	// 		case 6: fmt.Println("Six")
	// 		case 7: fmt.Println("Seven")
	// 		case 8: fmt.Println("Eight")
	// 		case 9: fmt.Println("Nine")
	// 		default: fmt.Println("Unknown Number")
	// 	}
	// }

	// divThree()
	fizzBuzz()
}

func divThree() {
	for i := 1; i < 100; i++ {
		if (i % 3 == 0) {
			fmt.Println(i)
		}
	}
}

func fizzBuzz() {
	for i := 1; i < 100; i++ {
		if (i % 3 == 0 && i % 5 == 0) {
			fmt.Println("FizzBuzz")
		} else if (i % 3 == 0) {
			fmt.Println("Fizz")
		} else if (i % 5 == 0) {
			fmt.Println("Buzz")
		}
	}
}
