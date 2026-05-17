// Copyright 2026 Daniel
// Licensed under the GNU Affero General Public License v3.0.
// Copying or distributing this file requires compliance with AGPLv3.

package server

import (
	"context"
	"encoding/json"
	"strings"
	"sync"
	"time"

	"github.com/m3t41/goblaze/pkg/goblaze"
	"nhooyr.io/websocket"
)

type Session struct {
	conn    *websocket.Conn
	root    goblaze.Component
	oldTree goblaze.Node
	mu      sync.Mutex
}

func NewSession(conn *websocket.Conn, root goblaze.Component) *Session {
	return &Session{conn: conn, root: root}
}

func (s *Session) Run(ctx context.Context) error {
	s.renderAndSend()

	// start ticker to push periodic updates (e.g., server-side time)
	tickCtx, cancel := context.WithCancel(ctx)
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-tickCtx.Done():
				return
			case <-ticker.C:
				s.renderAndSend()
			}
		}
	}()

	for {
		_, data, err := s.conn.Read(ctx)
		if err != nil {
			cancel()
			return err
		}

		msg := string(data)
		s.handleEvent(msg)
	}
}

func (s *Session) handleEvent(msg string) {
	dispatchEvent(s.root, msg)
	s.renderAndSend()
}

func dispatchEvent(c goblaze.Component, msg string) {
	name := msg
	payload := ""
	if i := strings.Index(msg, "|"); i >= 0 {
		name = msg[:i]
		payload = msg[i+1:]
	}

	if evt, ok := c.(interface {
		GetEvents() map[string]goblaze.EventHandler
	}); ok {
		if h, ok := evt.GetEvents()[name]; ok {
			h(payload)
		}
	}

	if parent, ok := c.(interface{ Children() []goblaze.Component }); ok {
		for _, child := range parent.Children() {
			dispatchEvent(child, msg)
		}
	}
}

func (s *Session) renderAndSend() {
	s.mu.Lock()
	defer s.mu.Unlock()

	newTree := s.root.Render()

	patches := []map[string]any{
		{
			"op":   "replace",
			"path": []int{},
			"node": newTree,
		},
	}

	data, _ := json.Marshal(patches)

	_ = s.conn.Write(context.Background(), websocket.MessageText, data)

	s.oldTree = newTree
}
