package elf

import (
	"fmt"
	"regexp"
)

var rgxHomeDir = regexp.MustCompile(`/home/[A-Za-z0-9\-_\\.]+`)

func (m *Metadata) checkDisclosure() {
	if rodata := m.ELF.Section(".rodata"); rodata != nil {
		if data, err := rodata.Data(); err == nil {
			usernames := make(map[string]bool)
			matches := rgxHomeDir.FindAll(data, -1)
			for _, match := range matches {
				usernames[string(match[6:])] = true
			}
			for username := range usernames {
				m.Notes = append(m.Notes,
					Note{
						Heading: "Information Disclosure",
						Content: fmt.Sprintf("Username of compiling user <yellow>%s</yellow> discovered through mention of home directory.", username),
					},
				)
			}
		}
	}
}
