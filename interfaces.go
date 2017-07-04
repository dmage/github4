package github4

import (
	"encoding/json"
	"fmt"
)

// Represents an object which can take actions on GitHub. Typically a User or Bot.
type Actor interface {
	// A URL pointing to the actor's public avatar.
	AvatarUrl() URI

	// The username of the actor.
	Login() string

	// The HTTP path for this actor.
	ResourcePath() URI

	// The HTTP URL for this actor.
	Url() URI
}

type Untyped_Actor map[string]interface{}

func (m Untyped_Actor) AvatarUrl() URI {
	return m["avatarUrl"].(URI)
}

func (m Untyped_Actor) Login() string {
	return m["login"].(string)
}

func (m Untyped_Actor) ResourcePath() URI {
	return m["resourcePath"].(URI)
}

func (m Untyped_Actor) Url() URI {
	return m["url"].(URI)
}

func Actor_UnmarshalJSON(data []byte) (Actor, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "Bot":
		var v Bot
		err = json.Unmarshal(data, &v)
		return v, err
	case "Organization":
		var v Organization
		err = json.Unmarshal(data, &v)
		return v, err
	case "User":
		var v User
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v Untyped_Actor
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "Actor", t.Typename)
}

// Represents an owner of a Project.
type ProjectOwner interface {
	Id() ID

	// Find project by number.
	Project() *Project

	// A list of projects under the owner.
	Projects() ProjectConnection

	// The HTTP path listing owners projects
	ProjectsResourcePath() URI

	// The HTTP URL listing owners projects
	ProjectsUrl() URI

	// Can the current viewer create new projects on this owner.
	ViewerCanCreateProjects() bool
}

type Untyped_ProjectOwner map[string]interface{}

func (m Untyped_ProjectOwner) Id() ID {
	return m["id"].(ID)
}

func (m Untyped_ProjectOwner) Project() *Project {
	return m["project"].(*Project)
}

func (m Untyped_ProjectOwner) Projects() ProjectConnection {
	return m["projects"].(ProjectConnection)
}

func (m Untyped_ProjectOwner) ProjectsResourcePath() URI {
	return m["projectsResourcePath"].(URI)
}

func (m Untyped_ProjectOwner) ProjectsUrl() URI {
	return m["projectsUrl"].(URI)
}

func (m Untyped_ProjectOwner) ViewerCanCreateProjects() bool {
	return m["viewerCanCreateProjects"].(bool)
}

func ProjectOwner_UnmarshalJSON(data []byte) (ProjectOwner, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "Organization":
		var v Organization
		err = json.Unmarshal(data, &v)
		return v, err
	case "Repository":
		var v Repository
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v Untyped_ProjectOwner
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "ProjectOwner", t.Typename)
}

// Represents a Git object.
type GitObject interface {
	// An abbreviated version of the Git object ID
	AbbreviatedOid() string

	// The HTTP path for this Git object
	CommitResourcePath() URI

	// The HTTP URL for this Git object
	CommitUrl() URI

	Id() ID

	// The Git object ID
	Oid() GitObjectID

	// The Repository the Git object belongs to
	Repository() Repository
}

type Untyped_GitObject map[string]interface{}

func (m Untyped_GitObject) AbbreviatedOid() string {
	return m["abbreviatedOid"].(string)
}

func (m Untyped_GitObject) CommitResourcePath() URI {
	return m["commitResourcePath"].(URI)
}

func (m Untyped_GitObject) CommitUrl() URI {
	return m["commitUrl"].(URI)
}

func (m Untyped_GitObject) Id() ID {
	return m["id"].(ID)
}

func (m Untyped_GitObject) Oid() GitObjectID {
	return m["oid"].(GitObjectID)
}

func (m Untyped_GitObject) Repository() Repository {
	return m["repository"].(Repository)
}

func GitObject_UnmarshalJSON(data []byte) (GitObject, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "Blob":
		var v Blob
		err = json.Unmarshal(data, &v)
		return v, err
	case "Commit":
		var v Commit
		err = json.Unmarshal(data, &v)
		return v, err
	case "Tag":
		var v Tag
		err = json.Unmarshal(data, &v)
		return v, err
	case "Tree":
		var v Tree
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v Untyped_GitObject
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "GitObject", t.Typename)
}

