package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenDiff(t *testing.T) {
	result, err := GenDiff("testdata/fixture/file1.json", "testdata/fixture/file2.json", "stylish")

	assert.NoError(t, err)
	assert.Equal(t, `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`, result)
}

func TestGenDiffYaml(t *testing.T) {
	result, err := GenDiff("testdata/fixture/file1.yml", "testdata/fixture/file2.yml", "stylish")

	assert.NoError(t, err)
	assert.Equal(t, `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`, result)
}

func TestGenDiffUnsupportedFormat(t *testing.T) {
	_, err := GenDiff("testdata/fixture/file1.txt", "testdata/fixture/file2.txt", "stylish")
	assert.Error(t, err)
}
