package trade

type Flow interface {
	Execute(*Trade)
	SetNext(Flow)
}
