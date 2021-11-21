package elf

func (m *Metadata) analyse() error {

	m.findCompiler()

	m.checkHardened()

	return nil
}
