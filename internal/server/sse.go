package server

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "sync/atomic"
    "time"

    "github.com/m3t41/goblaze/pkg/goblaze"
)

var sidCounter uint64

type sseSession struct {
    sid  string
    root goblaze.Component
    ch   chan []byte
}

func (h *Hub) HandleSSE(w http.ResponseWriter, r *http.Request) {
    flusher, ok := w.(http.Flusher)
    if !ok {
        http.Error(w, "streaming unsupported", http.StatusInternalServerError)
        return
    }

    ctx := r.Context()

    sid := fmt.Sprintf("%d", atomic.AddUint64(&sidCounter, 1))

    // headers for SSE
    w.Header().Set("Content-Type", "text/event-stream")
    w.Header().Set("Cache-Control", "no-cache")
    w.Header().Set("Connection", "keep-alive")

    s := &sseSession{
        sid:  sid,
        root: h.rootFactory(),
        ch:   make(chan []byte, 1),
    }

    h.sseSessions.Store(sid, s)
    defer h.sseSessions.Delete(sid)

    // send initial session id to client
    initMsg := map[string]any{"type": "init", "sid": sid}
    if b, err := json.Marshal(initMsg); err == nil {
        _, _ = io.WriteString(w, "event: init\n")
        _, _ = io.WriteString(w, "data: ")
        _, _ = w.Write(b)
        _, _ = io.WriteString(w, "\n\n")
        flusher.Flush()
    }

    // send initial render
    if b, err := json.Marshal([]map[string]any{{"op": "replace", "path": []int{}, "node": s.root.Render()}}); err == nil {
        _, _ = io.WriteString(w, "data: ")
        _, _ = w.Write(b)
        _, _ = io.WriteString(w, "\n\n")
        flusher.Flush()
    }

    // listen for render updates to send
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ctx.Done():
            return
        case b := <-s.ch:
            _, _ = io.WriteString(w, "data: ")
            _, _ = w.Write(b)
            _, _ = io.WriteString(w, "\n\n")
            flusher.Flush()
        case <-ticker.C:
            // periodic full render
            patches := []map[string]any{{"op": "replace", "path": []int{}, "node": s.root.Render()}}
            if b, err := json.Marshal(patches); err == nil {
                _, _ = io.WriteString(w, "data: ")
                _, _ = w.Write(b)
                _, _ = io.WriteString(w, "\n\n")
                flusher.Flush()
            }
        }
    }
}

// HandleEventPost accepts POSTs from client for a given session id.
// URL: /_goblaze/event?sid=<sid>
func (h *Hub) HandleEventPost(w http.ResponseWriter, r *http.Request) {
    sid := r.URL.Query().Get("sid")
    if sid == "" {
        http.Error(w, "missing sid", http.StatusBadRequest)
        return
    }

    v, ok := h.sseSessions.Load(sid)
    if !ok {
        http.Error(w, "session not found", http.StatusNotFound)
        return
    }
    s := v.(*sseSession)

    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "invalid body", http.StatusBadRequest)
        return
    }
    msg := string(body)

    dispatchEvent(s.root, msg)

    // render and send patch
    patches := []map[string]any{{"op": "replace", "path": []int{}, "node": s.root.Render()}}
    b, _ := json.Marshal(patches)

    // non-blocking send to session channel
    select {
    case s.ch <- b:
    default:
    }

    w.WriteHeader(http.StatusNoContent)
}
