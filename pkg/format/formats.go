package format

import "fmt"

type Format uint8

const (
	Unknown Format = iota
	ELF
	MachO32
	MachO64
	PE
)

func (f Format) Short() string {
	switch f {
	case ELF:
		return "ELF"
	case MachO32:
		return "Mach-O 32"
	case MachO64:
		return "Mach-O 64"
	case PE:
		return "PE"
	}
	return fmt.Sprintf("0x%x", uint8(f))
}

func (f Format) Long() string {
	switch f {
	case ELF:
		return "Executable and Linkable Format"
	case MachO32:
		return "32-bit Mach Object File"
	case MachO64:
		return "64-bit Mach Object File"
	case PE:
		return "Portable Executable"
	}
	return "unknown"
}

func (f Format) String() string {
	return fmt.Sprintf("%s (%s)", f.Long(), f.Short())
}
