package hardening

import "debug/elf"

func checkStackProtected(e *elf.File) bool {
	symbols, _ := e.Symbols()
	dynSymbols, _ := e.DynamicSymbols()
	for _, symbol := range append(symbols, dynSymbols...) {
		if symbol.Name == "__stack_chk_fail" {
			return true
		}
	}
	return false
}
