package output

import (
	"fmt"

	"github.com/liamg/extrude/pkg/report"
	"github.com/liamg/tml"
)

func Terminal(rep report.Report) error {
	for _, section := range rep.Sections() {
		tml.Printf("<underline>%s</underline>\n", section.Heading())
		for _, keyVal := range section.KeyValues() {
			tml.Printf("%20s: %s\n", keyVal.Key(), keyVal.Value())
		}
		fmt.Println("")
	}
	return nil
}
