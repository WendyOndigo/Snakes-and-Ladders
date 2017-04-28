package slgame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	input := `board 4 4
players 3
dice 3 2 2
powerup antivenom 7
ladder 3 7
snake 9 1
powerup double 4
powerup escalator 6 8
ladder 10 15
turns 3`
	output := `+---+---+---+---+
| 16| 15| 14| 13|
|B  |   |   |   |
+---+---+---+---+
|  9| 10| 11| 12|
|C S|  L|   |   |
+---+---+---+---+
|  8|  7|  6|  5|
| e | a | e |   |
+---+---+---+---+
|  1|  2|  3|  4|
|   |   |  L|Ad |
+---+---+---+---+
Player B won
`
	assert.Equal(t, Print(ReadFrom(input)), output, "New players should have just a name")
}
