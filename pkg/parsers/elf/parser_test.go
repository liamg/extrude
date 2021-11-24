package elf

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/liamg/extrude/pkg/format"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestELFDerivingCompiler(t *testing.T) {

	tests := []struct {
		file     string
		expected string
	}{
		{
			file:     "ocaml",
			expected: "ocamlc",
		},
		{
			file:     "nim",
			expected: "nim",
		},
	}

	for _, test := range tests {
		path := filepath.Join("_testdata", test.file)
		t.Run(test.expected, func(t *testing.T) {

			f, err := os.Open(path)
			require.NoError(t, err)

			reporter, err := New().Parse(f, path, format.ELF)
			require.NoError(t, err)

			assert.Equal(t, test.expected, reporter.(*Metadata).CompilerInfo.Name)
		})
	}

}
