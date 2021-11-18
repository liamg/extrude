package parser

import (
	"io"

	"github.com/liamg/extrude/pkg/format"
	"github.com/liamg/extrude/pkg/report"
)

type ParseFunc func(seeker io.ReadSeeker, filename string, format format.Format) (report.Report, error)
