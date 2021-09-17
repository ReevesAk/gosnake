package main

import term "github.com/JoelOtter/termloop"

type Coordinates struct {
	x, y int
}

// Border is the edge of the playing area. If the Snake collides with it,
// it dies.
type Border struct {
	*term.Entity
	width, height int
	coords        map[Coordinates]int
}

// NewBorder creates a Border with the given dimensions.
func NewBorder(w, h int) *Border {
	arenaBorder := new(Border)
	arenaBorder.Entity = term.NewEntity(1, 1, 1, 1)
	// Subtract one to account for bottom and right border
	arenaBorder.width, arenaBorder.height = w-1, h-1

	arenaBorder.coords = make(map[Coordinates]int)

	// Top and bottom
	for x := 0; x < arenaBorder.width; x++ {
		arenaBorder.coords[Coordinates{x, 0}] = 1
		arenaBorder.coords[Coordinates{x, arenaBorder.height}] = 1
	}

	// Left and right
	for y := 0; y < arenaBorder.height+1; y++ {
		arenaBorder.coords[Coordinates{0, y}] = 1
		arenaBorder.coords[Coordinates{arenaBorder.width, y}] = 1
	}

	return arenaBorder
}

// Contains returns true if a Coord is part of the border, else false.
// Used for collision detection.
func (b *Border) Contains(coord Coordinates) bool {
	_, exists := b.coords[coord]
	return exists
}

// Draw draws the border on the screen. A default color is used.
func (b *Border) Draw(screen *term.Screen) {
	if b == nil {
		return
	}

	for c := range b.coords {
		screen.RenderCell(c.x, c.y, &term.Cell{
			Bg: term.ColorBlack,
		})
	}
}
