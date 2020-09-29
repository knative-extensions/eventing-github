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
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"k8s.io/utils/pointer"

	"github.com/google/go-github/v31/github"
	"github.com/gorilla/mux"
)

type hookHandlers struct {
	lastId int64
	hooks  map[string]map[int64]github.Hook
}

func NewHookHandlers() *hookHandlers {
	return &hookHandlers{
		lastId: 0,
		hooks:  make(map[string]map[int64]github.Hook),
	}
}

type createHookRequest struct {
	// Config is required.
	Name   string                 `json:"name"`
	Config map[string]interface{} `json:"config,omitempty"`
	Events []string               `json:"events,omitempty"`
	Active *bool                  `json:"active,omitempty"`
}

func (h *hookHandlers) CreateHook(w http.ResponseWriter, r *http.Request) {
	var hr createHookRequest
	err := json.NewDecoder(r.Body).Decode(&hr)
	if err != nil {
		badRequest(w, err.Error())
		return
	}

	now := time.Now()
	h.lastId += 1

	hook := github.Hook{
		ID:        pointer.Int64Ptr(h.lastId),
		CreatedAt: &now,
		UpdatedAt: &now,
		URL:       &baseURL,
		Config:    nil,
		Events:    hr.Events,
		Active:    pointer.BoolPtr(true),
	}

	vars := mux.Vars(r)
	or := fmt.Sprintf("%s/%s", vars["org"], vars["repo"])

	ids, ok := h.hooks[or]
	if !ok {
		ids = make(map[int64]github.Hook)
		h.hooks[or] = ids
	}

	ids[*hook.ID] = hook

	if err := json.NewEncoder(w).Encode(hook); err != nil {
		internalError(w, err.Error())
		return
	}
}

func (h *hookHandlers) EditHook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		badRequest(w, err.Error())
	}

	var hook github.Hook
	if err := json.NewDecoder(r.Body).Decode(&hook); err != nil {
		badRequest(w, err.Error())
		return
	}

	now := time.Now()
	hook.UpdatedAt = &now

	or := fmt.Sprintf("%s/%s", vars["org"], vars["repo"])

	ids, ok := h.hooks[or]
	if !ok {
		ids = make(map[int64]github.Hook)
		h.hooks[or] = ids
	}

	ids[int64(id)] = hook

	if err := json.NewEncoder(w).Encode(hook); err != nil {
		internalError(w, err.Error())
		return
	}
}

func (h *hookHandlers) GetHook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	or := fmt.Sprintf("%s/%s", vars["org"], vars["repo"])
	ids, ok := h.hooks[or]
	if !ok {
		notFound(w)
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		badRequest(w, err.Error())
	}

	hook, ok := ids[int64(id)]
	if !ok {
		notFound(w)
		return
	}

	if err := json.NewEncoder(w).Encode(hook); err != nil {
		internalError(w, err.Error())
		return
	}
}
