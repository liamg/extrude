package elf

import (
	"fmt"
	"regexp"
	"strings"
)

type Compiler struct {
	Name     string
	Version  string
	Language string
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
	rgxGCC          = regexp.MustCompile(`GCC \(GNU\) [\d\.]+`)
	rgxGHC          = regexp.MustCompile(`GHC \d+\.\d+(\.\d+)?`)
	rgxGo           = regexp.MustCompile(`go\d+\.\d+(\.\d+)?`)
	rgxRust         = regexp.MustCompile(`rustc version [^(]+`)
	rgxRustStripped = regexp.MustCompile(`rustc version \d+\.\d+(\.\d+)(\-[a-z]+)`)
	rgxOCaml        = regexp.MustCompile(`OCaml.*version \d+\.(\d+\.)?\d+`)
	rgxNim          = regexp.MustCompile(`system\.nim`)
)

func (m *Metadata) findCompiler() {

	m.Compiler.Name = "unknown"
	m.Compiler.Version = "unknown"
	m.Compiler.Language = "unknown"

	// Go
	// example: "go1.17"
	if m.ELF.Section(".gosymtab") != nil {
		m.Compiler.Name = "go"
		m.Compiler.Language = "Go"
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

	// Rust
	if debugData := m.ELF.Section(".debug_str"); debugData != nil {
		if raw, err := debugData.Data(); err == nil {
			if version := rgxRust.Find(raw); version != nil {
				m.Compiler.Name = "rustc"
				m.Compiler.Language = "Rust"
				m.Compiler.Version = strings.Split(string(version), " ")[2]
				return
			}
		}
	}

	if roData := m.ELF.Section(".rodata"); roData != nil {
		if raw, err := roData.Data(); err == nil {

			// Rust stripped
			if version := rgxRustStripped.Find(raw); version != nil {
				m.Compiler.Name = "rustc"
				m.Compiler.Version = string(version)
				m.Compiler.Language = "Rust"
				return
			}

			// OCaml
			if version := rgxOCaml.Find(raw); version != nil {
				m.Compiler.Name = "ocamlc"
				m.Compiler.Version = strings.Split(string(version), "version ")[1]
				m.Compiler.Language = "OCaml"
				return
			}

			// Nim
			if version := rgxNim.Find(raw); version != nil {
				m.Compiler.Name = "nim"
				m.Compiler.Language = "Nim"
				return
			}
		}
	}

	// .comment
	if commentData := m.ELF.Section(".comment"); commentData != nil {
		if raw, err := commentData.Data(); err == nil {

			// ghc
			if ghcVersion := rgxGHC.Find(raw); ghcVersion != nil {
				m.Compiler.Name = "ghc"
				m.Compiler.Language = "Haskell"
				m.Compiler.Version = strings.Split(string(ghcVersion), " ")[1]
				return
			}

			// gcc
			// example: "GCC: (GNU) 4.8.5"
			if gccVersion := rgxGCC.Find(raw); gccVersion != nil {
				m.Compiler.Name = "gcc"
				m.Compiler.Language = "C/C++ (likely)"
				m.Compiler.Version = strings.Split(string(gccVersion), " ")[2]
				return
			}
		}
	}

	if m.ELF.Section(".note.ABI-tag") != nil {
		m.Compiler.Name = "gcc"
		m.Compiler.Language = "C/C++ (likely)"
	} else if m.ELF.Section(".rodata.cst4") != nil {
		m.Compiler.Name = "tcc"
		m.Compiler.Language = "C/C++"
	}
}
