package output

import "github.com/liamg/extrude/pkg/report"

type Outputter func(report.Report) error
