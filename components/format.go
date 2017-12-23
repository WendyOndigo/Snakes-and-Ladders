package components

import "strconv"

func Print(game Game) string {
	output := ""
	printRows(game, &output, game.BoardHeight, game.MaxCell)
	printHorizontalEdge(game.BoardWidth, &output)
	printWinner(game, &output)
	return output
}

// Print a horizontal edge of +--- for the given width of the Board
func printHorizontalEdge(width int, outputString *string) {
	for i := 0; i < width; i++ {
		*outputString += "+---"
	}
	*outputString += "+\n"
}

// Print all rows of the Board
func printRows(game Game, outputString *string, rowNumber int, label int) {
	currLabel := label
	for i := rowNumber; i >= 1; i-- {
		make_row(game, i, currLabel, outputString)
		currLabel -= game.BoardWidth
	}
}

// Prints an entire row of cells including cell number, player on
// the cell, powerup on the cell, and snake or ladder on the cell
func make_row(game Game, rowNumber int, startLabel int, outputString *string) {
	cellNumbers := getRowLabel(game.BoardWidth, rowNumber, startLabel)
	printHorizontalEdge(game.BoardWidth, outputString)
	printCellNumbers(cellNumbers, outputString)
	printCellSpecialties(game, cellNumbers, outputString)
}

// Get the cell number for a given row in order from left to right
func getRowLabel(totalCol int, rowNumber int, startNumber int) []int {
	labels := []int{}
	rightToLeft := rowNumber%2 == 0
	for i := 0; i < totalCol; i++ {
		if rightToLeft {
			labels = append(labels, startNumber-i)
		} else {
			labels = append([]int{startNumber - i}, labels...)
		}
	}
	return labels
}

// Prints the top half of the cell
func printCellNumbers(workingColumns []int, outputString *string) {
	for _, cellNumber := range workingColumns {
		label := strconv.Itoa(cellNumber)
		*outputString += "|"
		printSpacing(len(label), outputString)
		*outputString += label
	}
	*outputString += "|\n"
}

// Prints the bottom half of the cell
func printCellSpecialties(game Game, workingColumns []int, outputString *string) {
	for _, col := range workingColumns {
		cell := game.Board[col-1]
		*outputString += "|"
		printPlayer(cell, outputString)
		printPowerup(cell, outputString)
		printTranslate(cell, outputString)
	}
	*outputString += "|\n"
}

// Print the player if a player exists for the cell
func printPlayer(cell Cell, outputString *string) {
	if cell.Player == "" {
		*outputString += " "
	} else {
		*outputString += cell.Player
	}
}

// Print the powerup if a powerup exists for the cell
func printPowerup(cell Cell, outputString *string) {
	if cell.Powerup == "" {
		*outputString += " "
	} else {
		*outputString += cell.Powerup
	}
}

// Print the ladder or snake if one exists for the cell
func printTranslate(cell Cell, outputString *string) {
	if cell.IsLadder() {
		*outputString += "L"
	} else if cell.IsSnake() {
		*outputString += "S"
	} else {
		*outputString += " "
	}
}

// Announce the winner if one exists
func printWinner(game Game, outputString *string) {
	playerName := game.winner()
	if playerName != "" {
		*outputString += "Player " + playerName + " won\n"
	}
}

// Prints the preceeding spaces for the cell number to make
// the cell number right justified
func printSpacing(num int, outputString *string) {
	for i := num; i < 3; i++ {
		*outputString += " "
	}
}
