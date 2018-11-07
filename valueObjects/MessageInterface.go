package valueObjects

type MessageInterface interface {
	Message(field string, text string) error
	getField() string
	getText() string
}
