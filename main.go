package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var PlayerSprite = mustLoadImage("assets/PNG/Players/Tiles/tile_0000.png")

type Vector struct {
	X float64
	Y float64
}

type Game struct {
	player *Player
}

type Player struct {
	position Vector
	sprite   *ebiten.Image
}

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

// Player movement logic
func (p *Player) Update() {
	speed := 5.0

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.position.Y += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.position.Y -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += speed
	}
}

// drawing the sprites
func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello world")

	g := &Game{
		player: NewPlayer(),
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func NewPlayer() *Player {
	return &Player{
		position: Vector{X: 100, Y: 100},
		sprite:   PlayerSprite,
	}
}

// Sprite loading func
func mustLoadImage(name string) *ebiten.Image {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}
