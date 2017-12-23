package internal

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

func TestSnakeLadderCombo(t *testing.T) {
	input := `board 3 3
players 1
dice 2
ladder 3 8
snake 8 1
turns 1
`
	output := `+---+---+---+
|  7|  8|  9|
|   |  S|   |
+---+---+---+
|  6|  5|  4|
|   |   |   |
+---+---+---+
|  1|  2|  3|
|A  |   |  L|
+---+---+---+
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "Should be a simple small board with snake and ladder")
	playerA := Player{
		Name:     "A",
		Powers:   mapset.NewSet(),
		Location: 1,
	}
	assert.Equal(t, playerA, game.Players["A"], "A should go up ladder and down snake to 1")
}

func TestShootForTheMoon(t *testing.T) {
	input := `board 2 3
powerup escalator 1
players 1
dice 1
ladder 2 5
turns 1
`
	output := `+---+---+
|  5|  6|
|   |A  |
+---+---+
|  4|  3|
|   |   |
+---+---+
|  1|  2|
| e |  L|
+---+---+
Player A won
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "Should be a simple small board with powerups and ladder")
	playerA := Player{
		Name:     "A",
		Powers:   mapset.NewSet(),
		Location: 6,
	}
	assert.Equal(t, playerA, game.Players["A"], "A should fly off the board, but lands on the last cell")
}

func TestBumpingAfterSnake(t *testing.T) {
	input := `board 3 4
players 2
dice 3 2
snake 8 3
turns 2
`
	output := `+---+---+---+
| 12| 11| 10|
|   |   |   |
+---+---+---+
|  7|  8|  9|
|   |  S|   |
+---+---+---+
|  6|  5|  4|
|B  |   |   |
+---+---+---+
|  1|  2|  3|
|   |   |A  |
+---+---+---+
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "B should be bumped after A slides down to it's spot")
}

func TestTurnsAfterBoardUpdate(t *testing.T) {
	input := `board 5 5
players 3
dice 1 2 2
ladder 5 12
ladder 12 22
powerup escalator 4
snake 22 1
turns 2
powerup antivenom 21
turns 3
`
	output := `+---+---+---+---+---+
| 21| 22| 23| 24| 25|
| a |  S|   |A  |B  |
+---+---+---+---+---+
| 20| 19| 18| 17| 16|
|   |   |   |   |   |
+---+---+---+---+---+
| 11| 12| 13| 14| 15|
|   |  L|   |   |   |
+---+---+---+---+---+
| 10|  9|  8|  7|  6|
|   |   |   |   |   |
+---+---+---+---+---+
|  1|  2|  3|  4|  5|
|C  |   |   | e |  L|
+---+---+---+---+---+
Player B won
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "B should win")
}

func TestEscapaingLadderAndTroublingDouble(t *testing.T) {
	input := `board 5 5
players 2
dice 2
powerup escalator 4
ladder 6 8
snake 10 2
ladder 5 8
turns 3
powerup antivenom 8
turns 3
ladder 15 23
powerup double 23
turns 3
`
	output := `+---+---+---+---+---+
| 21| 22| 23| 24| 25|
|   |   | d |   |B  |
+---+---+---+---+---+
| 20| 19| 18| 17| 16|
|   |   |   |   |   |
+---+---+---+---+---+
| 11| 12| 13| 14| 15|
|   |   |   |   |  L|
+---+---+---+---+---+
| 10|  9|  8|  7|  6|
|  S|   | a |   |  L|
+---+---+---+---+---+
|  1|  2|  3|  4|  5|
|   |   |   |Ae |  L|
+---+---+---+---+---+
Player B won
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "B should win")
}

func TestMassBumping(t *testing.T) {
	input := `board 5 5
players 23
dice 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 1 2
turns 1
`
	output := `+---+---+---+---+---+
| 21| 22| 23| 24| 25|
|E  |D  |C  |B  |A  |
+---+---+---+---+---+
| 20| 19| 18| 17| 16|
|F  |G  |H  |I  |J  |
+---+---+---+---+---+
| 11| 12| 13| 14| 15|
|O  |N  |M  |L  |K  |
+---+---+---+---+---+
| 10|  9|  8|  7|  6|
|P  |Q  |R  |S  |T  |
+---+---+---+---+---+
|  1|  2|  3|  4|  5|
|W  |   |   |V  |U  |
+---+---+---+---+---+
Player A won
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "A should win")
}

