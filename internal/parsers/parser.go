package parsers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// ParseFile reads a file and returns its contents as a map.
// The format is determined by the file extension (.json, .yml, .yaml).
func ParseFile(path string) (map[string]any, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve path %q: %w", path, err)
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file %q: %w", absPath, err)
	}

	switch filepath.Ext(absPath) {
	case ".json":
		return parseJSON(data)
	case ".yml", ".yaml":
		return parseYAML(data)
	default:
		return nil, fmt.Errorf("unsupported file format: %q", filepath.Ext(absPath))
	}
}

func parseJSON(data []byte) (map[string]any, error) {
	var result map[string]any
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	return result, nil
}

func parseYAML(data []byte) (map[string]any, error) {
	var result map[string]any
	if err := yaml.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("invalid YAML: %w", err)
	}
	return result, nil
}
