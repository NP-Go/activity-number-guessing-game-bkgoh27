package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// This code creates random number and place as variable
	rand.Seed(time.Now().UnixNano())
	hiddenNumber := int64(rand.Intn(100))

	var number int64
	fmt.Println("Guess a number from 1 to 100")
	fmt.Scanln(&number)

	if number > 100 {
		fmt.Println("Game over because the secret number is", hiddenNumber)
	}

	if number == hiddenNumber {
		fmt.Println("Nice")
	} else {
		fmt.Println("Wrong, the number is", hiddenNumber)
	}

}
