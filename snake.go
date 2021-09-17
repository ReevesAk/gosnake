package main

import tl "github.com/JoelOtter/termloop"

type direction int

const (
	right direction = iota
	left
	up
	down
)

// Snake is the snake.
type Snake struct {
	*tl.Entity
	body      []Coordinates
	length   int
	direction direction
}

// NewSnake creates a new Snake with a default length and position.
func NewSnake() *Snake {
	s := new(Snake)
	s.Entity = tl.NewEntity(3, 3, 1, 1)
	s.body = []Coordinates{
		{3, 5},
		{4, 5},
		{5, 5}, // head
	}

	s.length = len(s.body)
	s.direction = right
	return s
}

func (s *Snake) head() *Coordinates {
	return &s.body[len(s.body)-1]
}

func (s *Snake) grow(amount int) {
	s.length += amount
}

func (s *Snake) isGrowing() bool {
	return s.length > len(s.body)
}

func (s *Snake) isCollidingWithSelf() bool {
	for i := 0; i < len(s.body)-1; i++ {
		if *s.head() == s.body[i] {
			return true
		}
	}
	return false
}

func (s *Snake) isCollidingWithBorder() bool {
	return border.Contains(*s.head())
}


func (s *Snake) Draw(screen *tl.Screen) {
	// Update position based on direction
	newHead := *s.head()
	switch s.direction {
	case right:
		newHead.x++
	case left:
		newHead.x--
	case up:
		newHead.y--
	case down:
		newHead.y++
	}

	if s.isGrowing() {
		// We must be growing
		s.body = append(s.body, newHead)
	} else {
		s.body = append(s.body[1:], newHead)
	}

	s.SetPosition(newHead.x, newHead.y)

	if s.isCollidingWithSelf() || s.isCollidingWithBorder() {
		EndGame()
	}

	// Draw snake
	for _, c := range s.body {
		screen.RenderCell(c.x, c.y, &tl.Cell{
			Fg: tl.ColorGreen,
			Ch: 'o',
		})
	}
}

// Tick handles keypress events
func (s *Snake) Tick(event tl.Event) {
	// Find new direction - but you can't go
	// back from where you came.
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyArrowRight:
			if s.direction != left {
				s.direction = right
			}
		case tl.KeyArrowLeft:
			if s.direction != right {
				s.direction = left
			}
		case tl.KeyArrowUp:
			if s.direction != down {
				s.direction = up
			}
		case tl.KeyArrowDown:
			if s.direction != up {
				s.direction = down
			}
		case 0:
			// Vim mode!
			switch event.Ch {
			case 'h', 'H':
				if s.direction != right {
					s.direction = left
				}
			case 'j', 'J':
				if s.direction != up {
					s.direction = down
				}
			case 'k', 'K':
				if s.direction != down {
					s.direction = up
				}
			case 'l', 'L':
				if s.direction != left {
					s.direction = right
				}
			}
		}
	}
}

// Collide is called when a collision occurs, since this Snake is a
// DynamicPhysical that can handle its own collisions. Here we check what
// we're colliding with and handle it accordingly.
func (s *Snake) Collide(collision tl.Physical) {
	switch collision.(type) {
	case *Food:
		s.handleFoodCollision()
	case *Border:
		s.handleBorderCollision()
	}
}

func (s *Snake) handleFoodCollision() {
	s.grow(5)
}

func (s *Snake) handleBorderCollision() {
	EndGame()
}