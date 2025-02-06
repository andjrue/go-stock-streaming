package internal

type QuoteUpdate struct {
	Symbol    string
	Price     float32
	OpenPrice float32
}

type Subscriber chan QuoteUpdate

func (p *Publisher) Subscribe() Subscriber {
	ch := make(Subscriber, 10) // Need to add buffer
	p.subscribers = append(p.subscribers, ch)
	return ch
}
