package hardening

import (
	"debug/macho"
)

func checkStackProtected(f *macho.File) bool {
	symbols, err := f.ImportedSymbols()
	if err != nil {
		return false
	}
	for _, imp := range symbols {
		if imp == "___stack_chk_fail" || imp == "___stack_chk_guard" {
			return true
		}
	}
	return false
}
