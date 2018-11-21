package valueObjects

type ValidationInterface interface {
	AddMessage(message MessageInterface)
	GetMessages() []MessageInterface
}
