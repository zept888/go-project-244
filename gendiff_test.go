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
