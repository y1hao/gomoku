package client

import (
	"fmt"
)

const (
	blackF, blackB     = 30, 40
	redF, redB         = 31, 41
	greenF, greenB     = 32, 42
	yellowF, yellowB   = 33, 43
	blueF, blueB       = 34, 44
	magentaF, magentaB = 35, 45
	cyanF, cyanB       = 36, 46
	whiteF, whiteB     = 37, 47
)

const (
	mainF, mainB           = blueF, blueB
	secondaryF, secondaryB = cyanF, cyanB
	infoF, infoB           = whiteF, whiteB
	errorF, errorB         = redF, redB
	youF, youB             = greenF, greenB
	oppF, oppB             = magentaF, magentaB
	highlightF, highlightB               = yellowF, yellowB
)

type color int

func pushPosition()            { fmt.Print("\033[s") }
func popPosition()             { fmt.Print("\033[u") }
func setPosition(row, col int) { fmt.Printf("\033[%d;%dH", row, col) }
func setColor(fg, bg color)    { fmt.Printf("\033[%d;%dm", fg, bg) }
func setDimColor(fg, bg color) { fmt.Printf("\033[2;%d;%dm", fg, bg) }
func resetColor()              { fmt.Printf("\033[m") }

func print(fg, bg color, m string) {
	setColor(fg, bg)
	defer resetColor()
	fmt.Print(m)
}

func printDim(fg, bg color, m string) {
	setDimColor(fg, bg)
	defer resetColor()
	fmt.Print(m)
}

func printf(fg, bg color, format string, a ...interface{}) {
	print(fg, bg, fmt.Sprintf(format, a...))
}

func printDimf(fg, bg color, format string, a ...interface{}) {
	printDim(fg, bg, fmt.Sprintf(format, a...))
}

func clear() {
	fmt.Print("\033[2J")
	setPosition(1, 1)
}
