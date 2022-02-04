package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

const (
	screenWidth, screenHeight = 650, 360
	birdsCount                = 500
)

type Game struct{}

var green = color.RGBA{
	R: 10,
	G: 255,
	B: 50,
	A: 255,
}

var birds [birdsCount]*Bird

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, bird := range birds {
		screen.Set(int(bird.Position.X+1), int(bird.Position.Y), green)
		screen.Set(int(bird.Position.X-1), int(bird.Position.Y), green)
		screen.Set(int(bird.Position.X), int(bird.Position.Y-1), green)
		screen.Set(int(bird.Position.X), int(bird.Position.Y+1), green)
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	for i := 0; i < birdsCount; i++ {
		createBirds(i)
	}
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Birds in the Sky !")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatalln(err)
	}
}
