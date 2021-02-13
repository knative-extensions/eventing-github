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

package common

import (
	"strconv"
	"strings"

	"go.uber.org/zap"
	gh "gopkg.in/go-playground/webhooks.v5/github"
)

const (
	GHHeaderEvent    = "X-GitHub-Event"
	GHHeaderDelivery = "X-GitHub-Delivery"
)

var ValidEvents = []gh.Event{
	gh.CheckRunEvent,
	gh.CheckSuiteEvent,
	gh.CommitCommentEvent,
	gh.CreateEvent,
	gh.DeleteEvent,
	gh.DeploymentEvent,
	gh.DeploymentStatusEvent,
	gh.ForkEvent,
	gh.GollumEvent,
	gh.InstallationEvent,
	gh.InstallationRepositoriesEvent,
	gh.IntegrationInstallationEvent,
	gh.IntegrationInstallationRepositoriesEvent,
	gh.IssueCommentEvent,
	gh.IssuesEvent,
	gh.LabelEvent,
	gh.MemberEvent,
	gh.MembershipEvent,
	gh.MilestoneEvent,
	gh.OrganizationEvent,
	gh.OrgBlockEvent,
	gh.PageBuildEvent,
	gh.PingEvent,
	gh.ProjectCardEvent,
	gh.ProjectColumnEvent,
	gh.ProjectEvent,
	gh.PublicEvent,
	gh.PullRequestEvent,
	gh.PullRequestReviewCommentEvent,
	gh.PullRequestReviewEvent,
	gh.PushEvent,
	gh.ReleaseEvent,
	gh.RepositoryEvent,
	gh.RepositoryVulnerabilityAlertEvent,
	gh.SecurityAdvisoryEvent,
	gh.StatusEvent,
	gh.TeamAddEvent,
	gh.TeamEvent,
	gh.WatchEvent,
}

