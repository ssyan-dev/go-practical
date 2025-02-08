package main

import (
	"fmt"
	"math/rand"
)

type Level struct {
	name string
	to   int
}

var levels = [4]Level{
	{},
	{
		name: "easy",
		to:   25,
	},
	{
		name: "normal",
		to:   100,
	},
	{
		name: "hard",
		to:   250,
	},
}

func main() {
	fmt.Println("============================================")
	fmt.Println("Guess the number game")
	fmt.Println("============================================")

	fmt.Println("Levels:")
	for i := 1; i < len(levels); i++ {
		lvl := levels[i]
		fmt.Printf("    - %s (0-%d): %d\n", lvl.name, lvl.to, i)
	}

	var lvlInput int
	fmt.Print("Enter the lvl: ")
	fmt.Scan(&lvlInput)

	if lvlInput < 1 || lvlInput >= len(levels) {
		fmt.Println("incorrect number!")
		return
	}

	lvl := levels[lvlInput]
	fmt.Printf("Ok. Guess the number from 0 to %d\n", lvl.to)

	reqNum := rand.Intn(lvl.to - 0)
	isGuessed := false
	attempts := 0

	for !isGuessed {
		attempts++
		fmt.Print("Your choice is..: ")
		var numInput int
		fmt.Scan(&numInput)

		if numInput < 0 {
			fmt.Println("Number is bigger than 0!")
		}

		if numInput < reqNum {
			fmt.Printf("%d is less than required!\n", numInput)
		}

		if numInput > reqNum {
			fmt.Printf("%d is bigger than required!\n", numInput)
		}

		if numInput == reqNum {
			fmt.Printf("\n\nYou win in %d attempts! %d = %d!", attempts, numInput, reqNum)
			return
		}
	}
}
