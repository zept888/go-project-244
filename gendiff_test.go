package code

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expected = `{
    common: {
      + follow: false
        setting1: Value 1
      - setting2: 200
      - setting3: true
      + setting3: null
      + setting4: blah blah
      + setting5: {
            key5: value5
        }
        setting6: {
            doge: {
              - wow: 
              + wow: so much
            }
            key: value
          + ops: vops
        }
    }
    group1: {
      - baz: bas
      + baz: bars
        foo: bar
      - nest: {
            key: value
        }
      + nest: str
    }
  - group2: {
        abc: 12345
        deep: {
            id: 45
        }
    }
  + group3: {
        deep: {
            id: {
                number: 45
            }
        }
        fee: 100500
    }
}`

func TestGenDiffNested(t *testing.T) {
	result, err := GenDiff("testdata/fixture/file1_nested.json", "testdata/fixture/file2_nested.json", "stylish")
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGenDiffNestedYaml(t *testing.T) {
	result, err := GenDiff("testdata/fixture/file1_nested.yml", "testdata/fixture/file2_nested.yml", "stylish")
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGenDiffNestedPlain(t *testing.T) {
	expected := `Property 'common.follow' was added with value: false
Property 'common.setting2' was removed
Property 'common.setting3' was updated. From true to null
Property 'common.setting4' was added with value: 'blah blah'
Property 'common.setting5' was added with value: [complex value]
Property 'common.setting6.doge.wow' was updated. From '' to 'so much'
Property 'common.setting6.ops' was added with value: 'vops'
Property 'group1.baz' was updated. From 'bas' to 'bars'
Property 'group1.nest' was updated. From [complex value] to 'str'
Property 'group2' was removed
Property 'group3' was added with value: [complex value]`

	result, err := GenDiff("testdata/fixture/file1_nested.json", "testdata/fixture/file2_nested.json", "plain")
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGenDiffNestedJSON(t *testing.T) {
	result, err := GenDiff("testdata/fixture/file1_nested.json", "testdata/fixture/file2_nested.json", "json")
	assert.NoError(t, err)

	// проверяем что вывод валидный json
	var parsed []any
	assert.NoError(t, json.Unmarshal([]byte(result), &parsed))
	assert.NotEmpty(t, parsed)
}
