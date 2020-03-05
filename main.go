package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func newGame() [][]string {
	// create main slice of slices
	field := make([][]string, 5)
	// variable for set numbers on game field
	n := 1
	// iterate over every element in slice
	for i := 0; i < len(field); i++ {
		// create underlying slice
		field[i] = make([]string, 5)
		// if i is odd (eg. 1,3,5,7...etc) fill slice with separators - horizontal line
		if i%2 != 0 {
			for j := 0; j < len(field[i]); j++ {
				field[i][j] = "—"
			}
		} else {
			// if j is odd (eg. 1,3,5,7...etc) fill element of slice with separators - vertical lines
			for j := 0; j < len(field[i]); j++ {
				if j%2 != 0 {
					field[i][j] = "|"
				} else {
					field[i][j] = fmt.Sprint(n)
					n++
				}
			}
		}
	}

	return field
}

func printField(field [][]string) {
	// iterate over every element of slice and print them on newlines
	for _, element := range field {
		fmt.Println(element)
	}
}

func setSign() (p1 string, p2 string, err error) {
	// ask player one for input, and scan it
	fmt.Printf("Player one - choose your sign: X or O?\n")
	fmt.Fscan(os.Stdin, &p1)
	// check input - if it's incorrect throw error
	switch {
	case p1 == "X" || p1 == "x":
		p1, p2 = "X", "O"
	case p1 == "O" || p1 == "o":
		p1, p2 = "O", "X"
	default:
		err = errors.New("incorrect input - please use only english letters o or x")
	}

	return p1, p2, err
}

func makeTurn(field [][]string, playerSign string, position string) {
	// iterate over every element of parent slice
	for i := 0; i < len(field); i++ {
		// iterate over every element of children slice
		for j := 0; j < len(field[i]); j++ {
			// check if value of input equal with value of element in slice
			if field[i][j] == position {
				// set element equal to player's sign
				field[i][j] = playerSign
				break
			}
		}
	}
}

func gameFinished(field [][]string, p1 string, p2 string) (result bool) {
	// checking loop
	for i := 0; i < len(field); i++ {
		// check only even elements of slice
		// set horizontal and vertical to none
		horizontal, vertical, diagonal1, diagonal2 := "", "", "", ""
		n := 0
		if i%2 == 0 {
			// check directions
			for j := 0; j < len(field); j++ {
				if j%2 == 0 {
					vertical = vertical + field[j][i]
					horizontal = horizontal + field[i][j]
					diagonal1 = diagonal1 + field[j][j]
				}
			}
			// check diagonal from left bottom to upper right
			for j := len(field) - 1; j >= 0; j-- {
				if j%2 == 0 {
					diagonal2 = diagonal2 + field[n][j]
					n += 2
				}
			}

			if vertical == strings.Repeat(p1, 3) || horizontal == strings.Repeat(p1, 3) || diagonal1 == strings.Repeat(p1, 3) || diagonal2 == strings.Repeat(p1, 3) {
				result = true
				fmt.Printf("Player One won!\n")
				break
			} else if vertical == strings.Repeat(p2, 3) || horizontal == strings.Repeat(p2, 3) || diagonal1 == strings.Repeat(p2, 3) || diagonal2 == strings.Repeat(p2, 3) {
				result = true
				fmt.Printf("Player Two won!\n")
				break
			} else {
				continue
			}

		}
	}

	return result

}

func main() {
	// counter for turns
	var turnCount int = 0
	// number of cell, which will be checked
	var turnNumber string
	// holder for loop, while game isn't finished
	var isFinished bool = false
	// Welcome message!
	fmt.Printf("Welcome to tic-tac-toe game!\n")
	// init new game field
	game := newGame()
	// set sign for players
	playerOne, playerTwo, err := setSign()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		// print signs of players
		fmt.Printf("Player One chosed %v, Player Two will use %v\n", playerOne, playerTwo)
	}
	// print game field for the first time
	printField(game)
	// main game loop
	for isFinished == false {
		// increase turn count
		turnCount++
		// print turn number and signs of palyers
		fmt.Printf("___________\nTurn №%v, (P1 = %v | P2 = %v)\n", turnCount, playerOne, playerTwo)
		// logic when player one choosed X
		if playerOne == "X" {
			if turnCount%2 != 0 {
				// player one always starts first and if turnCount is odd - it is player's one turn
				fmt.Printf("Player One make your turn\n")
				fmt.Fscan(os.Stdin, &turnNumber)
				makeTurn(game, playerOne, turnNumber)
			} else {
				// player one always starts first and if turnCount is odd - it is player's one turn
				fmt.Printf("Player Two make your turn\n")
				fmt.Fscan(os.Stdin, &turnNumber)
				makeTurn(game, playerTwo, turnNumber)
			}
			// logic when player one choosed O
		} else {
			if turnCount%2 != 0 {
				// player one always starts first and if turnCount is odd - it is player's one turn
				fmt.Printf("Player Two make your turn\n")
				fmt.Fscan(os.Stdin, &turnNumber)
				makeTurn(game, playerTwo, turnNumber)
			} else {
				// player one always starts first and if turnCount is odd - it is player's one turn
				fmt.Printf("Player One make your turn\n")
				fmt.Fscan(os.Stdin, &turnNumber)
				makeTurn(game, playerOne, turnNumber)
			}
		}
		// check if game is already finished
		isFinished = gameFinished(game, playerOne, playerTwo)
		// print game field
		printField(game)
	}
	// bye message!
	fmt.Printf("Game finished in %v turns!\nBye-Bye!\n\nWindow will close automatically in 5 seconds!", turnCount)
	// sleep before close window
	time.Sleep(5 * time.Second)
}
