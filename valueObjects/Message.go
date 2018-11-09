package valueObjects

type Message struct {
	field, text string
}

func (m Message) Message(field string, text string) MessageInterface {
	m.field, m.text = field, text
	return m
}
func (m Message) GetField() string {
	return m.field
}
func (m Message) GetText() string {
	return m.text
}
