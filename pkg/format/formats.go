package format

import "fmt"

type Format uint8

const (
	Unknown Format = iota
	ELF
	MachO
	PE
)

func (f Format) Short() string {
	switch f {
	case ELF:
		return "ELF"
	case MachO:
		return "Mach-O"
	case PE:
		return "PE"
	}
	return fmt.Sprintf("0x%x", uint8(f))
}

func (f Format) Long() string {
	switch f {
	case ELF:
		return "Executable and Linkable Format"
	case MachO:
		return "Mach Object File"
	case PE:
		return "Portable Executable"
	}
	return "unknown"
}

func (f Format) String() string {
	return fmt.Sprintf("%s (%s)", f.Long(), f.Short())
}
