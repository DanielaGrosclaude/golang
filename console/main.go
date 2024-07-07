package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappuccino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocca"
	coffees[5] = "Macchiato"
	coffees[6] = "Quit"

	fmt.Println("MENU:")
	fmt.Println("-----")
	fmt.Println("1- Cappucino")
	fmt.Println("2- Latte")
	fmt.Println("3- Americano")
	fmt.Println("4- Mocca")
	fmt.Println("5- Macchiato")
	fmt.Println("Q- Quit the program")

	fmt.Println("Press any key on the keyboard. Press ESC to quit.")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You chose %s", coffees[i]))

	}
	fmt.Println("Program exiting...")
}
