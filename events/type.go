package events

type Fetcher interface {
	Fecth(limit int) ([]Event, error)
}

type Proccessor interface {
	Process(e Event) error
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
}
