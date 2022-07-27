package state

type State struct {
	Watermark string
	BatchSeq  int
}

func (s *State) NextBatch() {}
