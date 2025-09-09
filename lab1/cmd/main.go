package main

import (
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

//Вариант  - 7

type Game struct {
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	Cellsize := 32
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			var col *color.RGBA
			if (i+j)%2 == 0 {
				col = &color.RGBA{0, 0, 0, 255}
			} else {
				col = &color.RGBA{255, 255, 255, 255}
			}

			ebitenutil.DrawRect(screen,
				float64(i),
				float64(j), float64(Cellsize), float64(Cellsize), col)
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func initOSVariables() {
	os.Setenv("LIBGL_ALWAYS_SOFTWARE", "1")
	os.Setenv("GALLIUM_DRIVER", "llvmpipe")
	os.Setenv("MESA_LOADER_DRIVER_OVERRIDE", "swrast")
	os.Setenv("EBITENGINE_GRAPHICS_LIBRARY", "opengl")
	os.Setenv("EBITENGINE_OPENGL", "es")

}

func main() {
	initOSVariables()
	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified
	ebiten.SetWindowTitle("Your game's title")
	ebiten.SetVsyncEnabled(false)
	ebiten.SetMaxTPS(30)
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
