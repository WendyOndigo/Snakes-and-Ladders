package main

import (
	"Snakes-and-Ladders/slgame"
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil && e != io.EOF {
		panic(e)
	}
}
func main() {

	reader := bufio.NewReader(os.Stdin)
	var input string
	config := ""

	input, err := reader.ReadString('\n')
	check(err)
	for err != io.EOF {
		config += input
		input, err = reader.ReadString('\n')
		check(err)
	}
	game := slgame.ReadFrom(config)
	//fmt.Println(game)
	fmt.Println(slgame.Print(game))
}
