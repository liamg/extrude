package format

import (
	"bytes"
	"io"
)

func Sniff(r io.ReaderAt) (Format, error) {
	buffer := make([]byte, 4)
	if _, err := r.ReadAt(buffer, 0); err != nil {
		return Unknown, err
	}
	for _, definition := range definitions {
		for _, signature := range definition.signatures {
			if bytes.Compare(buffer[:len(signature)], signature) == 0 {
				return definition.format, nil
			}
		}
	}
	return Unknown, nil
}
