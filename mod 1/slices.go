package main

import "fmt"

func main() {
	var arr1 [6]int
	arr1 = [6]int{1,2,3,4,5,6}
	slice := arr1[0:4] // initializing our slice
	fmt.Println(slice)
	slice = arr1[:5] // extended slice, unlike arrays which cannot be modified

	fmt.Println(slice)
	
	for x:= 0; x < 10; x++ {
		slice = append(slice, x)
	}
	fmt.Println(slice)
	// making slice with prior capacity to create an array like slice first so it is less time consuming to work with
	sliceN := make([]int, 10, 20)
	sliceN = []int{1,2,3,4,5,6,6,7,8,56,4,3,3,4,2,3,4,53,}
	fmt.Println(sliceN)
	for j, value := range sliceN {
		fmt.Println(j, value)
	}
	var stringSlice []string
	stringSlice = []string{"Hello", "hi"}
	fmt.Println(stringSlice)
	//using range loop to demonstrate individual slice
	for i, r := range stringSlice {
		fmt.Println(i, r)
	}
}	