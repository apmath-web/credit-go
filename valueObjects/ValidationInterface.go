package valueObjects

type ValidationInterface interface {
	AddMessages(message []MessageInterface) error
	GetMessages() []MessageInterface
}
