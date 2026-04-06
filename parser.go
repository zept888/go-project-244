package code

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ParseFile(path string) (map[string]any, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve path %q: %w", path, err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file %q: %w", absPath, err)
	}

	ext := filepath.Ext(absPath)

	switch ext {
	case ".json":
		return parseJSON(data)
	default:
		return nil, fmt.Errorf("unsupported file format: %q", ext)
	}
}

func parseJSON(data []byte) (map[string]any, error) {
	var result map[string]any
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	return result, nil
}
