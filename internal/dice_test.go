package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleDice(t *testing.T) {
	dice := NewDice([]string{"1"})
	assert.Equal(t, dice, SLDice{Position: 0, Faces: []int{1}}, "Dice should have position 0 and single face")
}

func TestMultipleDice(t *testing.T) {
	dice := NewDice([]string{"1", "2", "2"})
	assert.Equal(t, dice, SLDice{Position: 0, Faces: []int{1, 2, 2}}, "Dice should have position 0 and mulitple faces")
}

func TestSingleRoll(t *testing.T) {
	dice := NewDice([]string{"1", "2", "2"})
	rolled := dice.Roll()
	assert.Equal(t, dice, SLDice{Position: 1, Faces: []int{1, 2, 2}}, "Dice should have position 1 and mulitple faces")
	assert.Equal(t, rolled, 1, "Should have rolled the first face")
}

func TestMultiRoll(t *testing.T) {
	dice := NewDice([]string{"1", "2", "2"})
	rolled := dice.Roll()
	rolled = dice.Roll()
	assert.Equal(t, dice, SLDice{Position: 2, Faces: []int{1, 2, 2}}, "Dice should have position 2 and mulitple faces")
	assert.Equal(t, rolled, 2, "Should have rolled the second face")
}

func TestModRoll(t *testing.T) {
	dice := NewDice([]string{"1", "2"})
	rolled := dice.Roll()
	rolled = dice.Roll()
	assert.Equal(t, dice, SLDice{Position: 0, Faces: []int{1, 2}}, "Dice position should have wrapped back to 0")
	assert.Equal(t, rolled, 2, "Should have rolled the second face")
}
