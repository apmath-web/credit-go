package valueObjects

type MessageInterface interface {
	Message(field string, text string) error
	GetField() string
	GetText() string
}
