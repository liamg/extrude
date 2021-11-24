package compiler

import (
	"fmt"
	"strings"
)

type Info struct {
	Compiler Compiler
	Language Language
	Version  string
}

func (c Info) String() string {
	if c.Compiler == CompilerUnknown {
		return "unknown"
	}
	return strings.TrimSpace(fmt.Sprintf("%s %s", c.Compiler, c.Version))
}
