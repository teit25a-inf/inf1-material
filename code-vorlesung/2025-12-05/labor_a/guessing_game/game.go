package guessinggame

import (
	"fmt"
	"math/rand/v2"
)

func GuessingGame() {
	correct := rand.IntN(100) + 1
	for i := 0; i < 3; i++ {
		guess := ReadNumber()
		if NumberGood(guess, correct) {
			fmt.Println("Richtig geraten :-)")
			return
		}
	}
	fmt.Printf("Zu viele falsche Zahlen, die richtige Zahl war %d.", correct)
}
