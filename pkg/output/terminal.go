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

var rgxHtml = regexp.MustCompile(`<[^>]+>`)

func stripTags(input string) string {
	return rgxHtml.ReplaceAllString(input, "")
}

func (t *terminalOutput) limitSize(input string, size int) []string {
	var word string
	var words []string
	var inTag bool
	var lastRune rune
	var hasClosing bool
	for _, r := range []rune(input) {
		if inTag {
			word += string(r)
			if hasClosing && r == '>' {
				inTag = false
				hasClosing = false
				words = append(words, word)
				word = ""
			} else if lastRune == '<' && r == '/' {
				hasClosing = true
			}
		} else {
			if r == '<' {
				if word != "" {
					words = append(words, word)
				}
				word = "<"
				inTag = true
			} else if r == ' ' {
				words = append(words, word)
			} else {
				word += string(r)
			}
		}
		lastRune = r
	}
	if word != "" {
		words = append(words, word)
	}

	var line string
	var currentSize int
	var lines []string

	for _, word := range words {
		if word[0] == '<' {
			line += word
			continue
		}
		if currentSize+len(word) > size {
			lines = append(lines, line)
			line = word
			currentSize = len(word)
		} else {
			if line != "" {
				line += " "
				currentSize++
			}
			line += word
			currentSize += len(word)
		}
	}

	if line != "" {
		lines = append(lines, line)
	}

	return lines
}

func (t *terminalOutput) printIn(format string, args ...interface{}) {

	realStr := stripTags(fmt.Sprintf(format, args...))
	str := tml.Sprintf(format, args...)
	repeat := t.width - 4 - len([]rune(realStr))
	padded := str
	if repeat > 0 {
		padded = str + strings.Repeat(" ", repeat)
	}
	_ = tml.Printf("<dim>%c</dim> ", borderVertical)
	fmt.Printf("%s", padded)
	_ = tml.Printf(" <dim>%c</dim>\n", borderVertical)
}

func (t *terminalOutput) Output(rep report.Report) error {

	for _, section := range rep.Sections() {
		t.printHeader(section.Heading())
		var maxKeyLen int
		for _, keyVal := range section.KeyValues() {
			if len(keyVal.Key()) > maxKeyLen {
				maxKeyLen = len(keyVal.Key())
			}
		}
		for _, keyVal := range section.KeyValues() {
			paddedKey := strings.Repeat(" ", 1+maxKeyLen-len(keyVal.Key())) + keyVal.Key()
			t.printIn("%s  <blue>%s</blue>", paddedKey, keyVal.Value())
		}
		t.printFooter()
	}

	t.printHeader("Security")
	for i, test := range rep.ResultsTable().Tests() {

		if i > 0 {
			t.printBlank()
		}

		switch test.Result {
		case report.Pass:
			t.printIn("<green>✔ %s", test.Name)
		case report.PartialFail:
			t.printIn("<yellow>⚠ %s", test.Name)
		case report.Fail:
			t.printIn("<red>× %s", test.Name)
		}
		t.printIn("  FUCK")
	}
	t.printFooter()

	return nil
}
