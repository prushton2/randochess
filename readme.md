# Randochess

A chess game that implements weird rules to shake things up ingame. The game doesnt include check or checkmate, so winning requires taking the opponents king. En Passant is not implemented either. Some examples of rules include:

### Random
Choose a random rule to play

### Default
The base way to play the game

### Open World
The board is 16x16

### Oops! All Knights!
Every piece moves like a knight

### PREPARE THYSELF
On their first move, pawns can move or take up to 5 spaces infront of them

### Have a plan to kill everyone you meet
Bisops can move like a king, but no longer need line of sight to take other pieces. They cannot take cardinally, so be careful.

### Atomic Chess
When you take a piece, both pieces disappear

### Knook
Rooks can move as knights aswell as their default moveset

# Development
* Install npm and go
* Run the backend with `cd backend && go run .`
* Run the frontend with `cd frontend && npm run dev`

## Docker
copy `compose.yaml.example` to `compose.yaml` and modify as desired