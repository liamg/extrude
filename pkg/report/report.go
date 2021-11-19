package report

type Report interface {
	Sections() []Section
	Issues() []Issue
	AddSection(Section)
	AddIssue(Issue)
}

type Reporter interface {
	CreateReport() (Report, error)
}

type report struct {
	sections []Section
	issues   []Issue
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

func (r *report) Issues() []Issue {
	return r.issues
}

func (r *report) AddIssue(i Issue) {
	r.issues = append(r.issues, i)
}
