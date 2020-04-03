package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"	
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to the guessing game.")

	// return random number from createRandomInteger
	createNumber := createRandomInteger(1, 100)
	fmt.Println("the secret number is", createNumber)

	// accept and read user input
	var attempts int
	for {
		attempts++
		fmt.Println("Please guess a number between 1 and 100")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			return
		}
		// if/else statements comparing user input vs machine output
		fmt.Println("Your guess is", guess)
		if guess > createNumber {
			fmt.Println("You're guess is too high. Try again.")
		} else if guess < createNumber {
			fmt.Println("Your guess is too low. Try again. ")
		} else {
			fmt.Println("That is correct! Good job!")
			break
		}
	}
}
func createRandomInteger(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
