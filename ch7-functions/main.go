package main

import "fmt"

func average(xs []float64) float64 {
	// panic("not implemented")
	total := 0.0
	for _, v := range xs {
		total += v
	}

	return total / float64(len(xs))
}

func multipleReturn() (int, int) {
	return 5, 6
}

//variadic functions
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

//closures
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

//recursion
func factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//defer,panic, recover
//defer
func first() {
	fmt.Println("first")
}
func second() {
	fmt.Println("second")
}

// problems
func addInts(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func halves(num int) (int, bool) {
	half := num / 2
	if half%2 == 0 {
		return half, true
	}
	return 0, false
}

func findLargest(nums ...int) int {
	largest := nums[0]
	for _, n := range nums {
		if n > largest {
			largest = n
		}
	}
	return largest
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

//main
func main() {
	// fmt.Println("find average:")
	// fmt.Println(average([]float64{2.2, 4.4, 1.0}))
	// x, y := multipleReturn()
	// fmt.Print(x, y, "\n")
	// fmt.Println(add(1, 2, 3, 4))
	// closures := func(x, y int) int {
	// 	return x + y
	// }
	// fmt.Println(closures(1, 1))
	// nextEven := makeEvenGenerator()
	// fmt.Println(nextEven())
	// fmt.Println(nextEven())
	// fmt.Println(nextEven())
	// fmt.Println("factorial", factorial(5))
	// defer second()
	// first()
	// defer func() {
	// 	str := recover()
	// 	fmt.Println(str)
	// }()
	// panic("panic!!!")

	fmt.Println(addInts([]int{3, 5, 7}))
	fmt.Println(halves(10))
	fmt.Println(halves(4))
	fmt.Println(findLargest(40, 2, 500, 15))
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(fib(10))
}
