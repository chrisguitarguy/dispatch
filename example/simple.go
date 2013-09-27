// chrisguitarguy/disptach
// Copyright: 2013 Christopher Davis <http://christopherdavis.me>
// License: MIT

package main

import (
	"dispatch"
	"fmt"
)

func main() {
	d := dispatch.NewDispatcher()
	d.AddListener("an_event", dispatch.NewListener(10, func (e dispatch.Event, d dispatch.Dispatcher) {
		e.Set("val", 10)
	}))

	e := dispatch.NewEvent("a_name")

	d.Dispatch("an_event", e)

	val, _ := e.Get("val")

	fmt.Println(val)
}
