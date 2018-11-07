package valueObjects

type ValidationInterface interface {
	Validation(message []MessageInterface) error
	getMessages() []MessageInterface
}
