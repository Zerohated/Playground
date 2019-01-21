package main

import (
	"bufio"
	"davidchou93/playground/ballcollision"
	"davidchou93/playground/lifegame"
	"fmt"
	"os"

	"strconv"
)

func main() {
	fmt.Println("Input the number to launch certain gadget:")
	fmt.Println("[1]: Ball collision")
	fmt.Println("[2]: Conway's life game")
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	index, err := strconv.Atoi(string(input[0]))
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	switch index {
	case 1:
		ballcollision.Run()
	case 2:
		lifegame.Run()
	}

}
