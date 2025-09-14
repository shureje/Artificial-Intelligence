package main

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Вариант  - 7
type Position struct {
	X int
	Y int
}

type Game struct {
	Board         [8][8]int // -1 = не посещена, 0+ = номер хода
	Path          []Position
	Current       Position
	MoveCount     int
	Solved        bool
	Cellsize      float64
	StartPosition Position
}

var knightMoves = [8]Position{
	{2, 1}, {2, -1}, {-2, 1}, {-2, -1},
	{1, 2}, {1, -2}, {-1, 2}, {-1, -2},
}

func (g *Game) isValidMove(x, y int) bool {
	return x >= 0 && x < 8 && y >= 0 && y < 8 && g.Board[x][y] == -1
}

func (g *Game) countMoves(x, y int) int {
	count := 0
	for _, move := range knightMoves {
		nx, ny := x+move.X, y+move.Y
		if g.isValidMove(nx, ny) {
			count++
		}
	}
	return count
}

func (g *Game) SolveKnightTour(x, y int) bool {
	g.Board[x][y] = g.MoveCount
	g.Path = append(g.Path, Position{x, y})
	g.MoveCount++

	time.Sleep(100 * time.Millisecond)

	if g.MoveCount == 64 {
		g.Solved = true
		return true
	}

	var moves []Position
	for _, move := range knightMoves {
		nx, ny := x+move.X, y+move.Y
		if g.isValidMove(nx, ny) {
			moves = append(moves, Position{nx, ny})
		}
	}

	for i := 0; i < len(moves)-1; i++ {
		for j := i + 1; j < len(moves); j++ {
			if g.countMoves(moves[i].X, moves[i].Y) > g.countMoves(moves[j].X, moves[j].Y) {
				moves[i], moves[j] = moves[j], moves[i]
			}
		}
	}

	for _, move := range moves {
		if g.SolveKnightTour(move.X, move.Y) {
			return true
		}
	}

	g.Board[x][y] = -1
	g.Path = g.Path[:len(g.Path)-1]
	g.MoveCount--

	return false
}
func (g *Game) Draw(screen *ebiten.Image) {
	x, y := 640.0, 480.0
	border := 20.0

	ebitenutil.DrawRect(
		screen,
		0,
		0,
		640,
		480,
		color.RGBA{255, 228, 181, 255},
	)
	ebitenutil.DrawRect(
		screen,
		x/2-g.Cellsize*4-border/2,
		y/2-g.Cellsize*4-border/2,
		g.Cellsize*8+border,
		g.Cellsize*8+border,
		color.RGBA{160, 82, 45, 255},
	)

	ebitenutil.DrawRect(
		screen,
		x/2-g.Cellsize*4-border/4,
		y/2-g.Cellsize*4-border/4,
		g.Cellsize*8+border/2,
		g.Cellsize*8+border/2,
		color.RGBA{0, 0, 0, 255},
	)

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			var col *color.RGBA
			if (i+j)%2 == 0 {
				col = &color.RGBA{139, 69, 19, 255}
			} else {
				col = &color.RGBA{245, 222, 179, 255}
			}

			ebitenutil.DrawRect(screen,
				float64(i)*g.Cellsize+x/2-g.Cellsize*4,
				float64(j)*g.Cellsize+y/2-g.Cellsize*4,
				g.Cellsize,
				g.Cellsize, col)
		}
	}

	if len(g.Path) > 0 {
		lastPos := g.Path[len(g.Path)-1]
		knightX := float64(lastPos.X)*g.Cellsize + x/2 - g.Cellsize*4 + g.Cellsize/2
		knightY := float64(lastPos.Y)*g.Cellsize + y/2 - g.Cellsize*4 + g.Cellsize/2

		ebitenutil.DrawRect(screen,
			knightX-8, knightY-8,
			16, 16,
			color.RGBA{255, 0, 0, 255})

	}

	for i := 0; i < len(g.Path)-1; i++ {
		ebitenutil.DrawLine(screen,
			float64(g.Path[i].X)*g.Cellsize+x/2-g.Cellsize*4+g.Cellsize/2,
			float64(g.Path[i].Y)*g.Cellsize+y/2-g.Cellsize*4+g.Cellsize/2,
			float64(g.Path[i+1].X)*g.Cellsize+x/2-g.Cellsize*4+g.Cellsize/2,
			float64(g.Path[i+1].Y)*g.Cellsize+y/2-g.Cellsize*4+g.Cellsize/2,
			color.RGBA{0, 0, 0, 255})
	}
}

func NewGame() *Game {
	g := &Game{
		Cellsize:      48.0,
		StartPosition: Position{1, 3},
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			g.Board[i][j] = -1
		}
	}
	go g.SolveKnightTour(g.StartPosition.X, g.StartPosition.Y)
	return g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func initOSVariables() {

}

func main() {
	initOSVariables()
	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified
	ebiten.SetWindowTitle("Your game's title")
	ebiten.SetVsyncEnabled(false)
	ebiten.SetMaxTPS(30)
	ebiten.SetWindowSize(game.Layout(320, 240))
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