// An object with an ID.
type Node interface {
	// ID of the object.
	Id() ID
}

type Untyped_Node map[string]interface{}

func (m Untyped_Node) Id() ID {
	return m["id"].(ID)
}

func Node_UnmarshalJSON(data []byte) (Node, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "AssignedEvent":
		var v AssignedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "BaseRefForcePushedEvent":
		var v BaseRefForcePushedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "Blob":
		var v Blob
		err = json.Unmarshal(data, &v)
		return v, err
	case "Bot":
		var v Bot
		err = json.Unmarshal(data, &v)
		return v, err
	case "ClosedEvent":
		var v ClosedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "Commit":
		var v Commit
		err = json.Unmarshal(data, &v)
		return v, err
	case "CommitComment":
		var v CommitComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "DemilestonedEvent":
		var v DemilestonedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "DeployedEvent":
		var v DeployedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "Deployment":
		var v Deployment
		err = json.Unmarshal(data, &v)
		return v, err
	case "DeploymentStatus":
		var v DeploymentStatus
		err = json.Unmarshal(data, &v)
		return v, err
	case "ExternalIdentity":
		var v ExternalIdentity
		err = json.Unmarshal(data, &v)
		return v, err
	case "Gist":
		var v Gist
		err = json.Unmarshal(data, &v)
		return v, err
	case "GistComment":
		var v GistComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "HeadRefDeletedEvent":
		var v HeadRefDeletedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "HeadRefForcePushedEvent":
		var v HeadRefForcePushedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "HeadRefRestoredEvent":
		var v HeadRefRestoredEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "Issue":
		var v Issue
		err = json.Unmarshal(data, &v)
		return v, err
	case "IssueComment":
		var v IssueComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "Label":
		var v Label
		err = json.Unmarshal(data, &v)
		return v, err
	case "LabeledEvent":
		var v LabeledEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "Language":
		var v Language
		err = json.Unmarshal(data, &v)
		return v, err
	case "LockedEvent":
		var v LockedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "MergedEvent":
		var v MergedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "Milestone":
		var v Milestone
		err = json.Unmarshal(data, &v)
		return v, err
	case "MilestonedEvent":
		var v MilestonedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "Organization":
		var v Organization
		err = json.Unmarshal(data, &v)
		return v, err
	case "OrganizationIdentityProvider":
		var v OrganizationIdentityProvider
		err = json.Unmarshal(data, &v)
		return v, err
	case "Project":
		var v Project
		err = json.Unmarshal(data, &v)
		return v, err
	case "ProjectCard":
		var v ProjectCard
		err = json.Unmarshal(data, &v)
		return v, err
	case "ProjectColumn":
		var v ProjectColumn
		err = json.Unmarshal(data, &v)
		return v, err
	case "ProtectedBranch":
		var v ProtectedBranch
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequest":
		var v PullRequest
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestCommit":
		var v PullRequestCommit
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestReview":
		var v PullRequestReview
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestReviewComment":
		var v PullRequestReviewComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestReviewThread":
		var v PullRequestReviewThread
		err = json.Unmarshal(data, &v)
		return v, err
	case "PushAllowance":
		var v PushAllowance
		err = json.Unmarshal(data, &v)
		return v, err
	case "Reaction":
		var v Reaction
		err = json.Unmarshal(data, &v)
		return v, err
	case "Ref":
		var v Ref
		err = json.Unmarshal(data, &v)
		return v, err
	case "ReferencedEvent":
		var v ReferencedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "Release":
		var v Release
		err = json.Unmarshal(data, &v)
		return v, err
	case "ReleaseAsset":
		var v ReleaseAsset
		err = json.Unmarshal(data, &v)
		return v, err
	case "RenamedTitleEvent":
		var v RenamedTitleEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "ReopenedEvent":
		var v ReopenedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "Repository":
		var v Repository
		err = json.Unmarshal(data, &v)
		return v, err
	case "RepositoryInvitation":
		var v RepositoryInvitation
		err = json.Unmarshal(data, &v)
		return v, err
	case "RepositoryTopic":
		var v RepositoryTopic
		err = json.Unmarshal(data, &v)
		return v, err
	case "ReviewDismissalAllowance":
		var v ReviewDismissalAllowance
		err = json.Unmarshal(data, &v)
		return v, err
	case "ReviewDismissedEvent":
		var v ReviewDismissedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "ReviewRequest":
		var v ReviewRequest
		err = json.Unmarshal(data, &v)
		return v, err
	case "ReviewRequestRemovedEvent":
		var v ReviewRequestRemovedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "ReviewRequestedEvent":
		var v ReviewRequestedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "Status":
		var v Status
		err = json.Unmarshal(data, &v)
		return v, err
	case "StatusContext":
		var v StatusContext
		err = json.Unmarshal(data, &v)
		return v, err
	case "SubscribedEvent":
		var v SubscribedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "Tag":
		var v Tag
		err = json.Unmarshal(data, &v)
		return v, err
	case "Team":
		var v Team
		err = json.Unmarshal(data, &v)
		return v, err
	case "Topic":
		var v Topic
		err = json.Unmarshal(data, &v)
		return v, err
	case "Tree":
		var v Tree
		err = json.Unmarshal(data, &v)
		return v, err
	case "UnassignedEvent":
		var v UnassignedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "UnlabeledEvent":
		var v UnlabeledEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "UnlockedEvent":
		var v UnlockedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "UnsubscribedEvent":
		var v UnsubscribedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "User":
		var v User
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v Untyped_Node
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "Node", t.Typename)
}

