package pubsub

type Subscriber struct {
	ps      *Pubsub
	message chan []byte
}

func NewSubscriber(ps *Pubsub) *Subscriber {
	sub := &Subscriber{
		ps:      ps,
		message: make(chan []byte, 1024),
	}
	sub.ps.register <- sub
	return sub
}

func (s *Subscriber) Read() <-chan []byte {
	return s.message
}
