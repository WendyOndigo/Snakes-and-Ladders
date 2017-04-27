package slgame

import "strconv"

type Cell struct {
	Label   int
	Player  string
	Powerup string
	Snake   int
	Ladder  int
}

func (cell *Cell) AddPowerup(powerup string) {
	cell.Powerup = powerup
}

func (cell *Cell) AddSnakeTo(newLoc string) {
	num, err := strconv.Atoi(newLoc)
	if err != nil {
		panic(err)
	}
	cell.Snake = num
}

func (cell *Cell) AddLadderTo(newLoc string) {
	num, err := strconv.Atoi(newLoc)
	if err != nil {
		panic(err)
	}
	cell.Ladder = num
}

func (cell Cell) IsLadder() bool {
	return cell.Ladder != 0
}

func (cell Cell) IsSnake() bool {
	return cell.Snake != 0
}

func (cell Cell) TransportTo() int {
	if cell.IsLadder() {
		return cell.Ladder
	} else if cell.IsSnake() {
		return cell.Snake
	}
	return 0
}

func (cell *Cell) SetPlayer(name string) {
	cell.Player = name
}

func (cell *Cell) RemovePlayer() string {
	oldPlayer := cell.Player
	cell.Player = ""
	return oldPlayer
}

func NewCell(newLabel int) Cell {
	return Cell{
		Label: newLabel,
	}
}
