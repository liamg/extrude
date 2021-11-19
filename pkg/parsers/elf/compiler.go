package elf

import (
	"fmt"
	"regexp"
	"strings"
)

type Compiler struct {
	Name    string
	Version string
}

func (c Compiler) String() string {
	if c.Name == "" {
		return "unknown"
	}
	if c.Version == "" {
		return c.Name
	}
	return fmt.Sprintf("%s %s", c.Name, c.Version)
}

var (
	rgxGCC = regexp.MustCompile(`GCC: \(GNU\) [\d\.]+`)
	rgxGo  = regexp.MustCompile(`go\d+\.\d+(\.\d+)?`)
)

func (m *Metadata) findCompiler() {

	m.Compiler.Name = "unknown"
	m.Compiler.Version = "unknown"

	// Go
	// example: "go1.17"
	if m.ELF.Section(".gosymtab") != nil {
		m.Compiler.Name = "Go"
		if roData := m.ELF.Section(".rodata"); roData != nil {
			if raw, err := roData.Data(); err == nil {
				goVersion := rgxGo.Find(raw)
				if goVersion != nil {
					m.Compiler.Version = string(goVersion[2:])
				}
			}
		}
		return
	}

	// .comment
	if commentData := m.ELF.Section(".comment"); commentData != nil {
		if raw, err := commentData.Data(); err == nil {

			// gcc
			// example: "GCC: (GNU) 4.8.5"
			if gccVersion := rgxGCC.Find(raw); gccVersion != nil {
				m.Compiler.Name = "gcc"
				m.Compiler.Version = strings.Split(string(raw), " ")[2]
				return
			}
		}
	}

}
