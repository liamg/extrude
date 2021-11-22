package output

import (
	"strings"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/liamg/extrude/pkg/report"
	"github.com/liamg/tml"
)

const (
	borderTopLeft     = '╭'
	borderTopRight    = '╮'
	borderBottomLeft  = '╰'
	borderBottomRight = '╯'
	borderVertical    = '│'
	borderHorizontal  = '─'
)

func Terminal(rep report.Report) error {
	width, _, err := terminal.GetSize(0)
	if err != nil {
		width = 80
	}
	return (&terminalOutput{width: width}).Output(rep)
}

type terminalOutput struct {
	width int
}

func (t *terminalOutput) printHeader(heading string) {
	_ = tml.Printf(
		"\r\n<dim>%c%c</dim> %s <dim>%s%c</dim>\n",
		borderTopLeft, borderHorizontal,
		heading,
		strings.Repeat(string(borderHorizontal), t.width-5-len(heading)),
		borderTopRight,
	)
	t.printBlank()
}

func (t *terminalOutput) printFooter() {
	t.printBlank()
	_ = tml.Printf("<dim>%c%s%c</dim>\n", borderBottomLeft, strings.Repeat(string(borderHorizontal), t.width-2), borderBottomRight)
}

func (t *terminalOutput) printBlank() {
	t.printIn("")
}

func (t *terminalOutput) printIn(format string, args ...interface{}) {
	str := tml.Sprintf(format, args...)
	repeat := t.width - 2 - len(str)
	padded := str
	if repeat > 0 {
		padded = str + strings.Repeat(" ", t.width-2-len(str))
	}
	_ = tml.Printf("<dim>%c</dim> %s <dim>%c</dim>\n", borderVertical, padded, borderVertical)
}

func (t *terminalOutput) Output(rep report.Report) error {

	_ = tml.Printf(`<dim>
`)

	for _, section := range rep.Sections() {
		t.printHeader(section.Heading())
		var maxKeyLen int
		for _, keyVal := range section.KeyValues() {
			if len(keyVal.Key()) > maxKeyLen {
				maxKeyLen = len(keyVal.Key())
			}
		}
		// add extra padding
		for _, keyVal := range section.KeyValues() {
			paddedKey := strings.Repeat(" ", maxKeyLen-len(keyVal.Key())) + keyVal.Key()
			t.printIn("%s  <blue>%s</blue>", paddedKey, keyVal.Value())
		}
		t.printFooter()
	}

	t.printHeader("Security")
	for _, row := range rep.ResultsTable().Rows() {
		var rowStr string
		for _, col := range row {
			rowStr += tml.Sprintf("%c%s", borderVertical, col)
		}
		t.printIn("%s%c", rowStr, borderVertical)
	}
	t.printFooter()

	return nil
}
