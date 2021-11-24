package compiler

import (
	"debug/elf"
	"regexp"
	"strings"
)

var (
	rgxGCC          = regexp.MustCompile(`GCC \(GNU\) [\d\.]+`)
	rgxGHC          = regexp.MustCompile(`GHC \d+\.\d+(\.\d+)?`)
	rgxGo           = regexp.MustCompile(`go\d+\.\d+(\.\d+)?`)
	rgxRust         = regexp.MustCompile(`rustc version [^(]+`)
	rgxRustStripped = regexp.MustCompile(`rustc version \d+\.\d+(\.\d+)(\-[a-z]+)`)
	rgxOCaml        = regexp.MustCompile(`OCaml.*version \d+\.(\d+\.)?\d+`)
	rgxNim          = regexp.MustCompile(`system\.nim`)
)

func Identify(e *elf.File) Info {

	// Go
	// example: "go1.17"
	if e.Section(".gosymtab") != nil {
		info := Info{
			Compiler: CompilerGo,
			Language: "Go",
		}
		if roData := e.Section(".rodata"); roData != nil {
			if raw, err := roData.Data(); err == nil {
				goVersion := rgxGo.Find(raw)
				if goVersion != nil {
					info.Version = string(goVersion[2:])
				}
			}
		}
		return info
	}

	// Rust
	if debugData := e.Section(".debug_str"); debugData != nil {
		if raw, err := debugData.Data(); err == nil {
			if version := rgxRust.Find(raw); version != nil {
				return Info{
					Compiler: CompilerRustC,
					Language: LanguageRust,
					Version:  strings.Split(string(version), " ")[2],
				}
			}
		}
	}

	if roData := e.Section(".rodata"); roData != nil {
		if raw, err := roData.Data(); err == nil {

			// Rust stripped
			if version := rgxRustStripped.Find(raw); version != nil {
				return Info{
					Compiler: CompilerRustC,
					Language: LanguageRust,
					Version:  string(version),
				}
			}

			// OCaml
			if version := rgxOCaml.Find(raw); version != nil {
				return Info{
					Compiler: CompilerOCaml,
					Language: LanguageOCaml,
					Version:  strings.Split(string(version), "version ")[1],
				}
			}

			// Nim
			if version := rgxNim.Find(raw); version != nil {
				return Info{
					Compiler: CompilerNim,
					Language: LanguageNim,
				}
			}
		}
	}

	// .comment
	if commentData := e.Section(".comment"); commentData != nil {
		if raw, err := commentData.Data(); err == nil {

			// ghc
			if ghcVersion := rgxGHC.Find(raw); ghcVersion != nil {
				return Info{
					Compiler: CompilerGHC,
					Language: LanguageHaskell,
					Version:  strings.Split(string(ghcVersion), " ")[1],
				}
			}

			// gcc
			// example: "GCC: (GNU) 4.8.5"
			if gccVersion := rgxGCC.Find(raw); gccVersion != nil {
				return Info{
					Compiler: CompilerGCC,
					Language: LanguageCCPP,
					Version:  strings.Split(string(gccVersion), " ")[2],
				}
			}
		}
	}

	if e.Section(".note.ABI-tag") != nil {
		return Info{
			Compiler: CompilerGCC,
			Language: LanguageCCPP,
		}
	} else if e.Section(".rodata.cst4") != nil {
		return Info{
			Compiler: CompilerTCC,
			Language: LanguageCCPP,
		}
	}

	return Info{}
}
