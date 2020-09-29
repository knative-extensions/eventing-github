/*
Copyright 2019 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package github

import (
	"path"

	"github.com/blang/semver"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	"knative.dev/reconciler-test/pkg/config"
	"knative.dev/reconciler-test/pkg/framework"
	"knative.dev/reconciler-test/pkg/manifest"
	"knative.dev/reconciler-test/pkg/release"
)

var (
	Component = &githubComponent{}

	contribRelease = release.Release{
		Owner:      "knative",
		Repository: "eventing-contrib",
		Artifacts:  []string{"github.yaml"},
	}

	githubRelease = release.Release{
		Owner:      "knative-sandbox",
		Repository: "eventing-github",
		Artifacts:  []string{"github.yaml"},
	}
)

type githubComponent struct {
}

var _ framework.Component = (*githubComponent)(nil)

func (s *githubComponent) QName() string {
	return "components/eventing/sources/github"
}

func (s *githubComponent) InstalledVersion(rc framework.ResourceContext) string {
	rc = framework.NewResourceContext(rc, "knative-sources")
	var obj apiextensionsv1.CustomResourceDefinition
	_, err := rc.GetOrError("customresourcedefinitions", "githubsources.sources.knative.dev", &obj)

	if err != nil {
		return ""
	}
	if v, ok := obj.Labels["contrib.eventing.knative.dev/release"]; ok {
		return v
	}
	return ""
}

func (s *githubComponent) Install(rc framework.ResourceContext, gcfg config.Config) {
	cfg, ok := gcfg.(*Config)
	if !ok {
		rc.Errorf("invalid configuration type for %s", s.QName())
	}

	if cfg.Version == "devel" {
		rc.Apply(manifest.FromURL(path.Join(cfg.Path, "config")))
		return
	}

	if semver.MustParse(cfg.Version).LT(semver.MustParse("0.18.0")) {
		contribRelease.Install(rc, cfg.Version)
	} else {
		githubRelease.Install(rc, cfg.Version)
	}
}
