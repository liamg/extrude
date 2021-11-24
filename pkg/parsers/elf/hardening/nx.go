package hardening

import "debug/elf"

func checkNonExecutableStackProgHeader(e *elf.File) bool {
	for _, prog := range e.Progs {
		if prog.Type == elf.PT_GNU_STACK {
			if prog.Flags&elf.PF_X > 0 {
				return false
			}
		}
	}
	return true
}
