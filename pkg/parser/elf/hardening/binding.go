package hardening

import "debug/elf"

func checkImmediateBinding(e *elf.File) bool {
	if dynamic := e.SectionByType(elf.SHT_DYNAMIC); dynamic != nil {
		if dynData, err := dynamic.Data(); err == nil {
			inc := 16
			if e.Class == elf.ELFCLASS32 {
				inc = 8
			}
			var tag elf.DynTag
			var value uint64
			for i := 0; i+inc < len(dynData); i += inc {
				switch e.Class {
				case elf.ELFCLASS32:
					tag = elf.DynTag(e.ByteOrder.Uint32(dynData[i : i+4]))
					value = uint64(e.ByteOrder.Uint32(dynData[i+4 : i+8]))
				case elf.ELFCLASS64:
					tag = elf.DynTag(e.ByteOrder.Uint64(dynData[i : i+8]))
					value = e.ByteOrder.Uint64(dynData[i+8 : i+16])
				}
				if tag == elf.DT_FLAGS {
					if elf.DynTag(value)&elf.DT_BIND_NOW > 0 {
						return true
					}
				}
			}
		}
	}

	return false

}
