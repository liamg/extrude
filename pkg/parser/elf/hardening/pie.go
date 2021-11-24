package hardening

import "debug/elf"

func checkPIE(e *elf.File) bool {
	for _, prog := range e.Progs {
		switch prog.Type {
		case elf.PT_DYNAMIC:
			return true
		}
	}
	return false
}
