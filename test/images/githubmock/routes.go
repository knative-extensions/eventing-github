/*
Copyright 2020 The Knative Authors

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

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	baseURL = "http://mock.api.github.com/"
)

func newRouter() http.Handler {
	r := mux.NewRouter()

	hh := NewHookHandlers()
	r.Methods("GET").Path("/api/v3/repos/{org}/{repo}/hooks").HandlerFunc(notImplemented)
	r.Methods("POST").Path("/api/v3/repos/{org}/{repo}/hooks").HandlerFunc(hh.CreateHook)
	r.Methods("GET").Path("/api/v3/repos/{org}/{repo}/hooks/{id}").HandlerFunc(hh.GetHook)
	r.Methods("PATCH").Path("/api/v3/repos/{org}/{repo}/hooks/{id}").HandlerFunc(hh.EditHook)

	return r
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
}

func notFound(w http.ResponseWriter) {
	w.Write([]byte(`{"message":"Not Found"}`))
	w.WriteHeader(404)
}

func internalError(w http.ResponseWriter, msg string) {
	w.Write([]byte(msg))
	w.WriteHeader(500)
}

func badRequest(w http.ResponseWriter, msg string) {
	w.Write([]byte(msg))
	w.WriteHeader(400)
}
