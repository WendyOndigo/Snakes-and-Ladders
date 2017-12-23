package components

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCell(t *testing.T) {
	cell := NewCell(42)
	expectedCell := Cell{
		Label:   42,
		Player:  "",
		Powerup: "",
		Snake:   0,
		Ladder:  0,
	}
	assert.Equal(t, cell, expectedCell, "New cells should have just a label")
}

func TestPowerupCellWithPlayer(t *testing.T) {
	cell := NewCell(42)
	cell.AddPowerup("d")
	cell.SetPlayer("A")
	expectedCell := Cell{
		Label:   42,
		Player:  "A",
		Powerup: "d",
		Snake:   0,
		Ladder:  0,
	}
	assert.Equal(t, cell, expectedCell, "This cells should have a player and powerup")
}

func TestSnakeCell(t *testing.T) {
	cell := NewCell(42)
	cell.AddSnakeTo("5")
	assert.Equal(t, cell.IsLadder(), false, "This cells is not a ladder")
	assert.Equal(t, cell.IsSnake(), true, "This cells is a snake")
	assert.Equal(t, cell.TransportTo(), 5, "This cell should transport to cell 5")
}

func TestLadderCell(t *testing.T) {
	cell := NewCell(42)
	cell.AddLadderTo("50")
	assert.Equal(t, cell.IsLadder(), true, "This cells is not a ladder")
	assert.Equal(t, cell.IsSnake(), false, "This cells is a snake")
	assert.Equal(t, cell.TransportTo(), 50, "This cell should transport to cell 50")
}

func TestRemovePlayerFromCell(t *testing.T) {
	cell := NewCell(42)
	cell.SetPlayer("A")
	assert.Equal(t, cell.Player, "A", "This cells should have a player")
	cell.RemovePlayer()
	assert.Equal(t, cell.Player, "", "This cells should not have a player")
}
