package png

import (
	"errors"
)

func Decode(content []byte) (string, error) {
	err := isPngFile(content)
	if err != nil {
		return "", err
	}

	return "PNG file decoded", nil
}

func isPngFile(content []byte) error {
	pngBytes := []byte{137, 80, 78, 71, 13, 10, 26, 10}
	for i := 0; i < len(pngBytes); i++ {
		if content[i] != pngBytes[i] {
			return errors.New("File is not a PNG")
		}
	}

	return nil
}
