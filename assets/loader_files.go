//go:build !wasm

package assets

import (
	"context"
	"embed"
	"fmt"

	_ "image/png" // needed to correctly load PNG files
)

//go:embed files
var assetFS embed.FS

func getFileData(ctx context.Context, path string) ([]byte, error) {
	fileData, err := assetFS.ReadFile("files/" + path)
	if err != nil {
		return nil, fmt.Errorf("failed to read path [%s]: %w", path, err)
	}
	return fileData, nil
}
