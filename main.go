package main

import (
	"Snakes-and-Ladders/slgame"
	"fmt"
)

func main() {
	game := slgame.ReadFrom("board 3 4\nplayers 2\ndice 3 2 2\npowerup double 1 2 3\npowerup escalator 4 5\npowerup antivenom 10")
	fmt.Println(slgame.Print(game))

	/*
	       cell := slgame.NewCell(5)

	   	fmt.Println(cell)
	   	cell.AddPowerup("a")
	   	cell.AddSnakeTo("5")
	   	cell.SetPlayer("B")
	   	fmt.Println(cell)
	   	fmt.Printf("snake: %v, ladder: %v\n", cell.IsSnake(), cell.IsLadder())
	   	op := cell.RemovePlayer()
	   	fmt.Printf("op: %v, cell: %v\n", op, cell)

	   	player := slgame.NewPlayer("A")
	   	player.MoveTo(5)
	   	player.ObtainPowerup("d")
	   	player.ObtainPowerup("d")
	   	player.UsePowerup("p")
	   	fmt.Println(player)
	   	player.UsePowerup("d")
	   	fmt.Println(player)

	   		dice := slgame.NewDice("1", "2", "3")
	   		fmt.Println(dice)
	   		roll := dice.Roll()
	   		fmt.Printf("rolled: %v, updated dice: %v\n", roll, dice)
	   		roll = dice.Roll()
	   		fmt.Printf("rolled: %v, updated dice: %v\n", roll, dice)
	   		roll = dice.Roll()
	   		fmt.Printf("rolled: %v, updated dice: %v\n", roll, dice)*/
}
