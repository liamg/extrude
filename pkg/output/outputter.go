package output

import "github.com/liamg/extrude/pkg/report"

type Options struct {
	IncludePassingTests bool
}

type Outputter func(report.Report, *Options) error