// Represents a subject that can be reacted on.
type Reactable interface {
	// Identifies the primary key from the database.
	DatabaseId() *int32

	Id() ID

	// A list of reactions grouped by content left on the subject.
	ReactionGroups() []ReactionGroup

	// A list of Reactions left on the Issue.
	Reactions() ReactionConnection

	// Can user react to this subject
	ViewerCanReact() bool
}

type Untyped_Reactable map[string]interface{}

func (m Untyped_Reactable) DatabaseId() *int32 {
	return m["databaseId"].(*int32)
}

func (m Untyped_Reactable) Id() ID {
	return m["id"].(ID)
}

func (m Untyped_Reactable) ReactionGroups() []ReactionGroup {
	return m["reactionGroups"].([]ReactionGroup)
}

func (m Untyped_Reactable) Reactions() ReactionConnection {
	return m["reactions"].(ReactionConnection)
}

func (m Untyped_Reactable) ViewerCanReact() bool {
	return m["viewerCanReact"].(bool)
}

func Reactable_UnmarshalJSON(data []byte) (Reactable, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "CommitComment":
		var v CommitComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "Issue":
		var v Issue
		err = json.Unmarshal(data, &v)
		return v, err
	case "IssueComment":
		var v IssueComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequest":
		var v PullRequest
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestReviewComment":
		var v PullRequestReviewComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v Untyped_Reactable
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "Reactable", t.Typename)
}

// Represents a comment.
type Comment interface {
	// The actor who authored the comment.
	Author() Actor

	// Author's association with the subject of the comment.
	AuthorAssociation() CommentAuthorAssociation

	// The comment body as Markdown.
	Body() string

	// The comment body rendered to HTML.
	BodyHTML() HTML

	// Identifies the date and time when the object was created.
	CreatedAt() DateTime

	// Check if this comment was created via an email reply.
	CreatedViaEmail() bool

	// The actor who edited the comment.
	Editor() Actor

	Id() ID

	// The moment the editor made the last edit
	LastEditedAt() *DateTime

	// Identifies when the comment was published at.
	PublishedAt() *DateTime

	// Identifies the date and time when the object was last updated.
	UpdatedAt() DateTime

	// Did the viewer author this comment.
	ViewerDidAuthor() bool
}

type Untyped_Comment map[string]interface{}

func (m Untyped_Comment) Author() Actor {
	return m["author"].(Actor)
}

func (m Untyped_Comment) AuthorAssociation() CommentAuthorAssociation {
	return m["authorAssociation"].(CommentAuthorAssociation)
}

func (m Untyped_Comment) Body() string {
	return m["body"].(string)
}

func (m Untyped_Comment) BodyHTML() HTML {
	return m["bodyHTML"].(HTML)
}

func (m Untyped_Comment) CreatedAt() DateTime {
	return m["createdAt"].(DateTime)
}

func (m Untyped_Comment) CreatedViaEmail() bool {
	return m["createdViaEmail"].(bool)
}

