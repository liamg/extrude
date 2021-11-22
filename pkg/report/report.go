package report

type Report interface {
	Sections() []Section
	ResultsTable() *Table
	AddSection(Section)
	SetResultsTable(t *Table)
}

type Reporter interface {
	CreateReport() (Report, error)
}

type report struct {
	sections []Section
	table    Table
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

func (r *report) SetResultsTable(t *Table) {
	r.table = *t
}

func (r *report) ResultsTable() *Table {
	return &r.table
}
