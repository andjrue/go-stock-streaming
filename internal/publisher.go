package internal

type Publisher struct {
	subscribers []Subscriber
}

func NewPublisher() *Publisher {
	return &Publisher{
		subscribers: make([]Subscriber, 0),
	}
}

func (p *Publisher) Publish(update QuoteUpdate) {
	for _, sub := range p.subscribers {
		select {
		case sub <- update:
		default:
		}
	}
}
