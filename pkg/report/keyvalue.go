package report

type KeyValue [2]string

func (k KeyValue) Key() string {
	return k[0]
}

func (k KeyValue) Value() string {
	return k[1]
}
