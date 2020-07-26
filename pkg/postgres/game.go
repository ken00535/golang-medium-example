package postgres

// Game is game
type Game struct {
	ID       int
	WinnerID int
	Location string
	Prize    int
}

// JoinResult is join result
type JoinResult struct {
	Player
	Game
}