func TestCrazySnakesAndLadders(t *testing.T) {
	input := `board 5 5
players 2
ladder 3 14
ladder 5 15
snake 15 6
ladder 6 13
snake 14 7
ladder 7 12
snake 13 8
ladder 8 11
snake 12 9
ladder 9 21
snake 11 10
ladder 10 21
dice 3 2 2 4
powerup double 22 19 16
powerup antivenom 17
powerup escalator 18 4
turns 2
`
	output := `+---+---+---+---+---+
| 21| 22| 23| 24| 25|
|   |Ad |   |   |B  |
+---+---+---+---+---+
| 20| 19| 18| 17| 16|
|   | d | e | a | d |
+---+---+---+---+---+
| 11| 12| 13| 14| 15|
|  S|  S|  S|  S|  S|
+---+---+---+---+---+
| 10|  9|  8|  7|  6|
|  L|  L|  L|  L|  L|
+---+---+---+---+---+
|  1|  2|  3|  4|  5|
|   |   |  L| e |  L|
+---+---+---+---+---+
Player B won
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "B should win")
}

func TestComplexConfiguration(t *testing.T) {
	input := `board 4 4
players 3
dice 3 2 1
powerup escalator 4
powerup double 6
snake 12 1
ladder 8 10
turns 1
powerup antivenom 1
snake 10 2
turns 3
powerup antivenom 9
powerup double 11
turns 4
powerup double 14
turns 2
`
	output := `+---+---+---+---+
| 16| 15| 14| 13|
|C  |A  | d |   |
+---+---+---+---+
|  9| 10| 11| 12|
| a |  S| d |  S|
+---+---+---+---+
|  8|  7|  6|  5|
|  L|   | d |   |
+---+---+---+---+
|  1|  2|  3|  4|
| a |   |   |Be |
+---+---+---+---+
Player C won
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "C should win")
}

func TestSnakeSurvivor(t *testing.T) {
	input := `board 4 4
players 2
dice 3
powerup antivenom 5
ladder 8 14
snake 14 2
ladder 7 14
turns 4
`
	output := `+---+---+---+---+
| 16| 15| 14| 13|
|   |A  |B S|   |
+---+---+---+---+
|  9| 10| 11| 12|
|   |   |   |   |
+---+---+---+---+
|  8|  7|  6|  5|
|  L|  L|   | a |
+---+---+---+---+
|  1|  2|  3|  4|
|   |   |   |   |
+---+---+---+---+
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "No one should win")
}

func TestDoubleOvershoot(t *testing.T) {
	input := `board 4 4
players 2
dice 2 1
powerup double 14 4
ladder 10 14
turns 5
`
	output := `+---+---+---+---+
| 16| 15| 14| 13|
|A  |   | d |   |
+---+---+---+---+
|  9| 10| 11| 12|
|   |  L|   |   |
+---+---+---+---+
|  8|  7|  6|  5|
|   |   |B  |   |
+---+---+---+---+
|  1|  2|  3|  4|
|   |   |   | d |
+---+---+---+---+
Player A won
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "A should win")
}

func TestWinByDouble(t *testing.T) {
	input := `board 4 4
players 2
dice 2
powerup double 12
turns 6
`
	output := `+---+---+---+---+
| 16| 15| 14| 13|
|A  |   |   |   |
+---+---+---+---+
|  9| 10| 11| 12|
|   |   |B  | d |
+---+---+---+---+
|  8|  7|  6|  5|
|   |   |   |   |
+---+---+---+---+
|  1|  2|  3|  4|
|   |   |   |   |
+---+---+---+---+
Player A won
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "A should win")
}

func TestSnakeSurvivor2(t *testing.T) {
	input := `board 4 4
players 2
dice 1 2 2 2
ladder 3 13
snake 13 10
powerup antivenom 10
powerup double 12
snake 11 1
turns 4
`
	output := `+---+---+---+---+
| 16| 15| 14| 13|
|B  |   |A  |  S|
+---+---+---+---+
|  9| 10| 11| 12|
|   | a |  S| d |
+---+---+---+---+
|  8|  7|  6|  5|
|   |   |   |   |
+---+---+---+---+
|  1|  2|  3|  4|
|   |   |  L|   |
+---+---+---+---+
Player B won
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "B should win")
}

func TestDoubleMania(t *testing.T) {
	input := `board 4 4
players 2
dice 8 1 3 8 8
powerup double 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15
snake 12 8
ladder 3 6
turns 7
`
	output := `+---+---+---+---+
| 16| 15| 14| 13|
|A  | d | d | d |
+---+---+---+---+
|  9| 10| 11| 12|
| d | d | d | dS|
+---+---+---+---+
|  8|  7|  6|  5|
| d | d | d | d |
+---+---+---+---+
|  1|  2|  3|  4|
| d |Bd | dL| d |
+---+---+---+---+
Player A won
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "A should win")
}

func TestNoConsentWin(t *testing.T) {
	input := `board 4 4
players 2
dice 2 1 2 2
ladder 6 11
ladder 9 14
snake 11 8
powerup escalator 8
powerup double 2
turns 2
`
	output := `+---+---+---+---+
| 16| 15| 14| 13|
|A  |   |   |   |
+---+---+---+---+
|  9| 10| 11| 12|
|  L|   |  S|   |
+---+---+---+---+
|  8|  7|  6|  5|
|Be |   |  L|   |
+---+---+---+---+
|  1|  2|  3|  4|
|   | d |   |   |
+---+---+---+---+
Player A won
`
	game := ReadFrom(input)
	assert.Equal(t, output, Print(game), "A should win")
}
