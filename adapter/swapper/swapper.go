package swapper

type Swapper interface {
	OpenOrder() error
	Status() error
	Settled() bool
}
