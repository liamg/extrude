package hardening

import "debug/elf"

func checkRELRO(e *elf.File) bool {
	for _, prog := range e.Progs {
		switch prog.Type {
		case elf.PT_GNU_RELRO:
			return true
		}
	}
	return false
}
