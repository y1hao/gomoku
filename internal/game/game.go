package game

import (
	"errors"
)

const Size = 15
const (
	None = iota
	Black
	White
	Draw
)

type Player uint8
type Board [Size][Size]Player

type Piece struct {
	Row    int    `json:"row"`
	Col    int    `json:"col"`
	Player Player `json:"player"`
}

type Game struct {
	Board         Board    `json:"board"`
	Player        Player   `json:"player"`
	Winner        Player   `json:"winner,omitempty"`
	LastMove      *Piece   `json:"lastMove,omitempty"`
	WinningPieces []*Piece `json:"winningPieces,omitempty"`
}

func New() *Game {
	board := [Size][Size]Player{}
	for i := range board {
		board[i] = [Size]Player{}
	}

	return &Game{
		Board:  board,
		Player: Black,
	}
}

func (g *Game) Move(p *Piece) error {
	if g.Winner != None {
		return errors.New("the game has ended")
	}

	r, c := p.Row, p.Col
	if r < 0 || r >= Size ||
		c < 0 || c >= Size ||
		g.Board[r][c] != 0 ||
		p.Player != g.Player {
		return errors.New("invalid position")
	}
	g.LastMove = p
	player := g.Player
	if player == Black {
		g.Player = White
	} else {
		g.Player = Black
	}
	g.Board[r][c] = player
	g.calcWinning(p)

	if g.Winner == None {
		g.testDraw()
	}

	return nil
}

func (g *Game) testDraw() {
	for i := range g.Board {
		for j := range g.Board[i] {
			if g.Board[i][j] == None {
				return
			}
		}
	}
	g.Winner = Draw
}

func (g *Game) calcWinning(p *Piece) {
	var beg, end int

	// vertical
	for beg = p.Row; beg >= 0; beg-- {
		if g.Board[beg][p.Col] != p.Player {
			break
		}
	}
	for end = p.Row; end < Size; end++ {
		if g.Board[end][p.Col] != p.Player {
			break
		}
	}
	if end-beg-1 >= 5 {
		g.Winner = p.Player
		for i := beg + 1; i < end; i++ {
			g.WinningPieces = append(g.WinningPieces, &Piece{Row: i, Col: p.Col, Player: p.Player})
		}
	}

	// horizontal
	for beg = p.Col; beg >= 0; beg-- {
		if g.Board[p.Row][beg] != p.Player {
			break
		}
	}
	for end = p.Col; end < Size; end++ {
		if g.Board[p.Row][end] != p.Player {
			break
		}
	}
	if end-beg-1 >= 5 {
		for i := beg + 1; i < end; i++ {
			g.Winner = p.Player
			g.WinningPieces = append(g.WinningPieces, &Piece{Row: p.Row, Col: i, Player: p.Player})
		}
	}

	var begR, begC, endR, endC int
	// forward diagonal
	for begR, begC = p.Row, p.Col; begR >= 0 && begC < Size; begR, begC = begR-1, begC+1 {
		if g.Board[begR][begC] != p.Player {
			break
		}
	}
	for endR, endC = p.Row, p.Col; endR < Size && endC >= 0; endR, endC = endR+1, endC-1 {
		if g.Board[endR][endC] != p.Player {
			break
		}
	}
	if endR-begR-1 >= 5 {
		for i, j := begR+1, begC-1; i < endR && j > endC; i, j = i+1, j-1 {
			g.Winner = p.Player
			g.WinningPieces = append(g.WinningPieces, &Piece{Row: i, Col: j, Player: p.Player})
		}
	}

	// backward diagonal
	for begR, begC = p.Row, p.Col; begR >= 0 && begC >= 0; begR, begC = begR-1, begC-1 {
		if g.Board[begR][begC] != p.Player {
			break
		}
	}
	for endR, endC = p.Row, p.Col; endR < Size && endC < Size; endR, endC = endR+1, endC+1 {
		if g.Board[endR][endC] != p.Player {
			break
		}
	}
	if endR-begR-1 >= 5 {
		for i, j := begR+1, begC+1; i < endR && j < endC; i, j = i+1, j+1 {
			g.Winner = p.Player
			g.WinningPieces = append(g.WinningPieces, &Piece{Row: i, Col: j, Player: p.Player})
		}
	}
}
