package code

import (
	"code/internal/diff"
	"code/internal/formatters"
	"code/internal/parsers"
	"sort"
)

func buildDiff(data1, data2 map[string]any) []diff.Node {
	// собираем все ключи
	keysMap := make(map[string]struct{})
	for k := range data1 {
		keysMap[k] = struct{}{}
	}
	for k := range data2 {
		keysMap[k] = struct{}{}
	}

	keys := make([]string, 0, len(keysMap))
	for k := range keysMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var nodes []diff.Node
	for _, k := range keys {
		val1, inFile1 := data1[k]
		val2, inFile2 := data2[k]

		// оба значения — вложенные мапы, уходим в рекурсию
		map1, isMap1 := val1.(map[string]any)
		map2, isMap2 := val2.(map[string]any)

		switch {
		case inFile1 && inFile2 && isMap1 && isMap2:
			nodes = append(nodes, diff.Node{
				Key:      k,
				Type:     diff.Nested,
				Children: buildDiff(map1, map2),
			})
		case inFile1 && inFile2 && val1 == val2:
			nodes = append(nodes, diff.Node{
				Key:      k,
				Type:     diff.Unchanged,
				OldValue: val1,
			})
		case inFile1 && inFile2:
			nodes = append(nodes, diff.Node{
				Key:      k,
				Type:     diff.Updated,
				OldValue: val1,
				NewValue: val2,
			})
		case inFile1:
			nodes = append(nodes, diff.Node{
				Key:      k,
				Type:     diff.Removed,
				OldValue: val1,
			})
		default:
			nodes = append(nodes, diff.Node{
				Key:      k,
				Type:     diff.Added,
				NewValue: val2,
			})
		}
	}
	return nodes
}

// GenDiff compares two configuration files and returns the difference as a string.
// The format parameter specifies the output format (stylish, plain, json).
func GenDiff(filepath1, filepath2, format string) (string, error) {
	data1, err := parsers.ParseFile(filepath1)
	if err != nil {
		return "", err
	}

	data2, err := parsers.ParseFile(filepath2)
	if err != nil {
		return "", err
	}

	nodes := buildDiff(data1, data2)

	return formatters.Format(nodes, format)
}
