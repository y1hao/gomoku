package game

const Size = 15
const (
	Empty = iota
	Black
	White
)

type Board [Size][Size]uint8
type Player uint8

type Piece struct {
	Row    int    `json:"row"`
	Col    int    `json:"col"`
	Player Player `json:"player"`
}

type Game struct {
	Board         Board
	Player        Player
	WinningPieces []Piece
}

func New() *Game {
	board := [Size][Size]uint8{}
	for i := range board {
		board[i] = [Size]uint8{}
	}
	return &Game{
		Board: board,
	}
}

func (g *Game) Move(i, j int) (Player, error) {
	return 0, nil
}
