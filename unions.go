package github4

import (
	"encoding/json"
	"fmt"
)

// Types that can be inside Project Cards.
type ProjectCardItem interface{}

func ProjectCardItem_UnmarshalJSON(data []byte) (ProjectCardItem, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "Issue":
		var v Issue
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequest":
		var v PullRequest
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v interface{}
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for union %s: %s", "ProjectCardItem", t.Typename)
}

// An item in an pull request timeline
type PullRequestTimelineItem interface{}

func PullRequestTimelineItem_UnmarshalJSON(data []byte) (PullRequestTimelineItem, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "Commit":
		var v Commit
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestReview":
		var v PullRequestReview
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestReviewThread":
		var v PullRequestReviewThread
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestReviewComment":
		var v PullRequestReviewComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "IssueComment":
		var v IssueComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "ClosedEvent":
		var v ClosedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "ReopenedEvent":
		var v ReopenedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "SubscribedEvent":
		var v SubscribedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "UnsubscribedEvent":
		var v UnsubscribedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "MergedEvent":
		var v MergedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "ReferencedEvent":
		var v ReferencedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "AssignedEvent":
		var v AssignedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "UnassignedEvent":
		var v UnassignedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "LabeledEvent":
		var v LabeledEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "UnlabeledEvent":
		var v UnlabeledEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "MilestonedEvent":
		var v MilestonedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "DemilestonedEvent":
		var v DemilestonedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "RenamedTitleEvent":
		var v RenamedTitleEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "LockedEvent":
		var v LockedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "UnlockedEvent":
		var v UnlockedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "DeployedEvent":
		var v DeployedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "HeadRefDeletedEvent":
		var v HeadRefDeletedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "HeadRefRestoredEvent":
		var v HeadRefRestoredEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "HeadRefForcePushedEvent":
		var v HeadRefForcePushedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "BaseRefForcePushedEvent":
		var v BaseRefForcePushedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "ReviewRequestedEvent":
		var v ReviewRequestedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "ReviewRequestRemovedEvent":
		var v ReviewRequestRemovedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "ReviewDismissedEvent":
		var v ReviewDismissedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v interface{}
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for union %s: %s", "PullRequestTimelineItem", t.Typename)
}

// Any referencable object
type ReferencedSubject interface{}

func ReferencedSubject_UnmarshalJSON(data []byte) (ReferencedSubject, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "Issue":
		var v Issue
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequest":
		var v PullRequest
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v interface{}
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for union %s: %s", "ReferencedSubject", t.Typename)
}

// Types that can be inside a Milestone.
type MilestoneItem interface{}

func MilestoneItem_UnmarshalJSON(data []byte) (MilestoneItem, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "Issue":
		var v Issue
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequest":
		var v PullRequest
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v interface{}
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for union %s: %s", "MilestoneItem", t.Typename)
}

// An object which has a renamable title
type RenamedTitleSubject interface{}

func RenamedTitleSubject_UnmarshalJSON(data []byte) (RenamedTitleSubject, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "Issue":
		var v Issue
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequest":
		var v PullRequest
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v interface{}
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for union %s: %s", "RenamedTitleSubject", t.Typename)
}

// An item in an issue timeline
type IssueTimelineItem interface{}

func IssueTimelineItem_UnmarshalJSON(data []byte) (IssueTimelineItem, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "Commit":
		var v Commit
		err = json.Unmarshal(data, &v)
		return v, err
	case "IssueComment":
		var v IssueComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "ClosedEvent":
		var v ClosedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "ReopenedEvent":
		var v ReopenedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "SubscribedEvent":
		var v SubscribedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "UnsubscribedEvent":
		var v UnsubscribedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "ReferencedEvent":
		var v ReferencedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "AssignedEvent":
		var v AssignedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "UnassignedEvent":
		var v UnassignedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "LabeledEvent":
		var v LabeledEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "UnlabeledEvent":
		var v UnlabeledEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "MilestonedEvent":
		var v MilestonedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "DemilestonedEvent":
		var v DemilestonedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "RenamedTitleEvent":
		var v RenamedTitleEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "LockedEvent":
		var v LockedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "UnlockedEvent":
		var v UnlockedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v interface{}
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for union %s: %s", "IssueTimelineItem", t.Typename)
}

// Used for return value of Repository.issueOrPullRequest.
type IssueOrPullRequest interface{}

func IssueOrPullRequest_UnmarshalJSON(data []byte) (IssueOrPullRequest, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "Issue":
		var v Issue
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequest":
		var v PullRequest
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v interface{}
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for union %s: %s", "IssueOrPullRequest", t.Typename)
}

// Types that can be an actor.
type ReviewDismissalAllowanceActor interface{}

func ReviewDismissalAllowanceActor_UnmarshalJSON(data []byte) (ReviewDismissalAllowanceActor, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "User":
		var v User
		err = json.Unmarshal(data, &v)
		return v, err
	case "Team":
		var v Team
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v interface{}
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for union %s: %s", "ReviewDismissalAllowanceActor", t.Typename)
}

// Types that can be an actor.
type PushAllowanceActor interface{}

func PushAllowanceActor_UnmarshalJSON(data []byte) (PushAllowanceActor, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "User":
		var v User
		err = json.Unmarshal(data, &v)
		return v, err
	case "Team":
		var v Team
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v interface{}
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for union %s: %s", "PushAllowanceActor", t.Typename)
}

// The results of a search.
type SearchResultItem interface{}

func SearchResultItem_UnmarshalJSON(data []byte) (SearchResultItem, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "Issue":
		var v Issue
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequest":
		var v PullRequest
		err = json.Unmarshal(data, &v)
		return v, err
	case "Repository":
		var v Repository
		err = json.Unmarshal(data, &v)
		return v, err
	case "User":
		var v User
		err = json.Unmarshal(data, &v)
		return v, err
	case "Organization":
		var v Organization
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v interface{}
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for union %s: %s", "SearchResultItem", t.Typename)
}
