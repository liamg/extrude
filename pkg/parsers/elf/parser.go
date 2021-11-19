package elf

import (
	"debug/elf"
	"io"
	"path/filepath"

	"github.com/liamg/extrude/pkg/format"
	"github.com/liamg/extrude/pkg/report"
)

type parser struct{}

func New() *parser {
	return &parser{}
}

func (*parser) Parse(r io.ReaderAt, path string, format format.Format) (report.Reporter, error) {

	var metadata Metadata

	metadata.File.Path = path
	metadata.File.Name = filepath.Base(path)
	metadata.File.Format = format

	f, err := elf.NewFile(r)
	if err != nil {
		return nil, err
	}
	metadata.ELF = f
	if err := metadata.analyse(); err != nil {
		return nil, err
	}
	return &metadata, nil
}
