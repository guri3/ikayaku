package ikayaku

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"time"

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
	IkaSize      = 96
)

var background *ebiten.Image
var mplusNormalFont font.Face
var score *Score
var timer *Timer
var ikas []*Ika

func init() {
	var err error

	background, _, err = ebitenutil.NewImageFromFile("assets/images/shitirin.jpg")
	if err != nil {
		log.Fatal(err)
	}

	score = NewScore()
	timer = NewTimer()

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

	rand.Seed(time.Now().UnixNano())
}

type Game struct {
	counter int
}

func NewGame() *Game {
	return &Game{counter: 0}
}

func (g *Game) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		ikas = IkasClick(ikas, x, y)
	}

	if timer.GetTime() != 0 && g.counter%ebiten.TPS() == 0 {
		if timer.GetTime() != 0 {
			timer.SubTime(1)
		}

		ikas = IkasUpdate(ikas)

		if rand.Intn(9) < 5 {
			ikas = append(ikas, NewIka(float64(rand.Intn(ScreenWidth-IkaSize)), float64(rand.Intn(ScreenHeight-IkaSize)), score))
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
	for _, ika := range ikas {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(ika.x, ika.y)
		screen.DrawImage(ika.image, op)
	}

	scoreMsg := fmt.Sprintf("Score: %d", score.GetScore())
	text.Draw(screen, scoreMsg, mplusNormalFont, 10, 30, color.White)
	timerMsg := fmt.Sprintf("Time: %d", timer.GetTime())
	text.Draw(screen, timerMsg, mplusNormalFont, 10, 60, color.White)
}
