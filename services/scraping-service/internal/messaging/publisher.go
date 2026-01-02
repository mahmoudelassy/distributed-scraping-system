package messaging

type Publisher interface {
	Publish(topic string, message []byte) error
}
