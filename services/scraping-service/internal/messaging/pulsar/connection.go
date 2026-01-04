package pulsar

import (
	"github.com/apache/pulsar-client-go/pulsar"
)

type Pulsar struct {
	client pulsar.Client
}

func NewPulsar(url string) (*Pulsar, error) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: url,
	})
	if err != nil {
		return nil, err
	}

	return &Pulsar{
		client: client,
	}, nil
}

func (p *Pulsar) Close() {
	p.client.Close()
}
