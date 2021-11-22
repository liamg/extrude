package report

type Test struct {
	Name        string
	Result      Result
	Description string
}

type Table struct {
	tests []Test
}

type Result uint8

const (
	Fail Result = iota
	PartialFail
	Pass
)

func NewTable() *Table {
	return &Table{}
}

func (t *Table) AddTest(name string, result Result, description string) {
	t.tests = append(t.tests, Test{
		Name:        name,
		Result:      result,
		Description: description,
	})
}

func (t *Table) Tests() []Test {
	return t.tests
}
