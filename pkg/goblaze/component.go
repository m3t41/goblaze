// Copyright 2026 Daniel\n// Licensed under the GNU Affero General Public License v3.0.\n// Copying or distributing this file requires compliance with AGPLv3.\n
package goblaze

import "fmt"

type Counter struct {
	ComponentBase
	pubTitle string
	Count    int
}

func NewCounter(title string) *Counter {
	c := &Counter{pubTitle: title}
	c.RegisterEvents(map[string]EventHandler{
		"inc": c.Increment,
	})
	return c
}

func (c *Counter) Increment() {
	c.Count++
}

func (c *Counter) Render() Node {
	return Div(
		H1(c.pubTitle),
		Div(Text(fmt.Sprintf("Count: %d", c.Count))),
		Button("Increment", "onclick", "inc"),
	)
}
