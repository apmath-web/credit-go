package valueObjects

type ValidationInterface interface {
	AddMessages(messages []MessageInterface)
	GetMessages() []MessageInterface
}
