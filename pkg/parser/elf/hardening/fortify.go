package hardening

import (
	"debug/elf"
	"fmt"
)

func checkSourceFortified(e *elf.File) (info FortifySourceFunctions) {
	symbols, _ := e.Symbols()
	dynSymbols, _ := e.DynamicSymbols()

	for _, symbol := range append(symbols, dynSymbols...) {
		for _, libcFunc := range libcFunctions {
			if fmt.Sprintf("__%s_chk", libcFunc) == symbol.Name {
				info.Fortified++
			} else if symbol.Name == libcFunc {
				info.Total++
			}
		}
	}

	return
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
