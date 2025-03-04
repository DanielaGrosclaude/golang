package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

/*b
->one way - declare, then assign (two steps)
var firstNumber int
firstNumber = 2
-> another way, declare type and name and assign value
var secondNumber = rand.Intn(8) + 2
-> one step variable:  declare name, assign value and let Go
subtraction := 7
*/
const prompt = "press ENTER when ready!."

func main() {
	rand.Seed(time.Now().UnixNano())
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction
	playTheGame(firstNumber, secondNumber, subtraction, answer)

}

func playTheGame(firstNumber, secondNumber, subtraction, answer int) {
	reader := bufio.NewReader(os.Stdin)

	//display a welcom/instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10 and don't type your number in, just", prompt)
	reader.ReadString('\n')

	//take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now substract ", subtraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	fmt.Println("The answer is:", answer, "!")

}
