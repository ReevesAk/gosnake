package main 

import (
	"fmt"
	"flag"

	"github.com/nsf/termbox-go"
	term "github.com/JoelOtter/termloop"
)

var score = 0
var game *term.Game
var border *Border
var scoreText *term.Text
var isFullscreen *bool

type endgameScreen struct {
	*term.BaseLevel
}

// Handle events on the endLevel. Enter resets.
func (eg *endgameScreen) Tick(event term.Event) {
	if event.Type == term.EventKey { // Is it a keyboard event?
		if event.Key == term.KeyEnter {
			score = 0
			game.Screen().SetLevel(newMainLevel(isFullscreen))
		}
	}
}


func IncrementScore(amount int) {
	score += amount
	scoreText.SetText(fmt.Sprint(" Score: ", score, " "))
}


func EndGame() {
	endLevel := term.NewBaseLevel(term.Cell{
		Bg: term.ColorRed,
	})
	gameScreen := new(endgameScreen)
	gameScreen.BaseLevel = endLevel
	
	game.Screen().SetLevel(gameScreen)
}

func newMainLevel(isFullscreen *bool) term.Level{

	mainLevel := term.NewBaseLevel(term.Cell{
		Bg: term.ColorBlack,
	})

	width, height := 80, 30
	if *isFullscreen {
		// Must initialize Termbox before getting the terminal size
		termbox.Init()
		width, height = termbox.Size()
	}
	border = NewBorder(width, height)

	snake := NewSnake()
	food := NewFood()
	scoreText = term.NewText(0, 0, " Score: 0", term.ColorBlack, term.ColorWhite)

	mainLevel.AddEntity(border)
	mainLevel.AddEntity(snake)
	mainLevel.AddEntity(food)
	mainLevel.AddEntity(scoreText)
	return mainLevel
}

func main () {
	game :=	term.NewGame()

	isFullscreen = flag.Bool("fullscreen", false, "Play fullscreen!")
	flag.Parse()
	mainLevel := newMainLevel(isFullscreen)
	game.Screen().SetLevel(mainLevel)
	game.Screen().SetFps(10)

	game.Start()
}