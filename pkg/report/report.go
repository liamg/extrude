package report

type Report interface {
	Sections() []Section
	AddSection(Section)
	Status() Result
}

type Reporter interface {
	CreateReport() (Report, error)
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

func (r *report) Status() Result {
	status := Pass
	for _, section := range r.sections {
		for _, test := range section.Tests() {
			switch test.Result {
			case Fail:
				return Fail
			case Warning:
				status = Warning
			}
		}
	}
	return status
}
