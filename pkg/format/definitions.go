package format

type definition struct {
	format     Format
	signatures [][]byte
}

var definitions = []definition{
	{
		format: ELF,
		signatures: [][]byte{
			{0x7f, 'E', 'L', 'F'},
		},
	},
	{
		format: PE,
		signatures: [][]byte{
			{'M', 'Z'},
			{'Z', 'M'},
		},
	},
	{
		format: MachO32,
		signatures: [][]byte{
			{0xfe, 0xed, 0xfa, 0xce},
			{0xce, 0xfa, 0xed, 0xfe},
		},
	},
	{
		format: MachO64,
		signatures: [][]byte{
			{0xfe, 0xed, 0xfa, 0xcf},
			{0xcf, 0xfa, 0xed, 0xfe},
		},
	},
}
