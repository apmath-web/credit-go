package valueObjects

type Validation struct {
	messages []MessageInterface
}

func (v *Validation) AddMessage(message MessageInterface) {
	v.messages = append(v.messages, message)
}

func (v *Validation) GetMessages() []MessageInterface {
	return v.messages
}
