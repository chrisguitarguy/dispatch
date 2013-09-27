// chrisguitarguy/disptach
// Copyright: 2013 Christopher Davis <http://christopherdavis.me>
// License: MIT

package dispatch

type EventContext map[string]interface{}

// Events are used by the dispatcher. They're just a container for
// "stuff" -- whatever your listeners
type Event interface {
	// Get the name of the event. Just a string, whatever you like. This is not
	// used internally, but may be helpful in listener
	GetName() string

	// Stop the propotation of the event. If a listener calls this, no futher
	// listeners will be notified.
	StopPropogation()

	// Checks whether not the event can continue. Listeners usual have no use
	// for this, but it's used by the dispatcher.
	IsStopped() bool

	// Set an value in the event context
	Set(string, interface{})

	// Get a value from the event context
	Get(string) (interface{}, bool)

	// Replace the entire event context
	SetContext(ctx EventContext)

	// Get the entire event context
	GetContext() EventContext
}

// A default event type that implements Event (see above)
type DefaultEvent struct {
	name string
	stopped bool
	ctx EventContext
}

// Make a new event
func NewEvent(name string) Event {
	return &DefaultEvent{name, false, make(EventContext)}
}

func (e *DefaultEvent) GetName() string {
	return e.name
}

func (e *DefaultEvent) StopPropogation() {
	e.stopped = true
}

func (e *DefaultEvent) IsStopped() bool {
	return e.stopped
}

func (e *DefaultEvent) Set(key string, val interface{}) {
	e.ctx[key] = val
}

func (e *DefaultEvent) Get(key string) (val interface{}, ok bool) {
	val, ok = e.ctx[key]
	return
}

func (e *DefaultEvent) SetContext(ctx EventContext) {
	e.ctx = ctx
}

func (e *DefaultEvent) GetContext() EventContext {
	return e.ctx
}
