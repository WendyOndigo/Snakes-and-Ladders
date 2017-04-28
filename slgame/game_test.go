package slgame

import (
	"testing"

	"github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
)

func TestAllFunctions(t *testing.T) {
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
	assert.Equal(t, output, Print(ReadFrom(input)), "Player B should Win")
}

func TestSmallBoard(t *testing.T) {
	input := `board 3 2`
	output := `+---+---+---+
|  6|  5|  4|
|   |   |   |
+---+---+---+
|  1|  2|  3|
|   |   |   |
+---+---+---+
`
	assert.Equal(t, output, Print(ReadFrom(input)), "Should be a simple small board")
}

func TestFullSmallBoard(t *testing.T) {
	input := `board 3 2
players 6
`
	output := `+---+---+---+
|  6|  5|  4|
|A  |B  |C  |
+---+---+---+
|  1|  2|  3|
|F  |E  |D  |
+---+---+---+
Player A won
`
	assert.Equal(t, output, Print(ReadFrom(input)), "Should be a fully occupied board")
}

func TestInitialPlayers(t *testing.T) {
	input := `board 3 2
players 2
`
	output := `+---+---+---+
|  6|  5|  4|
|   |   |   |
+---+---+---+
|  1|  2|  3|
|B  |A  |   |
+---+---+---+
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "Should be a simple small board with 2 players")
	playerA := Player{
		Name:     "A",
		Powers:   mapset.NewSet(),
		Location: 2,
	}
	assert.Equal(t, playerA, game.Players["A"], "Player A stats")
	playerB := Player{
		Name:     "B",
		Powers:   mapset.NewSet(),
		Location: 1,
	}
	assert.Equal(t, playerB, game.Players["B"], "Player B stats")
}

func TestTurns(t *testing.T) {
	input := `board 3 3
players 2
dice 1
turns 4`
	output := `+---+---+---+
|  7|  8|  9|
|   |   |   |
+---+---+---+
|  6|  5|  4|
|A  |B  |   |
+---+---+---+
|  1|  2|  3|
|   |   |   |
+---+---+---+
`
	assert.Equal(t, output, Print(ReadFrom(input)), "Each turn should make B follow A")
}

func TestMultiFaceDice(t *testing.T) {
	input := `board 3 4
players 2
dice 1 2
turns 5`
	output := `+---+---+---+
| 12| 11| 10|
|A  |B  |   |
+---+---+---+
|  7|  8|  9|
|   |   |   |
+---+---+---+
|  6|  5|  4|
|   |   |   |
+---+---+---+
|  1|  2|  3|
|   |   |   |
+---+---+---+
Player A won
`
	assert.Equal(t, output, Print(ReadFrom(input)), "Each turn should bump A")
}

func TestLongTurns(t *testing.T) {
	input := `board 3 4
players 2
dice 1 2 2 2 2
ladder 5 11
snake 8 4
powerup escalator 6 9
powerup antivenom 7
powerup double 4
turns 10`
	output := `+---+---+---+
| 12| 11| 10|
|B  |   |   |
+---+---+---+
|  7|  8|  9|
| a |  S| e |
+---+---+---+
|  6|  5|  4|
| e |  L|Ad |
+---+---+---+
|  1|  2|  3|
|   |   |   |
+---+---+---+
Player B won
`
	assert.Equal(t, output, Print(ReadFrom(input)), "Player B wins no longer than 10 turns")
}

func TestRollingOverflow(t *testing.T) {
	input := `board 3 4
players 2
dice 1 2 3
ladder 3 9
turns 4`
	output := `+---+---+---+
| 12| 11| 10|
|A  |   |B  |
+---+---+---+
|  7|  8|  9|
|   |   |   |
+---+---+---+
|  6|  5|  4|
|   |   |   |
+---+---+---+
|  1|  2|  3|
|   |   |  L|
+---+---+---+
Player A won
`
	assert.Equal(t, output, Print(ReadFrom(input)), "Player B should not be able to win")
}

func TestLadderAndPickup(t *testing.T) {
	input := `board 3 2
players 1
dice 1
ladder 2 5
powerup double 5
turns 1
`
	output := `+---+---+---+
|  6|  5|  4|
|   |Ad |   |
+---+---+---+
|  1|  2|  3|
|   |  L|   |
+---+---+---+
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "Should be a simple small board with 1 player")
	playerA := Player{
		Name:     "A",
		Powers:   mapset.NewSetWith("d"),
		Location: 5,
	}
	assert.Equal(t, playerA, game.Players["A"], "A should move up ladder and pick up a double")
}

