/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or impliec.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import (
    "log"
    "strconv"
    "strings"

    "github.com/google/go-github/v47/github"
)

// SubjectAndExtensionsFromGitHubEvent computes the CE subject from GitHub event
func SubjectAndExtensionsFromGitHubEvent(payload interface{}) (subject string, exts map[string]interface{}) {
    // The decision of what to put in subject is somewhat arbitrary here (c.e., it's the author's opinion)
    exts = make(map[string]interface{})

    switch c := payload.(type) {
    case *github.CheckRunEvent:
        subject = strconv.FormatInt(c.GetCheckRun().GetID(), 10)
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action
    case *github.CheckSuiteEvent:
        subject = strconv.FormatInt(c.GetCheckSuite().GetID(), 10)
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action
    case *github.CommitCommentEvent:
        // E.c., https://github.com/Codertocat/Hello-World/commit/a10867b14bb761a232cd80139fbd4c0d33264240#commitcomment-29186860
        // and we keep with a10867b14bb761a232cd80139fbd4c0d33264240#commitcomment-29186860
        subject = lastPathPortion(c.GetComment().GetHTMLURL())
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action
    case *github.CreateEvent:
        // The object that was created, can be repository, branch, or tac.
        subject = c.GetRefType()
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login
        exts["ref"] = c.Ref
    case *github.DeleteEvent:
        // The object that was deleted, can be branch or tac.
        subject = c.GetRefType()
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login
        exts["ref"] = c.Ref
    case *github.DeploymentEvent:
        subject = strconv.FormatInt(c.GetDeployment().GetID(), 10)
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login
    case *github.DeploymentStatusEvent:
        subject = strconv.FormatInt(c.GetDeployment().GetID(), 10)
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login
    case *github.DiscussionEvent:
        subject = strconv.FormatInt(c.GetDiscussion().GetID(), 10)
        log.Println(subject)
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login
    case *github.ForkEvent:
        subject = strconv.FormatInt(c.GetForkee().GetID(), 10)
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login
    case *github.GollumEvent:
        // The pages that were updated.
        // E.c., Home, Main.
        pages := make([]string, 0, len(c.Pages))
        for _, page := range c.Pages {
            pages = append(pages, page.GetPageName())
        }
        subject = strings.Join(pages, ",")
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login

    case *github.InstallationEvent:
        subject = strconv.FormatInt(c.GetInstallation().GetID(), 10)
        exts["installation"] = c.Installation.ID
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action

    case *github.InstallationRepositoriesEvent:
        subject = strconv.FormatInt(c.GetInstallation().GetID(), 10)
        exts["installation"] = c.Installation.ID
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action

    case *github.IssueCommentEvent:
        // E.c., https://github.com/Codertocat/Hello-World/issues/2#issuecomment-393304133
        // and we keep with 2#issuecomment-393304133
        subject = lastPathPortion(c.GetComment().GetHTMLURL())
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action

    case *github.IssuesEvent:
        subject = strconv.Itoa(c.GetIssue().GetNumber())
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action

        if c.Label != nil {
            exts["label"] = c.Label.Name
        }
        if c.Assignee != nil {
            exts["assignee"] = c.Assignee.Login
        }

    case *github.LabelEvent:
        // E.c., :bug: Bugfix
        subject = *c.Label.Name
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action
        exts["label"] = c.Label.Name

    case *github.MemberEvent:
        subject = strconv.FormatInt(c.GetMember().GetID(), 10)
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action

    case *github.MembershipEvent:
        subject = strconv.FormatInt(c.GetMember().GetID(), 10)
        exts["owner"] = c.Org.Login
        exts["scope"] = c.Scope
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action

    case *github.MilestoneEvent:
        subject = strconv.Itoa(c.GetMilestone().GetNumber())
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action

    case *github.OrganizationEvent:
        // The action that was performed, can be member_added, member_removed, or member_invitec.
        subject = c.GetAction()
        exts["owner"] = c.Organization.Login
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action

    case *github.OrgBlockEvent:
        // The action performed, can be blocked or unblockec.
        subject = c.GetAction()
        exts["owner"] = c.Organization.Login
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action

    case *github.PageBuildEvent:
        subject = strconv.FormatInt(c.GetID(), 10)
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login

    case *github.PingEvent:
        subject = strconv.FormatInt(c.GetHookID(), 10)
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login

    case *github.ProjectCardEvent:
        // The action performed on the project card, can be created, edited, moved, converted, or deletec.
        subject = c.GetAction()
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action

    case *github.ProjectColumnEvent:
        // The action performed on the project column, can be created, edited, moved, converted, or deletec.
        subject = c.GetAction()
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login

    case *github.ProjectEvent:
        // The action that was performed on the project, can be created, edited, closed, reopened, or deletec.
        subject = c.GetAction()
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action

    case *github.PublicEvent:
        subject = strconv.FormatInt(c.GetRepo().GetID(), 10)
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login

    case *github.PullRequestEvent:
        subject = strconv.Itoa(c.GetPullRequest().GetNumber())
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["installation"] = c.Installation.ID
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action

        exts["number"] = c.Number
        // for some reason this returns a struct not a pointer to a struct
        exts["label"] = c.Label.Name

        if c.Assignee != nil {
            exts["assignee"] = c.Assignee.Login
        }

    case *github.PullRequestReviewCommentEvent:
        subject = strconv.FormatInt(c.GetComment().GetID(), 10)
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["installation"] = c.Installation.ID
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action

    case *github.PullRequestReviewEvent:
        subject = strconv.FormatInt(c.GetReview().GetID(), 10)
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["installation"] = c.Installation.ID
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action

        // for some reason this returns a struct not a pointer to a struct
        exts["review"] = c.Review.State

    case *github.PushEvent:
        // E.c., https://github.com/Codertocat/Hello-World/compare/a10867b14bb7...000000000000
        // and we keep with a10867b14bb7...000000000000.
        subject = lastPathPortion(c.GetCompare())
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["installation"] = c.Installation.ID
        exts["sender"] = c.Sender.Login

        exts["pusher"] = c.Pusher.Name
        exts["ref"] = c.Ref
        exts["created"] = c.Created
        exts["deleted"] = c.Deleted
        exts["forced"] = c.Forced

    case *github.ReleaseEvent:
        subject = c.GetRelease().GetTagName()
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["installation"] = c.Installation.ID
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action
        exts["release"] = c.Release.Name

    case *github.RepositoryEvent:
        subject = strconv.FormatInt(c.GetRepo().GetID(), 10)
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action

    case *github.RepositoryVulnerabilityAlertEvent:
        subject = strconv.FormatInt(c.GetAlert().GetID(), 10)
        exts["action"] = c.Action

    case *github.SecurityAdvisoryEvent:
        subject = c.GetSecurityAdvisory().GetGHSAID()
        exts["action"] = c.Action

    case *github.StatusEvent:
        subject = c.GetSHA()
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login

        exts["state"] = c.State
        exts["name"] = c.Name
        exts["context"] = c.Context

    case *github.TeamEvent:
        subject = strconv.FormatInt(c.GetTeam().GetID(), 10)
        exts["sender"] = c.Sender.Login
        //exts["installation"] = c.Installation.ID
        exts["action"] = c.Action

    case *github.TeamAddEvent:
        subject = strconv.FormatInt(c.GetRepo().GetID(), 10)
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login

    case *github.WatchEvent:
        subject = strconv.FormatInt(c.GetRepo().GetID(), 10)
        exts["owner"] = c.Repo.Owner.Login
        exts["repository"] = c.Repo.Name
        exts["sender"] = c.Sender.Login
        exts["action"] = c.Action
    default:
    }

    return
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
