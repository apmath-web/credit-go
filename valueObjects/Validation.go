package valueObjects

type Validation struct {
	messages []MessageInterface
}

func (v *Validation) AddMessage(messages MessageInterface) {
	v.messages = append(v.messages, messages)
}

func (v *Validation) GetMessages() []MessageInterface {
	return v.messages
}
