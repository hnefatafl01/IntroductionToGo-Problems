package main

import "fmt"

func main() {
	// arrays()
	// slices()
	// maps()
	// probs()
	findSmallest()
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
	grades := [5]float64{98, 93, 77, 82, 83}
	var total float64 = 0
	for _, value := range grades {
		total += value
	}

	fmt.Println(total / float64(len(grades)))
}

func slices() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println("slice2 ", slice2)
	slice3 := make([]int, 2)
	fmt.Println("slice3", slice3)
	copy(slice3, slice1)
	fmt.Println("slice1 copy", slice3)
}

func maps() {
	//declaration only =>
	// var x map[string]int
	//initialize map with key type string =>
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)
	// initialize map with key type int
	y := make(map[int]int)
	y[0] = 11
	fmt.Println(y)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println("elements contains ", elements["Li"])

	// check for element in map
	name, ok := elements["Un"]
	fmt.Println(name, ok)

	//alternative map creation
	elements2 := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	for _, value := range elements2 {
		fmt.Println("elements2", value)
	}

	elements3 := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}
	if el, ok := elements3["Li"]; ok {
		fmt.Println("elements3", el["name"], el["state"])
	}
}

func probs() {
	arr := [4]int{1, 2, 3, 4}
	third := arr[3]
	fmt.Println(third)

	//make(type, slice length, array capacity)
	slc := make([]int, 3, 9)
	// fmt.Println(len(slc))
	fmt.Println(slc)

	slc2 := arr[0:2]
	fmt.Println(slc2)

	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5])
}

func findSmallest() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	smallest := x[0]
	for i := 0; i < len(x); i++ {
		if x[i] < smallest {
			smallest = x[i]
		}
	}
	fmt.Println(smallest)
}
