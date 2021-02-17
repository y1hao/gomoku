package game

import "errors"

const Size = 15
const (
	Empty = iota
	Black
	White
)

type Player uint8
type Board [Size][Size]Player

type Piece struct {
	Row    int    `json:"row"`
	Col    int    `json:"col"`
	Player Player `json:"player"`
}

type Game struct {
	Board         Board   `json:"board"`
	Player        Player  `json:"player"`
	WinningPieces []Piece `json:"winningPieces,omitempty"`
}

func New() *Game {
	board := [Size][Size]Player{}
	for i := range board {
		board[i] = [Size]Player{}
	}
	return &Game{
		Board: board,
		Player: Black,
	}
}

func (g *Game) Move(p *Piece) error {
	r, c := p.Row, p.Col
	if r < 0 || r >= Size || c < 0 || c >= Size || g.Board[r][c] != 0 || p.Player != g.Player{
		return errors.New("invalid position")
	}
	player := g.Player
	if player == Black {
		g.Player = White
	} else {
		g.Player = Black
	}
	g.Board[r][c] = player
	g.calcWinning()
	return nil
}

func (g *Game) calcWinning() {

}
