package format

import (
	"bytes"
	"io"
)

func Sniff(seeker io.ReadSeeker) (Format, error) {
	if _, err := seeker.Seek(0, 0); err != nil {
		return Unknown, err
	}
	buffer := make([]byte, 4)
	if _, err := seeker.Read(buffer); err != nil {
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
