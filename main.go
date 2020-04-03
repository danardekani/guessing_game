package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Guess a number between 1 and 100")

	createNumber := createRandomInteger(1, 100)
	fmt.Println("The random number is ", createNumber)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')	
	input = strings.TrimSuffix(input, "\n")

	guess, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Invalid input. Please enter an integer value")
		return
	}

	fmt.Println("Your guess is", guess)
}

func createRandomInteger(min, max int) int  {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}