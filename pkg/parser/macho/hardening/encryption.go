package hardening

import "debug/macho"

const (
	EncInfo32 = 0x21
	EncInfo64 = 0x2c
)

func checkEncrypted(f *macho.File) bool {
	return f.Symtab.Cmd&EncInfo32 > 0 || f.Symtab.Cmd&EncInfo64 > 0
}
