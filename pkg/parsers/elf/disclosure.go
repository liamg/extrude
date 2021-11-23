package elf

import (
	"debug/elf"
	"fmt"
	"regexp"
)

var rgxHomeDir = regexp.MustCompile(`/home/[A-Za-z0-9\-_\\.]+`)

func (m *Metadata) checkDisclosure() {
	for _, sectionName := range []string{".rodata", ".gopclntab", ".strtab"} {
		if section := m.ELF.Section(sectionName); section != nil {
			m.checkSectionForDisclosure(section)
		}
	}
}

func (m *Metadata) checkSectionForDisclosure(section *elf.Section) {
	if data, err := section.Data(); err == nil {
		usernames := make(map[string]bool)
		matches := rgxHomeDir.FindAll(data, -1)
		for _, match := range matches {
			usernames[string(match[6:])] = true
		}
		for username := range usernames {
			m.Notes = append(m.Notes,
				Note{
					Heading: "Information Disclosure",
					Content: fmt.Sprintf("Username of compiling user <yellow>%s</yellow> discovered through reference to home directory.", username),
				},
			)
		}
	}
}
