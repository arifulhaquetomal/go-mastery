package main

import "fmt"

func main() {

	var arr1 [2]int
	var arr2 [3]int
	arr1 = [2]int{1,2}
	arr2 = [3]int{2,2,4}
	fmt.Println(arr1[0])
	var arr3 [][]int //slices that are not fixed, dynamic arrays
	var slice1, slice2 []int
	slice1 = []int{1,2,3,4,5}
	slice2 = []int{2,3,4,5,5,3}
	fmt.Println(slice1, "--slice--", slice2)
	arr3 = [][]int{{1,2,3,4,5,6},{1,2,3,4}} //slices
	arr4 := [...][2]int{{1,200},{1,2}, {2,3}} // [...][...]int asd illegal
	fmt.Println(arr1, arr2, arr3)
	fmt.Printf("%T", arr4)

}