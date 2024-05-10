package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
	"strings"
)

// len of user input minus len of comp hardcoded who wins
var userMinusComp = map[string][]int{
	"User won": {1, 2, -3},
	"Comp won": {-1, -2, 3},
}

func main() {
	fmt.Println("Welcome to Rock, Paper, Scissor! Please enter either of the three to play the game!")
	fmt.Println("What do you choose?")

	//read user input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	input = strings.ToLower(input)

	//validate user input
	validInputs := []string{"rock", "paper", "scissor"}

	if !slices.Contains(validInputs, input) {
		fmt.Println("your input is not valid!")
	}

	//let computer play
	randNumer := rand.IntN(len(validInputs))
	compChoice := validInputs[randNumer]

	fmt.Println("You've chosen", input, "and the computer", compChoice)

	diff := len(input) - len(compChoice)

	if diff == 0 {
		fmt.Println("You've tied with the compter")
	} else if slices.Contains(userMinusComp["User won"], diff) {
		fmt.Println("Congrats, you won!")
	} else if slices.Contains(userMinusComp["Comp won"], diff) {
		fmt.Println("You've lost :(")
	} else {
		fmt.Println("Ooops, sth went wrong!")
	}

}
