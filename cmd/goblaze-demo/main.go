// Copyright 2026 Daniel\n// Licensed under the GNU Affero General Public License v3.0.\n// Copying or distributing this file requires compliance with AGPLv3.\n
package main

import (
	"github.com/yourname/goblaze/internal/server"
	"github.com/yourname/goblaze/pkg/goblaze"
)

type App struct {
	goblaze.ComponentBase
	Counter *goblaze.Counter
}

func (a *App) Render() goblaze.Node {
	return goblaze.Div(
		goblaze.H1("GoBlaze Demo"),
		a.Counter.Render(),
	)
}

func (a *App) Children() []goblaze.Component {
	return []goblaze.Component{a.Counter}
}

func main() {
	_ = server.StartServer(func() goblaze.Component {
		return &App{
			Counter: goblaze.NewCounter("Server-Side Counter"),
		}
	})
}
