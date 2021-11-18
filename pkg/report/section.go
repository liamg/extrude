package report

type Section interface {
	Heading() string
	KeyValues() []KeyValue
	AddKeyValue(key, value string)
}

func NewSection(heading string) Section {
	return &section{
		heading: heading,
	}
}

type section struct {
	heading   string
	keyValues []KeyValue
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
