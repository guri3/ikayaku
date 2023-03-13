package ikayaku

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var ika1Image, ika2Image, ika3Image, ika4Image *ebiten.Image

func init() {
	var err error

	ika1Image, _, err = ebitenutil.NewImageFromFile("assets/images/ika_1.png")
	if err != nil {
		log.Fatal(err)
	}

	ika2Image, _, err = ebitenutil.NewImageFromFile("assets/images/ika_2.png")
	if err != nil {
		log.Fatal(err)
	}

	ika3Image, _, err = ebitenutil.NewImageFromFile("assets/images/ika_3.png")
	if err != nil {
		log.Fatal(err)
	}

	ika4Image, _, err = ebitenutil.NewImageFromFile("assets/images/ika_4.gif")
	if err != nil {
		log.Fatal(err)
	}
}

type Ika struct {
	x     float64
	y     float64
	image *ebiten.Image
	score *Score
}

func NewIka(x, y float64, score *Score) *Ika {
	return &Ika{
		x:     x,
		y:     y,
		image: ika1Image,
		score: score,
	}
}

func IkasUpdate(ikas []*Ika) []*Ika {
	var newIkas []*Ika
	for _, ika := range ikas {
		isUpdated := ika.checkUpdate()
		if isUpdated {
			newIkas = append(newIkas, ika)
		}
	}
	return newIkas
}

func (i *Ika) checkUpdate() bool {
	if i.image == ika4Image {
		i.score.SubScore(3)
		return false
	}

	if i.image == ika1Image {
		i.updateToIka2()
	} else if i.image == ika2Image {
		i.updateToIka3()
	} else if i.image == ika3Image {
		i.updateToIka4()
	}

	return true
}

func (i *Ika) updateToIka2() {
	i.image = ika2Image
}

func (i *Ika) updateToIka3() {
	i.image = ika3Image
}

func (i *Ika) updateToIka4() {
	i.image = ika4Image
}

func IkasClick(ikas []*Ika, x, y int) []*Ika {
	var newIkas []*Ika
	for _, ika := range ikas {
		if ika.checkClick(x, y) {
			if ika.image == ika1Image {
				ika.score.AddScore(1)
			} else if ika.image == ika2Image {
				ika.score.AddScore(2)
			} else if ika.image == ika3Image {
				ika.score.AddScore(3)
			} else if ika.image == ika4Image {
				ika.score.SubScore(1)
			}
		} else {
			newIkas = append(newIkas, ika)
		}
	}
	return newIkas
}

func (i *Ika) checkClick(x, y int) bool {
	if x > (int(i.x)+i.image.Bounds().Min.X) && x < (int(i.x)+i.image.Bounds().Max.X) &&
		y > (int(i.y)+i.image.Bounds().Min.Y) && y < (int(i.y)+i.image.Bounds().Max.Y) {
		return true
	}
	return false
}
