package ports

type MessageService interface {
	PublishMessage(message string) error
	PublishDelayedMessage(message string) error
}
