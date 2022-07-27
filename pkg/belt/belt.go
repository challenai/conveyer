package belt

import "github.com/challenai/conveyer/pkg/codec"

type Belt struct {
	// default capacity is 1.
	capacity int
	// channel for bytes
	ch chan [][]codec.Bytes
	// tell if the channel is half-closed.
	timeWaitPeriod bool
}

func NewBelt(capacity int) *Belt {
	return &Belt{
		capacity:       capacity,
		ch:             make(chan [][]codec.Bytes, capacity),
		timeWaitPeriod: false,
	}
}

func (b *Belt) Deliever(by [][]codec.Bytes) error {
	b.ch <- by
	return nil
}

func (b *Belt) Metrics() (Stat, error) {
	return 0, nil
}

func (b *Belt) Accept() [][]codec.Bytes {
	return <-b.ch
}

// func (b *Belt) Next() {
// }

func (b *Belt) Close() {
	b.timeWaitPeriod = true
}

func (b *Belt) HasNext() bool {
	return b.timeWaitPeriod
}
