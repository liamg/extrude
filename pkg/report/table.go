package report

type Table struct {
	rows [][]string
}

func NewTable() *Table {
	return &Table{}
}

func (t *Table) AddRow(fields ...string) {
	t.rows = append(t.rows, fields)
}

func (t *Table) Rows() [][]string {
	return t.rows
}
