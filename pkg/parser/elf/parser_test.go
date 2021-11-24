package elf

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/liamg/extrude/pkg/format"
	"github.com/liamg/extrude/pkg/parser/elf/compiler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestELFDerivingCompiler(t *testing.T) {

	tests := []struct {
		file     string
		expected compiler.Compiler
	}{
		{
			file:     "ocaml",
			expected: compiler.CompilerOCaml,
		},
		{
			file:     "nim",
			expected: compiler.CompilerNim,
		},
	}

	for _, test := range tests {
		path := filepath.Join("_testdata", test.file)
		t.Run(test.expected.String(), func(t *testing.T) {

			f, err := os.Open(path)
			require.NoError(t, err)

			reporter, err := New().Parse(f, path, format.ELF)
			require.NoError(t, err)

			assert.Equal(t, test.expected, reporter.(*Metadata).CompilerInfo.Compiler)
		})
	}

}
