package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	var guess int
	correctNumber := random.Intn(100) + 1 // The number to guess

	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("Guess a number between 1 and 100:")

	for {
		fmt.Println("Try and guess the number!")
		fmt.Scan(&guess)
		if guess == correctNumber {
			fmt.Println("Congratulations! You guessed the number.")
			break
		} else if guess < correctNumber {
			fmt.Println("Too low! Try again.")
		} else {
			fmt.Println("Too high! Try again.")
		}
	}
	fmt.Printf("Execution Time: %s\n", time.Since(start))
}
