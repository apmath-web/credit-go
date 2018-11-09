package valueObjects

type MessageInterface interface {
	Message(field string, text string) MessageInterface
	GetField() string
	GetText() string
}
