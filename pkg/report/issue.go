package report

type Issue interface {
	Message() string
	Severity() Severity
}

type Severity uint8

const (
	SeverityLow Severity = iota
	SeverityMedium
	SeverityHigh
)

type issue struct {
	message  string
	severity Severity
}

func NewIssue(message string, severity Severity) Issue {
	return &issue{
		message:  message,
		severity: severity,
	}
}

func (i *issue) Message() string {
	return i.message
}

func (i *issue) Severity() Severity {
	return i.severity
}
