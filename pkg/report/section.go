package report

type Section interface {
	Heading() string
	KeyValues() []KeyValue
	AddKeyValue(key, value string)
	AddTest(name string, result Result, description string)
	Tests() []Test
}

type Test struct {
	Name        string
	Result      Result
	Description string
}

func NewSection(heading string) Section {
	return &section{
		heading: heading,
	}
}

type section struct {
	heading   string
	keyValues []KeyValue
	tests     []Test
}

type Result uint8

const (
	Fail Result = iota
	Warning
	Pass
)

func (s *section) AddTest(name string, result Result, description string) {
	s.tests = append(s.tests, Test{
		Name:        name,
		Result:      result,
		Description: description,
	})
}

func (s *section) Tests() []Test {
	return s.tests
}
func (s *section) Heading() string {
	return s.heading
}

func (s *section) KeyValues() []KeyValue {
	return s.keyValues
}

func (s *section) AddKeyValue(key, value string) {
	s.keyValues = append(s.keyValues, KeyValue{key, value})
}
