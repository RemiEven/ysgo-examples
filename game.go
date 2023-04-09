package main

import (
	"context"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"

	"github.com/remieven/ysgo"
	"github.com/remieven/ysgo-examples/assets"
)

type Game struct {
	backgroundImage         *ebiten.Image
	fontFace                font.Face
	dialogueRunner          *ysgo.DialogueRunner
	currentDialogueElement  *ysgo.DialogueElement
	currentlySelectedOption int
}

func (g *Game) Init() error {
	ctx := context.Background()
	backgroundImage, err := assets.GetImage(ctx, "background.png")
	if err != nil {
		return fmt.Errorf("failed to load background image: %w", err)
	}
	g.backgroundImage = backgroundImage

	fontFace, err := assets.GetFontFace(ctx)
	if err != nil {
		return fmt.Errorf("failed to load font face: %w", err)
	}
	g.fontFace = fontFace

	dr, err := assets.GetDialogueRunner(ctx, "script.yarn")
	if err != nil {
		return fmt.Errorf("failed to load dialogue: %w", err)
	}
	g.dialogueRunner = dr

	de, err := dr.Next(0)
	if err != nil {
		return fmt.Errorf("failed to get first dialogue element: %w", err)
	}
	g.currentDialogueElement = de

	return nil
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		de, err := g.dialogueRunner.Next(g.currentlySelectedOption)
		if err != nil {
			return fmt.Errorf("failed to get next dialogue element: %w", err)
		}
		g.currentDialogueElement = de
		g.currentlySelectedOption = 0
	}

	if g.currentDialogueElement == nil {
		return nil
	}

	if options := g.currentDialogueElement.Options; options != nil {
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			g.currentlySelectedOption++
			if g.currentlySelectedOption >= len(options) {
				g.currentlySelectedOption = 0
			}
		} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			g.currentlySelectedOption--
			if g.currentlySelectedOption < 0 {
				g.currentlySelectedOption += len(options)
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	dio := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.backgroundImage, dio)

	if g.currentDialogueElement == nil {
		text.Draw(screen, "END", g.fontFace, 190, 650, color.Black)
	} else if g.currentDialogueElement.Line != nil {
		text.Draw(screen, g.currentDialogueElement.Line.Text, g.fontFace, 190, 650, color.Black)
	} else if options := g.currentDialogueElement.Options; options != nil {
		for i := range options {
			line := options[i].Line.Text
			if i == g.currentlySelectedOption {
				line = "-> " + line
			} else {
				line = "   " + line
			}
			text.Draw(screen, line, g.fontFace, 190, 650+30*i, color.Black)
		}
	}
}

// Layout is used to implement the ebiten.Game interface
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1024, 768
}
