package originator

type OriginatorType struct {
	state string
}

// Set sets the state of the originator and is publically accessible
func (o *OriginatorType) Set(s string) {
	o.state = s
}

// Get returns the state of the originator
func (o *OriginatorType) Get() string {
	return o.state
}

// SetFromMemento sets the state of the originator from a memento
func (o *OriginatorType) SetFromMemento(m MementoType) {
	o.state = m.getState()
}

// CreateMemento creates a new memento with the state of the
func (o *OriginatorType) CreateMemento() MementoType {
	m := MementoType{state: o.state}
	return m
}

// MementoType is the Memento of the application
type MementoType struct {
	state string
}

// getState is used by the Originator to restore state
func (m *MementoType) getState() string {
	return m.state
}

// setState is used by the Originator to initialize the memento with the
// object state
func (m *MementoType) setState(s string) MementoType {
	m.state = s
	return *m
}
