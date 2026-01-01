package messaging

type Subscriber interface {
	Subscribe(topic string, handler func([]byte)) error
}