func (m Untyped_Comment) Editor() Actor {
	return m["editor"].(Actor)
}

func (m Untyped_Comment) Id() ID {
	return m["id"].(ID)
}

func (m Untyped_Comment) LastEditedAt() *DateTime {
	return m["lastEditedAt"].(*DateTime)
}

func (m Untyped_Comment) PublishedAt() *DateTime {
	return m["publishedAt"].(*DateTime)
}

func (m Untyped_Comment) UpdatedAt() DateTime {
	return m["updatedAt"].(DateTime)
}

func (m Untyped_Comment) ViewerDidAuthor() bool {
	return m["viewerDidAuthor"].(bool)
}

func Comment_UnmarshalJSON(data []byte) (Comment, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "CommitComment":
		var v CommitComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "GistComment":
		var v GistComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "Issue":
		var v Issue
		err = json.Unmarshal(data, &v)
		return v, err
	case "IssueComment":
		var v IssueComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequest":
		var v PullRequest
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestReview":
		var v PullRequestReview
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestReviewComment":
		var v PullRequestReviewComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v Untyped_Comment
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "Comment", t.Typename)
}

// Entities that can be deleted.
type Deletable interface {
	// Check if the current viewer can delete this object.
	ViewerCanDelete() bool
}

type Untyped_Deletable map[string]interface{}

func (m Untyped_Deletable) ViewerCanDelete() bool {
	return m["viewerCanDelete"].(bool)
}

func Deletable_UnmarshalJSON(data []byte) (Deletable, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "CommitComment":
		var v CommitComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "GistComment":
		var v GistComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "IssueComment":
		var v IssueComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestReview":
		var v PullRequestReview
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestReviewComment":
		var v PullRequestReviewComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v Untyped_Deletable
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "Deletable", t.Typename)
}

// Entities that can be updated.
type Updatable interface {
	// Check if the current viewer can update this object.
	ViewerCanUpdate() bool
}

type Untyped_Updatable map[string]interface{}

func (m Untyped_Updatable) ViewerCanUpdate() bool {
	return m["viewerCanUpdate"].(bool)
}

func Updatable_UnmarshalJSON(data []byte) (Updatable, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "CommitComment":
		var v CommitComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "GistComment":
		var v GistComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "Issue":
		var v Issue
		err = json.Unmarshal(data, &v)
		return v, err
	case "IssueComment":
		var v IssueComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "Project":
		var v Project
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequest":
		var v PullRequest
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestReview":
		var v PullRequestReview
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestReviewComment":
		var v PullRequestReviewComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v Untyped_Updatable
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "Updatable", t.Typename)
}

// Comments that can be updated.
type UpdatableComment interface {
	// Reasons why the current viewer can not update this comment.
	ViewerCannotUpdateReasons() []CommentCannotUpdateReason
}

type Untyped_UpdatableComment map[string]interface{}

func (m Untyped_UpdatableComment) ViewerCannotUpdateReasons() []CommentCannotUpdateReason {
	return m["viewerCannotUpdateReasons"].([]CommentCannotUpdateReason)
}

func UpdatableComment_UnmarshalJSON(data []byte) (UpdatableComment, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "CommitComment":
		var v CommitComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "GistComment":
		var v GistComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "Issue":
		var v Issue
		err = json.Unmarshal(data, &v)
		return v, err
	case "IssueComment":
		var v IssueComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequest":
		var v PullRequest
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestReview":
		var v PullRequestReview
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestReviewComment":
		var v PullRequestReviewComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v Untyped_UpdatableComment
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "UpdatableComment", t.Typename)
}

// Represents a object that belongs to a repository.
type RepositoryNode interface {
	// The repository associated with this node.
	Repository() Repository
}

type Untyped_RepositoryNode map[string]interface{}

func (m Untyped_RepositoryNode) Repository() Repository {
	return m["repository"].(Repository)
}

func RepositoryNode_UnmarshalJSON(data []byte) (RepositoryNode, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "CommitComment":
		var v CommitComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "Issue":
		var v Issue
		err = json.Unmarshal(data, &v)
		return v, err
	case "IssueComment":
		var v IssueComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequest":
		var v PullRequest
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestReview":
		var v PullRequestReview
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestReviewComment":
		var v PullRequestReviewComment
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v Untyped_RepositoryNode
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "RepositoryNode", t.Typename)
}