// SubjectAndExtensionsFromGitHubEvent computes the CE subject from GitHub event
func SubjectAndExtensionsFromGitHubEvent(gitHubEvent gh.Event, payload interface{}, logger *zap.SugaredLogger) (string, map[string]interface{}) {
	// The decision of what to put in subject is somewhat arbitrary here (i.e., it's the author's opinion)
	// TODO check if we should be setting subject to these values.
	var subject string
	var exts = make(map[string]interface{})
	var ok bool
	switch gitHubEvent {
	case gh.CheckRunEvent:
		var c gh.CheckRunPayload
		if c, ok = payload.(gh.CheckRunPayload); ok {
			subject = strconv.FormatInt(c.CheckRun.ID, 10)
			exts["owner"] = c.Repository.Owner.Login
			exts["repository"] = c.Repository.Name
			//exts["installation"] = c.Installation.ID
			exts["sender"] = c.Sender.Login
			exts["action"] = c.Action
		}
	case gh.CheckSuiteEvent:
		var c gh.CheckSuitePayload
		if c, ok = payload.(gh.CheckSuitePayload); ok {
			subject = strconv.FormatInt(c.CheckSuite.ID, 10)
			exts["owner"] = c.Repository.Owner.Login
			exts["repository"] = c.Repository.Name
			//exts["installation"] = c.Installation.ID
			exts["sender"] = c.Sender.Login
			exts["action"] = c.Action
		}
	case gh.CommitCommentEvent:
		var c gh.CommitCommentPayload
		if c, ok = payload.(gh.CommitCommentPayload); ok {
			// E.g., https://github.com/Codertocat/Hello-World/commit/a10867b14bb761a232cd80139fbd4c0d33264240#commitcomment-29186860
			// and we keep with a10867b14bb761a232cd80139fbd4c0d33264240#commitcomment-29186860
			subject = lastPathPortion(c.Comment.HTMLURL)
			exts["owner"] = c.Repository.Owner.Login
			exts["repository"] = c.Repository.Name
			//exts["installation"] = c.Installation.ID
			exts["sender"] = c.Sender.Login
			exts["action"] = c.Action
		}
	case gh.CreateEvent:
		var c gh.CreatePayload
		if c, ok = payload.(gh.CreatePayload); ok {
			// The object that was created, can be repository, branch, or tag.
			subject = c.RefType
			exts["owner"] = c.Repository.Owner.Login
			exts["repository"] = c.Repository.Name
			//exts["installation"] = c.Installation.ID
			exts["sender"] = c.Sender.Login

			exts["ref"] = c.Ref
		}
	case gh.DeleteEvent:
		var d gh.DeletePayload
		if d, ok = payload.(gh.DeletePayload); ok {
			// The object that was deleted, can be branch or tag.
			subject = d.RefType
			exts["owner"] = d.Repository.Owner.Login
			exts["repository"] = d.Repository.Name
			//exts["installation"] = d.Installation.ID
			exts["sender"] = d.Sender.Login

			exts["ref"] = d.Ref
		}
	case gh.DeploymentEvent:
		var d gh.DeploymentPayload
		if d, ok = payload.(gh.DeploymentPayload); ok {
			subject = strconv.FormatInt(d.Deployment.ID, 10)
			exts["owner"] = d.Repository.Owner.Login
			exts["repository"] = d.Repository.Name
			//exts["installation"] = d.Installation.ID
			exts["sender"] = d.Sender.Login
		}
	case gh.DeploymentStatusEvent:
		var d gh.DeploymentStatusPayload
		if d, ok = payload.(gh.DeploymentStatusPayload); ok {
			subject = strconv.FormatInt(d.Deployment.ID, 10)
			exts["owner"] = d.Repository.Owner.Login
			exts["repository"] = d.Repository.Name
			//exts["installation"] = d.Installation.ID
			exts["sender"] = d.Sender.Login
		}
	case gh.ForkEvent:
		var f gh.ForkPayload
		if f, ok = payload.(gh.ForkPayload); ok {
			subject = strconv.FormatInt(f.Forkee.ID, 10)
			exts["owner"] = f.Repository.Owner.Login
			exts["repository"] = f.Repository.Name
			//exts["installation"] = f.Installation.ID
			exts["sender"] = f.Sender.Login
		}
	case gh.GollumEvent:
		var g gh.GollumPayload
		if g, ok = payload.(gh.GollumPayload); ok {
			// The pages that were updated.
			// E.g., Home, Main.
			pages := make([]string, 0, len(g.Pages))
			for _, page := range g.Pages {
				pages = append(pages, page.PageName)
			}
			subject = strings.Join(pages, ",")
			exts["owner"] = g.Repository.Owner.Login
			exts["repository"] = g.Repository.Name
			//exts["installation"] = g.Installation.ID
			exts["sender"] = g.Sender.Login
		}
	case gh.InstallationEvent, gh.IntegrationInstallationEvent:
		var i gh.InstallationPayload
		if i, ok = payload.(gh.InstallationPayload); ok {
			subject = strconv.FormatInt(i.Installation.ID, 10)
			exts["installation"] = i.Installation.ID
			exts["sender"] = i.Sender.Login
			exts["action"] = i.Action
		}
	case gh.InstallationRepositoriesEvent, gh.IntegrationInstallationRepositoriesEvent:
		var i gh.InstallationRepositoriesPayload
		if i, ok = payload.(gh.InstallationRepositoriesPayload); ok {
			subject = strconv.FormatInt(i.Installation.ID, 10)
			exts["installation"] = i.Installation.ID
			exts["sender"] = i.Sender.Login
			exts["action"] = i.Action
		}
	case gh.IssueCommentEvent:
		var i gh.IssueCommentPayload
		if i, ok = payload.(gh.IssueCommentPayload); ok {
			// E.g., https://github.com/Codertocat/Hello-World/issues/2#issuecomment-393304133
			// and we keep with 2#issuecomment-393304133
			subject = lastPathPortion(i.Comment.HTMLURL)
			exts["owner"] = i.Repository.Owner.Login
			exts["repository"] = i.Repository.Name
			//exts["installation"] = i.Installation.ID
			exts["sender"] = i.Sender.Login
			exts["action"] = i.Action
		}
	case gh.IssuesEvent:
		var i gh.IssuesPayload
		if i, ok = payload.(gh.IssuesPayload); ok {
			subject = strconv.FormatInt(i.Issue.Number, 10)
			exts["owner"] = i.Repository.Owner.Login
			exts["repository"] = i.Repository.Name
			//exts["installation"] = i.Installation.ID
			exts["sender"] = i.Sender.Login
			exts["action"] = i.Action

			if i.Label != nil {
				exts["label"] = i.Label.Name
			}
			if i.Assignee != nil {
				exts["assignee"] = i.Assignee.Login
			}
		}
	case gh.LabelEvent:
		var l gh.LabelPayload
		if l, ok = payload.(gh.LabelPayload); ok {
			// E.g., :bug: Bugfix
			subject = l.Label.Name
			exts["owner"] = l.Repository.Owner.Login
			exts["repository"] = l.Repository.Name
			//exts["installation"] = l.Installation.ID
			exts["sender"] = l.Sender.Login
			exts["action"] = l.Action

			exts["label"] = l.Label.Name
		}
	case gh.MemberEvent:
		var m gh.MemberPayload
		if m, ok = payload.(gh.MemberPayload); ok {
			subject = strconv.FormatInt(m.Member.ID, 10)
			exts["owner"] = m.Repository.Owner.Login
			exts["repository"] = m.Repository.Name
			//exts["installation"] = m.Installation.ID
			exts["sender"] = m.Sender.Login
			exts["action"] = m.Action
		}
	case gh.MembershipEvent:
		var m gh.MembershipPayload
		if m, ok = payload.(gh.MembershipPayload); ok {
			subject = strconv.FormatInt(m.Member.ID, 10)
			exts["owner"] = m.Organization.Login
			exts["scope"] = m.Scope
			//exts["installation"] = m.Installation.ID
			exts["sender"] = m.Sender.Login
			exts["action"] = m.Action
		}
	case gh.MilestoneEvent:
		var m gh.MilestonePayload
		if m, ok = payload.(gh.MilestonePayload); ok {
			subject = strconv.FormatInt(m.Milestone.Number, 10)
			exts["owner"] = m.Repository.Owner.Login
			exts["repository"] = m.Repository.Name
			//exts["installation"] = m.Installation.ID
			exts["sender"] = m.Sender.Login
			exts["action"] = m.Action
		}
	case gh.OrganizationEvent:
		var o gh.OrganizationPayload
		if o, ok = payload.(gh.OrganizationPayload); ok {
			// The action that was performed, can be member_added, member_removed, or member_invited.
			subject = o.Action
			exts["owner"] = o.Organization.Login
			//exts["installation"] = o.Installation.ID
			exts["sender"] = o.Sender.Login
			exts["action"] = o.Action
		}
	case gh.OrgBlockEvent:
		var o gh.OrgBlockPayload
		if o, ok = payload.(gh.OrgBlockPayload); ok {
			// The action performed, can be blocked or unblocked.
			subject = o.Action
			exts["owner"] = o.Organization.Login
			//exts["installation"] = o.Installation.ID
			exts["sender"] = o.Sender.Login
			exts["action"] = o.Action
		}
	case gh.PageBuildEvent:
		var p gh.PageBuildPayload
		if p, ok = payload.(gh.PageBuildPayload); ok {
			subject = strconv.FormatInt(p.ID, 10)
			exts["owner"] = p.Repository.Owner.Login
			exts["repository"] = p.Repository.Name
			//exts["installation"] = p.Installation.ID
			exts["sender"] = p.Sender.Login
		}
	case gh.PingEvent:
		var p gh.PingPayload
		if p, ok = payload.(gh.PingPayload); ok {
			subject = strconv.Itoa(p.HookID)
			exts["owner"] = p.Repository.Owner.Login
			exts["repository"] = p.Repository.Name
			exts["sender"] = p.Sender.Login
		}
	case gh.ProjectCardEvent:
		var p gh.ProjectCardPayload
		if p, ok = payload.(gh.ProjectCardPayload); ok {
			// The action performed on the project card, can be created, edited, moved, converted, or deleted.
			subject = p.Action
			exts["owner"] = p.Repository.Owner.Login
			exts["repository"] = p.Repository.Name
			//exts["installation"] = p.Installation.ID
			exts["sender"] = p.Sender.Login
			exts["action"] = p.Action
		}
	case gh.ProjectColumnEvent:
		var p gh.ProjectColumnPayload
		if p, ok = payload.(gh.ProjectColumnPayload); ok {
			// The action performed on the project column, can be created, edited, moved, converted, or deleted.
			subject = p.Action
			exts["owner"] = p.Repository.Owner.Login
			exts["repository"] = p.Repository.Name
			//exts["installation"] = p.Installation.ID
			exts["sender"] = p.Sender.Login
			exts["action"] = p.Action
		}
	case gh.ProjectEvent:
		var p gh.ProjectPayload
		if p, ok = payload.(gh.ProjectPayload); ok {
			// The action that was performed on the project, can be created, edited, closed, reopened, or deleted.
			subject = p.Action
			exts["owner"] = p.Repository.Owner.Login
			exts["repository"] = p.Repository.Name
			//exts["installation"] = p.Installation.ID
			exts["sender"] = p.Sender.Login
			exts["action"] = p.Action
		}
	case gh.PublicEvent:
		var p gh.PublicPayload
		if p, ok = payload.(gh.PublicPayload); ok {
			subject = strconv.FormatInt(p.Repository.ID, 10)
			exts["owner"] = p.Repository.Owner.Login
			exts["repository"] = p.Repository.Name
			//exts["installation"] = p.Installation.ID
			exts["sender"] = p.Sender.Login
		}
	case gh.PullRequestEvent:
		var p gh.PullRequestPayload
		if p, ok = payload.(gh.PullRequestPayload); ok {
			subject = strconv.FormatInt(p.PullRequest.Number, 10)
			exts["owner"] = p.Repository.Owner.Login
			exts["repository"] = p.Repository.Name
			exts["installation"] = p.Installation.ID
			exts["sender"] = p.Sender.Login
			exts["action"] = p.Action

			exts["number"] = p.Number
			// for some reason this returns a struct not a pointer to a struct
			exts["label"] = p.Label.Name

			if p.Assignee != nil {
				exts["assignee"] = p.Assignee.Login
			}
		}
	case gh.PullRequestReviewCommentEvent:
		var p gh.PullRequestReviewCommentPayload
		if p, ok = payload.(gh.PullRequestReviewCommentPayload); ok {
			subject = strconv.FormatInt(p.Comment.ID, 10)
			exts["owner"] = p.Repository.Owner.Login
			exts["repository"] = p.Repository.Name
			exts["installation"] = p.Installation.ID
			exts["sender"] = p.Sender.Login
			exts["action"] = p.Action
		}
	case gh.PullRequestReviewEvent:
		var p gh.PullRequestReviewPayload
		if p, ok = payload.(gh.PullRequestReviewPayload); ok {
			subject = strconv.FormatInt(p.Review.ID, 10)
			exts["owner"] = p.Repository.Owner.Login
			exts["repository"] = p.Repository.Name
			exts["installation"] = p.Installation.ID
			exts["sender"] = p.Sender.Login
			exts["action"] = p.Action

			// for some reason this returns a struct not a pointer to a struct
			exts["review"] = p.Review.State
		}
	case gh.PushEvent:
		var p gh.PushPayload
		if p, ok = payload.(gh.PushPayload); ok {
			// E.g., https://github.com/Codertocat/Hello-World/compare/a10867b14bb7...000000000000
			// and we keep with a10867b14bb7...000000000000.
			subject = lastPathPortion(p.Compare)
			exts["owner"] = p.Repository.Owner.Login
			exts["repository"] = p.Repository.Name
			exts["installation"] = p.Installation.ID
			exts["sender"] = p.Sender.Login

			exts["pusher"] = p.Pusher.Name
			exts["ref"] = p.Ref
			exts["created"] = p.Created
			exts["deleted"] = p.Deleted
			exts["forced"] = p.Forced
		}
	case gh.ReleaseEvent:
		var r gh.ReleasePayload
		if r, ok = payload.(gh.ReleasePayload); ok {
			subject = r.Release.TagName
			exts["owner"] = r.Repository.Owner.Login
			exts["repository"] = r.Repository.Name
			exts["installation"] = r.Installation.ID
			exts["sender"] = r.Sender.Login
			exts["action"] = r.Action
			exts["release"] = r.Release.Name
		}
	case gh.RepositoryEvent:
		var r gh.RepositoryPayload
		if r, ok = payload.(gh.RepositoryPayload); ok {
			subject = strconv.FormatInt(r.Repository.ID, 10)
			exts["owner"] = r.Repository.Owner.Login
			exts["repository"] = r.Repository.Name
			//exts["installation"] = r.Installation.ID
			exts["sender"] = r.Sender.Login
			exts["action"] = r.Action
		}
	case gh.RepositoryVulnerabilityAlertEvent:
		var r gh.RepositoryVulnerabilityAlertPayload
		if r, ok = payload.(gh.RepositoryVulnerabilityAlertPayload); ok {
			subject = strconv.FormatInt(r.Alert.ID, 10)
			exts["action"] = r.Action
		}
	case gh.SecurityAdvisoryEvent:
		var s gh.SecurityAdvisoryPayload
		if s, ok = payload.(gh.SecurityAdvisoryPayload); ok {
			subject = s.SecurityAdvisory.GHSAID
			exts["action"] = s.Action
		}
	case gh.StatusEvent:
		var s gh.StatusPayload
		if s, ok = payload.(gh.StatusPayload); ok {
			subject = s.Sha
			exts["owner"] = s.Repository.Owner.Login
			exts["repository"] = s.Repository.Name
			//exts["installation"] = s.Installation.ID
			exts["sender"] = s.Sender.Login

			exts["state"] = s.State
			exts["name"] = s.Name
			exts["context"] = s.Context
		}
	case gh.TeamEvent:
		var t gh.TeamPayload
		if t, ok = payload.(gh.TeamPayload); ok {
			subject = strconv.FormatInt(t.Team.ID, 10)
			exts["sender"] = t.Sender.Login
			//exts["installation"] = t.Installation.ID
			exts["action"] = t.Action
		}
	case gh.TeamAddEvent:
		var t gh.TeamAddPayload
		if t, ok = payload.(gh.TeamAddPayload); ok {
			subject = strconv.FormatInt(t.Repository.ID, 10)
			exts["owner"] = t.Repository.Owner.Login
			exts["repository"] = t.Repository.Name
			//exts["installation"] = t.Installation.ID
			exts["sender"] = t.Sender.Login
		}
	case gh.WatchEvent:
		var w gh.WatchPayload
		if w, ok = payload.(gh.WatchPayload); ok {
			subject = strconv.FormatInt(w.Repository.ID, 10)
			exts["owner"] = w.Repository.Owner.Login
			exts["repository"] = w.Repository.Name
			//exts["installation"] = w.Installation.ID
			exts["sender"] = w.Sender.Login
			exts["action"] = w.Action
		}
	}
	if !ok {
		logger.Errorf("Invalid payload in gitHub event %s", gitHubEvent)
	} else if subject == "" {
		logger.Warnf("No subject found in gitHub event %s", gitHubEvent)
	}
	return subject, exts
}

func lastPathPortion(url string) string {
	var subject string
	index := strings.LastIndex(url, "/")
	if index != -1 {
		// Keep the last part.
		subject = url[index+1:]
	}
	return subject
}
