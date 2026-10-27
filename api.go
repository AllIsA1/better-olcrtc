// Copyright (c) 2026 AllIsA1.
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// coreUpdater is what the plugin's calls need from the updater.
type coreUpdater interface {
	snapshot(ctx context.Context) snapshot
	ensure(ctx context.Context, force bool) error
}

// newMux serves the calls of the plugin: what it keeps and a forced install. It has no page (the panel's own
// inbound form shows what the core can do); the panel's plugin gateway only reaches /api/...
func newMux(u coreUpdater) http.Handler {
	inner := http.NewServeMux()
	inner.HandleFunc("GET /status", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 20*time.Second)
		defer cancel()
		writeJSON(w, http.StatusOK, u.snapshot(ctx))
	})
	inner.HandleFunc("POST /install", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), callBudget)
		defer cancel()
		if err := u.ensure(ctx, true); err != nil {
			writeJSON(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, u.snapshot(ctx))
	})
	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", inner))
	return mux
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
