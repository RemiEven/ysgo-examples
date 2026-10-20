package assets

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"syscall/js"

	_ "image/png" // needed to correctly load PNG files
)

func getFileData(ctx context.Context, path string) ([]byte, error) {
	href := js.Global().Get("location").Get("href").String()
	href = strings.Split(href, "/index.html")[0]
	url := href + "/assets/" + path

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for url [%s]: %w", url, err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch file for url [%s]: %w", url, err)
	}
	defer response.Body.Close()

	if response.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("received an error HTTP status [%v] for url [%s]", response.StatusCode, url)
	}

	fileData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read fetched data for path [%s]: %w", path, err)
	}

	return fileData, nil
}
