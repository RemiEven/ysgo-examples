package assets

import (
	"bytes"
	"context"
	"fmt"
	"image"

	_ "image/png" // needed to correctly load PNG files

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/remieven/ysgo"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
)

func GetImage(ctx context.Context, path string) (*ebiten.Image, error) {
	content, err := getFileData(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("failed to load image [%s]: %w", path, err)
	}

	img, _, err := image.Decode(bytes.NewReader(content))
	if err != nil {
		return nil, fmt.Errorf("failed to decode image [%s]: %w", path, err)
	}

	return ebiten.NewImageFromImage(img), nil
}

func GetFontFace(ctx context.Context) (font.Face, error) {
	parsedGoFont, err := truetype.Parse(gomono.TTF)
	if err != nil {
		return nil, fmt.Errorf("failed to parse gomono font: %w", err)
	}

	fontFace := truetype.NewFace(parsedGoFont, &truetype.Options{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	return fontFace, nil
}

func GetDialogueRunner(ctx context.Context, path string) (*ysgo.DialogueRunner, error) {
	dialogueFileData, err := getFileData(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("failed to load dialogue file [%s]: %w", path, err)
	}

	dr, err := ysgo.NewDialogueRunner(nil, "", bytes.NewReader(dialogueFileData))
	if err != nil {
		return nil, fmt.Errorf("failed to parse dialogue [%s]: %w", path, err)
	}

	return dr, nil
}
