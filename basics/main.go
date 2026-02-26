package main

import "fmt"

func main() {
	var var1, var2 int

	var1 = 244
	var2 = 112
	result := var1 + var2 // one way to declare variable so that go auto infers the type

	fmt.Println(result)
	fmt.Println(return_string())
	guessGame(24, 5)

}

// defining a function
func return_string() string {
	return "hello"
}

func guessGame(secret int, limit int) bool {
	var guess int
	fmt.Println("----GUESSING GAME----")

	for i := 0; i <= limit-1; i++ {
		fmt.Printf("Guess the number. Attempts left: %d\n", limit-i) // here the %s is not going to do anything, instead we use Printf
		fmt.Scanln(&guess)
		if guess == secret {
			fmt.Println("You guessed right")
			return true
		} else {
			fmt.Println("Guess again")
			if guess < secret {
				fmt.Println("Your number is less than the secret")
			} else {
				fmt.Println("your number is more than the secret")
			}
		}
	}
	fmt.Println("you failed")
	return false
}
