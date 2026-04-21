package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
