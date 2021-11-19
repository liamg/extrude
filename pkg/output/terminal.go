package output

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/liamg/extrude/pkg/report"
	"github.com/liamg/tml"
)

func Terminal(rep report.Report) error {

	width, _, err := terminal.GetSize(0)
	if err != nil {
		width = 80
	}

	_ = tml.Printf(`<dim>
`)

	for _, section := range rep.Sections() {
		fmt.Println("")
		_ = tml.Printf("<dim>┏━━</dim> %s <dim>%s┓</dim>\n", section.Heading(), strings.Repeat("━", width-6-len(section.Heading())))
		_ = tml.Printf("<dim>┃%s┃</dim>\n", strings.Repeat(" ", width-2))
		var maxKeyLen int
		for _, keyVal := range section.KeyValues() {
			if len(keyVal.Key()) > maxKeyLen {
				maxKeyLen = len(keyVal.Key())
			}
		}
		// add extra padding
		for _, keyVal := range section.KeyValues() {
			pad := width - maxKeyLen - 6 - len(keyVal.Value())
			paddedKey := strings.Repeat(" ", maxKeyLen-len(keyVal.Key())) + keyVal.Key()
			_ = tml.Printf("<dim>┃</dim>  %s  <blue>%s</blue>%s<dim>┃</dim>\n", paddedKey, keyVal.Value(), strings.Repeat(" ", pad))
		}
		_ = tml.Printf("<dim>┃%s┃</dim>\n", strings.Repeat(" ", width-2))
		_ = tml.Printf("<dim>┗%s┛</dim>\n", strings.Repeat("━", width-2))
	}

	for _, issue := range rep.Issues() {
		fmt.Printf("- %s\n", issue.Message())
	}
	return nil
}
