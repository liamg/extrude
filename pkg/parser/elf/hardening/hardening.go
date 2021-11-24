package hardening

import (
	"debug/elf"
)

type Attributes struct {
	StackProtected           bool
	FortifySourceFunctions   FortifySourceFunctions
	PositionIndependent      bool
	ReadOnlyRelocations      bool
	ImmediateBinding         bool
	NonExecutableStackHeader bool
}

type FortifySourceFunctions struct {
	Total     int
	Fortified int
}

func IdentifyAttributes(e *elf.File) Attributes {
	return Attributes{
		StackProtected:           checkStackProtected(e),
		FortifySourceFunctions:   checkSourceFortified(e),
		PositionIndependent:      checkPIE(e),
		ReadOnlyRelocations:      checkRELRO(e),
		ImmediateBinding:         checkImmediateBinding(e),
		NonExecutableStackHeader: checkNonExecutableStackProgHeader(e),
	}
}

// TODO: see https://github.com/nya3jp/tast-tests/blob/9fd02c2b27c3d2ec52299a95bc4b26a7e662b034/src/chromiumos/tast/local/bundles/cros/security/toolchain/verify.go#L22
