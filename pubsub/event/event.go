package event

type Event interface {
	// Identifier returns the ID of pubsub
	Identifier() string

	// Name returns pubsub name of current pubsub
	Name() string

	// Payload returns pubsub payload of current pubsub
	Payload() interface{}

	// String convert pubsub data to string
	String() (string, error)
}
