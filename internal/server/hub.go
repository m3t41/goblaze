// Copyright 2026 Daniel\n// Licensed under the GNU Affero General Public License v3.0.\n// Copying or distributing this file requires compliance with AGPLv3.\n
package server

import (
	"context"
	"net/http"

	"github.com/yourname/goblaze/pkg/goblaze"
	"nhooyr.io/websocket"
)

type Hub struct {
	rootFactory func() goblaze.Component
}

func NewHub(rootFactory func() goblaze.Component) *Hub {
	return &Hub{rootFactory: rootFactory}
}

func (h *Hub) HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close(websocket.StatusNormalClosure, "")

	root := h.rootFactory()
	sess := NewSession(conn, root)

	_ = sess.Run(context.Background())
}
