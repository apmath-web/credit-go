package valueObjects

type ValidationInterface interface {
	ValidationInterface(message MessageInterface) error
	getMessages() ([]MessageInterface, error)
}
