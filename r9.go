package main

import (
	_ "embed"
	"log"
	"r9/es"
	"r9/sb"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	W *es.World
}

func (g *Game) Update() error {
	g.W.Canvas.Clear(sb.ColorBlack)
	g.W.Run()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.ReplacePixels(g.W.Canvas.Buf)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return es.Width, es.Height
}

func main() {
	gg := &Game{W: es.NewWorld(20)}
	ebiten.SetWindowSize(es.Width, es.Height)
	ebiten.SetWindowTitle("r9")
	if err := ebiten.RunGame(gg); err != nil {
		log.Fatal(err)
	}
}
