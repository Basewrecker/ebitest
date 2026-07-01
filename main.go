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
	ScreenWidth  = 1000
	ScreenHeight = 1000
	SpriteScale  = 4.0
)

// constructors
type Game struct {
	player *Player
	enemy  *Enemy
}

type Player struct {
	position         Vector
	previousPosition Vector
	sprite           *ebiten.Image
}

type Enemy struct {
	position Vector
	sprite   *ebiten.Image
}

func (g *Game) Update() error {
	g.player.Update()

	w, h := g.Layout(0, 0)

	clampToScreen(&g.player.position, g.player.sprite, w, h)
	clampToScreen(&g.enemy.position, g.enemy.sprite, w, h)

	// scale-aware collision dimensions
	pw := float64(g.player.sprite.Bounds().Dx()) * SpriteScale
	ph := float64(g.player.sprite.Bounds().Dy()) * SpriteScale
	ew := float64(g.enemy.sprite.Bounds().Dx()) * SpriteScale
	eh := float64(g.enemy.sprite.Bounds().Dy()) * SpriteScale

	if rectsOverlap(
		g.player.position.X, g.player.position.Y, pw, ph,
		g.enemy.position.X, g.enemy.position.Y, ew, eh,
	) {
		g.player.position = g.player.previousPosition
	}

	return nil
}

// Player movement logic
func (p *Player) Update() {
	p.previousPosition = p.position

	speed := 9.0
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
	op.GeoM.Scale(SpriteScale, SpriteScale)
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.sprite, op)
}

// drawing the enemy sprite
func (e *Enemy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(SpriteScale, SpriteScale)
	op.GeoM.Translate(e.position.X, e.position.Y)
	screen.DrawImage(e.sprite, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
	g.enemy.Draw(screen)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	ebiten.SetWindowSize(1000, 1000)
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
		position: Vector{X: 300, Y: 300},
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

func clampToScreen(position *Vector, sprite *ebiten.Image, w, h int) {
	spriteW := float64(sprite.Bounds().Dx()) * SpriteScale
	spriteH := float64(sprite.Bounds().Dy()) * SpriteScale

	if position.X < 0 {
		position.X = 0
	} else if position.X > float64(w)-spriteW {
		position.X = float64(w) - spriteW
	}

	if position.Y < 0 {
		position.Y = 0
	} else if position.Y > float64(h)-spriteH {
		position.Y = float64(h) - spriteH
	}
}

// collision detection using axis aligned bounding box
func rectsOverlap(ax, ay, aw, ah, bx, by, bw, bh float64) bool {
	return ax < bx+bw &&
		ax+aw > bx &&
		ay < by+bh &&
		ay+ah > by
}
