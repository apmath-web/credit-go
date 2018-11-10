package valueObjects

type Validation struct {
	messages []MessageInterface
}

func (v *Validation) AddMessages(messages []MessageInterface) {
	v.messages = append(v.messages, messages...)
}
func (v *Validation) GetMessages() []MessageInterface {
	return v.messages
}
