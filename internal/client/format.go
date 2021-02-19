package client

import (
	"fmt"
)

const (
	blackF, blackB = 30, 40
	redF, redB = 31, 41
	greenF, greenB = 32, 42
	yellowF, yellowB = 33, 43
	blueF, blueB = 34, 44
	magentaF, magentaB = 35, 45
	cyanF, cyanB = 36, 46
	whiteF, whiteB = 37, 47
)

type color int

type colored struct {
	fg, bg color
	text string
}

type formatted struct {
	row, col int
	colored
}

func print(output ...*formatted) {
	pushPosition()
	defer popPosition()

	defer resetColor()
	for _, f := range output {
		setPosition(f.row, f.col)
		setColor(f.fg, f.bg)
		fmt.Print(f.text)
	}
}

func clear() { fmt.Print("\033[J") }

func clearArea(fromR, fromC, toR, toC int) {

}

func pushPosition() { fmt.Print("\033[s") }
func popPosition() { fmt.Print("\330[u") }
func setPosition(row, col int) { fmt.Printf("\033[%d;%dH", row, col) }
func setColor(fg, bg color) { fmt.Printf("\033[%d;%dm", fg, bg) }
func resetColor() { fmt.Printf("\033[m") }
