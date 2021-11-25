package hardening

import "debug/macho"

type Attributes struct {
	init                          bool
	PositionIndependentExecutable bool
	StackExecutionNotAllowed      bool
	HeapExecutionNotAllowed       bool
	StackProtected                bool
	AutomaticReferenceCounting    bool
	Encrypted                     bool
}

func (a Attributes) Merge(b Attributes) Attributes {
	if !a.init {
		return b
	}
	if !b.init {
		return a
	}
	return Attributes{
		PositionIndependentExecutable: a.PositionIndependentExecutable && b.PositionIndependentExecutable,
		StackExecutionNotAllowed:      a.StackExecutionNotAllowed && b.StackExecutionNotAllowed,
		HeapExecutionNotAllowed:       a.HeapExecutionNotAllowed && b.HeapExecutionNotAllowed,
		StackProtected:                a.StackProtected && b.StackProtected,
		AutomaticReferenceCounting:    a.AutomaticReferenceCounting && b.AutomaticReferenceCounting,
		Encrypted:                     a.Encrypted && b.Encrypted,
	}
}

func IdentifyAttributes(f *macho.File) Attributes {
	return Attributes{
		init:                          true,
		PositionIndependentExecutable: f.Flags&macho.FlagPIE > 0,
		StackExecutionNotAllowed:      f.Flags&macho.FlagAllowStackExecution == 0,
		HeapExecutionNotAllowed:       f.Flags&macho.FlagNoHeapExecution > 0,
		StackProtected:                checkStackProtected(f),
		AutomaticReferenceCounting:    checkAutomaticReferenceCounting(f),
		Encrypted:                     checkEncrypted(f),
	}
}
