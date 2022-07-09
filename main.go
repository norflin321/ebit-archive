package main

// https://ebiten.org/documents/packages.html
import (
	"log"

	"github.com/hajimehoshi/ebiten/v2" // provides functions for drawing and inputs
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// has necessary functions for an Ebitengine game: Update, Draw and Layout
type Game struct{}

// that is called every tick. Tick is a time unit for logical updating.
// The default value is 1/60 [s], then Update is called 60 times per second by default
// Update updates the game's logical state.
func (g *Game) Update() error {
	// when updating function returns a non-nil error, the Ebitengine game suspends.
	return nil
}

// that is called every frame. Frame is a time unit for rendering and this depends on the display's refresh rate.
// If the monitor's refresh rate is 60 [Hz], Draw is called 60 times per second.
// (screen *ebiten.Image) argument is the final destination of rendering.
// The window shows the final state of screen every frame.
// As screen is reset (or cleared in another word) whenever Draw is called.
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

// Layout accepts an outside size, which is a window size on desktop, and returns the game's logical screen size.
// it will be more meaningful e.g., when the window is resizable.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	// the function to run the main loop of a game.
	// The argument is a Game object, that implements ebiten.Game.
	// ebiten.RunGame returns a non-nil error value when an error happen
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
