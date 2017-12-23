package internal

import (
	"testing"

	"github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
)

func TestNewPlayer(t *testing.T) {
	player := NewPlayer("A")
	expectedPlayer := Player{
		Name:     "A",
		Powers:   mapset.NewSet(),
		Location: 0,
	}
	assert.Equal(t, player, expectedPlayer, "New players should have just a name")
}

func TestPlayerPowerups(t *testing.T) {
	player := NewPlayer("A")
	assert.Equal(t, player.HasPowerup("d"), false, "New players should not have a double powerup")
	player.ObtainPowerup("d")
	assert.Equal(t, player.HasPowerup("d"), true, "player should have picked up a double powerup")
	player.UsePowerup("d")
	assert.Equal(t, player.HasPowerup("d"), false, "player should not have a double powerup")
}

func TestPlayerFullSetPowerups(t *testing.T) {
	player := NewPlayer("A")
	player.ObtainPowerup("d")
	player.ObtainPowerup("a")
	player.ObtainPowerup("e")
	assert.Equal(t, player.HasPowerup("d"), true, "player should have picked up a double powerup")
	assert.Equal(t, player.HasPowerup("a"), true, "player should have picked up an antivenom powerup")
	assert.Equal(t, player.HasPowerup("e"), true, "player should have picked up an escalator powerup")
}

func TestUpdatingPlayerLocation(t *testing.T) {
	player := NewPlayer("A")
	assert.Equal(t, player.Location, 0, "New players should have a default location of 0")
	player.MoveTo(42)
	assert.Equal(t, player.Location, 42, "New players should have moved to location 42")
}
