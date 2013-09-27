// chrisguitarguy/disptach
// Copyright: 2013 Christopher Davis <http://christopherdavis.me>
// License: MIT

package dispatch

import (
	"testing"
	"reflect"
)

func TestGetName(t *testing.T) {
	e := NewEvent("an_event");

	if "an_event" != e.GetName() {
		t.Errorf(`e.GetName should have returned "an_event", got "%s" instead`, e.GetName())
	}
}

func TestPropogation(t *testing.T) {
	e := NewEvent("an_event");

	if e.IsStopped() {
		t.Error("Event propogation should not be stopped by default")
	}

	e.StopPropogation()

	if !e.IsStopped() {
		t.Error("Expected event propogation to be called after Event.StopPropogation call")
	}
}

func TestGetSetContextValues(t *testing.T) {
	var val interface{}
	var ok bool

	e := NewEvent("an_event")

	val, ok = e.Get("val")
	if ok {
		t.Error(`Key "val" should not be set in event context`)
	}

	e.Set("val", 10)

	val, ok = e.Get("val")
	if !ok {
		t.Error(`Should have set "val" in event context`)
	}

	if val != 10 {
		t.Error(`Should have gotten back "10" from event context key "val"`)
	}
}

func TestGetSetContext(t *testing.T) {
	e := NewEvent("an_event")
	ctx := make(EventContext)
	ctx["key"] = 123

	if reflect.DeepEqual(e.GetContext(), ctx) {
		t.Error("Event context and new context shouldn't be equal yet!")
	}

	e.SetContext(ctx)

	if !reflect.DeepEqual(e.GetContext(), ctx) {
		t.Error("Event contexts should be equal")
	}
}