// Information about a signature (GPG or S/MIME) on a Commit or Tag.
type GitSignature interface {
	// Email used to sign this object.
	Email() string

	// True if the signature is valid and verified by GitHub.
	IsValid() bool

	// Payload for GPG signing object. Raw ODB object without the signature header.
	Payload() string

	// ASCII-armored signature header from object.
	Signature() string

	// GitHub user corresponding to the email signing this commit.
	Signer() *User

	// The state of this signature. `VALID` if signature is valid and verified by GitHub, otherwise represents reason why signature is considered invalid.
	State() GitSignatureState
}

type Untyped_GitSignature map[string]interface{}

func (m Untyped_GitSignature) Email() string {
	return m["email"].(string)
}

func (m Untyped_GitSignature) IsValid() bool {
	return m["isValid"].(bool)
}

func (m Untyped_GitSignature) Payload() string {
	return m["payload"].(string)
}

func (m Untyped_GitSignature) Signature() string {
	return m["signature"].(string)
}

func (m Untyped_GitSignature) Signer() *User {
	return m["signer"].(*User)
}

func (m Untyped_GitSignature) State() GitSignatureState {
	return m["state"].(GitSignatureState)
}

func GitSignature_UnmarshalJSON(data []byte) (GitSignature, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "GpgSignature":
		var v GpgSignature
		err = json.Unmarshal(data, &v)
		return v, err
	case "SmimeSignature":
		var v SmimeSignature
		err = json.Unmarshal(data, &v)
		return v, err
	case "UnknownSignature":
		var v UnknownSignature
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v Untyped_GitSignature
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "GitSignature", t.Typename)
}

// Entities that can be subscribed to for web and email notifications.
type Subscribable interface {
	// Check if the viewer is able to change their subscription status for the repository.
	ViewerCanSubscribe() bool

	// Identifies if the viewer is watching, not watching, or ignoring the repository.
	ViewerSubscription() SubscriptionState
}

type Untyped_Subscribable map[string]interface{}

func (m Untyped_Subscribable) ViewerCanSubscribe() bool {
	return m["viewerCanSubscribe"].(bool)
}

func (m Untyped_Subscribable) ViewerSubscription() SubscriptionState {
	return m["viewerSubscription"].(SubscriptionState)
}

func Subscribable_UnmarshalJSON(data []byte) (Subscribable, error) {
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
	case "":
		var v Untyped_Subscribable
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "Subscribable", t.Typename)
}

// Represents a type that can be retrieved by a URL.
type UniformResourceLocatable interface {
	// The HTML path to this resource.
	ResourcePath() URI

	// The URL to this resource.
	Url() URI
}

type Untyped_UniformResourceLocatable map[string]interface{}

func (m Untyped_UniformResourceLocatable) ResourcePath() URI {
	return m["resourcePath"].(URI)
}

func (m Untyped_UniformResourceLocatable) Url() URI {
	return m["url"].(URI)
}

func UniformResourceLocatable_UnmarshalJSON(data []byte) (UniformResourceLocatable, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "Bot":
		var v Bot
		err = json.Unmarshal(data, &v)
		return v, err
	case "Issue":
		var v Issue
		err = json.Unmarshal(data, &v)
		return v, err
	case "MergedEvent":
		var v MergedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "Organization":
		var v Organization
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequest":
		var v PullRequest
		err = json.Unmarshal(data, &v)
		return v, err
	case "PullRequestCommit":
		var v PullRequestCommit
		err = json.Unmarshal(data, &v)
		return v, err
	case "Release":
		var v Release
		err = json.Unmarshal(data, &v)
		return v, err
	case "Repository":
		var v Repository
		err = json.Unmarshal(data, &v)
		return v, err
	case "RepositoryTopic":
		var v RepositoryTopic
		err = json.Unmarshal(data, &v)
		return v, err
	case "ReviewDismissedEvent":
		var v ReviewDismissedEvent
		err = json.Unmarshal(data, &v)
		return v, err
	case "User":
		var v User
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v Untyped_UniformResourceLocatable
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "UniformResourceLocatable", t.Typename)
}

