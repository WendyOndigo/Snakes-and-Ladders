package slgame

import "strconv"

type Dice struct {
	Position int
	Faces    []int
}

func (dice *Dice) Roll() int {
	rolledFace := dice.Faces[dice.Position]
	dice.Position = (dice.Position + 1) % len(dice.Faces)
	return rolledFace
}

func NewDice(nums ...string) Dice {
	var newFaces []int

	for _, strFace := range nums {
		intFace, err := strconv.Atoi(strFace)
		if err != nil {
			panic(err)
		}
		newFaces = append(newFaces, intFace)

	}

	return Dice{
		Position: 0,
		Faces:    newFaces,
	}
}
