package components

import "github.com/deckarep/golang-set"

type Player struct {
	Name     string
	Powers   mapset.Set
	Location int
}

func (player Player) HasPowerup(powerup string) bool {
	return player.Powers.Contains(powerup)
}

func (player Player) ObtainPowerup(powerup string) {
	player.Powers.Add(powerup)
}

func (player Player) UsePowerup(powerup string) {
	player.Powers.Remove(powerup)
}

func (player *Player) MoveTo(newLoc int) int {
	oldLoc := player.Location
	player.Location = newLoc
	return oldLoc
}

func NewPlayer(newName string) Player {
	return Player{
		Name:   newName,
		Powers: mapset.NewSet(),
	}
}
