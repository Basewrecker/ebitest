package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

// declaring the sprites

var PlayerSprite = mustLoadImage("assets/PNG/Players/Tiles/tile_0001.png")
var EnemySprite = mustLoadImage("assets/PNG/Players/Tiles/tile_0012.png")

// constraints

type Vector struct {
	X float64
	Y float64
}

const (
	ScreenWidth  = 530
	ScreenHeight = 480
)

// constructors

type Game struct {
	player *Player
	enemy  *Enemy
}

type Player struct {
	position Vector
	sprite   *ebiten.Image
}

type Enemy struct {
	position Vector
	sprite   *ebiten.Image
}

func (g *Game) Update() error {
	g.player.Update()

	w, h := g.Layout(0, 0)

	if g.player.position.X < 0 {
		g.player.position.X = 0
	} else if g.player.position.X > float64(w-g.player.sprite.Bounds().Dx()) {
		g.player.position.X = float64(w - g.player.sprite.Bounds().Dx())
	}

	if g.player.position.Y < 0 {
		g.player.position.Y = 0
	} else if g.player.position.Y > float64(h-g.player.sprite.Bounds().Dy()) {
		g.player.position.Y = float64(h - g.player.sprite.Bounds().Dy())
	}

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

// drawing the player sprite
func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
}

// drawing the enemy sprite

func (p *Enemy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
	g.enemy.Draw(screen)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	ebiten.SetWindowSize(530, 480)
	ebiten.SetWindowTitle("Hello world")

	g := &Game{
		player: NewPlayer(),
		enemy:  NewEnemy(),
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

func NewEnemy() *Enemy {
	return &Enemy{
		position: Vector{X: 150, Y: 150},
		sprite:   EnemySprite,
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
