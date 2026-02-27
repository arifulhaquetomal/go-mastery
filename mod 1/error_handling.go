package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(error_handling(2, 0))

	fmt.Println(error_handling(2, 1))
	
	// idiomatic expression
	result, err := error_handling(2, 0)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)

}
func error_handling(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return a / b, nil
}
