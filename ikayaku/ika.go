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
	count int
}

func NewIka(x, y float64) *Ika {
	return &Ika{
		x:     x,
		y:     y,
		image: ika1Image,
		count: 0,
	}
}

func IkasUpdate(ikas []*Ika) []*Ika {
	var newIkas []*Ika
	for _, ika := range ikas {
		isUpdated := ika.CheckUpdate()
		if isUpdated {
			newIkas = append(newIkas, ika)
		}
	}
	return newIkas
}

func (i *Ika) CheckUpdate() bool {
	if i.image == ika4Image {
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
