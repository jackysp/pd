// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package uiserver

import (
	"io"
	"net/http"

	"github.com/pingcap/pd/pkg/dashboard/uiserver/assets"
)

// Handler returns an http.Handler that serves the dashboard UI.
func Handler() http.Handler {
	fs := assets.AssetFS()
	if fs != nil {
		fileServer := http.FileServer(fs)
		return fileServer
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "Dashboard UI is not built.\n")
	})
}
