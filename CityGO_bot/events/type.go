package events

import commandshistory "graduation-project/CityGO_bot/lib/commandsHistory"

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event, commandHistory *[]commandshistory.CommandHistoryItem) error
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
	Meta interface{}
}
