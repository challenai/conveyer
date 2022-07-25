package belt

import "github.com/challenai/conveyer/pkg/codec"

type Belt struct {
	// default capacity is 1.
	capacity int
	ch       chan [][]codec.Bytes
	// tell if the channel is half-closed.
	timeWait bool
}

func (b *Belt) Deliever([][]codec.Bytes) error {
	return nil
}

func (b *Belt) Metrics(dsl string) (Stat, error) {
	return 0, nil
}

func (b *Belt) Accept() []codec.Bytes {
	return nil
}

func (b *Belt) Next() {

}

func (b *Belt) HasNext() bool {
	return false
}