func TestSnakeAndPickup(t *testing.T) {
	input := `board 3 2
players 1
dice 2
snake 5 1
powerup antivenom 1
turns 2
`
	output := `+---+---+---+
|  6|  5|  4|
|   |  S|   |
+---+---+---+
|  1|  2|  3|
|Aa |   |   |
+---+---+---+
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "Should be a simple small board with 1 player")
	playerA := Player{
		Name:     "A",
		Powers:   mapset.NewSetWith("a"),
		Location: 1,
	}
	assert.Equal(t, playerA, game.Players["A"], "A should slide down a snake and pick up an antivenom")
}

func TestDoubleLadder(t *testing.T) {
	input := `board 3 3
players 1
dice 2
ladder 3 6
ladder 6 9
turns 2
`
	output := `+---+---+---+
|  7|  8|  9|
|   |   |A  |
+---+---+---+
|  6|  5|  4|
|  L|   |   |
+---+---+---+
|  1|  2|  3|
|   |   |  L|
+---+---+---+
Player A won
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "Should be a simple small board with 2 ladders")
	playerA := Player{
		Name:     "A",
		Powers:   mapset.NewSet(),
		Location: 9,
	}
	assert.Equal(t, playerA, game.Players["A"], "A should move up 2 ladders and win")
}

func TestDoubleSnake(t *testing.T) {
	input := `board 3 3
players 1
dice 2 3
snake 4 1
snake 6 4
turns 2
`
	output := `+---+---+---+
|  7|  8|  9|
|   |   |   |
+---+---+---+
|  6|  5|  4|
|  S|   |  S|
+---+---+---+
|  1|  2|  3|
|A  |   |   |
+---+---+---+
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "Should be a simple small board with 2 snakes")
	playerA := Player{
		Name:     "A",
		Powers:   mapset.NewSet(),
		Location: 1,
	}
	assert.Equal(t, playerA, game.Players["A"], "A should slide down 2 snakes")
}

func TestUsingAntivenom(t *testing.T) {
	input := `board 3 3
players 1
dice 2 3
powerup antivenom 3
snake 6 4
turns 2
`
	output := `+---+---+---+
|  7|  8|  9|
|   |   |   |
+---+---+---+
|  6|  5|  4|
|A S|   |   |
+---+---+---+
|  1|  2|  3|
|   |   | a |
+---+---+---+
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "Should be a simple small board with snake and antivenom")
	playerA := Player{
		Name:     "A",
		Powers:   mapset.NewSet(),
		Location: 6,
	}
	assert.Equal(t, playerA, game.Players["A"], "A should not slide down snakes")
}

func TestUsingEscalator(t *testing.T) {
	input := `board 3 3
players 1
dice 1
powerup escalator 2
powerup double 7
ladder 3 5
turns 2
`
	output := `+---+---+---+
|  7|  8|  9|
|Ad |   |   |
+---+---+---+
|  6|  5|  4|
|   |   |   |
+---+---+---+
|  1|  2|  3|
|   | e |  L|
+---+---+---+
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "Should be a simple small board with powerups and ladder")
	playerA := Player{
		Name:     "A",
		Powers:   mapset.NewSetWith("d"),
		Location: 7,
	}
	assert.Equal(t, playerA, game.Players["A"], "A should go twice the ladder distance and pick up double")
}
