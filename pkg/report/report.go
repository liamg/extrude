package report

type Report interface {
	Sections() []Section
	AddSection(Section)
}

type report struct {
	sections []Section
}

func New() Report {
	return &report{}
}

func (r *report) Sections() []Section {
	return r.sections
}

func (r *report) AddSection(s Section) {
	r.sections = append(r.sections, s)
}
