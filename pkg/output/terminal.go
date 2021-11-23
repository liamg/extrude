package output

import (
	"fmt"
	"regexp"
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
	borderLeftT       = '├'
	borderRightT      = '┤'
)

func Terminal(rep report.Report) error {
	width, _, err := terminal.GetSize(0)
	if err != nil {
		width = 80
	}
	return (&terminalOutput{width: width}).Output(rep)
}

func safeRepeat(input string, repeat int) string {
	if repeat <= 0 {
		return ""
	}
	return strings.Repeat(input, repeat)
}

type terminalOutput struct {
	width int
}

func (t *terminalOutput) printHeader(heading string) {
	_ = tml.Printf(
		"\r\n<dim>%c%c%c</dim> <bold>%s</bold> <dim>%c%s%c</dim>\n",
		borderTopLeft,
		borderHorizontal,
		borderRightT,
		heading,
		borderLeftT,
		safeRepeat(string(borderHorizontal), t.width-7-len([]rune(heading))),
		borderTopRight,
	)
	t.printBlank()
}

func (t *terminalOutput) printDivider(heading string) {
	t.printBlank()
	_ = tml.Printf(
		"<dim>%c%c%c</dim> <bold>%s</bold> <dim>%c%s%c</dim>\n",
		borderLeftT,
		borderHorizontal,
		borderRightT,
		heading,
		borderLeftT,
		safeRepeat(string(borderHorizontal), t.width-7-len([]rune(heading))),
		borderRightT,
	)
	t.printBlank()
}

func (t *terminalOutput) printFooter() {
	t.printBlank()
	_ = tml.Printf("<dim>%c%s%c</dim>\n", borderBottomLeft, safeRepeat(string(borderHorizontal), t.width-2), borderBottomRight)
}

func (t *terminalOutput) printBlank() {
	t.printIn(0, "")
}

var rgxHtml = regexp.MustCompile(`<[^>]+>`)

func stripTags(input string) string {
	return rgxHtml.ReplaceAllString(input, "")
}

func (t *terminalOutput) limitSize(input string, size int) []string {
	var word string
	var words []string
	var inTag bool
	for _, r := range []rune(input) {
		if inTag {
			word += string(r)
			if r == '>' {
				inTag = false
				if word != "" {
					words = append(words, word)
				}
				word = ""
			}
		} else {
			if r == '<' {
				if word != "" {
					words = append(words, word)
				}
				word = "<"
				inTag = true
			} else if r == ' ' {
				if word != "" {
					words = append(words, word)
					word = ""
				}
			} else if r == '\n' {
				if word != "" {
					words = append(words, word)
					word = ""
				}
				words = append(words, "\n")
			} else {
				word += string(r)
			}
		}
	}
	if word != "" {
		words = append(words, word)
	}

	var line string
	var currentSize int
	var lines []string
	var hasContent bool

	for _, word := range words {
		if word == "\n" {
			lines = append(lines, line)
			line = ""
			continue
		}
		if word[0] == '<' {
			line += word
			continue
		}
		if currentSize+len([]rune(word))+1 > size {
			lines = append(lines, line)
			line = word
			hasContent = true
			currentSize = len([]rune(word))
		} else {
			if line != "" && hasContent {
				line += " "
				currentSize++
			}
			line += word
			hasContent = true
			currentSize += len([]rune(word))
		}
	}

	if line != "" {
		lines = append(lines, line)
	}

	if len(lines) == 0 {
		return []string{""}
	}

	return lines
}

func (t *terminalOutput) printIn(indent int, format string, args ...interface{}) {

	lines := t.limitSize(fmt.Sprintf(format, args...), t.width-indent-4)

	for _, line := range lines {
		realStr := stripTags(line)
		repeat := t.width - 4 - indent - len([]rune(realStr))
		padded := tml.Sprintf(line) + safeRepeat(" ", repeat)
		_ = tml.Printf("<dim>%c</dim> %s", borderVertical, safeRepeat(" ", indent))
		fmt.Printf("%s", padded)
		_ = tml.Printf(" <dim>%c</dim>\n", borderVertical)
	}
}

func (t *terminalOutput) Output(rep report.Report) error {

	for i, section := range rep.Sections() {
		if i == 0 {
			t.printHeader(section.Heading())
		} else {
			t.printDivider(section.Heading())
		}
		var maxKeyLen int
		for _, keyVal := range section.KeyValues() {
			if len(keyVal.Key()) > maxKeyLen {
				maxKeyLen = len(keyVal.Key())
			}
		}
		for _, keyVal := range section.KeyValues() {
			t.printIn(1+maxKeyLen-len(keyVal.Key()), "%s  <blue>%s</blue>", keyVal.Key(), keyVal.Value())
		}
	}

	t.printDivider("Security")
	for i, test := range rep.ResultsTable().Tests() {

		if i > 0 {
			t.printBlank()
		}

		switch test.Result {
		case report.Pass:
			t.printIn(0, "<green>✔ %s", test.Name)
		case report.PartialFail:
			t.printIn(0, "<yellow>⚠ %s", test.Name)
		case report.Fail:
			t.printIn(0, "<red>× %s", test.Name)
		}
		t.printIn(2, test.Description)
	}
	t.printFooter()

	return nil
}
