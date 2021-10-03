package trade

type Flow interface {
	Execute(*Trade) error
	SetNext(Flow)
}
