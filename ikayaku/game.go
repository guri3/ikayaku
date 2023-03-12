package ikayaku

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 360
)

var background *ebiten.Image
var ika1 *ebiten.Image
var ika2 *ebiten.Image
var ika3 *ebiten.Image
var ika4 *ebiten.Image
var mplusNormalFont font.Face
var score *Score
var timer *Timer

func init() {
	var err error
	background, _, err = ebitenutil.NewImageFromFile("assets/images/shitirin.jpg")
	if err != nil {
		log.Fatal(err)
	}

	ika1, _, err = ebitenutil.NewImageFromFile("assets/images/ika_1.png")
	if err != nil {
		log.Fatal(err)
	}

	ika2, _, err = ebitenutil.NewImageFromFile("assets/images/ika_2.png")
	if err != nil {
		log.Fatal(err)
	}

	ika3, _, err = ebitenutil.NewImageFromFile("assets/images/ika_3.png")
	if err != nil {
		log.Fatal(err)
	}

	ika4, _, err = ebitenutil.NewImageFromFile("assets/images/ika_4.gif")
	if err != nil {
		log.Fatal(err)
	}

	score = NewScore()

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	timer = NewTimer()
}

type Game struct {
	counter int
}

func NewGame() *Game {
	return &Game{counter: 0}
}

func (g *Game) Update() error {
	if g.counter%ebiten.TPS() == 0 {
		if timer.GetTime() != 0 {
			timer.SubTime(1)
		}
	}
	g.counter++
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(background, nil)

	scoreMsg := fmt.Sprintf("Score: %d", score.GetScore())
	text.Draw(screen, scoreMsg, mplusNormalFont, 10, 30, color.White)
	timerMsg := fmt.Sprintf("Time : %d", timer.GetTime())
	text.Draw(screen, timerMsg, mplusNormalFont, 10, 60, color.White)
}