// Represents an owner of a Repository.
type RepositoryOwner interface {
	// A URL pointing to the owner's public avatar.
	AvatarUrl() URI

	Id() ID

	// The username used to login.
	Login() string

	// A list of repositories that the user owns.
	Repositories() RepositoryConnection

	// Find Repository.
	Repository() *Repository

	// The HTTP URL for the owner.
	ResourcePath() URI

	// The HTTP URL for the owner.
	Url() URI
}

type Untyped_RepositoryOwner map[string]interface{}

func (m Untyped_RepositoryOwner) AvatarUrl() URI {
	return m["avatarUrl"].(URI)
}

func (m Untyped_RepositoryOwner) Id() ID {
	return m["id"].(ID)
}

func (m Untyped_RepositoryOwner) Login() string {
	return m["login"].(string)
}

func (m Untyped_RepositoryOwner) Repositories() RepositoryConnection {
	return m["repositories"].(RepositoryConnection)
}

func (m Untyped_RepositoryOwner) Repository() *Repository {
	return m["repository"].(*Repository)
}

func (m Untyped_RepositoryOwner) ResourcePath() URI {
	return m["resourcePath"].(URI)
}

func (m Untyped_RepositoryOwner) Url() URI {
	return m["url"].(URI)
}

func RepositoryOwner_UnmarshalJSON(data []byte) (RepositoryOwner, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "Organization":
		var v Organization
		err = json.Unmarshal(data, &v)
		return v, err
	case "User":
		var v User
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v Untyped_RepositoryOwner
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "RepositoryOwner", t.Typename)
}

// An object that can be closed
type Closable interface {
	// `true` if the object is closed (definition of closed may depend on type)
	Closed() bool
}

type Untyped_Closable map[string]interface{}

func (m Untyped_Closable) Closed() bool {
	return m["closed"].(bool)
}

func Closable_UnmarshalJSON(data []byte) (Closable, error) {
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
		var v Untyped_Closable
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "Closable", t.Typename)
}

// An object that can have users assigned to it.
type Assignable interface {
	// A list of Users assigned to this object.
	Assignees() UserConnection
}

type Untyped_Assignable map[string]interface{}

func (m Untyped_Assignable) Assignees() UserConnection {
	return m["assignees"].(UserConnection)
}

func Assignable_UnmarshalJSON(data []byte) (Assignable, error) {
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
		var v Untyped_Assignable
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "Assignable", t.Typename)
}

// An object that can have labels assigned to it.
type Labelable interface {
	// A list of labels associated with the object.
	Labels() *LabelConnection
}

type Untyped_Labelable map[string]interface{}

func (m Untyped_Labelable) Labels() *LabelConnection {
	return m["labels"].(*LabelConnection)
}

func Labelable_UnmarshalJSON(data []byte) (Labelable, error) {
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
		var v Untyped_Labelable
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "Labelable", t.Typename)
}

// An object that can be locked.
type Lockable interface {
	// `true` if the object is locked
	Locked() bool
}

type Untyped_Lockable map[string]interface{}

func (m Untyped_Lockable) Locked() bool {
	return m["locked"].(bool)
}

func Lockable_UnmarshalJSON(data []byte) (Lockable, error) {
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
		var v Untyped_Lockable
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "Lockable", t.Typename)
}

// Things that can be starred.
type Starrable interface {
	Id() ID

	// A list of users who have starred this starrable.
	Stargazers() StargazerConnection

	// Returns a boolean indicating whether the viewing user has starred this starrable.
	ViewerHasStarred() bool
}

type Untyped_Starrable map[string]interface{}

func (m Untyped_Starrable) Id() ID {
	return m["id"].(ID)
}

func (m Untyped_Starrable) Stargazers() StargazerConnection {
	return m["stargazers"].(StargazerConnection)
}

func (m Untyped_Starrable) ViewerHasStarred() bool {
	return m["viewerHasStarred"].(bool)
}

func Starrable_UnmarshalJSON(data []byte) (Starrable, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "Gist":
		var v Gist
		err = json.Unmarshal(data, &v)
		return v, err
	case "Repository":
		var v Repository
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v Untyped_Starrable
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "Starrable", t.Typename)
}

