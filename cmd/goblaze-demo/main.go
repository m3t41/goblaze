// Copyright 2026 Daniel
// Licensed under the GNU Affero General Public License v3.0.
// Copying or distributing this file requires compliance with AGPLv3.

package main

import (
	"github.com/m3t41/goblaze/internal/server"
	"github.com/m3t41/goblaze/pkg/goblaze"
)

func main() {
	// StartServer takes a factory function that creates a fresh component tree for each client session.
	// This ensures every client has their own independent state (e.g., separate time offsets).
	_ = server.StartServer(func() goblaze.Component {
		// Step 1: Create stateful child components
		// These components manage their own state (Counter.Count, TimeView.Offset)
		counter := goblaze.NewCounter("Server-Side Counter")
		tv := goblaze.NewTimeView()

		// Step 2: Build the root component using goblaze.Func
		// goblaze.Func() takes:
		//   - a render function (returns the UI Node tree)
		//   - an events map (nil here, root handles no direct events)
		//   - child components (counter, tv) so they receive event dispatches
		root := goblaze.Func(func() goblaze.Node {
			// Render hierarchy: root -> H1 + TimeView + Counter
			// Each child component (tv, counter) re-renders independently
			return goblaze.Div(
				goblaze.H1("GoBlaze Demo"),
				tv.Render(),      // TimeView displays real time, session time, and buttons
				counter.Render(), // Counter displays click count and increment button
			)
		}, nil, tv, counter)

		// Return the root component to the server
		// The server will call root.Render() periodically and send patches to the client
		return root
	})
}
