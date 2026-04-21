package code

import (
	"fmt"
	"sort"
	"strings"
)

func GenDiff(filepath1, filepath2, format string) (string, error) {
	data1, err := ParseFile(filepath1)
	if err != nil {
		return "", err
	}

	data2, err := ParseFile(filepath2)
	if err != nil {
		return "", err
	}

	_ = format

	// собираем все ключи из обоих файлов
	keysMap := make(map[string]struct{})
	for k := range data1 {
		keysMap[k] = struct{}{}
	}
	for k := range data2 {
		keysMap[k] = struct{}{}
	}

	// сортируем
	keys := make([]string, 0, len(keysMap))
	for k := range keysMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// строим дифф
	var lines []string
	for _, k := range keys {
		val1, inFile1 := data1[k]
		val2, inFile2 := data2[k]

		switch {
		case inFile1 && inFile2 && val1 == val2:
			// ключ есть в обоих, значения совпадают
			lines = append(lines, fmt.Sprintf("    %s: %v", k, val1))
		case inFile1 && inFile2:
			// ключ есть в обоих, значения разные
			lines = append(lines, fmt.Sprintf("  - %s: %v", k, val1))
			lines = append(lines, fmt.Sprintf("  + %s: %v", k, val2))
		case inFile1:
			// ключ только в первом файле
			lines = append(lines, fmt.Sprintf("  - %s: %v", k, val1))
		default:
			// ключ только во втором файле
			lines = append(lines, fmt.Sprintf("  + %s: %v", k, val2))
		}
	}

	return "{\n" + strings.Join(lines, "\n") + "\n}", nil
}
