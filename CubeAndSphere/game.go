package src

import (
	"context"
	"fmt"

	"github.com/RemiEven/ysgo/tree"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type Game struct {
	backgroundImage *ebiten.Image
	fontFace        font.Face
	dialogue        *tree.Dialogue
}

func (g *Game) Init() error {
	ctx := context.Background()
	backgroundImage, err := GetImage(ctx, "background.png")
	if err != nil {
		return fmt.Errorf("failed to load background image: %w", err)
	}
	g.backgroundImage = backgroundImage

	fontFace, err := GetFontFace(ctx, "JupiteroidRegular.ttf")
	if err != nil {
		return fmt.Errorf("failed to load font face: %w", err)
	}
	g.fontFace = fontFace

	dialogueTree, err := GetDialogue(ctx, "script.yarn")
	if err != nil {
		return fmt.Errorf("failed to load dialogue: %w", err)
	}
	g.dialogue = dialogueTree

	return nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	dio := &ebiten.DrawImageOptions{}
	dio.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(g.backgroundImage, dio)
}

// Layout is used to implement the ebiten.Game interface
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1024, 768
}
