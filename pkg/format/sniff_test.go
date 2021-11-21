package format

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormatSniffing(t *testing.T) {
	tests := []struct {
		content  []byte
		expected Format
	}{
		{
			content:  []byte{0xfe, 0xed, 0xfa, 0xce},
			expected: MachO32,
		},
		{
			content:  []byte{0xce, 0xfa, 0xed, 0xfe},
			expected: MachO32,
		},
		{
			content:  []byte{0xfe, 0xed, 0xfa, 0xcf},
			expected: MachO64,
		},
		{
			content:  []byte{0xcf, 0xfa, 0xed, 0xfe},
			expected: MachO64,
		},
		{
			content:  []byte{0x7f, 'E', 'L', 'F'},
			expected: ELF,
		},
		{
			content:  []byte{'M', 'Z'},
			expected: PE,
		},
		{
			content:  []byte{'Z', 'M'},
			expected: PE,
		},
		{
			content:  []byte{0x00, 0x00, 0x00, 0x00},
			expected: Unknown,
		},
	}
	for _, test := range tests {
		t.Run(
			fmt.Sprintf("%q -> %s", string(test.content), test.expected),
			func(t *testing.T) {
				buffer := bytes.NewReader(append(test.content, make([]byte, 32)...))
				format, err := Sniff(buffer)
				require.NoError(t, err)
				assert.Equal(t, test.expected, format)
			},
		)
	}
}
