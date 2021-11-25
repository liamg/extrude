package hardening

import (
	"debug/macho"
)

func checkAutomaticReferenceCounting(f *macho.File) bool {
	symbols, err := f.ImportedSymbols()
	if err != nil {
		return false
	}
	for _, imp := range symbols {
		if imp == "_objc_release" {
			return true
		}
	}
	return false
}
