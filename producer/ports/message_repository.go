package ports

type MessageRepository interface {
	PublishMessage(message string) error
	PublishDelayedMessage(message string) error
}
