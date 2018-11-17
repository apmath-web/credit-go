package valueObjects

type ValidationInterface interface {
	AddMessage(messages MessageInterface)
	GetMessages() []MessageInterface
}
