# ms-object-minesweeper
This microservice is designed to handle the minesweeper game behaviour.
It has 2 endpoints, first one to setup a new board game.
The second one is used to process the cell clicked and validates if the game ends or keeps running.

##Install Process

For make this project to run you must have the GOA package installed.

1. Run `go get -u github.com/goadesign/goa/...`
2. Run `make` command on the project's root folder.
3. Run `make run` for staritng the application.

##New Board Setup
For setting up a new board game you must invoke the endpoint GET `/minesweeper/new-game?height=3&width=3&mines=2` where:
* Height sets the height of the board game.
* Width sets the width of the board game.
* Mines sets the amaunt of mines in the board game.

##Clicked Cell Endpoint
When user clicks on a cell then must invoke the endpoint POST `minesweeper/clicked-cell/:row/:cell` where:
* Row: is the row where the cell was clicked.
* Cell: is the cell where user clicked.
* Also the board with the previous move is expected, i.e.
{
  "grid": [
    [
      {
        "clicked": false,
        "mine": false,
        "value": 0
      },
      {
        "clicked": false,
        "mine": false,
        "value": 1
      },
      {
        "clicked": false,
        "mine": false,
        "value": 1
      }
    ],
    [
      {
        "clicked": false,
        "mine": false,
        "value": 1
      },
      {
        "clicked": false,
        "mine": false,
        "value": 2
      },
      {
        "clicked": false,
        "mine": true,
        "value": 0
      }
    ],
    [
      {
        "clicked": false,
        "mine": true,
        "value": 0
      },
      {
        "clicked": true,
        "mine": false,
        "value": 2
      },
      {
        "clicked": false,
        "mine": false,
        "value": 1
      }
    ]
  ],
  "height": 3,
  "mineNum": 2,
  "width": 3
}
