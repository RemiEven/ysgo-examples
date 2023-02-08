package src

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"io"
	"net/http"
	"strings"

	// "syscall/js"

	_ "image/png" // needed to correctly load PNG files

	"github.com/RemiEven/ysgo/tree"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
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

func GetFontFace(ctx context.Context, path string) (font.Face, error) {
	fontFileData, err := getFileData(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("failed to load font [%s]: %w", path, err)
	}

	parsedFont, err := opentype.Parse(fontFileData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font [%s]: %w", path, err)
	}

	const dpi = 72
	fontFace, err := opentype.NewFace(parsedFont, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create face from parsed font [%s]: %w", path, err)
	}

	return fontFace, nil
}

func GetDialogue(ctx context.Context, path string) (*tree.Dialogue, error) {
	dialogueFileData, err := getFileData(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("failed to load dialogue [%s]: %w", path, err)
	}

	dialogueTree, err := tree.FromString(string(dialogueFileData))
	if err != nil {
		return nil, fmt.Errorf("failed to parse dialogue [%s]: %w", path, err)
	}

	return dialogueTree, nil
}

func getFileData(ctx context.Context, path string) ([]byte, error) {
	// href := js.Global().Get("location").Get("href").String()
	href := ""
	href = strings.TrimSuffix(href, "/index.html")

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, href+"/assets/"+path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for path [%s]: %w", path, err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch file for path [%s]: %w", path, err)
	}
	defer response.Body.Close()

	fileData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read fetched data for path [%s]: %w", path, err)
	}

	return fileData, nil
}
