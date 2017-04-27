package slgame

import (
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	Board       []Cell
	Players     map[string]Player
	Dice        SLDice
	BoardWidth  int
	BoardHeight int
	MaxCell     int
}

func NewGame() Game {
	return Game{
		Board:   make([]Cell, 0, 0),
		Players: make(map[string]Player),
		Dice:    NewDice([]string{"1"}),
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func ReadFrom(input string) Game {
	game := NewGame()
	for _, command := range strings.Split(input, "\n") {
		fmt.Println("Command:", command)
		doCommand(command, &game)
	}
	return game
}

func doCommand(cmd string, game *Game) {
	args := strings.Split(cmd, " ")
	switch args[0] {
	case "board":
		game.makeBoard(args[1], args[2])
	case "players":
		game.makePlayers(args[1])
	case "dice":
		game.makeDice(args[1:])
	case "powerup":
		game.makePowerups(args[1], args[2:])
	default:
	}
}

func (game *Game) makeBoard(col string, row string) {
	width, err := strconv.Atoi(col)
	check(err)
	height, err := strconv.Atoi(row)
	check(err)
	game.BoardWidth = width
	game.BoardHeight = height
	game.MaxCell = width * height

	for i := 1; i <= game.MaxCell; i++ {
		game.Board = append(game.Board, NewCell(i))
	}
}

func (game *Game) makePlayers(num string) {
	numberOfPlayers, err := strconv.Atoi(num)
	check(err)
	for i := 0; i < numberOfPlayers; i++ {
		name := string(65 + i)
		player := NewPlayer(name)
		game.Players[name] = player
	}
}

func (game *Game) makeDice(faces []string) {
	game.Dice = NewDice(faces)
}

func (game *Game) makePowerups(powerupType string, affectedCells []string) {
	var powerupLabel string
	switch powerupType {
	case "double":
		powerupLabel = "d"
	case "antivenom":
		powerupLabel = "a"
	case "escalator":
		powerupLabel = "e"
	default:
		panic("Unknown powerup was inputted")
	}
	game.addPowerup(powerupLabel, affectedCells)
}

func (game *Game) addPowerup(label string, cells []string) {
	for _, cellNumber := range cells {
		index, err := strconv.Atoi(cellNumber)
		check(err)
		game.Board[index-1].AddPowerup(label)
	}
}

func (game Game) winner() string {
	lastCell := game.Board[game.MaxCell-1]
	return lastCell.Player
}
