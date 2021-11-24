package compiler

type Compiler string

const (
	CompilerUnknown Compiler = ""
	CompilerGCC     Compiler = "gcc"
	CompilerGHC     Compiler = "ghc"
	CompilerGo      Compiler = "go"
	CompilerRustC   Compiler = "rustc"
	CompilerOCaml   Compiler = "ocaml"
	CompilerNim     Compiler = "nim"
	CompilerTCC     Compiler = "tcc"
)

func (c Compiler) String() string {
	if c == CompilerUnknown {
		return "unknown"
	}
	return string(c)
}
