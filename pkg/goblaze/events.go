// Copyright 2026 Daniel
// Licensed under the GNU Affero General Public License v3.0.
// Copying or distributing this file requires compliance with AGPLv3.


package goblaze

import "github.com/yourname/goblaze/internal/vdom"

type Node = vdom.Node

type EventHandler func()

type Component interface {
	Render() Node
	RegisterEvents(map[string]EventHandler)
}

type ComponentBase struct {
	Events map[string]EventHandler
}

func (c *ComponentBase) RegisterEvents(m map[string]EventHandler) {
	if c.Events == nil {
		c.Events = map[string]EventHandler{}
	}
	for k, v := range m {
		c.Events[k] = v
	}
}

func (c *ComponentBase) GetEvents() map[string]EventHandler {
	return c.Events
}

func Text(text string) Node {
	return vdom.Text(text)
}

func Div(children ...Node) Node {
	return vdom.Element("div", nil, children...)
}

func H1(text string) Node {
	return vdom.Element("h1", nil, vdom.Text(text))
}

func Button(label, eventAttr, eventName string) Node {
	return vdom.Element("button", map[string]string{eventAttr: eventName}, vdom.Text(label))
}
