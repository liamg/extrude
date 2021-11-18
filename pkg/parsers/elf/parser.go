package elf

import (
	"io"

	"github.com/liamg/extrude/pkg/format"
	"github.com/liamg/extrude/pkg/report"
)

func Parse(
	seeker io.ReadSeeker,
	filename string,
	format format.Format,
) (report.Report, error) {

	rep := report.New()

	overview := report.NewSection("Overview")

	overview.AddKeyValue("Filename", filename)
	overview.AddKeyValue("Format", format.String())

	rep.AddSection(overview)

	return rep, nil
}
