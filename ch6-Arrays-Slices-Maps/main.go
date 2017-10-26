package main

import "fmt"

func main() {
	// arrays()
	// slices()
	maps()
}

func arrays() {
	// var grades [5]float64
	// grades[0] = 98
	// grades[1] = 93
	// grades[2] = 77
	// grades[3] = 82
	// grades[4] = 83

	// var total float64 = 0
	// for i := 0; i < len(grades); i++ {
	// 	total += grades[i]
	// }
	
	//another way...
	grades := [5]float64{ 98, 93, 77, 82, 83, }
	var total float64 = 0
	for _ , value := range grades {
		total += value
	}
	
	fmt.Println(total / float64(len(grades)))
}

func slices() {
	slice1 := []int{ 1, 2, 3 }
	slice2 := append(slice1, 4, 5)
	fmt.Println("slice2 ",slice2)
	slice3 := make([]int, 2)
	fmt.Println("slice3", slice3)
	copy(slice3, slice1)
	fmt.Println("slice1 copy", slice3)
}

func maps() {

}