package valueObjects

type Message struct {
	field, text string
}

func (m *Message) Message(field string, text string) {
	m.field, m.text = field, text
}
func (m *Message) GetField() string {
	return m.field
}
func (m *Message) GetText() string {
	return m.text
}
