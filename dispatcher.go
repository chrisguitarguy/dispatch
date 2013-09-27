// chrisguitarguy/disptach
// Copyright: 2013 Christopher Davis <http://christopherdavis.me>
// License: MIT

package dispatch

import "sort"

type ListenerFunc func(Event, Dispatcher)

type Listener struct {
	Priority int
	Callback ListenerFunc
}

type Dispatcher interface {
	AddListener(eventName string, listener *Listener)
	HasListeners(eventName string) bool
	//RemoveListener(eventName string, listener Listener) error
	Dispatch(name string, event Event)
}

type ListenerSlice []*Listener

func (s ListenerSlice) Len() int {
	return len(s)
}

func (s ListenerSlice) Less(i, j int) bool {
	return s[i].Priority > s[j].Priority
}

func (s ListenerSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type DefaultDispatcher struct {
	listeners map[string]ListenerSlice
}

func NewDispatcher() Dispatcher {
	return &DefaultDispatcher{make(map[string]ListenerSlice)}
}

func NewListener(prio int, fn ListenerFunc) *Listener {
	return &Listener{prio, fn}
}

func (d *DefaultDispatcher) AddListener(eventName string, listener *Listener) {
	_, ok := d.listeners[eventName];

	if !ok {
		d.listeners[eventName] = make(ListenerSlice, 0, 8) // XXX maybe 8 is too small?
	}

	d.listeners[eventName] = append(d.listeners[eventName], listener)
}

func (d *DefaultDispatcher) HasListeners(eventName string) bool {
	l, ok := d.listeners[eventName]
	return ok && len(l) >= 1
}

func (d *DefaultDispatcher) Dispatch(eventName string, event Event) {
	listeners, ok := d.listeners[eventName];
	if !ok {
		return
	}

	sort.Sort(listeners)

	for _, l := range listeners {
		l.Callback(event, d)
		if event.IsStopped() {
			break
		}
	}
}
