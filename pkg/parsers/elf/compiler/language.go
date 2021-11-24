package compiler

type Language string

const (
	LanguageUnknown Language = ""
	LanguageGo      Language = "Go"
	LanguageCCPP    Language = "C/C++"
	LanguageRust    Language = "Rust"
	LanguageHaskell Language = "Haskell"
	LanguageOCaml   Language = "OCaml"
	LanguageNim     Language = "Nim"
)

func (l Language) String() string {
	if l == LanguageUnknown {
		return "unknown"
	}
	return string(l)
}
