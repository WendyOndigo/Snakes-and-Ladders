package components

import (
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
		//fmt.Println("Command:", command)
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
	case "ladder":
		game.makeLadder(args[1], args[2])
	case "snake":
		game.makeSnake(args[1], args[2])
	case "turns":
		game.makeTurn(args[1])
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
		game.movePlayer(name, "1")
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

func (game *Game) makeLadder(startCell string, endCell string) {
	startIndex, sErr := strconv.Atoi(startCell)
	check(sErr)
	game.Board[startIndex-1].AddLadderTo(endCell)
}

func (game *Game) makeSnake(startCell string, endCell string) {
	startIndex, sErr := strconv.Atoi(startCell)
	check(sErr)
	game.Board[startIndex-1].AddSnakeTo(endCell)
}

func clearPlayer(cellLoc int, game *Game) {
	if cellLoc != 0 {
		game.Board[cellLoc-1].RemovePlayer()
	}
}

func setPlayer(cellLoc int, game *Game, playerName string) {
	if cellLoc != 0 {
		game.Board[cellLoc-1].SetPlayer(playerName)
	}
}

func updateCellPlayer(newCellLoc string, playerName string, game *Game) {
	if newCellLoc != "" {
		player := game.Players[playerName]
		newLoc, err := strconv.Atoi(newCellLoc)
		check(err)
		oldLoc := player.MoveTo(newLoc)
		clearPlayer(oldLoc, game)
		setPlayer(newLoc, game, playerName)
		game.Players[playerName] = player
	}
}

func (game Game) winner() string {
	lastCell := game.Board[game.MaxCell-1]
	return lastCell.Player
}

func bumpPlayer(cellLoc int, cell Cell, game *Game) {
	if cell.Player != "" {
		game.movePlayer(cell.Player, strconv.Itoa(cellLoc+1))
	}
}

func pickUpPowerup(cellLoc int, playerName string, game *Game) {
	cell := game.Board[cellLoc-1]
	if cell.Powerup != "" {
		player := game.Players[playerName]
		player.ObtainPowerup(cell.Powerup)
		game.Players[playerName] = player
	}
}

func teleportPlayer(cellLoc int, playerName string, game *Game) {
	cell := game.Board[cellLoc-1]
	if cell.IsLadder() {
		difference := cell.TransportTo() - cellLoc
		player := game.Players[playerName]
		if player.HasPowerup("e") {
			player.UsePowerup("e")
			game.Players[playerName] = player
			difference += difference
		}
		if finalLoc := difference + cellLoc; finalLoc >= game.MaxCell {
			game.movePlayer(playerName, strconv.Itoa(game.MaxCell))
		} else {
			game.movePlayer(playerName, strconv.Itoa(finalLoc))
		}
	} else if cell.IsSnake() {
		player := game.Players[playerName]
		if player.HasPowerup("a") {
			player.UsePowerup("a")
			game.Players[playerName] = player
			updateCellPlayer(strconv.Itoa(cellLoc), playerName, game)
		} else {
			game.movePlayer(playerName, strconv.Itoa(cell.TransportTo()))
		}
	} else {
		updateCellPlayer(strconv.Itoa(cellLoc), playerName, game)
	}
}

func (game *Game) movePlayer(playerName string, newCellLoc string) {
	index, err := strconv.Atoi(newCellLoc)
	check(err)
	destCell := game.Board[index-1]
	//fmt.Printf("moving player %v to %v\n", playerName, newCellLoc)
	updateCellPlayer(newCellLoc, playerName, game)

	if newCellLoc != strconv.Itoa(game.MaxCell) {
		// Bump player if necessary
		bumpPlayer(index, destCell, game)
		// Pick up powerup if one is available
		pickUpPowerup(index, playerName, game)
		// Handle snakes and ladders as appropriate
		teleportPlayer(index, playerName, game)
	}
}

func (game Game) AllPlayerNames() []string {
	names := []string{}
	for i := 0; i < len(game.Players); i++ {
		names = append(names, string(65+i))
	}
	//fmt.Println("names:", names)
	return names
}

func takeTurn(turnNumber int, pendingPlayers []string, game *Game) {
	if len(pendingPlayers) == 0 {
		takeTurn(turnNumber-1, game.AllPlayerNames(), game)
	} else if turnNumber > 0 {
		if game.winner() == "" {
			player := game.Players[pendingPlayers[0]]
			moves := game.Dice.Roll()
			newLoc := player.Location + moves
			if player.HasPowerup("d") {
				player.UsePowerup("d")
				game.Players[pendingPlayers[0]] = player
				newLoc += moves
			}
			if newLoc <= game.MaxCell {
				game.movePlayer(pendingPlayers[0], strconv.Itoa(newLoc))
			}
			takeTurn(turnNumber, pendingPlayers[1:], game)
		}
	}
}

func (game *Game) makeTurn(totalTurns string) {
	numberOfTurns, err := strconv.Atoi(totalTurns)
	check(err)
	takeTurn(numberOfTurns, game.AllPlayerNames(), game)
}
