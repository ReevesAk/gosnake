package main

import (
	"math/rand"

	term "github.com/JoelOtter/termloop"
)

type Food struct {
	*term.Entity
	coord Coordinates
}

// NewFood creates a new Food at a random position.
func NewFood() *Food {
	food := new(Food)
	food.Entity = term.NewEntity(1, 1, 1, 1)
	food.moveToRandomPosition()
	return food
}

// Draw draws the Food as a default character.
func (f *Food) Draw(screen *term.Screen) {
	screen.RenderCell(f.coord.x, f.coord.y, &term.Cell{
		Fg: term.ColorRed,
		Ch: '$',
	})
}

// Position returns the x,y position of this Food.
func (f Food) Position() (int, int) {
	return f.coord.x, f.coord.y
}

// Size returns the size of this Food - always 1x1.
func (f Food) Size() (int, int) {
	return 1, 1
}

// Collide handles collisions with the Snake. It updates the score and places
// the Food randomly on the screen again.
func (f *Food) Collide(collision term.Physical) {
	switch collision.(type) {
	case *Snake:
		// It better be a snake that we're colliding with...
		f.handleSnakeCollision()
	}
}

func (f *Food) moveToRandomPosition() {
	newX := randInRange(1, border.width-1)
	newY := randInRange(1, border.height-1)
	f.coord.x, f.coord.y = newX, newY
	f.SetPosition(newX, newY)

}

func (f *Food) handleSnakeCollision() {
	f.moveToRandomPosition()
	IncrementScore(5)
}

func randInRange(min, max int) int {
	return rand.Intn(max-min) + min
}

