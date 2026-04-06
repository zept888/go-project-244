package code

import "fmt"

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

	return fmt.Sprintf("file1: %v\nfile2: %v", data1, data2), nil
}
