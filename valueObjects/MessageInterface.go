package valueObjects

type MessageInterface interface {
	Message(field string, text string)
	GetField() string
	GetText() string
}
