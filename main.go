package main

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/guri3/ikayaku/ikayaku"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(ikayaku.ScreenWidth, ikayaku.ScreenHeight)
	ebiten.SetWindowTitle("Ikayaku")
	if err := ebiten.RunGame(ikayaku.NewGame()); err != nil {
		log.Fatal(err)
	}
}
