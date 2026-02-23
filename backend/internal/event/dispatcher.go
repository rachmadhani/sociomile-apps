package event

import "errors"

type Dispatcher struct {
	channel chan Event
}

func NewDispatcher(buffer int) *Dispatcher {
	return &Dispatcher{
		channel: make(chan Event, buffer),
	}
}

func (d *Dispatcher) Dispatch(event Event) error {
	select {
	case d.channel <- event:
		return nil
	default:
		return errors.New("channel is full")
	}
}

func (d *Dispatcher) Channel() <-chan Event {
	return d.channel
}