// A subset of repository info.
type RepositoryInfo interface {
	// Identifies the date and time when the object was created.
	CreatedAt() DateTime

	// The description of the repository.
	Description() *string

	// The description of the repository rendered to HTML.
	DescriptionHTML() HTML

	// Indicates if the repository has issues feature enabled.
	HasIssuesEnabled() bool

	// Indicates if the repository has wiki feature enabled.
	HasWikiEnabled() bool

	// The repository's URL.
	HomepageUrl() *URI

	// Identifies if the repository is a fork.
	IsFork() bool

	// Indicates if the repository has been locked or not.
	IsLocked() bool

	// Identifies if the repository is a mirror.
	IsMirror() bool

	// Identifies if the repository is private.
	IsPrivate() bool

	// The license associated with the repository
	License() *string

	// The reason the repository has been locked.
	LockReason() *RepositoryLockReason

	// The repository's original mirror URL.
	MirrorUrl() *URI

	// The name of the repository.
	Name() string

	// The repository's name with owner.
	NameWithOwner() string

	// The User owner of the repository.
	Owner() RepositoryOwner

	// Identifies when the repository was last pushed to.
	PushedAt() *DateTime

	// The HTTP path for this repository
	ResourcePath() URI

	// Identifies the date and time when the object was last updated.
	UpdatedAt() DateTime

	// The HTTP URL for this repository
	Url() URI
}

type Untyped_RepositoryInfo map[string]interface{}

func (m Untyped_RepositoryInfo) CreatedAt() DateTime {
	return m["createdAt"].(DateTime)
}

func (m Untyped_RepositoryInfo) Description() *string {
	return m["description"].(*string)
}

func (m Untyped_RepositoryInfo) DescriptionHTML() HTML {
	return m["descriptionHTML"].(HTML)
}

func (m Untyped_RepositoryInfo) HasIssuesEnabled() bool {
	return m["hasIssuesEnabled"].(bool)
}

func (m Untyped_RepositoryInfo) HasWikiEnabled() bool {
	return m["hasWikiEnabled"].(bool)
}

func (m Untyped_RepositoryInfo) HomepageUrl() *URI {
	return m["homepageUrl"].(*URI)
}

func (m Untyped_RepositoryInfo) IsFork() bool {
	return m["isFork"].(bool)
}

func (m Untyped_RepositoryInfo) IsLocked() bool {
	return m["isLocked"].(bool)
}

func (m Untyped_RepositoryInfo) IsMirror() bool {
	return m["isMirror"].(bool)
}

func (m Untyped_RepositoryInfo) IsPrivate() bool {
	return m["isPrivate"].(bool)
}

func (m Untyped_RepositoryInfo) License() *string {
	return m["license"].(*string)
}

func (m Untyped_RepositoryInfo) LockReason() *RepositoryLockReason {
	return m["lockReason"].(*RepositoryLockReason)
}

func (m Untyped_RepositoryInfo) MirrorUrl() *URI {
	return m["mirrorUrl"].(*URI)
}

func (m Untyped_RepositoryInfo) Name() string {
	return m["name"].(string)
}

func (m Untyped_RepositoryInfo) NameWithOwner() string {
	return m["nameWithOwner"].(string)
}

func (m Untyped_RepositoryInfo) Owner() RepositoryOwner {
	return m["owner"].(RepositoryOwner)
}

func (m Untyped_RepositoryInfo) PushedAt() *DateTime {
	return m["pushedAt"].(*DateTime)
}

func (m Untyped_RepositoryInfo) ResourcePath() URI {
	return m["resourcePath"].(URI)
}

func (m Untyped_RepositoryInfo) UpdatedAt() DateTime {
	return m["updatedAt"].(DateTime)
}

func (m Untyped_RepositoryInfo) Url() URI {
	return m["url"].(URI)
}

func RepositoryInfo_UnmarshalJSON(data []byte) (RepositoryInfo, error) {
	var t struct {
		Typename string `json:"__typename"`
	}
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	switch t.Typename {
	case "Repository":
		var v Repository
		err = json.Unmarshal(data, &v)
		return v, err
	case "RepositoryInvitationRepository":
		var v RepositoryInvitationRepository
		err = json.Unmarshal(data, &v)
		return v, err
	case "":
		var v Untyped_RepositoryInfo
		err = json.Unmarshal(data, &v)
		return v, err
	}
	return nil, fmt.Errorf("unexpected __typename for interface %s: %s", "RepositoryInfo", t.Typename)
}
