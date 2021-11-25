package macho

import (
	"debug/macho"
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

	fat, err := macho.NewFatFile(r)
	if err != nil {
		if err != macho.ErrNotFat {
			return nil, err
		}
		thin, err := macho.NewFile(r)
		if err != nil {
			return nil, err
		}
		defer func() { _ = thin.Close() }()
		metadata.thin = thin
	} else {
		defer func() { _ = fat.Close() }()
		metadata.fat = fat
	}

	if err := metadata.analyse(); err != nil {
		return nil, err
	}
	return &metadata, nil
}
