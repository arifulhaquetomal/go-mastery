package main

import "fmt"

func main() {
	mp := map[int][]int{1:{1,2,3,4}, 2:{4,5,4,3,}}
	fmt.Println(mp[1])

	mp2 := map[string][]int{"a":{1,2,3}, "b":{4,5,6}}
	fmt.Println(mp2["a"])
	mp2["a"] = []int{1,2,34,5}
	fmt.Println(mp2["a"])
	delete(mp2, "a")
	fmt.Println(mp2)
	fmt.Println(ifExists("d"))
}

func ifExists(k string) ([]int, bool){
	dict := map[string][]int{"a":{1,2,3,}, "b":{2,3,4,5}}
	value, ok := dict[k]
	return value, ok
}