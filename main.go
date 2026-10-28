// Copyright (c) 2026 AllIsA1.
// SPDX-License-Identifier: MIT

package main

import (
	"log"
	"net/http"
)

// Since lucx.325 the panel ships the ghostlane-project olcRTC fork itself, so there is nothing left for this
// plugin to do. It only answers the plugin system's health check until it is removed.
func main() {
	log.Print("better-olcrtc plugin ready (nothing to do: the panel brings the olcRTC fork v0.12.1 itself)")
	if err := Serve(http.NotFoundHandler()); err != nil {
		log.Fatal(err)
	}
}
