# Snakes-and-Ladders
This is an upgraded snakes and ladders game. The functionality of this game is inspired from an Assignment from [Dr. Dave Mason](http://cps506.sarg.ryerson.ca/current/)

## How It Works
You are to supply a series of **commands/configurations** to setup and run the game. The program will run those commands and output a string representation of the state of the game board. You may save multiple sets of game configurations into a JSON formatted file and save each of those results inside another JSON formatted file.

### JSON Input
To load a set of configurations to the program, make sure to have a json file formatted as such:  
```json
{
    "author": "Your name",
    "configurations": [
        {
            "title": "Name of the configuration",
            "commands": "commands as a single string"
        }
    ]
}
```
For example,
```json
{
    "author": "John Smith",
    "configurations": [
        {
            "title": "Plain Board",
            "commands": "board 2 3"
        },
        {
            "title": "Complex Game",
            "commands": "board 4 4\nplayers 3\ndice 3 2 1\npowerup escalator 4\npowerup double 6\nsnake 12 1\nladder 8 10\nturns 1\npowerup antivenom 1\nsnake 10 2\nturns 3\npowerup antivenom 9\npowerup double 11\nturns 4\npowerup double 14\nturns 2"
        }
    ]
}
``` 

### JSON ouput
The outputted json file will have the following formatting:
```json
{
    "requestedBy": "Your Name",
    "resultSet": [
        {
            "title": "Name of the configuration",
            "result": "+---+---+\n|  5|  6|\n|   |   |\n+---+---+\n|  4|  3|\n|   |   |\n+---+---+\n|  1|  2|\n|   |   |\n+---+---+\n"
        }
    ]
}
```
For example,
```json
{
    "requestedBy": "John Smith",
    "resultSet": [
        {
            "title": "Plain Board",
            "result": "+---+---+\n|  5|  6|\n|   |   |\n+---+---+\n|  4|  3|\n|   |   |\n+---+---+\n|  1|  2|\n|   |   |\n+---+---+\n"
        },
        {
            "title": "Complex Game",
            "result": "+---+---+---+---+\n| 16| 15| 14| 13|\n|C  |A  | d |   |\n+---+---+---+---+\n|  9| 10| 11| 12|\n| a |  S| d |  S|\n+---+---+---+---+\n|  8|  7|  6|  5|\n|  L|   | d |   |\n+---+---+---+---+\n|  1|  2|  3|  4|\n| a |   |   |Be |\n+---+---+---+---+\nPlayer C won\n"
        }
    ]
}
```

## Commands/Configurations
A command is a keyword followed by one or more parameters separated by a single space.  
1. `board 3 4` command:  
  - Specifies the number of columns and rows for the board. The total number of cells cannot exceed 999.
  - The command above will produce a board with 3 columns and 4 rows.
2. `players 2` command:  
  - Specifies the number of players, who are named: A, B,... There can be up to 26 players.  
3. `dice 1 2 2` command:  
  - Specifies the sequence of dice rolls. The sequence will repeat indefinitely.
  - The above command will produce the sequence: 1,2,2,1,2,2,1,2,...
4. `ladder 5 11` command:
  - Creates a ladder that starts at the first number and ends at the second number.
  - If the players lands on the first cell, they are transported to land on the second cell.
5. `snake 8 4` command:
  - creates a snake that starts at the first number and ends at the second number.
  - If the players lands on the first cell, they are transported to land on the second cell.
6. `powerup type cells` command:
  - Describes a powerup that is applied to a series of cells. When a player lands on a powerup cell, they acquire that powerup and retain it until they use it. Using the powerup removes it from the player. A powerup cell can be triggered any number of times by any player, but they do not accumulate - a player either has the powerup or they don't.
  - `powerup escalator 6 9` sub-command:
    - Makes the next ladder cell a player steps onto twice as boosting - ie. they move twice as far up the board. If the boost takes them past the end of the board, they get moved to the last cell, and wins.
  - `powerup antivenom 7` sub-command:
    - Makes the player immune to the next snake cell they step onto - ie. they don't slide down the snake.
  - `powerup double 5` sub-command:
    - Doubles the next dice roll.
7. `turns 10` command:
  - Plays the specified number of turns (or until a player wins the game). A turn means each player, in order, rolls the dice, and then moves that many cells (or double if they have the powerup).
  - If the (possibly doubled) roll would take them past the end of the board, they don't move, and play proceeds to the next player.
  - As soon as a player wins, the gamne is over and the turns stop.
  
## Rules
1. The last cell is the winning cell. If a player lands on it, by any means, they win.
2. There can be any number of `snake`, `ladder`, or `powerup` commands. There can also be any number of `turn` commands, each of which will run in turn. Each `turn` commands runs on the current state of the board.
3. A given cell can only have a single special property: winning, snake start, ladder start, or powerup. Note that the end of a snake or ladder could have a special property.
4. A given cell can only have one player on it. When a plyer lands on a cell (including initial positioning), if there is already a player on the cell, that player gets bumped one cell ahead. When a bumped player lands on a cell, they get the action associated with that cell, including winning, powerups, snake, ladders, or bumping yet another player.
  
