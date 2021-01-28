/*
Copyright 2021 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package router

import (
	"net/http"
	"sync"

	"knative.dev/eventing-github/pkg/client/listers/sources/v1alpha1"
)

// keyedHandler associates a http.Handler to a source object identified by
// namespace/name.
type keyedHandler struct {
	handler   http.Handler
	namespace string
	name      string
}

// Router is a GitHub webhook router which delegates webhook events received
// over HTTP to sub-routers.
type Router struct {
	routersMu sync.RWMutex
	routers   map[string]keyedHandler

	lister v1alpha1.GitHubSourceLister
}

// Check that Router implements http.Handler.
var _ http.Handler = (*Router)(nil)

// New returns a new Router.
func New(lister v1alpha1.GitHubSourceLister) *Router {
	return &Router{
		routers: make(map[string]keyedHandler),
		lister:  lister,
	}
}

// ServeHTTP implements http.Handler.
func (h *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Path-based dispatch
	h.routersMu.RLock()
	keyedHandler, ok := h.routers[r.URL.Path]
	h.routersMu.RUnlock()

	if ok {
		// Check if source still exists.
		_, err := h.lister.GitHubSources(keyedHandler.namespace).Get(keyedHandler.name)
		if err == nil {
			keyedHandler.handler.ServeHTTP(w, r)
			return
		}

		h.Unregister(r.URL.Path)
	}
	http.NotFoundHandler().ServeHTTP(w, r)
}

// Register adds a new GitHub event handler for the given GitHubSource.
func (h *Router) Register(name, namespace, path string, handler http.Handler) {
	h.routersMu.Lock()
	defer h.routersMu.Unlock()
	h.routers[path] = keyedHandler{
		handler:   handler,
		namespace: namespace,
		name:      name,
	}
}

// Unregister removes the GitHubSource served at the given path.
func (h *Router) Unregister(path string) {
	h.routersMu.Lock()
	defer h.routersMu.Unlock()
	delete(h.routers, path)
}
