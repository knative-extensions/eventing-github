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
	"context"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"

	"k8s.io/utils/pointer"

	"github.com/google/go-github/v31/github"
	"golang.org/x/oauth2"
)

func setup(t *testing.T) (*github.Client, func()) {
	server := httptest.NewServer(newRouter())
	// client is the GitHub client being tested and is
	// configured to use test server.
	client := github.NewClient(nil)
	url, _ := url.Parse(server.URL + "/api/v3/")
	client.BaseURL = url
	client.UploadURL = url

	return client, server.Close
}

//
func setupReal(t *testing.T) (*github.Client, func()) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	return github.NewClient(tc), func() {}
}

func TestHookHandlers(t *testing.T) {
	client, teardown := setup(t)
	defer teardown()

	hook := &github.Hook{
		URL:    pointer.StringPtr("dummy"),
		Events: []string{"dummy_event"},
	}
	h, _, err := client.Repositories.CreateHook(context.Background(), "o", "r", hook)
	if err != nil {
		t.Error(err)
	}
	if h == nil {
		t.Error("expected non-nil create hook response")
	}

	h2, _, err := client.Repositories.GetHook(context.Background(), "o", "r", *h.ID)
	if err != nil {
		t.Error(err)
	}
	if h2 == nil {
		t.Error("expected non-nil get hook response")
	}
	if *h2.ID != *h.ID {
		t.Errorf("expected ID to be identical (%d != %d)", *h2.ID, *h.ID)
	}

	h2.Events = []string{"dummy_event_2"}
	h3, _, err := client.Repositories.EditHook(context.Background(), "o", "r", h2.GetID(), h2)
	if err != nil {
		t.Error(err)
	}
	if h3 == nil {
		t.Error("expected non-nil get hook response")
	}
	if !reflect.DeepEqual(h3.Events, []string{"dummy_event_2"}) {
		t.Errorf("expected events to be identical (%+v != %+v)", h3.Events, []string{"dummy_event_2"})
	}

}
