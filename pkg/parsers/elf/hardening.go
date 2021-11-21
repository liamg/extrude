package elf

import "fmt"

type Hardening struct {
	StackProtected         bool
	FortifySourceFunctions bool
	PositionIndependent    bool
	ReadOnlyRelocations    bool
	ImmediateBinding       bool
}

var libcFunctions = []string{
	"asprintf",
	"confstr",
	"dprintf",
	"fgets",
	"fgets_unlocked",
	"fgetws",
	"fgetws_unlocked",
	"fprintf",
	"fread",
	"fread_unlocked",
	"fwprintf",
	"getcwd",
	"getdomainname",
	"getgroups",
	"gethostname",
	"getlogin_r",
	"gets",
	"getwd",
	"longjmp",
	"mbsnrtowcs",
	"mbsrtowcs",
	"mbstowcs",
	"memcpy",
	"memmove",
	"mempcpy",
	"memset",
	"obstack_printf",
	"obstack_vprintf",
	"pread64",
	"pread",
	"printf",
	"ptsname_r",
	"read",
	"readlink",
	"readlinkat",
	"realpath",
	"recv",
	"recvfrom",
	"snprintf",
	"sprintf",
	"stpcpy",
	"stpncpy",
	"strcat",
	"strcpy",
	"strncat",
	"strncpy",
	"swprintf",
	"syslog",
	"ttyname_r",
	"vasprintf",
	"vdprintf",
	"vfprintf",
	"vfwprintf",
	"vprintf",
	"vsnprintf",
	"vsprintf",
	"vswprintf",
	"vsyslog",
	"vwprintf",
	"wcpcpy",
	"wcpncpy",
	"wcrtomb",
	"wcscat",
	"wcscpy",
	"wcsncat",
	"wcsncpy",
	"wcsnrtombs",
	"wcsrtombs",
	"wcstombs",
	"wctomb",
	"wmemcpy",
	"wmemmove",
	"wmempcpy",
	"wmemset",
	"wprintf",
}

func (m *Metadata) checkHardened() {
	symbols, _ := m.ELF.Symbols()
	dynSymbols, _ := m.ELF.DynamicSymbols()

	var hasLibc bool
	var hasProtected bool

	for _, symbol := range append(symbols, dynSymbols...) {
		if symbol.Name == "__stack_chk_fail" {
			m.Hardening.StackProtected = true
		}
		for _, libcFunc := range libcFunctions {
			if fmt.Sprintf("__%s_chk", libcFunc) == symbol.Name {
				hasProtected = true
			} else if symbol.Name == libcFunc {
				hasLibc = true
			}
		}
	}

	m.Hardening.FortifySourceFunctions = !hasLibc || hasProtected
}
