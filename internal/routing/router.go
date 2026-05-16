// Copyright 2026 Daniel
// Licensed under the GNU Affero General Public License v3.0.
// Copying or distributing this file requires compliance with AGPLv3.


package routing

import (
    "net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Router struct {
    routes map[string]HandlerFunc
}

func New() *Router {
    return &Router{routes: make(map[string]HandlerFunc)}
}

func (r *Router) Handle(path string, h HandlerFunc) {
    r.routes[path] = h
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    if h, ok := r.routes[req.URL.Path]; ok {
        h(w, req)
        return
    }
    http.NotFound(w, req)
}
