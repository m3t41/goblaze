// Copyright 2026 Daniel
// Licensed under the GNU Affero General Public License v3.0.
// Copying or distributing this file requires compliance with AGPLv3.

package server

import (
	"net/http"

	"github.com/m3t41/goblaze/pkg/goblaze"
)

func StartServer(rootFactory func() goblaze.Component) error {
	hub := NewHub(rootFactory)

	// 1. WebSocket
	http.HandleFunc("/_goblaze", hub.HandleWS)

	// 1b. Server-Sent Events (receive via POST to /_goblaze/event)
	http.HandleFunc("/_goblaze/sse", hub.HandleSSE)
	http.HandleFunc("/_goblaze/event", hub.HandleEventPost)

	// 2. Static files (client.js)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web"))))

	// 3. index.html
	http.Handle("/", http.FileServer(http.Dir("web")))

	return http.ListenAndServe(":8080", nil)
}
