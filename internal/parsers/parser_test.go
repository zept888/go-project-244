package parsers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseFileUnsupported(t *testing.T) {
	_, err := ParseFile("../../testdata/fixture/file1.txt")
	assert.Error(t, err)
}

func TestParseFileJSON(t *testing.T) {
	result, err := ParseFile("../../testdata/fixture/file1.json")
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestParseFileYAML(t *testing.T) {
	result, err := ParseFile("../../testdata/fixture/file1.yml")
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestParseFileNotFound(t *testing.T) {
	_, err := ParseFile("nonexistent.json")
	assert.Error(t, err)
}

func TestParseFileInvalidJSON(t *testing.T) {
	tmpDir := t.TempDir()
	path := tmpDir + "/invalid.json"
	require.NoError(t, os.WriteFile(path, []byte("not a json"), 0644))

	_, err := ParseFile(path)
	assert.Error(t, err)
}

func TestParseFileInvalidYAML(t *testing.T) {
	tmpDir := t.TempDir()
	path := tmpDir + "/invalid.yaml"
	require.NoError(t, os.WriteFile(path, []byte("key: [invalid"), 0644))

	_, err := ParseFile(path)
	assert.Error(t, err)
}
