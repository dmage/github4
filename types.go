package github4

import (
	"encoding/json"
)

// A repository contains the content for a project.
type Repository struct {
	// Returns the code of conduct for this repository
	field_codeOfConduct *CodeOfConduct `json:"codeOfConduct"`

	// A list of commit comments associated with the repository.
	field_commitComments CommitCommentConnection `json:"commitComments"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// Identifies the primary key from the database.
	field_databaseId *int32 `json:"databaseId"`

	// The Ref associated with the repository's default branch.
	field_defaultBranchRef *Ref `json:"defaultBranchRef"`

	// Deployments associated with the repository
	field_deployments DeploymentConnection `json:"deployments"`

	// The description of the repository.
	field_description *string `json:"description"`

	// The description of the repository rendered to HTML.
	field_descriptionHTML HTML `json:"descriptionHTML"`

	// The number of kilobytes this repository occupies on disk.
	field_diskUsage *int32 `json:"diskUsage"`

	// A list of forked repositories.
	field_forks RepositoryConnection `json:"forks"`

	// Indicates if the repository has issues feature enabled.
	field_hasIssuesEnabled bool `json:"hasIssuesEnabled"`

	// Indicates if the repository has wiki feature enabled.
	field_hasWikiEnabled bool `json:"hasWikiEnabled"`

	// The repository's URL.
	field_homepageUrl *URI `json:"homepageUrl"`

	field_id ID `json:"id"`

	// Identifies if the repository is a fork.
	field_isFork bool `json:"isFork"`

	// Indicates if the repository has been locked or not.
	field_isLocked bool `json:"isLocked"`

	// Identifies if the repository is a mirror.
	field_isMirror bool `json:"isMirror"`

	// Identifies if the repository is private.
	field_isPrivate bool `json:"isPrivate"`

	// Returns a single issue from the current repository by number.
	field_issue *Issue `json:"issue"`

	// Returns a single issue-like object from the current repository by number.
	field_issueOrPullRequest IssueOrPullRequest `json:"issueOrPullRequest"`

	// A list of issues that have been opened in the repository.
	field_issues IssueConnection `json:"issues"`

	// Returns a single label by name
	field_label *Label `json:"label"`

	// A list of labels associated with the repository.
	field_labels *LabelConnection `json:"labels"`

	// A list containing a breakdown of the language composition of the repository.
	field_languages *LanguageConnection `json:"languages"`

	// The license associated with the repository
	field_license *string `json:"license"`

	// The reason the repository has been locked.
	field_lockReason *RepositoryLockReason `json:"lockReason"`

	// A list of Users that can be mentioned in the context of the repository.
	field_mentionableUsers UserConnection `json:"mentionableUsers"`

	// Returns a single milestone from the current repository by number.
	field_milestone *Milestone `json:"milestone"`

	// A list of milestones associated with the repository.
	field_milestones *MilestoneConnection `json:"milestones"`

	// The repository's original mirror URL.
	field_mirrorUrl *URI `json:"mirrorUrl"`

	// The name of the repository.
	field_name string `json:"name"`

	// The repository's name with owner.
	field_nameWithOwner string `json:"nameWithOwner"`

	// A Git object in the repository
	field_object GitObject `json:"object"`

	// The User owner of the repository.
	field_owner RepositoryOwner `json:"owner"`

	// The repository parent, if this is a fork.
	field_parent *Repository `json:"parent"`

	// The primary language of the repository's code.
	field_primaryLanguage *Language `json:"primaryLanguage"`

	// Find project by number.
	field_project *Project `json:"project"`

	// A list of projects under the owner.
	field_projects ProjectConnection `json:"projects"`

	// The HTTP path listing repository's projects
	field_projectsResourcePath URI `json:"projectsResourcePath"`

	// The HTTP URL listing repository's projects
	field_projectsUrl URI `json:"projectsUrl"`

	// A list of protected branches that are on this repository.
	field_protectedBranches ProtectedBranchConnection `json:"protectedBranches"`

	// Returns a single pull request from the current repository by number.
	field_pullRequest *PullRequest `json:"pullRequest"`

	// A list of pull requests that have been opened in the repository.
	field_pullRequests PullRequestConnection `json:"pullRequests"`

	// Identifies when the repository was last pushed to.
	field_pushedAt *DateTime `json:"pushedAt"`

	// Fetch a given ref from the repository
	field_ref *Ref `json:"ref"`

	// Fetch a list of refs from the repository
	field_refs *RefConnection `json:"refs"`

	// List of releases which are dependent on this repository.
	field_releases ReleaseConnection `json:"releases"`

	// A list of applied repository-topic associations for this repository.
	field_repositoryTopics RepositoryTopicConnection `json:"repositoryTopics"`

	// The HTTP path for this repository
	field_resourcePath URI `json:"resourcePath"`

	// A list of users who have starred this starrable.
	field_stargazers StargazerConnection `json:"stargazers"`

	// Identifies the date and time when the object was last updated.
	field_updatedAt DateTime `json:"updatedAt"`

	// The HTTP URL for this repository
	field_url URI `json:"url"`

	// Indicates whether the viewer has admin permissions on this repository.
	field_viewerCanAdminister bool `json:"viewerCanAdminister"`

	// Can the current viewer create new projects on this owner.
	field_viewerCanCreateProjects bool `json:"viewerCanCreateProjects"`

	// Check if the viewer is able to change their subscription status for the repository.
	field_viewerCanSubscribe bool `json:"viewerCanSubscribe"`

	// Indicates whether the viewer can update the topics of this repository.
	field_viewerCanUpdateTopics bool `json:"viewerCanUpdateTopics"`

	// Returns a boolean indicating whether the viewing user has starred this starrable.
	field_viewerHasStarred bool `json:"viewerHasStarred"`

	// Identifies if the viewer is watching, not watching, or ignoring the repository.
	field_viewerSubscription SubscriptionState `json:"viewerSubscription"`

	// A list of users watching the repository.
	field_watchers UserConnection `json:"watchers"`
}

func (o Repository) CodeOfConduct() *CodeOfConduct {
	return o.field_codeOfConduct
}

func (o Repository) CommitComments() CommitCommentConnection {
	return o.field_commitComments
}

func (o Repository) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o Repository) DatabaseId() *int32 {
	return o.field_databaseId
}

func (o Repository) DefaultBranchRef() *Ref {
	return o.field_defaultBranchRef
}

func (o Repository) Deployments() DeploymentConnection {
	return o.field_deployments
}

func (o Repository) Description() *string {
	return o.field_description
}

func (o Repository) DescriptionHTML() HTML {
	return o.field_descriptionHTML
}

func (o Repository) DiskUsage() *int32 {
	return o.field_diskUsage
}

func (o Repository) Forks() RepositoryConnection {
	return o.field_forks
}

func (o Repository) HasIssuesEnabled() bool {
	return o.field_hasIssuesEnabled
}

func (o Repository) HasWikiEnabled() bool {
	return o.field_hasWikiEnabled
}

func (o Repository) HomepageUrl() *URI {
	return o.field_homepageUrl
}

func (o Repository) Id() ID {
	return o.field_id
}

func (o Repository) IsFork() bool {
	return o.field_isFork
}

func (o Repository) IsLocked() bool {
	return o.field_isLocked
}

func (o Repository) IsMirror() bool {
	return o.field_isMirror
}

func (o Repository) IsPrivate() bool {
	return o.field_isPrivate
}

func (o Repository) Issue() *Issue {
	return o.field_issue
}

func (o Repository) IssueOrPullRequest() IssueOrPullRequest {
	return o.field_issueOrPullRequest
}

func (o Repository) Issues() IssueConnection {
	return o.field_issues
}

func (o Repository) Label() *Label {
	return o.field_label
}

func (o Repository) Labels() *LabelConnection {
	return o.field_labels
}

func (o Repository) Languages() *LanguageConnection {
	return o.field_languages
}

func (o Repository) License() *string {
	return o.field_license
}

func (o Repository) LockReason() *RepositoryLockReason {
	return o.field_lockReason
}

func (o Repository) MentionableUsers() UserConnection {
	return o.field_mentionableUsers
}

func (o Repository) Milestone() *Milestone {
	return o.field_milestone
}

func (o Repository) Milestones() *MilestoneConnection {
	return o.field_milestones
}

func (o Repository) MirrorUrl() *URI {
	return o.field_mirrorUrl
}

func (o Repository) Name() string {
	return o.field_name
}

func (o Repository) NameWithOwner() string {
	return o.field_nameWithOwner
}

func (o Repository) Object() GitObject {
	return o.field_object
}

func (o Repository) Owner() RepositoryOwner {
	return o.field_owner
}

func (o Repository) Parent() *Repository {
	return o.field_parent
}

func (o Repository) PrimaryLanguage() *Language {
	return o.field_primaryLanguage
}

func (o Repository) Project() *Project {
	return o.field_project
}

func (o Repository) Projects() ProjectConnection {
	return o.field_projects
}

func (o Repository) ProjectsResourcePath() URI {
	return o.field_projectsResourcePath
}

func (o Repository) ProjectsUrl() URI {
	return o.field_projectsUrl
}

func (o Repository) ProtectedBranches() ProtectedBranchConnection {
	return o.field_protectedBranches
}

func (o Repository) PullRequest() *PullRequest {
	return o.field_pullRequest
}

func (o Repository) PullRequests() PullRequestConnection {
	return o.field_pullRequests
}

func (o Repository) PushedAt() *DateTime {
	return o.field_pushedAt
}

func (o Repository) Ref() *Ref {
	return o.field_ref
}

func (o Repository) Refs() *RefConnection {
	return o.field_refs
}

func (o Repository) Releases() ReleaseConnection {
	return o.field_releases
}

func (o Repository) RepositoryTopics() RepositoryTopicConnection {
	return o.field_repositoryTopics
}

func (o Repository) ResourcePath() URI {
	return o.field_resourcePath
}

func (o Repository) Stargazers() StargazerConnection {
	return o.field_stargazers
}

func (o Repository) UpdatedAt() DateTime {
	return o.field_updatedAt
}

func (o Repository) Url() URI {
	return o.field_url
}

func (o Repository) ViewerCanAdminister() bool {
	return o.field_viewerCanAdminister
}

func (o Repository) ViewerCanCreateProjects() bool {
	return o.field_viewerCanCreateProjects
}

func (o Repository) ViewerCanSubscribe() bool {
	return o.field_viewerCanSubscribe
}

func (o Repository) ViewerCanUpdateTopics() bool {
	return o.field_viewerCanUpdateTopics
}

func (o Repository) ViewerHasStarred() bool {
	return o.field_viewerHasStarred
}

func (o Repository) ViewerSubscription() SubscriptionState {
	return o.field_viewerSubscription
}

func (o Repository) Watchers() UserConnection {
	return o.field_watchers
}

func (o *Repository) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_codeOfConduct           *CodeOfConduct            `json:"codeOfConduct"`
		Field_commitComments          CommitCommentConnection   `json:"commitComments"`
		Field_createdAt               DateTime                  `json:"createdAt"`
		Field_databaseId              *int32                    `json:"databaseId"`
		Field_defaultBranchRef        *Ref                      `json:"defaultBranchRef"`
		Field_deployments             DeploymentConnection      `json:"deployments"`
		Field_description             *string                   `json:"description"`
		Field_descriptionHTML         HTML                      `json:"descriptionHTML"`
		Field_diskUsage               *int32                    `json:"diskUsage"`
		Field_forks                   RepositoryConnection      `json:"forks"`
		Field_hasIssuesEnabled        bool                      `json:"hasIssuesEnabled"`
		Field_hasWikiEnabled          bool                      `json:"hasWikiEnabled"`
		Field_homepageUrl             *URI                      `json:"homepageUrl"`
		Field_id                      ID                        `json:"id"`
		Field_isFork                  bool                      `json:"isFork"`
		Field_isLocked                bool                      `json:"isLocked"`
		Field_isMirror                bool                      `json:"isMirror"`
		Field_isPrivate               bool                      `json:"isPrivate"`
		Field_issue                   *Issue                    `json:"issue"`
		Field_issueOrPullRequest      json.RawMessage           `json:"issueOrPullRequest"`
		Field_issues                  IssueConnection           `json:"issues"`
		Field_label                   *Label                    `json:"label"`
		Field_labels                  *LabelConnection          `json:"labels"`
		Field_languages               *LanguageConnection       `json:"languages"`
		Field_license                 *string                   `json:"license"`
		Field_lockReason              *RepositoryLockReason     `json:"lockReason"`
		Field_mentionableUsers        UserConnection            `json:"mentionableUsers"`
		Field_milestone               *Milestone                `json:"milestone"`
		Field_milestones              *MilestoneConnection      `json:"milestones"`
		Field_mirrorUrl               *URI                      `json:"mirrorUrl"`
		Field_name                    string                    `json:"name"`
		Field_nameWithOwner           string                    `json:"nameWithOwner"`
		Field_object                  json.RawMessage           `json:"object"`
		Field_owner                   RepositoryOwner           `json:"owner"`
		Field_parent                  *Repository               `json:"parent"`
		Field_primaryLanguage         *Language                 `json:"primaryLanguage"`
		Field_project                 *Project                  `json:"project"`
		Field_projects                ProjectConnection         `json:"projects"`
		Field_projectsResourcePath    URI                       `json:"projectsResourcePath"`
		Field_projectsUrl             URI                       `json:"projectsUrl"`
		Field_protectedBranches       ProtectedBranchConnection `json:"protectedBranches"`
		Field_pullRequest             *PullRequest              `json:"pullRequest"`
		Field_pullRequests            PullRequestConnection     `json:"pullRequests"`
		Field_pushedAt                *DateTime                 `json:"pushedAt"`
		Field_ref                     *Ref                      `json:"ref"`
		Field_refs                    *RefConnection            `json:"refs"`
		Field_releases                ReleaseConnection         `json:"releases"`
		Field_repositoryTopics        RepositoryTopicConnection `json:"repositoryTopics"`
		Field_resourcePath            URI                       `json:"resourcePath"`
		Field_stargazers              StargazerConnection       `json:"stargazers"`
		Field_updatedAt               DateTime                  `json:"updatedAt"`
		Field_url                     URI                       `json:"url"`
		Field_viewerCanAdminister     bool                      `json:"viewerCanAdminister"`
		Field_viewerCanCreateProjects bool                      `json:"viewerCanCreateProjects"`
		Field_viewerCanSubscribe      bool                      `json:"viewerCanSubscribe"`
		Field_viewerCanUpdateTopics   bool                      `json:"viewerCanUpdateTopics"`
		Field_viewerHasStarred        bool                      `json:"viewerHasStarred"`
		Field_viewerSubscription      SubscriptionState         `json:"viewerSubscription"`
		Field_watchers                UserConnection            `json:"watchers"`
	}{
		Field_codeOfConduct:           o.field_codeOfConduct,
		Field_commitComments:          o.field_commitComments,
		Field_createdAt:               o.field_createdAt,
		Field_databaseId:              o.field_databaseId,
		Field_defaultBranchRef:        o.field_defaultBranchRef,
		Field_deployments:             o.field_deployments,
		Field_description:             o.field_description,
		Field_descriptionHTML:         o.field_descriptionHTML,
		Field_diskUsage:               o.field_diskUsage,
		Field_forks:                   o.field_forks,
		Field_hasIssuesEnabled:        o.field_hasIssuesEnabled,
		Field_hasWikiEnabled:          o.field_hasWikiEnabled,
		Field_homepageUrl:             o.field_homepageUrl,
		Field_id:                      o.field_id,
		Field_isFork:                  o.field_isFork,
		Field_isLocked:                o.field_isLocked,
		Field_isMirror:                o.field_isMirror,
		Field_isPrivate:               o.field_isPrivate,
		Field_issue:                   o.field_issue,
		Field_issues:                  o.field_issues,
		Field_label:                   o.field_label,
		Field_labels:                  o.field_labels,
		Field_languages:               o.field_languages,
		Field_license:                 o.field_license,
		Field_lockReason:              o.field_lockReason,
		Field_mentionableUsers:        o.field_mentionableUsers,
		Field_milestone:               o.field_milestone,
		Field_milestones:              o.field_milestones,
		Field_mirrorUrl:               o.field_mirrorUrl,
		Field_name:                    o.field_name,
		Field_nameWithOwner:           o.field_nameWithOwner,
		Field_owner:                   o.field_owner,
		Field_parent:                  o.field_parent,
		Field_primaryLanguage:         o.field_primaryLanguage,
		Field_project:                 o.field_project,
		Field_projects:                o.field_projects,
		Field_projectsResourcePath:    o.field_projectsResourcePath,
		Field_projectsUrl:             o.field_projectsUrl,
		Field_protectedBranches:       o.field_protectedBranches,
		Field_pullRequest:             o.field_pullRequest,
		Field_pullRequests:            o.field_pullRequests,
		Field_pushedAt:                o.field_pushedAt,
		Field_ref:                     o.field_ref,
		Field_refs:                    o.field_refs,
		Field_releases:                o.field_releases,
		Field_repositoryTopics:        o.field_repositoryTopics,
		Field_resourcePath:            o.field_resourcePath,
		Field_stargazers:              o.field_stargazers,
		Field_updatedAt:               o.field_updatedAt,
		Field_url:                     o.field_url,
		Field_viewerCanAdminister:     o.field_viewerCanAdminister,
		Field_viewerCanCreateProjects: o.field_viewerCanCreateProjects,
		Field_viewerCanSubscribe:      o.field_viewerCanSubscribe,
		Field_viewerCanUpdateTopics:   o.field_viewerCanUpdateTopics,
		Field_viewerHasStarred:        o.field_viewerHasStarred,
		Field_viewerSubscription:      o.field_viewerSubscription,
		Field_watchers:                o.field_watchers,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_codeOfConduct = v.Field_codeOfConduct
	o.field_commitComments = v.Field_commitComments
	o.field_createdAt = v.Field_createdAt
	o.field_databaseId = v.Field_databaseId
	o.field_defaultBranchRef = v.Field_defaultBranchRef
	o.field_deployments = v.Field_deployments
	o.field_description = v.Field_description
	o.field_descriptionHTML = v.Field_descriptionHTML
	o.field_diskUsage = v.Field_diskUsage
	o.field_forks = v.Field_forks
	o.field_hasIssuesEnabled = v.Field_hasIssuesEnabled
	o.field_hasWikiEnabled = v.Field_hasWikiEnabled
	o.field_homepageUrl = v.Field_homepageUrl
	o.field_id = v.Field_id
	o.field_isFork = v.Field_isFork
	o.field_isLocked = v.Field_isLocked
	o.field_isMirror = v.Field_isMirror
	o.field_isPrivate = v.Field_isPrivate
	o.field_issue = v.Field_issue
	if len(v.Field_issueOrPullRequest) > 0 {
		o.field_issueOrPullRequest, err = IssueOrPullRequest_UnmarshalJSON(v.Field_issueOrPullRequest)
		if err != nil {
			return err
		}
	}
	o.field_issues = v.Field_issues
	o.field_label = v.Field_label
	o.field_labels = v.Field_labels
	o.field_languages = v.Field_languages
	o.field_license = v.Field_license
	o.field_lockReason = v.Field_lockReason
	o.field_mentionableUsers = v.Field_mentionableUsers
	o.field_milestone = v.Field_milestone
	o.field_milestones = v.Field_milestones
	o.field_mirrorUrl = v.Field_mirrorUrl
	o.field_name = v.Field_name
	o.field_nameWithOwner = v.Field_nameWithOwner
	if len(v.Field_object) > 0 {
		o.field_object, err = GitObject_UnmarshalJSON(v.Field_object)
		if err != nil {
			return err
		}
	}
	o.field_owner = v.Field_owner
	o.field_parent = v.Field_parent
	o.field_primaryLanguage = v.Field_primaryLanguage
	o.field_project = v.Field_project
	o.field_projects = v.Field_projects
	o.field_projectsResourcePath = v.Field_projectsResourcePath
	o.field_projectsUrl = v.Field_projectsUrl
	o.field_protectedBranches = v.Field_protectedBranches
	o.field_pullRequest = v.Field_pullRequest
	o.field_pullRequests = v.Field_pullRequests
	o.field_pushedAt = v.Field_pushedAt
	o.field_ref = v.Field_ref
	o.field_refs = v.Field_refs
	o.field_releases = v.Field_releases
	o.field_repositoryTopics = v.Field_repositoryTopics
	o.field_resourcePath = v.Field_resourcePath
	o.field_stargazers = v.Field_stargazers
	o.field_updatedAt = v.Field_updatedAt
	o.field_url = v.Field_url
	o.field_viewerCanAdminister = v.Field_viewerCanAdminister
	o.field_viewerCanCreateProjects = v.Field_viewerCanCreateProjects
	o.field_viewerCanSubscribe = v.Field_viewerCanSubscribe
	o.field_viewerCanUpdateTopics = v.Field_viewerCanUpdateTopics
	o.field_viewerHasStarred = v.Field_viewerHasStarred
	o.field_viewerSubscription = v.Field_viewerSubscription
	o.field_watchers = v.Field_watchers
	return nil
}

// Projects manage issues, pull requests and notes within a project owner.
type Project struct {
	// The project's description body.
	field_body *string `json:"body"`

	// The projects description body rendered to HTML.
	field_bodyHTML HTML `json:"bodyHTML"`

	// Identifities the date and time when the project was closed.
	field_closedAt *DateTime `json:"closedAt"`

	// List of columns in the project
	field_columns ProjectColumnConnection `json:"columns"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// The actor who originally created the project.
	field_creator Actor `json:"creator"`

	// Identifies the primary key from the database.
	field_databaseId *int32 `json:"databaseId"`

	field_id ID `json:"id"`

	// The project's name.
	field_name string `json:"name"`

	// The project's number.
	field_number int32 `json:"number"`

	// The project's owner. Currently limited to repositories and organizations.
	field_owner ProjectOwner `json:"owner"`

	// The HTTP path for this project
	field_resourcePath URI `json:"resourcePath"`

	// Whether the project is open or closed.
	field_state ProjectState `json:"state"`

	// Identifies the date and time when the object was last updated.
	field_updatedAt DateTime `json:"updatedAt"`

	// The HTTP URL for this project
	field_url URI `json:"url"`

	// Check if the current viewer can update this object.
	field_viewerCanUpdate bool `json:"viewerCanUpdate"`
}

func (o Project) Body() *string {
	return o.field_body
}

func (o Project) BodyHTML() HTML {
	return o.field_bodyHTML
}

func (o Project) ClosedAt() *DateTime {
	return o.field_closedAt
}

func (o Project) Columns() ProjectColumnConnection {
	return o.field_columns
}

func (o Project) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o Project) Creator() Actor {
	return o.field_creator
}

func (o Project) DatabaseId() *int32 {
	return o.field_databaseId
}

func (o Project) Id() ID {
	return o.field_id
}

func (o Project) Name() string {
	return o.field_name
}

func (o Project) Number() int32 {
	return o.field_number
}

func (o Project) Owner() ProjectOwner {
	return o.field_owner
}

func (o Project) ResourcePath() URI {
	return o.field_resourcePath
}

func (o Project) State() ProjectState {
	return o.field_state
}

func (o Project) UpdatedAt() DateTime {
	return o.field_updatedAt
}

func (o Project) Url() URI {
	return o.field_url
}

func (o Project) ViewerCanUpdate() bool {
	return o.field_viewerCanUpdate
}

func (o *Project) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_body            *string                 `json:"body"`
		Field_bodyHTML        HTML                    `json:"bodyHTML"`
		Field_closedAt        *DateTime               `json:"closedAt"`
		Field_columns         ProjectColumnConnection `json:"columns"`
		Field_createdAt       DateTime                `json:"createdAt"`
		Field_creator         json.RawMessage         `json:"creator"`
		Field_databaseId      *int32                  `json:"databaseId"`
		Field_id              ID                      `json:"id"`
		Field_name            string                  `json:"name"`
		Field_number          int32                   `json:"number"`
		Field_owner           ProjectOwner            `json:"owner"`
		Field_resourcePath    URI                     `json:"resourcePath"`
		Field_state           ProjectState            `json:"state"`
		Field_updatedAt       DateTime                `json:"updatedAt"`
		Field_url             URI                     `json:"url"`
		Field_viewerCanUpdate bool                    `json:"viewerCanUpdate"`
	}{
		Field_body:            o.field_body,
		Field_bodyHTML:        o.field_bodyHTML,
		Field_closedAt:        o.field_closedAt,
		Field_columns:         o.field_columns,
		Field_createdAt:       o.field_createdAt,
		Field_databaseId:      o.field_databaseId,
		Field_id:              o.field_id,
		Field_name:            o.field_name,
		Field_number:          o.field_number,
		Field_owner:           o.field_owner,
		Field_resourcePath:    o.field_resourcePath,
		Field_state:           o.field_state,
		Field_updatedAt:       o.field_updatedAt,
		Field_url:             o.field_url,
		Field_viewerCanUpdate: o.field_viewerCanUpdate,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_body = v.Field_body
	o.field_bodyHTML = v.Field_bodyHTML
	o.field_closedAt = v.Field_closedAt
	o.field_columns = v.Field_columns
	o.field_createdAt = v.Field_createdAt
	if len(v.Field_creator) > 0 {
		o.field_creator, err = Actor_UnmarshalJSON(v.Field_creator)
		if err != nil {
			return err
		}
	}
	o.field_databaseId = v.Field_databaseId
	o.field_id = v.Field_id
	o.field_name = v.Field_name
	o.field_number = v.Field_number
	o.field_owner = v.Field_owner
	o.field_resourcePath = v.Field_resourcePath
	o.field_state = v.Field_state
	o.field_updatedAt = v.Field_updatedAt
	o.field_url = v.Field_url
	o.field_viewerCanUpdate = v.Field_viewerCanUpdate
	return nil
}

// A list of projects associated with the owner.
type ProjectConnection struct {
	// A list of edges.
	field_edges []*ProjectEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*Project `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o ProjectConnection) Edges() []*ProjectEdge {
	return o.field_edges
}

func (o ProjectConnection) Nodes() []*Project {
	return o.field_nodes
}

func (o ProjectConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o ProjectConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *ProjectConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*ProjectEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*Project, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type ProjectEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *Project `json:"node"`
}

func (o ProjectEdge) Cursor() string {
	return o.field_cursor
}

func (o ProjectEdge) Node() *Project {
	return o.field_node
}

func (o *ProjectEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string   `json:"cursor"`
		Field_node   *Project `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// Information about pagination in a connection.
type PageInfo struct {
	// When paginating forwards, the cursor to continue.
	field_endCursor *string `json:"endCursor"`

	// When paginating forwards, are there more items?
	field_hasNextPage bool `json:"hasNextPage"`

	// When paginating backwards, are there more items?
	field_hasPreviousPage bool `json:"hasPreviousPage"`

	// When paginating backwards, the cursor to continue.
	field_startCursor *string `json:"startCursor"`
}

func (o PageInfo) EndCursor() *string {
	return o.field_endCursor
}

func (o PageInfo) HasNextPage() bool {
	return o.field_hasNextPage
}

func (o PageInfo) HasPreviousPage() bool {
	return o.field_hasPreviousPage
}

func (o PageInfo) StartCursor() *string {
	return o.field_startCursor
}

func (o *PageInfo) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_endCursor       *string `json:"endCursor"`
		Field_hasNextPage     bool    `json:"hasNextPage"`
		Field_hasPreviousPage bool    `json:"hasPreviousPage"`
		Field_startCursor     *string `json:"startCursor"`
	}{
		Field_endCursor:       o.field_endCursor,
		Field_hasNextPage:     o.field_hasNextPage,
		Field_hasPreviousPage: o.field_hasPreviousPage,
		Field_startCursor:     o.field_startCursor,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_endCursor = v.Field_endCursor
	o.field_hasNextPage = v.Field_hasNextPage
	o.field_hasPreviousPage = v.Field_hasPreviousPage
	o.field_startCursor = v.Field_startCursor
	return nil
}

// The connection type for ProjectColumn.
type ProjectColumnConnection struct {
	// A list of edges.
	field_edges []*ProjectColumnEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*ProjectColumn `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o ProjectColumnConnection) Edges() []*ProjectColumnEdge {
	return o.field_edges
}

func (o ProjectColumnConnection) Nodes() []*ProjectColumn {
	return o.field_nodes
}

func (o ProjectColumnConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o ProjectColumnConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *ProjectColumnConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*ProjectColumnEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*ProjectColumn, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type ProjectColumnEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *ProjectColumn `json:"node"`
}

func (o ProjectColumnEdge) Cursor() string {
	return o.field_cursor
}

func (o ProjectColumnEdge) Node() *ProjectColumn {
	return o.field_node
}

func (o *ProjectColumnEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string         `json:"cursor"`
		Field_node   *ProjectColumn `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// A column inside a project.
type ProjectColumn struct {
	// List of cards in the column
	field_cards ProjectCardConnection `json:"cards"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// Identifies the primary key from the database.
	field_databaseId *int32 `json:"databaseId"`

	field_id ID `json:"id"`

	// The project column's name.
	field_name string `json:"name"`

	// The project that contains this column.
	field_project Project `json:"project"`

	// Identifies the date and time when the object was last updated.
	field_updatedAt DateTime `json:"updatedAt"`
}

func (o ProjectColumn) Cards() ProjectCardConnection {
	return o.field_cards
}

func (o ProjectColumn) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o ProjectColumn) DatabaseId() *int32 {
	return o.field_databaseId
}

func (o ProjectColumn) Id() ID {
	return o.field_id
}

func (o ProjectColumn) Name() string {
	return o.field_name
}

func (o ProjectColumn) Project() Project {
	return o.field_project
}

func (o ProjectColumn) UpdatedAt() DateTime {
	return o.field_updatedAt
}

func (o *ProjectColumn) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cards      ProjectCardConnection `json:"cards"`
		Field_createdAt  DateTime              `json:"createdAt"`
		Field_databaseId *int32                `json:"databaseId"`
		Field_id         ID                    `json:"id"`
		Field_name       string                `json:"name"`
		Field_project    Project               `json:"project"`
		Field_updatedAt  DateTime              `json:"updatedAt"`
	}{
		Field_cards:      o.field_cards,
		Field_createdAt:  o.field_createdAt,
		Field_databaseId: o.field_databaseId,
		Field_id:         o.field_id,
		Field_name:       o.field_name,
		Field_project:    o.field_project,
		Field_updatedAt:  o.field_updatedAt,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cards = v.Field_cards
	o.field_createdAt = v.Field_createdAt
	o.field_databaseId = v.Field_databaseId
	o.field_id = v.Field_id
	o.field_name = v.Field_name
	o.field_project = v.Field_project
	o.field_updatedAt = v.Field_updatedAt
	return nil
}

// The connection type for ProjectCard.
type ProjectCardConnection struct {
	// A list of edges.
	field_edges []*ProjectCardEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*ProjectCard `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o ProjectCardConnection) Edges() []*ProjectCardEdge {
	return o.field_edges
}

func (o ProjectCardConnection) Nodes() []*ProjectCard {
	return o.field_nodes
}

func (o ProjectCardConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o ProjectCardConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *ProjectCardConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*ProjectCardEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*ProjectCard, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type ProjectCardEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *ProjectCard `json:"node"`
}

func (o ProjectCardEdge) Cursor() string {
	return o.field_cursor
}

func (o ProjectCardEdge) Node() *ProjectCard {
	return o.field_node
}

func (o *ProjectCardEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string       `json:"cursor"`
		Field_node   *ProjectCard `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// A card in a project.
type ProjectCard struct {
	// The project column this card is associated under. A card may only belong to one
	// project column at a time. The column field will be null if the card is created
	// in a pending state and has yet to be associated with a column. Once cards are
	// associated with a column, they will not become pending in the future.
	//
	field_column *ProjectColumn `json:"column"`

	// The card content item
	field_content ProjectCardItem `json:"content"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// The actor who created this card
	field_creator Actor `json:"creator"`

	// Identifies the primary key from the database.
	field_databaseId *int32 `json:"databaseId"`

	field_id ID `json:"id"`

	// The card note
	field_note *string `json:"note"`

	// The project that contains this card.
	field_project Project `json:"project"`

	// The column that contains this card.
	field_projectColumn ProjectColumn `json:"projectColumn"`

	// The HTTP path for this card
	field_resourcePath URI `json:"resourcePath"`

	// The state of ProjectCard
	field_state *ProjectCardState `json:"state"`

	// Identifies the date and time when the object was last updated.
	field_updatedAt DateTime `json:"updatedAt"`

	// The HTTP URL for this card
	field_url URI `json:"url"`
}

func (o ProjectCard) Column() *ProjectColumn {
	return o.field_column
}

func (o ProjectCard) Content() ProjectCardItem {
	return o.field_content
}

func (o ProjectCard) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o ProjectCard) Creator() Actor {
	return o.field_creator
}

func (o ProjectCard) DatabaseId() *int32 {
	return o.field_databaseId
}

func (o ProjectCard) Id() ID {
	return o.field_id
}

func (o ProjectCard) Note() *string {
	return o.field_note
}

func (o ProjectCard) Project() Project {
	return o.field_project
}

func (o ProjectCard) ProjectColumn() ProjectColumn {
	return o.field_projectColumn
}

func (o ProjectCard) ResourcePath() URI {
	return o.field_resourcePath
}

func (o ProjectCard) State() *ProjectCardState {
	return o.field_state
}

func (o ProjectCard) UpdatedAt() DateTime {
	return o.field_updatedAt
}

func (o ProjectCard) Url() URI {
	return o.field_url
}

func (o *ProjectCard) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_column        *ProjectColumn    `json:"column"`
		Field_content       json.RawMessage   `json:"content"`
		Field_createdAt     DateTime          `json:"createdAt"`
		Field_creator       json.RawMessage   `json:"creator"`
		Field_databaseId    *int32            `json:"databaseId"`
		Field_id            ID                `json:"id"`
		Field_note          *string           `json:"note"`
		Field_project       Project           `json:"project"`
		Field_projectColumn ProjectColumn     `json:"projectColumn"`
		Field_resourcePath  URI               `json:"resourcePath"`
		Field_state         *ProjectCardState `json:"state"`
		Field_updatedAt     DateTime          `json:"updatedAt"`
		Field_url           URI               `json:"url"`
	}{
		Field_column:        o.field_column,
		Field_createdAt:     o.field_createdAt,
		Field_databaseId:    o.field_databaseId,
		Field_id:            o.field_id,
		Field_note:          o.field_note,
		Field_project:       o.field_project,
		Field_projectColumn: o.field_projectColumn,
		Field_resourcePath:  o.field_resourcePath,
		Field_state:         o.field_state,
		Field_updatedAt:     o.field_updatedAt,
		Field_url:           o.field_url,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_column = v.Field_column
	if len(v.Field_content) > 0 {
		o.field_content, err = ProjectCardItem_UnmarshalJSON(v.Field_content)
		if err != nil {
			return err
		}
	}
	o.field_createdAt = v.Field_createdAt
	if len(v.Field_creator) > 0 {
		o.field_creator, err = Actor_UnmarshalJSON(v.Field_creator)
		if err != nil {
			return err
		}
	}
	o.field_databaseId = v.Field_databaseId
	o.field_id = v.Field_id
	o.field_note = v.Field_note
	o.field_project = v.Field_project
	o.field_projectColumn = v.Field_projectColumn
	o.field_resourcePath = v.Field_resourcePath
	o.field_state = v.Field_state
	o.field_updatedAt = v.Field_updatedAt
	o.field_url = v.Field_url
	return nil
}

// An Issue is a place to discuss ideas, enhancements, tasks, and bugs for a project.
type Issue struct {
	// A list of Users assigned to this object.
	field_assignees UserConnection `json:"assignees"`

	// The actor who authored the comment.
	field_author Actor `json:"author"`

	// Author's association with the subject of the comment.
	field_authorAssociation CommentAuthorAssociation `json:"authorAssociation"`

	// Identifies the body of the issue.
	field_body string `json:"body"`

	// Identifies the body of the issue rendered to HTML.
	field_bodyHTML HTML `json:"bodyHTML"`

	// Identifies the body of the issue rendered to text.
	field_bodyText string `json:"bodyText"`

	// `true` if the object is closed (definition of closed may depend on type)
	field_closed bool `json:"closed"`

	// A list of comments associated with the Issue.
	field_comments IssueCommentConnection `json:"comments"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// Check if this comment was created via an email reply.
	field_createdViaEmail bool `json:"createdViaEmail"`

	// Identifies the primary key from the database.
	field_databaseId *int32 `json:"databaseId"`

	// The actor who edited the comment.
	field_editor Actor `json:"editor"`

	field_id ID `json:"id"`

	// A list of labels associated with the object.
	field_labels *LabelConnection `json:"labels"`

	// The moment the editor made the last edit
	field_lastEditedAt *DateTime `json:"lastEditedAt"`

	// `true` if the object is locked
	field_locked bool `json:"locked"`

	// Identifies the milestone associated with the issue.
	field_milestone *Milestone `json:"milestone"`

	// Identifies the issue number.
	field_number int32 `json:"number"`

	// A list of Users that are participating in the Issue conversation.
	field_participants UserConnection `json:"participants"`

	// Identifies when the comment was published at.
	field_publishedAt *DateTime `json:"publishedAt"`

	// A list of reactions grouped by content left on the subject.
	field_reactionGroups []ReactionGroup `json:"reactionGroups"`

	// A list of Reactions left on the Issue.
	field_reactions ReactionConnection `json:"reactions"`

	// Identifies the repository associated with the issue.
	field_repository Repository `json:"repository"`

	// The HTTP path for this issue
	field_resourcePath URI `json:"resourcePath"`

	// Identifies the state of the issue.
	field_state IssueState `json:"state"`

	// A list of events associated with an Issue.
	field_timeline IssueTimelineConnection `json:"timeline"`

	// Identifies the issue title.
	field_title string `json:"title"`

	// Identifies the date and time when the object was last updated.
	field_updatedAt DateTime `json:"updatedAt"`

	// The HTTP URL for this issue
	field_url URI `json:"url"`

	// Can user react to this subject
	field_viewerCanReact bool `json:"viewerCanReact"`

	// Check if the viewer is able to change their subscription status for the repository.
	field_viewerCanSubscribe bool `json:"viewerCanSubscribe"`

	// Check if the current viewer can update this object.
	field_viewerCanUpdate bool `json:"viewerCanUpdate"`

	// Reasons why the current viewer can not update this comment.
	field_viewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`

	// Did the viewer author this comment.
	field_viewerDidAuthor bool `json:"viewerDidAuthor"`

	// Identifies if the viewer is watching, not watching, or ignoring the repository.
	field_viewerSubscription SubscriptionState `json:"viewerSubscription"`
}

func (o Issue) Assignees() UserConnection {
	return o.field_assignees
}

func (o Issue) Author() Actor {
	return o.field_author
}

func (o Issue) AuthorAssociation() CommentAuthorAssociation {
	return o.field_authorAssociation
}

func (o Issue) Body() string {
	return o.field_body
}

func (o Issue) BodyHTML() HTML {
	return o.field_bodyHTML
}

func (o Issue) BodyText() string {
	return o.field_bodyText
}

func (o Issue) Closed() bool {
	return o.field_closed
}

func (o Issue) Comments() IssueCommentConnection {
	return o.field_comments
}

func (o Issue) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o Issue) CreatedViaEmail() bool {
	return o.field_createdViaEmail
}

func (o Issue) DatabaseId() *int32 {
	return o.field_databaseId
}

func (o Issue) Editor() Actor {
	return o.field_editor
}

func (o Issue) Id() ID {
	return o.field_id
}

func (o Issue) Labels() *LabelConnection {
	return o.field_labels
}

func (o Issue) LastEditedAt() *DateTime {
	return o.field_lastEditedAt
}

func (o Issue) Locked() bool {
	return o.field_locked
}

func (o Issue) Milestone() *Milestone {
	return o.field_milestone
}

func (o Issue) Number() int32 {
	return o.field_number
}

func (o Issue) Participants() UserConnection {
	return o.field_participants
}

func (o Issue) PublishedAt() *DateTime {
	return o.field_publishedAt
}

func (o Issue) ReactionGroups() []ReactionGroup {
	return o.field_reactionGroups
}

func (o Issue) Reactions() ReactionConnection {
	return o.field_reactions
}

func (o Issue) Repository() Repository {
	return o.field_repository
}

func (o Issue) ResourcePath() URI {
	return o.field_resourcePath
}

func (o Issue) State() IssueState {
	return o.field_state
}

func (o Issue) Timeline() IssueTimelineConnection {
	return o.field_timeline
}

func (o Issue) Title() string {
	return o.field_title
}

func (o Issue) UpdatedAt() DateTime {
	return o.field_updatedAt
}

func (o Issue) Url() URI {
	return o.field_url
}

func (o Issue) ViewerCanReact() bool {
	return o.field_viewerCanReact
}

func (o Issue) ViewerCanSubscribe() bool {
	return o.field_viewerCanSubscribe
}

func (o Issue) ViewerCanUpdate() bool {
	return o.field_viewerCanUpdate
}

func (o Issue) ViewerCannotUpdateReasons() []CommentCannotUpdateReason {
	return o.field_viewerCannotUpdateReasons
}

func (o Issue) ViewerDidAuthor() bool {
	return o.field_viewerDidAuthor
}

func (o Issue) ViewerSubscription() SubscriptionState {
	return o.field_viewerSubscription
}

func (o *Issue) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_assignees                 UserConnection              `json:"assignees"`
		Field_author                    json.RawMessage             `json:"author"`
		Field_authorAssociation         CommentAuthorAssociation    `json:"authorAssociation"`
		Field_body                      string                      `json:"body"`
		Field_bodyHTML                  HTML                        `json:"bodyHTML"`
		Field_bodyText                  string                      `json:"bodyText"`
		Field_closed                    bool                        `json:"closed"`
		Field_comments                  IssueCommentConnection      `json:"comments"`
		Field_createdAt                 DateTime                    `json:"createdAt"`
		Field_createdViaEmail           bool                        `json:"createdViaEmail"`
		Field_databaseId                *int32                      `json:"databaseId"`
		Field_editor                    json.RawMessage             `json:"editor"`
		Field_id                        ID                          `json:"id"`
		Field_labels                    *LabelConnection            `json:"labels"`
		Field_lastEditedAt              *DateTime                   `json:"lastEditedAt"`
		Field_locked                    bool                        `json:"locked"`
		Field_milestone                 *Milestone                  `json:"milestone"`
		Field_number                    int32                       `json:"number"`
		Field_participants              UserConnection              `json:"participants"`
		Field_publishedAt               *DateTime                   `json:"publishedAt"`
		Field_reactionGroups            json.RawMessage             `json:"reactionGroups"`
		Field_reactions                 ReactionConnection          `json:"reactions"`
		Field_repository                Repository                  `json:"repository"`
		Field_resourcePath              URI                         `json:"resourcePath"`
		Field_state                     IssueState                  `json:"state"`
		Field_timeline                  IssueTimelineConnection     `json:"timeline"`
		Field_title                     string                      `json:"title"`
		Field_updatedAt                 DateTime                    `json:"updatedAt"`
		Field_url                       URI                         `json:"url"`
		Field_viewerCanReact            bool                        `json:"viewerCanReact"`
		Field_viewerCanSubscribe        bool                        `json:"viewerCanSubscribe"`
		Field_viewerCanUpdate           bool                        `json:"viewerCanUpdate"`
		Field_viewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`
		Field_viewerDidAuthor           bool                        `json:"viewerDidAuthor"`
		Field_viewerSubscription        SubscriptionState           `json:"viewerSubscription"`
	}{
		Field_assignees:                 o.field_assignees,
		Field_authorAssociation:         o.field_authorAssociation,
		Field_body:                      o.field_body,
		Field_bodyHTML:                  o.field_bodyHTML,
		Field_bodyText:                  o.field_bodyText,
		Field_closed:                    o.field_closed,
		Field_comments:                  o.field_comments,
		Field_createdAt:                 o.field_createdAt,
		Field_createdViaEmail:           o.field_createdViaEmail,
		Field_databaseId:                o.field_databaseId,
		Field_id:                        o.field_id,
		Field_labels:                    o.field_labels,
		Field_lastEditedAt:              o.field_lastEditedAt,
		Field_locked:                    o.field_locked,
		Field_milestone:                 o.field_milestone,
		Field_number:                    o.field_number,
		Field_participants:              o.field_participants,
		Field_publishedAt:               o.field_publishedAt,
		Field_reactions:                 o.field_reactions,
		Field_repository:                o.field_repository,
		Field_resourcePath:              o.field_resourcePath,
		Field_state:                     o.field_state,
		Field_timeline:                  o.field_timeline,
		Field_title:                     o.field_title,
		Field_updatedAt:                 o.field_updatedAt,
		Field_url:                       o.field_url,
		Field_viewerCanReact:            o.field_viewerCanReact,
		Field_viewerCanSubscribe:        o.field_viewerCanSubscribe,
		Field_viewerCanUpdate:           o.field_viewerCanUpdate,
		Field_viewerCannotUpdateReasons: o.field_viewerCannotUpdateReasons,
		Field_viewerDidAuthor:           o.field_viewerDidAuthor,
		Field_viewerSubscription:        o.field_viewerSubscription,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_assignees = v.Field_assignees
	if len(v.Field_author) > 0 {
		o.field_author, err = Actor_UnmarshalJSON(v.Field_author)
		if err != nil {
			return err
		}
	}
	o.field_authorAssociation = v.Field_authorAssociation
	o.field_body = v.Field_body
	o.field_bodyHTML = v.Field_bodyHTML
	o.field_bodyText = v.Field_bodyText
	o.field_closed = v.Field_closed
	o.field_comments = v.Field_comments
	o.field_createdAt = v.Field_createdAt
	o.field_createdViaEmail = v.Field_createdViaEmail
	o.field_databaseId = v.Field_databaseId
	if len(v.Field_editor) > 0 {
		o.field_editor, err = Actor_UnmarshalJSON(v.Field_editor)
		if err != nil {
			return err
		}
	}
	o.field_id = v.Field_id
	o.field_labels = v.Field_labels
	o.field_lastEditedAt = v.Field_lastEditedAt
	o.field_locked = v.Field_locked
	o.field_milestone = v.Field_milestone
	o.field_number = v.Field_number
	o.field_participants = v.Field_participants
	o.field_publishedAt = v.Field_publishedAt
	if len(v.Field_reactionGroups) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_reactionGroups, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_reactionGroups = nil
		} else {
			o.field_reactionGroups = make([]ReactionGroup, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_reactionGroups[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_reactions = v.Field_reactions
	o.field_repository = v.Field_repository
	o.field_resourcePath = v.Field_resourcePath
	o.field_state = v.Field_state
	o.field_timeline = v.Field_timeline
	o.field_title = v.Field_title
	o.field_updatedAt = v.Field_updatedAt
	o.field_url = v.Field_url
	o.field_viewerCanReact = v.Field_viewerCanReact
	o.field_viewerCanSubscribe = v.Field_viewerCanSubscribe
	o.field_viewerCanUpdate = v.Field_viewerCanUpdate
	o.field_viewerCannotUpdateReasons = v.Field_viewerCannotUpdateReasons
	o.field_viewerDidAuthor = v.Field_viewerDidAuthor
	o.field_viewerSubscription = v.Field_viewerSubscription
	return nil
}

// The connection type for User.
type UserConnection struct {
	// A list of edges.
	field_edges []*UserEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*User `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o UserConnection) Edges() []*UserEdge {
	return o.field_edges
}

func (o UserConnection) Nodes() []*User {
	return o.field_nodes
}

func (o UserConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o UserConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *UserConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*UserEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*User, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type UserEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *User `json:"node"`
}

func (o UserEdge) Cursor() string {
	return o.field_cursor
}

func (o UserEdge) Node() *User {
	return o.field_node
}

func (o *UserEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string `json:"cursor"`
		Field_node   *User  `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// A user is an individual's account on GitHub that owns repositories and can make new content.
type User struct {
	// A URL pointing to the user's public avatar.
	field_avatarUrl URI `json:"avatarUrl"`

	// The user's public profile bio.
	field_bio *string `json:"bio"`

	// The user's public profile bio as HTML.
	field_bioHTML HTML `json:"bioHTML"`

	// The user's public profile company.
	field_company *string `json:"company"`

	// The user's public profile company as HTML.
	field_companyHTML HTML `json:"companyHTML"`

	// A list of repositories that the user recently contributed to.
	field_contributedRepositories RepositoryConnection `json:"contributedRepositories"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// Identifies the primary key from the database.
	field_databaseId *int32 `json:"databaseId"`

	// The user's publicly visible profile email.
	field_email string `json:"email"`

	// A list of users the given user is followed by.
	field_followers FollowerConnection `json:"followers"`

	// A list of users the given user is following.
	field_following FollowingConnection `json:"following"`

	// Find gist by repo name.
	field_gist *Gist `json:"gist"`

	// A list of the Gists the user has created.
	field_gists GistConnection `json:"gists"`

	field_id ID `json:"id"`

	// Whether or not this user is a participant in the GitHub Security Bug Bounty.
	field_isBountyHunter bool `json:"isBountyHunter"`

	// Whether or not this user is a participant in the GitHub Campus Experts Program.
	field_isCampusExpert bool `json:"isCampusExpert"`

	// Whether or not this user is a GitHub Developer Program member.
	field_isDeveloperProgramMember bool `json:"isDeveloperProgramMember"`

	// Whether or not this user is a GitHub employee.
	field_isEmployee bool `json:"isEmployee"`

	// Whether or not the user has marked themselves as for hire.
	field_isHireable bool `json:"isHireable"`

	// Is the account billed through invoices?
	field_isInvoiced bool `json:"isInvoiced"`

	// Whether or not this user is a site administrator.
	field_isSiteAdmin bool `json:"isSiteAdmin"`

	// Whether or not this user is the viewing user.
	field_isViewer bool `json:"isViewer"`

	// The user's public profile location.
	field_location *string `json:"location"`

	// The username used to login.
	field_login string `json:"login"`

	// The user's public profile name.
	field_name *string `json:"name"`

	// Find an organization by its login that the user belongs to.
	field_organization *Organization `json:"organization"`

	// A list of organizations the user belongs to.
	field_organizations OrganizationConnection `json:"organizations"`

	// A list of repositories this user has pinned to their profile
	field_pinnedRepositories RepositoryConnection `json:"pinnedRepositories"`

	// A list of pull requests assocated with this user.
	field_pullRequests PullRequestConnection `json:"pullRequests"`

	// A list of repositories that the user owns.
	field_repositories RepositoryConnection `json:"repositories"`

	// Find Repository.
	field_repository *Repository `json:"repository"`

	// The HTTP path for this user
	field_resourcePath URI `json:"resourcePath"`

	// Repositories the user has starred.
	field_starredRepositories StarredRepositoryConnection `json:"starredRepositories"`

	// Identifies the date and time when the object was last updated.
	field_updatedAt DateTime `json:"updatedAt"`

	// The HTTP URL for this user
	field_url URI `json:"url"`

	// Whether or not the viewer is able to follow the user.
	field_viewerCanFollow bool `json:"viewerCanFollow"`

	// Whether or not this user is followed by the viewer.
	field_viewerIsFollowing bool `json:"viewerIsFollowing"`

	// A list of repositories the given user is watching.
	field_watching RepositoryConnection `json:"watching"`

	// A URL pointing to the user's public website/blog.
	field_websiteUrl *URI `json:"websiteUrl"`
}

func (o User) AvatarUrl() URI {
	return o.field_avatarUrl
}

func (o User) Bio() *string {
	return o.field_bio
}

func (o User) BioHTML() HTML {
	return o.field_bioHTML
}

func (o User) Company() *string {
	return o.field_company
}

func (o User) CompanyHTML() HTML {
	return o.field_companyHTML
}

func (o User) ContributedRepositories() RepositoryConnection {
	return o.field_contributedRepositories
}

func (o User) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o User) DatabaseId() *int32 {
	return o.field_databaseId
}

func (o User) Email() string {
	return o.field_email
}

func (o User) Followers() FollowerConnection {
	return o.field_followers
}

func (o User) Following() FollowingConnection {
	return o.field_following
}

func (o User) Gist() *Gist {
	return o.field_gist
}

func (o User) Gists() GistConnection {
	return o.field_gists
}

func (o User) Id() ID {
	return o.field_id
}

func (o User) IsBountyHunter() bool {
	return o.field_isBountyHunter
}

func (o User) IsCampusExpert() bool {
	return o.field_isCampusExpert
}

func (o User) IsDeveloperProgramMember() bool {
	return o.field_isDeveloperProgramMember
}

func (o User) IsEmployee() bool {
	return o.field_isEmployee
}

func (o User) IsHireable() bool {
	return o.field_isHireable
}

func (o User) IsInvoiced() bool {
	return o.field_isInvoiced
}

func (o User) IsSiteAdmin() bool {
	return o.field_isSiteAdmin
}

func (o User) IsViewer() bool {
	return o.field_isViewer
}

func (o User) Location() *string {
	return o.field_location
}

func (o User) Login() string {
	return o.field_login
}

func (o User) Name() *string {
	return o.field_name
}

func (o User) Organization() *Organization {
	return o.field_organization
}

func (o User) Organizations() OrganizationConnection {
	return o.field_organizations
}

func (o User) PinnedRepositories() RepositoryConnection {
	return o.field_pinnedRepositories
}

func (o User) PullRequests() PullRequestConnection {
	return o.field_pullRequests
}

func (o User) Repositories() RepositoryConnection {
	return o.field_repositories
}

func (o User) Repository() *Repository {
	return o.field_repository
}

func (o User) ResourcePath() URI {
	return o.field_resourcePath
}

func (o User) StarredRepositories() StarredRepositoryConnection {
	return o.field_starredRepositories
}

func (o User) UpdatedAt() DateTime {
	return o.field_updatedAt
}

func (o User) Url() URI {
	return o.field_url
}

func (o User) ViewerCanFollow() bool {
	return o.field_viewerCanFollow
}

func (o User) ViewerIsFollowing() bool {
	return o.field_viewerIsFollowing
}

func (o User) Watching() RepositoryConnection {
	return o.field_watching
}

func (o User) WebsiteUrl() *URI {
	return o.field_websiteUrl
}

func (o *User) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_avatarUrl                URI                         `json:"avatarUrl"`
		Field_bio                      *string                     `json:"bio"`
		Field_bioHTML                  HTML                        `json:"bioHTML"`
		Field_company                  *string                     `json:"company"`
		Field_companyHTML              HTML                        `json:"companyHTML"`
		Field_contributedRepositories  RepositoryConnection        `json:"contributedRepositories"`
		Field_createdAt                DateTime                    `json:"createdAt"`
		Field_databaseId               *int32                      `json:"databaseId"`
		Field_email                    string                      `json:"email"`
		Field_followers                FollowerConnection          `json:"followers"`
		Field_following                FollowingConnection         `json:"following"`
		Field_gist                     *Gist                       `json:"gist"`
		Field_gists                    GistConnection              `json:"gists"`
		Field_id                       ID                          `json:"id"`
		Field_isBountyHunter           bool                        `json:"isBountyHunter"`
		Field_isCampusExpert           bool                        `json:"isCampusExpert"`
		Field_isDeveloperProgramMember bool                        `json:"isDeveloperProgramMember"`
		Field_isEmployee               bool                        `json:"isEmployee"`
		Field_isHireable               bool                        `json:"isHireable"`
		Field_isInvoiced               bool                        `json:"isInvoiced"`
		Field_isSiteAdmin              bool                        `json:"isSiteAdmin"`
		Field_isViewer                 bool                        `json:"isViewer"`
		Field_location                 *string                     `json:"location"`
		Field_login                    string                      `json:"login"`
		Field_name                     *string                     `json:"name"`
		Field_organization             *Organization               `json:"organization"`
		Field_organizations            OrganizationConnection      `json:"organizations"`
		Field_pinnedRepositories       RepositoryConnection        `json:"pinnedRepositories"`
		Field_pullRequests             PullRequestConnection       `json:"pullRequests"`
		Field_repositories             RepositoryConnection        `json:"repositories"`
		Field_repository               *Repository                 `json:"repository"`
		Field_resourcePath             URI                         `json:"resourcePath"`
		Field_starredRepositories      StarredRepositoryConnection `json:"starredRepositories"`
		Field_updatedAt                DateTime                    `json:"updatedAt"`
		Field_url                      URI                         `json:"url"`
		Field_viewerCanFollow          bool                        `json:"viewerCanFollow"`
		Field_viewerIsFollowing        bool                        `json:"viewerIsFollowing"`
		Field_watching                 RepositoryConnection        `json:"watching"`
		Field_websiteUrl               *URI                        `json:"websiteUrl"`
	}{
		Field_avatarUrl:                o.field_avatarUrl,
		Field_bio:                      o.field_bio,
		Field_bioHTML:                  o.field_bioHTML,
		Field_company:                  o.field_company,
		Field_companyHTML:              o.field_companyHTML,
		Field_contributedRepositories:  o.field_contributedRepositories,
		Field_createdAt:                o.field_createdAt,
		Field_databaseId:               o.field_databaseId,
		Field_email:                    o.field_email,
		Field_followers:                o.field_followers,
		Field_following:                o.field_following,
		Field_gist:                     o.field_gist,
		Field_gists:                    o.field_gists,
		Field_id:                       o.field_id,
		Field_isBountyHunter:           o.field_isBountyHunter,
		Field_isCampusExpert:           o.field_isCampusExpert,
		Field_isDeveloperProgramMember: o.field_isDeveloperProgramMember,
		Field_isEmployee:               o.field_isEmployee,
		Field_isHireable:               o.field_isHireable,
		Field_isInvoiced:               o.field_isInvoiced,
		Field_isSiteAdmin:              o.field_isSiteAdmin,
		Field_isViewer:                 o.field_isViewer,
		Field_location:                 o.field_location,
		Field_login:                    o.field_login,
		Field_name:                     o.field_name,
		Field_organization:             o.field_organization,
		Field_organizations:            o.field_organizations,
		Field_pinnedRepositories:       o.field_pinnedRepositories,
		Field_pullRequests:             o.field_pullRequests,
		Field_repositories:             o.field_repositories,
		Field_repository:               o.field_repository,
		Field_resourcePath:             o.field_resourcePath,
		Field_starredRepositories:      o.field_starredRepositories,
		Field_updatedAt:                o.field_updatedAt,
		Field_url:                      o.field_url,
		Field_viewerCanFollow:          o.field_viewerCanFollow,
		Field_viewerIsFollowing:        o.field_viewerIsFollowing,
		Field_watching:                 o.field_watching,
		Field_websiteUrl:               o.field_websiteUrl,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_avatarUrl = v.Field_avatarUrl
	o.field_bio = v.Field_bio
	o.field_bioHTML = v.Field_bioHTML
	o.field_company = v.Field_company
	o.field_companyHTML = v.Field_companyHTML
	o.field_contributedRepositories = v.Field_contributedRepositories
	o.field_createdAt = v.Field_createdAt
	o.field_databaseId = v.Field_databaseId
	o.field_email = v.Field_email
	o.field_followers = v.Field_followers
	o.field_following = v.Field_following
	o.field_gist = v.Field_gist
	o.field_gists = v.Field_gists
	o.field_id = v.Field_id
	o.field_isBountyHunter = v.Field_isBountyHunter
	o.field_isCampusExpert = v.Field_isCampusExpert
	o.field_isDeveloperProgramMember = v.Field_isDeveloperProgramMember
	o.field_isEmployee = v.Field_isEmployee
	o.field_isHireable = v.Field_isHireable
	o.field_isInvoiced = v.Field_isInvoiced
	o.field_isSiteAdmin = v.Field_isSiteAdmin
	o.field_isViewer = v.Field_isViewer
	o.field_location = v.Field_location
	o.field_login = v.Field_login
	o.field_name = v.Field_name
	o.field_organization = v.Field_organization
	o.field_organizations = v.Field_organizations
	o.field_pinnedRepositories = v.Field_pinnedRepositories
	o.field_pullRequests = v.Field_pullRequests
	o.field_repositories = v.Field_repositories
	o.field_repository = v.Field_repository
	o.field_resourcePath = v.Field_resourcePath
	o.field_starredRepositories = v.Field_starredRepositories
	o.field_updatedAt = v.Field_updatedAt
	o.field_url = v.Field_url
	o.field_viewerCanFollow = v.Field_viewerCanFollow
	o.field_viewerIsFollowing = v.Field_viewerIsFollowing
	o.field_watching = v.Field_watching
	o.field_websiteUrl = v.Field_websiteUrl
	return nil
}

// A release contains the content for a release.
type Release struct {
	// Identifies the description of the release.
	field_description *string `json:"description"`

	field_id ID `json:"id"`

	// Identifies the title of the release.
	field_name *string `json:"name"`

	// Identifies the date and time when the release was created.
	field_publishedAt *DateTime `json:"publishedAt"`

	// List of releases assets which are dependent on this release.
	field_releaseAsset ReleaseAssetConnection `json:"releaseAsset"`

	// List of releases assets which are dependent on this release.
	field_releaseAssets ReleaseAssetConnection `json:"releaseAssets"`

	// The HTTP path for this issue
	field_resourcePath URI `json:"resourcePath"`

	// The Git tag the release points to
	field_tag *Ref `json:"tag"`

	// The HTTP URL for this issue
	field_url URI `json:"url"`
}

func (o Release) Description() *string {
	return o.field_description
}

func (o Release) Id() ID {
	return o.field_id
}

func (o Release) Name() *string {
	return o.field_name
}

func (o Release) PublishedAt() *DateTime {
	return o.field_publishedAt
}

func (o Release) ReleaseAsset() ReleaseAssetConnection {
	return o.field_releaseAsset
}

func (o Release) ReleaseAssets() ReleaseAssetConnection {
	return o.field_releaseAssets
}

func (o Release) ResourcePath() URI {
	return o.field_resourcePath
}

func (o Release) Tag() *Ref {
	return o.field_tag
}

func (o Release) Url() URI {
	return o.field_url
}

func (o *Release) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_description   *string                `json:"description"`
		Field_id            ID                     `json:"id"`
		Field_name          *string                `json:"name"`
		Field_publishedAt   *DateTime              `json:"publishedAt"`
		Field_releaseAsset  ReleaseAssetConnection `json:"releaseAsset"`
		Field_releaseAssets ReleaseAssetConnection `json:"releaseAssets"`
		Field_resourcePath  URI                    `json:"resourcePath"`
		Field_tag           *Ref                   `json:"tag"`
		Field_url           URI                    `json:"url"`
	}{
		Field_description:   o.field_description,
		Field_id:            o.field_id,
		Field_name:          o.field_name,
		Field_publishedAt:   o.field_publishedAt,
		Field_releaseAsset:  o.field_releaseAsset,
		Field_releaseAssets: o.field_releaseAssets,
		Field_resourcePath:  o.field_resourcePath,
		Field_tag:           o.field_tag,
		Field_url:           o.field_url,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_description = v.Field_description
	o.field_id = v.Field_id
	o.field_name = v.Field_name
	o.field_publishedAt = v.Field_publishedAt
	o.field_releaseAsset = v.Field_releaseAsset
	o.field_releaseAssets = v.Field_releaseAssets
	o.field_resourcePath = v.Field_resourcePath
	o.field_tag = v.Field_tag
	o.field_url = v.Field_url
	return nil
}

// Represents a Git reference.
type Ref struct {
	// A list of pull requests with this ref as the head ref.
	field_associatedPullRequests PullRequestConnection `json:"associatedPullRequests"`

	field_id ID `json:"id"`

	// The ref name.
	field_name string `json:"name"`

	// The ref's prefix, such as `refs/heads/` or `refs/tags/`.
	field_prefix string `json:"prefix"`

	// The repository the ref belongs to.
	field_repository Repository `json:"repository"`

	// The object the ref points to.
	field_target GitObject `json:"target"`
}

func (o Ref) AssociatedPullRequests() PullRequestConnection {
	return o.field_associatedPullRequests
}

func (o Ref) Id() ID {
	return o.field_id
}

func (o Ref) Name() string {
	return o.field_name
}

func (o Ref) Prefix() string {
	return o.field_prefix
}

func (o Ref) Repository() Repository {
	return o.field_repository
}

func (o Ref) Target() GitObject {
	return o.field_target
}

func (o *Ref) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_associatedPullRequests PullRequestConnection `json:"associatedPullRequests"`
		Field_id                     ID                    `json:"id"`
		Field_name                   string                `json:"name"`
		Field_prefix                 string                `json:"prefix"`
		Field_repository             Repository            `json:"repository"`
		Field_target                 GitObject             `json:"target"`
	}{
		Field_associatedPullRequests: o.field_associatedPullRequests,
		Field_id:                     o.field_id,
		Field_name:                   o.field_name,
		Field_prefix:                 o.field_prefix,
		Field_repository:             o.field_repository,
		Field_target:                 o.field_target,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_associatedPullRequests = v.Field_associatedPullRequests
	o.field_id = v.Field_id
	o.field_name = v.Field_name
	o.field_prefix = v.Field_prefix
	o.field_repository = v.Field_repository
	o.field_target = v.Field_target
	return nil
}

// Represents a Git commit.
type Commit struct {
	// An abbreviated version of the Git object ID
	field_abbreviatedOid string `json:"abbreviatedOid"`

	// Authorship details of the commit.
	field_author *GitActor `json:"author"`

	// Check if the committer and the author match.
	field_authoredByCommitter bool `json:"authoredByCommitter"`

	// Fetches `git blame` information.
	field_blame Blame `json:"blame"`

	// Comments made on the commit.
	field_comments CommitCommentConnection `json:"comments"`

	// The HTTP path for this Git object
	field_commitResourcePath URI `json:"commitResourcePath"`

	// The HTTP URL for this Git object
	field_commitUrl URI `json:"commitUrl"`

	// The datetime when this commit was committed.
	field_committedDate DateTime `json:"committedDate"`

	// Check if commited via GitHub web UI.
	field_committedViaWeb bool `json:"committedViaWeb"`

	// Committership details of the commit.
	field_committer *GitActor `json:"committer"`

	// The linear commit history starting from (and including) this commit, in the same order as `git log`.
	field_history CommitHistoryConnection `json:"history"`

	field_id ID `json:"id"`

	// The Git commit message
	field_message string `json:"message"`

	// The Git commit message body
	field_messageBody string `json:"messageBody"`

	// The commit message body rendered to HTML.
	field_messageBodyHTML HTML `json:"messageBodyHTML"`

	// The Git commit message headline
	field_messageHeadline string `json:"messageHeadline"`

	// The commit message headline rendered to HTML.
	field_messageHeadlineHTML HTML `json:"messageHeadlineHTML"`

	// The Git object ID
	field_oid GitObjectID `json:"oid"`

	// The Repository this commit belongs to
	field_repository Repository `json:"repository"`

	// The HTTP path for this commit
	field_resourcePath URI `json:"resourcePath"`

	// Commit signing information, if present.
	field_signature GitSignature `json:"signature"`

	// Status information for this commit
	field_status *Status `json:"status"`

	// Commit's root Tree
	field_tree Tree `json:"tree"`

	// The HTTP path for the tree of this commit
	field_treeResourcePath URI `json:"treeResourcePath"`

	// The HTTP URL for the tree of this commit
	field_treeUrl URI `json:"treeUrl"`

	// The HTTP URL for this commit
	field_url URI `json:"url"`

	// Check if the viewer is able to change their subscription status for the repository.
	field_viewerCanSubscribe bool `json:"viewerCanSubscribe"`

	// Identifies if the viewer is watching, not watching, or ignoring the repository.
	field_viewerSubscription SubscriptionState `json:"viewerSubscription"`
}

func (o Commit) AbbreviatedOid() string {
	return o.field_abbreviatedOid
}

func (o Commit) Author() *GitActor {
	return o.field_author
}

func (o Commit) AuthoredByCommitter() bool {
	return o.field_authoredByCommitter
}

func (o Commit) Blame() Blame {
	return o.field_blame
}

func (o Commit) Comments() CommitCommentConnection {
	return o.field_comments
}

func (o Commit) CommitResourcePath() URI {
	return o.field_commitResourcePath
}

func (o Commit) CommitUrl() URI {
	return o.field_commitUrl
}

func (o Commit) CommittedDate() DateTime {
	return o.field_committedDate
}

func (o Commit) CommittedViaWeb() bool {
	return o.field_committedViaWeb
}

func (o Commit) Committer() *GitActor {
	return o.field_committer
}

func (o Commit) History() CommitHistoryConnection {
	return o.field_history
}

func (o Commit) Id() ID {
	return o.field_id
}

func (o Commit) Message() string {
	return o.field_message
}

func (o Commit) MessageBody() string {
	return o.field_messageBody
}

func (o Commit) MessageBodyHTML() HTML {
	return o.field_messageBodyHTML
}

func (o Commit) MessageHeadline() string {
	return o.field_messageHeadline
}

func (o Commit) MessageHeadlineHTML() HTML {
	return o.field_messageHeadlineHTML
}

func (o Commit) Oid() GitObjectID {
	return o.field_oid
}

func (o Commit) Repository() Repository {
	return o.field_repository
}

func (o Commit) ResourcePath() URI {
	return o.field_resourcePath
}

func (o Commit) Signature() GitSignature {
	return o.field_signature
}

func (o Commit) Status() *Status {
	return o.field_status
}

func (o Commit) Tree() Tree {
	return o.field_tree
}

func (o Commit) TreeResourcePath() URI {
	return o.field_treeResourcePath
}

func (o Commit) TreeUrl() URI {
	return o.field_treeUrl
}

func (o Commit) Url() URI {
	return o.field_url
}

func (o Commit) ViewerCanSubscribe() bool {
	return o.field_viewerCanSubscribe
}

func (o Commit) ViewerSubscription() SubscriptionState {
	return o.field_viewerSubscription
}

func (o *Commit) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_abbreviatedOid      string                  `json:"abbreviatedOid"`
		Field_author              *GitActor               `json:"author"`
		Field_authoredByCommitter bool                    `json:"authoredByCommitter"`
		Field_blame               Blame                   `json:"blame"`
		Field_comments            CommitCommentConnection `json:"comments"`
		Field_commitResourcePath  URI                     `json:"commitResourcePath"`
		Field_commitUrl           URI                     `json:"commitUrl"`
		Field_committedDate       DateTime                `json:"committedDate"`
		Field_committedViaWeb     bool                    `json:"committedViaWeb"`
		Field_committer           *GitActor               `json:"committer"`
		Field_history             CommitHistoryConnection `json:"history"`
		Field_id                  ID                      `json:"id"`
		Field_message             string                  `json:"message"`
		Field_messageBody         string                  `json:"messageBody"`
		Field_messageBodyHTML     HTML                    `json:"messageBodyHTML"`
		Field_messageHeadline     string                  `json:"messageHeadline"`
		Field_messageHeadlineHTML HTML                    `json:"messageHeadlineHTML"`
		Field_oid                 GitObjectID             `json:"oid"`
		Field_repository          Repository              `json:"repository"`
		Field_resourcePath        URI                     `json:"resourcePath"`
		Field_signature           json.RawMessage         `json:"signature"`
		Field_status              *Status                 `json:"status"`
		Field_tree                Tree                    `json:"tree"`
		Field_treeResourcePath    URI                     `json:"treeResourcePath"`
		Field_treeUrl             URI                     `json:"treeUrl"`
		Field_url                 URI                     `json:"url"`
		Field_viewerCanSubscribe  bool                    `json:"viewerCanSubscribe"`
		Field_viewerSubscription  SubscriptionState       `json:"viewerSubscription"`
	}{
		Field_abbreviatedOid:      o.field_abbreviatedOid,
		Field_author:              o.field_author,
		Field_authoredByCommitter: o.field_authoredByCommitter,
		Field_blame:               o.field_blame,
		Field_comments:            o.field_comments,
		Field_commitResourcePath:  o.field_commitResourcePath,
		Field_commitUrl:           o.field_commitUrl,
		Field_committedDate:       o.field_committedDate,
		Field_committedViaWeb:     o.field_committedViaWeb,
		Field_committer:           o.field_committer,
		Field_history:             o.field_history,
		Field_id:                  o.field_id,
		Field_message:             o.field_message,
		Field_messageBody:         o.field_messageBody,
		Field_messageBodyHTML:     o.field_messageBodyHTML,
		Field_messageHeadline:     o.field_messageHeadline,
		Field_messageHeadlineHTML: o.field_messageHeadlineHTML,
		Field_oid:                 o.field_oid,
		Field_repository:          o.field_repository,
		Field_resourcePath:        o.field_resourcePath,
		Field_status:              o.field_status,
		Field_tree:                o.field_tree,
		Field_treeResourcePath:    o.field_treeResourcePath,
		Field_treeUrl:             o.field_treeUrl,
		Field_url:                 o.field_url,
		Field_viewerCanSubscribe:  o.field_viewerCanSubscribe,
		Field_viewerSubscription:  o.field_viewerSubscription,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_abbreviatedOid = v.Field_abbreviatedOid
	o.field_author = v.Field_author
	o.field_authoredByCommitter = v.Field_authoredByCommitter
	o.field_blame = v.Field_blame
	o.field_comments = v.Field_comments
	o.field_commitResourcePath = v.Field_commitResourcePath
	o.field_commitUrl = v.Field_commitUrl
	o.field_committedDate = v.Field_committedDate
	o.field_committedViaWeb = v.Field_committedViaWeb
	o.field_committer = v.Field_committer
	o.field_history = v.Field_history
	o.field_id = v.Field_id
	o.field_message = v.Field_message
	o.field_messageBody = v.Field_messageBody
	o.field_messageBodyHTML = v.Field_messageBodyHTML
	o.field_messageHeadline = v.Field_messageHeadline
	o.field_messageHeadlineHTML = v.Field_messageHeadlineHTML
	o.field_oid = v.Field_oid
	o.field_repository = v.Field_repository
	o.field_resourcePath = v.Field_resourcePath
	if len(v.Field_signature) > 0 {
		o.field_signature, err = GitSignature_UnmarshalJSON(v.Field_signature)
		if err != nil {
			return err
		}
	}
	o.field_status = v.Field_status
	o.field_tree = v.Field_tree
	o.field_treeResourcePath = v.Field_treeResourcePath
	o.field_treeUrl = v.Field_treeUrl
	o.field_url = v.Field_url
	o.field_viewerCanSubscribe = v.Field_viewerCanSubscribe
	o.field_viewerSubscription = v.Field_viewerSubscription
	return nil
}

// Represents a Git tree.
type Tree struct {
	// An abbreviated version of the Git object ID
	field_abbreviatedOid string `json:"abbreviatedOid"`

	// The HTTP path for this Git object
	field_commitResourcePath URI `json:"commitResourcePath"`

	// The HTTP URL for this Git object
	field_commitUrl URI `json:"commitUrl"`

	// A list of tree entries.
	field_entries []TreeEntry `json:"entries"`

	field_id ID `json:"id"`

	// The Git object ID
	field_oid GitObjectID `json:"oid"`

	// The Repository the Git object belongs to
	field_repository Repository `json:"repository"`
}

func (o Tree) AbbreviatedOid() string {
	return o.field_abbreviatedOid
}

func (o Tree) CommitResourcePath() URI {
	return o.field_commitResourcePath
}

func (o Tree) CommitUrl() URI {
	return o.field_commitUrl
}

func (o Tree) Entries() []TreeEntry {
	return o.field_entries
}

func (o Tree) Id() ID {
	return o.field_id
}

func (o Tree) Oid() GitObjectID {
	return o.field_oid
}

func (o Tree) Repository() Repository {
	return o.field_repository
}

func (o *Tree) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_abbreviatedOid     string          `json:"abbreviatedOid"`
		Field_commitResourcePath URI             `json:"commitResourcePath"`
		Field_commitUrl          URI             `json:"commitUrl"`
		Field_entries            json.RawMessage `json:"entries"`
		Field_id                 ID              `json:"id"`
		Field_oid                GitObjectID     `json:"oid"`
		Field_repository         Repository      `json:"repository"`
	}{
		Field_abbreviatedOid:     o.field_abbreviatedOid,
		Field_commitResourcePath: o.field_commitResourcePath,
		Field_commitUrl:          o.field_commitUrl,
		Field_id:                 o.field_id,
		Field_oid:                o.field_oid,
		Field_repository:         o.field_repository,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_abbreviatedOid = v.Field_abbreviatedOid
	o.field_commitResourcePath = v.Field_commitResourcePath
	o.field_commitUrl = v.Field_commitUrl
	if len(v.Field_entries) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_entries, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_entries = nil
		} else {
			o.field_entries = make([]TreeEntry, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_entries[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_id = v.Field_id
	o.field_oid = v.Field_oid
	o.field_repository = v.Field_repository
	return nil
}

// Represents a Git tree entry.
type TreeEntry struct {
	// Entry file mode.
	field_mode int32 `json:"mode"`

	// Entry file name.
	field_name string `json:"name"`

	// Entry file object.
	field_object GitObject `json:"object"`

	// Entry file Git object ID.
	field_oid GitObjectID `json:"oid"`

	// The Repository the tree entry belongs to
	field_repository Repository `json:"repository"`

	// Entry file type.
	field_type string `json:"type"`
}

func (o TreeEntry) Mode() int32 {
	return o.field_mode
}

func (o TreeEntry) Name() string {
	return o.field_name
}

func (o TreeEntry) Object() GitObject {
	return o.field_object
}

func (o TreeEntry) Oid() GitObjectID {
	return o.field_oid
}

func (o TreeEntry) Repository() Repository {
	return o.field_repository
}

func (o TreeEntry) Type() string {
	return o.field_type
}

func (o *TreeEntry) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_mode       int32       `json:"mode"`
		Field_name       string      `json:"name"`
		Field_object     GitObject   `json:"object"`
		Field_oid        GitObjectID `json:"oid"`
		Field_repository Repository  `json:"repository"`
		Field_type       string      `json:"type"`
	}{
		Field_mode:       o.field_mode,
		Field_name:       o.field_name,
		Field_object:     o.field_object,
		Field_oid:        o.field_oid,
		Field_repository: o.field_repository,
		Field_type:       o.field_type,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_mode = v.Field_mode
	o.field_name = v.Field_name
	o.field_object = v.Field_object
	o.field_oid = v.Field_oid
	o.field_repository = v.Field_repository
	o.field_type = v.Field_type
	return nil
}

// Represents an actor in a Git commit (ie. an author or committer).
type GitActor struct {
	// A URL pointing to the author's public avatar.
	field_avatarUrl URI `json:"avatarUrl"`

	// The timestamp of the Git action (authoring or committing).
	field_date *GitTimestamp `json:"date"`

	// The email in the Git commit.
	field_email *string `json:"email"`

	// The name in the Git commit.
	field_name *string `json:"name"`

	// The GitHub user corresponding to the email field. Null if no such user exists.
	field_user *User `json:"user"`
}

func (o GitActor) AvatarUrl() URI {
	return o.field_avatarUrl
}

func (o GitActor) Date() *GitTimestamp {
	return o.field_date
}

func (o GitActor) Email() *string {
	return o.field_email
}

func (o GitActor) Name() *string {
	return o.field_name
}

func (o GitActor) User() *User {
	return o.field_user
}

func (o *GitActor) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_avatarUrl URI           `json:"avatarUrl"`
		Field_date      *GitTimestamp `json:"date"`
		Field_email     *string       `json:"email"`
		Field_name      *string       `json:"name"`
		Field_user      *User         `json:"user"`
	}{
		Field_avatarUrl: o.field_avatarUrl,
		Field_date:      o.field_date,
		Field_email:     o.field_email,
		Field_name:      o.field_name,
		Field_user:      o.field_user,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_avatarUrl = v.Field_avatarUrl
	o.field_date = v.Field_date
	o.field_email = v.Field_email
	o.field_name = v.Field_name
	o.field_user = v.Field_user
	return nil
}

// The connection type for Commit.
type CommitHistoryConnection struct {
	field_edges []*CommitEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*Commit `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`
}

func (o CommitHistoryConnection) Edges() []*CommitEdge {
	return o.field_edges
}

func (o CommitHistoryConnection) Nodes() []*Commit {
	return o.field_nodes
}

func (o CommitHistoryConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o *CommitHistoryConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges    json.RawMessage `json:"edges"`
		Field_nodes    json.RawMessage `json:"nodes"`
		Field_pageInfo PageInfo        `json:"pageInfo"`
	}{
		Field_pageInfo: o.field_pageInfo,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*CommitEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*Commit, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	return nil
}

// An edge in a connection.
type CommitEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *Commit `json:"node"`
}

func (o CommitEdge) Cursor() string {
	return o.field_cursor
}

func (o CommitEdge) Node() *Commit {
	return o.field_node
}

func (o *CommitEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string  `json:"cursor"`
		Field_node   *Commit `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// The connection type for CommitComment.
type CommitCommentConnection struct {
	// A list of edges.
	field_edges []*CommitCommentEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*CommitComment `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o CommitCommentConnection) Edges() []*CommitCommentEdge {
	return o.field_edges
}

func (o CommitCommentConnection) Nodes() []*CommitComment {
	return o.field_nodes
}

func (o CommitCommentConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o CommitCommentConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *CommitCommentConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*CommitCommentEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*CommitComment, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type CommitCommentEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *CommitComment `json:"node"`
}

func (o CommitCommentEdge) Cursor() string {
	return o.field_cursor
}

func (o CommitCommentEdge) Node() *CommitComment {
	return o.field_node
}

func (o *CommitCommentEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string         `json:"cursor"`
		Field_node   *CommitComment `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// Represents a comment on a given Commit.
type CommitComment struct {
	// The actor who authored the comment.
	field_author Actor `json:"author"`

	// Author's association with the subject of the comment.
	field_authorAssociation CommentAuthorAssociation `json:"authorAssociation"`

	// Identifies the comment body.
	field_body string `json:"body"`

	// Identifies the comment body rendered to HTML.
	field_bodyHTML HTML `json:"bodyHTML"`

	// Identifies the commit associated with the comment.
	field_commit Commit `json:"commit"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// Check if this comment was created via an email reply.
	field_createdViaEmail bool `json:"createdViaEmail"`

	// Identifies the primary key from the database.
	field_databaseId *int32 `json:"databaseId"`

	// The actor who edited the comment.
	field_editor Actor `json:"editor"`

	field_id ID `json:"id"`

	// The moment the editor made the last edit
	field_lastEditedAt *DateTime `json:"lastEditedAt"`

	// Identifies the file path associated with the comment.
	field_path *string `json:"path"`

	// Identifies the line position associated with the comment.
	field_position *int32 `json:"position"`

	// Identifies when the comment was published at.
	field_publishedAt *DateTime `json:"publishedAt"`

	// A list of reactions grouped by content left on the subject.
	field_reactionGroups []ReactionGroup `json:"reactionGroups"`

	// A list of Reactions left on the Issue.
	field_reactions ReactionConnection `json:"reactions"`

	// Identifies the repository associated with the comment.
	field_repository Repository `json:"repository"`

	// Identifies the date and time when the object was last updated.
	field_updatedAt DateTime `json:"updatedAt"`

	// Check if the current viewer can delete this object.
	field_viewerCanDelete bool `json:"viewerCanDelete"`

	// Can user react to this subject
	field_viewerCanReact bool `json:"viewerCanReact"`

	// Check if the current viewer can update this object.
	field_viewerCanUpdate bool `json:"viewerCanUpdate"`

	// Reasons why the current viewer can not update this comment.
	field_viewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`

	// Did the viewer author this comment.
	field_viewerDidAuthor bool `json:"viewerDidAuthor"`
}

func (o CommitComment) Author() Actor {
	return o.field_author
}

func (o CommitComment) AuthorAssociation() CommentAuthorAssociation {
	return o.field_authorAssociation
}

func (o CommitComment) Body() string {
	return o.field_body
}

func (o CommitComment) BodyHTML() HTML {
	return o.field_bodyHTML
}

func (o CommitComment) Commit() Commit {
	return o.field_commit
}

func (o CommitComment) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o CommitComment) CreatedViaEmail() bool {
	return o.field_createdViaEmail
}

func (o CommitComment) DatabaseId() *int32 {
	return o.field_databaseId
}

func (o CommitComment) Editor() Actor {
	return o.field_editor
}

func (o CommitComment) Id() ID {
	return o.field_id
}

func (o CommitComment) LastEditedAt() *DateTime {
	return o.field_lastEditedAt
}

func (o CommitComment) Path() *string {
	return o.field_path
}

func (o CommitComment) Position() *int32 {
	return o.field_position
}

func (o CommitComment) PublishedAt() *DateTime {
	return o.field_publishedAt
}

func (o CommitComment) ReactionGroups() []ReactionGroup {
	return o.field_reactionGroups
}

func (o CommitComment) Reactions() ReactionConnection {
	return o.field_reactions
}

func (o CommitComment) Repository() Repository {
	return o.field_repository
}

func (o CommitComment) UpdatedAt() DateTime {
	return o.field_updatedAt
}

func (o CommitComment) ViewerCanDelete() bool {
	return o.field_viewerCanDelete
}

func (o CommitComment) ViewerCanReact() bool {
	return o.field_viewerCanReact
}

func (o CommitComment) ViewerCanUpdate() bool {
	return o.field_viewerCanUpdate
}

func (o CommitComment) ViewerCannotUpdateReasons() []CommentCannotUpdateReason {
	return o.field_viewerCannotUpdateReasons
}

func (o CommitComment) ViewerDidAuthor() bool {
	return o.field_viewerDidAuthor
}

func (o *CommitComment) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_author                    json.RawMessage             `json:"author"`
		Field_authorAssociation         CommentAuthorAssociation    `json:"authorAssociation"`
		Field_body                      string                      `json:"body"`
		Field_bodyHTML                  HTML                        `json:"bodyHTML"`
		Field_commit                    Commit                      `json:"commit"`
		Field_createdAt                 DateTime                    `json:"createdAt"`
		Field_createdViaEmail           bool                        `json:"createdViaEmail"`
		Field_databaseId                *int32                      `json:"databaseId"`
		Field_editor                    json.RawMessage             `json:"editor"`
		Field_id                        ID                          `json:"id"`
		Field_lastEditedAt              *DateTime                   `json:"lastEditedAt"`
		Field_path                      *string                     `json:"path"`
		Field_position                  *int32                      `json:"position"`
		Field_publishedAt               *DateTime                   `json:"publishedAt"`
		Field_reactionGroups            json.RawMessage             `json:"reactionGroups"`
		Field_reactions                 ReactionConnection          `json:"reactions"`
		Field_repository                Repository                  `json:"repository"`
		Field_updatedAt                 DateTime                    `json:"updatedAt"`
		Field_viewerCanDelete           bool                        `json:"viewerCanDelete"`
		Field_viewerCanReact            bool                        `json:"viewerCanReact"`
		Field_viewerCanUpdate           bool                        `json:"viewerCanUpdate"`
		Field_viewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`
		Field_viewerDidAuthor           bool                        `json:"viewerDidAuthor"`
	}{
		Field_authorAssociation:         o.field_authorAssociation,
		Field_body:                      o.field_body,
		Field_bodyHTML:                  o.field_bodyHTML,
		Field_commit:                    o.field_commit,
		Field_createdAt:                 o.field_createdAt,
		Field_createdViaEmail:           o.field_createdViaEmail,
		Field_databaseId:                o.field_databaseId,
		Field_id:                        o.field_id,
		Field_lastEditedAt:              o.field_lastEditedAt,
		Field_path:                      o.field_path,
		Field_position:                  o.field_position,
		Field_publishedAt:               o.field_publishedAt,
		Field_reactions:                 o.field_reactions,
		Field_repository:                o.field_repository,
		Field_updatedAt:                 o.field_updatedAt,
		Field_viewerCanDelete:           o.field_viewerCanDelete,
		Field_viewerCanReact:            o.field_viewerCanReact,
		Field_viewerCanUpdate:           o.field_viewerCanUpdate,
		Field_viewerCannotUpdateReasons: o.field_viewerCannotUpdateReasons,
		Field_viewerDidAuthor:           o.field_viewerDidAuthor,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_author) > 0 {
		o.field_author, err = Actor_UnmarshalJSON(v.Field_author)
		if err != nil {
			return err
		}
	}
	o.field_authorAssociation = v.Field_authorAssociation
	o.field_body = v.Field_body
	o.field_bodyHTML = v.Field_bodyHTML
	o.field_commit = v.Field_commit
	o.field_createdAt = v.Field_createdAt
	o.field_createdViaEmail = v.Field_createdViaEmail
	o.field_databaseId = v.Field_databaseId
	if len(v.Field_editor) > 0 {
		o.field_editor, err = Actor_UnmarshalJSON(v.Field_editor)
		if err != nil {
			return err
		}
	}
	o.field_id = v.Field_id
	o.field_lastEditedAt = v.Field_lastEditedAt
	o.field_path = v.Field_path
	o.field_position = v.Field_position
	o.field_publishedAt = v.Field_publishedAt
	if len(v.Field_reactionGroups) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_reactionGroups, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_reactionGroups = nil
		} else {
			o.field_reactionGroups = make([]ReactionGroup, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_reactionGroups[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_reactions = v.Field_reactions
	o.field_repository = v.Field_repository
	o.field_updatedAt = v.Field_updatedAt
	o.field_viewerCanDelete = v.Field_viewerCanDelete
	o.field_viewerCanReact = v.Field_viewerCanReact
	o.field_viewerCanUpdate = v.Field_viewerCanUpdate
	o.field_viewerCannotUpdateReasons = v.Field_viewerCannotUpdateReasons
	o.field_viewerDidAuthor = v.Field_viewerDidAuthor
	return nil
}

// A group of emoji reactions to a particular piece of content.
type ReactionGroup struct {
	// Identifies the emoji reaction.
	field_content ReactionContent `json:"content"`

	// Identifies when the reaction was created.
	field_createdAt *DateTime `json:"createdAt"`

	// The subject that was reacted to.
	field_subject Reactable `json:"subject"`

	// Users who have reacted to the reaction subject with the emotion represented by this reaction group
	field_users ReactingUserConnection `json:"users"`

	// Whether or not the authenticated user has left a reaction on the subject.
	field_viewerHasReacted bool `json:"viewerHasReacted"`
}

func (o ReactionGroup) Content() ReactionContent {
	return o.field_content
}

func (o ReactionGroup) CreatedAt() *DateTime {
	return o.field_createdAt
}

func (o ReactionGroup) Subject() Reactable {
	return o.field_subject
}

func (o ReactionGroup) Users() ReactingUserConnection {
	return o.field_users
}

func (o ReactionGroup) ViewerHasReacted() bool {
	return o.field_viewerHasReacted
}

func (o *ReactionGroup) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_content          ReactionContent        `json:"content"`
		Field_createdAt        *DateTime              `json:"createdAt"`
		Field_subject          Reactable              `json:"subject"`
		Field_users            ReactingUserConnection `json:"users"`
		Field_viewerHasReacted bool                   `json:"viewerHasReacted"`
	}{
		Field_content:          o.field_content,
		Field_createdAt:        o.field_createdAt,
		Field_subject:          o.field_subject,
		Field_users:            o.field_users,
		Field_viewerHasReacted: o.field_viewerHasReacted,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_content = v.Field_content
	o.field_createdAt = v.Field_createdAt
	o.field_subject = v.Field_subject
	o.field_users = v.Field_users
	o.field_viewerHasReacted = v.Field_viewerHasReacted
	return nil
}

// A list of reactions that have been left on the subject.
type ReactionConnection struct {
	// A list of edges.
	field_edges []*ReactionEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*Reaction `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`

	// Whether or not the authenticated user has left a reaction on the subject.
	field_viewerHasReacted bool `json:"viewerHasReacted"`
}

func (o ReactionConnection) Edges() []*ReactionEdge {
	return o.field_edges
}

func (o ReactionConnection) Nodes() []*Reaction {
	return o.field_nodes
}

func (o ReactionConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o ReactionConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o ReactionConnection) ViewerHasReacted() bool {
	return o.field_viewerHasReacted
}

func (o *ReactionConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges            json.RawMessage `json:"edges"`
		Field_nodes            json.RawMessage `json:"nodes"`
		Field_pageInfo         PageInfo        `json:"pageInfo"`
		Field_totalCount       int32           `json:"totalCount"`
		Field_viewerHasReacted bool            `json:"viewerHasReacted"`
	}{
		Field_pageInfo:         o.field_pageInfo,
		Field_totalCount:       o.field_totalCount,
		Field_viewerHasReacted: o.field_viewerHasReacted,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*ReactionEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*Reaction, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	o.field_viewerHasReacted = v.Field_viewerHasReacted
	return nil
}

// An edge in a connection.
type ReactionEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *Reaction `json:"node"`
}

func (o ReactionEdge) Cursor() string {
	return o.field_cursor
}

func (o ReactionEdge) Node() *Reaction {
	return o.field_node
}

func (o *ReactionEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string    `json:"cursor"`
		Field_node   *Reaction `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// An emoji reaction to a particular piece of content.
type Reaction struct {
	// Identifies the emoji reaction.
	field_content ReactionContent `json:"content"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// Identifies the primary key from the database.
	field_databaseId *int32 `json:"databaseId"`

	field_id ID `json:"id"`

	// Identifies the user who created this reaction.
	field_user *User `json:"user"`
}

func (o Reaction) Content() ReactionContent {
	return o.field_content
}

func (o Reaction) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o Reaction) DatabaseId() *int32 {
	return o.field_databaseId
}

func (o Reaction) Id() ID {
	return o.field_id
}

func (o Reaction) User() *User {
	return o.field_user
}

func (o *Reaction) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_content    ReactionContent `json:"content"`
		Field_createdAt  DateTime        `json:"createdAt"`
		Field_databaseId *int32          `json:"databaseId"`
		Field_id         ID              `json:"id"`
		Field_user       *User           `json:"user"`
	}{
		Field_content:    o.field_content,
		Field_createdAt:  o.field_createdAt,
		Field_databaseId: o.field_databaseId,
		Field_id:         o.field_id,
		Field_user:       o.field_user,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_content = v.Field_content
	o.field_createdAt = v.Field_createdAt
	o.field_databaseId = v.Field_databaseId
	o.field_id = v.Field_id
	o.field_user = v.Field_user
	return nil
}

// The connection type for User.
type ReactingUserConnection struct {
	// A list of edges.
	field_edges []*ReactingUserEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*User `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o ReactingUserConnection) Edges() []*ReactingUserEdge {
	return o.field_edges
}

func (o ReactingUserConnection) Nodes() []*User {
	return o.field_nodes
}

func (o ReactingUserConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o ReactingUserConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *ReactingUserConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*ReactingUserEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*User, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// Represents a user that's made a reaction.
type ReactingUserEdge struct {
	field_cursor string `json:"cursor"`

	field_node User `json:"node"`

	// The moment when the user made the reaction.
	field_reactedAt DateTime `json:"reactedAt"`
}

func (o ReactingUserEdge) Cursor() string {
	return o.field_cursor
}

func (o ReactingUserEdge) Node() User {
	return o.field_node
}

func (o ReactingUserEdge) ReactedAt() DateTime {
	return o.field_reactedAt
}

func (o *ReactingUserEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor    string   `json:"cursor"`
		Field_node      User     `json:"node"`
		Field_reactedAt DateTime `json:"reactedAt"`
	}{
		Field_cursor:    o.field_cursor,
		Field_node:      o.field_node,
		Field_reactedAt: o.field_reactedAt,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	o.field_reactedAt = v.Field_reactedAt
	return nil
}

// Represents a commit status.
type Status struct {
	// The commit this status is attached to.
	field_commit *Commit `json:"commit"`

	// Looks up an individual status context by context name.
	field_context *StatusContext `json:"context"`

	// The individual status contexts for this commit.
	field_contexts []StatusContext `json:"contexts"`

	field_id ID `json:"id"`

	// The combined commit status.
	field_state StatusState `json:"state"`
}

func (o Status) Commit() *Commit {
	return o.field_commit
}

func (o Status) Context() *StatusContext {
	return o.field_context
}

func (o Status) Contexts() []StatusContext {
	return o.field_contexts
}

func (o Status) Id() ID {
	return o.field_id
}

func (o Status) State() StatusState {
	return o.field_state
}

func (o *Status) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_commit   *Commit         `json:"commit"`
		Field_context  *StatusContext  `json:"context"`
		Field_contexts []StatusContext `json:"contexts"`
		Field_id       ID              `json:"id"`
		Field_state    StatusState     `json:"state"`
	}{
		Field_commit:   o.field_commit,
		Field_context:  o.field_context,
		Field_contexts: o.field_contexts,
		Field_id:       o.field_id,
		Field_state:    o.field_state,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_commit = v.Field_commit
	o.field_context = v.Field_context
	o.field_contexts = v.Field_contexts
	o.field_id = v.Field_id
	o.field_state = v.Field_state
	return nil
}

// Represents an individual commit status context
type StatusContext struct {
	// This commit this status context is attached to.
	field_commit *Commit `json:"commit"`

	// The name of this status context.
	field_context string `json:"context"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// The actor who created this status context.
	field_creator Actor `json:"creator"`

	// The description for this status context.
	field_description *string `json:"description"`

	field_id ID `json:"id"`

	// The state of this status context.
	field_state StatusState `json:"state"`

	// The URL for this status context.
	field_targetUrl *URI `json:"targetUrl"`
}

func (o StatusContext) Commit() *Commit {
	return o.field_commit
}

func (o StatusContext) Context() string {
	return o.field_context
}

func (o StatusContext) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o StatusContext) Creator() Actor {
	return o.field_creator
}

func (o StatusContext) Description() *string {
	return o.field_description
}

func (o StatusContext) Id() ID {
	return o.field_id
}

func (o StatusContext) State() StatusState {
	return o.field_state
}

func (o StatusContext) TargetUrl() *URI {
	return o.field_targetUrl
}

func (o *StatusContext) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_commit      *Commit         `json:"commit"`
		Field_context     string          `json:"context"`
		Field_createdAt   DateTime        `json:"createdAt"`
		Field_creator     json.RawMessage `json:"creator"`
		Field_description *string         `json:"description"`
		Field_id          ID              `json:"id"`
		Field_state       StatusState     `json:"state"`
		Field_targetUrl   *URI            `json:"targetUrl"`
	}{
		Field_commit:      o.field_commit,
		Field_context:     o.field_context,
		Field_createdAt:   o.field_createdAt,
		Field_description: o.field_description,
		Field_id:          o.field_id,
		Field_state:       o.field_state,
		Field_targetUrl:   o.field_targetUrl,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_commit = v.Field_commit
	o.field_context = v.Field_context
	o.field_createdAt = v.Field_createdAt
	if len(v.Field_creator) > 0 {
		o.field_creator, err = Actor_UnmarshalJSON(v.Field_creator)
		if err != nil {
			return err
		}
	}
	o.field_description = v.Field_description
	o.field_id = v.Field_id
	o.field_state = v.Field_state
	o.field_targetUrl = v.Field_targetUrl
	return nil
}

// Represents a Git blame.
type Blame struct {
	// The list of ranges from a Git blame.
	field_ranges []BlameRange `json:"ranges"`
}

func (o Blame) Ranges() []BlameRange {
	return o.field_ranges
}

func (o *Blame) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_ranges []BlameRange `json:"ranges"`
	}{
		Field_ranges: o.field_ranges,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_ranges = v.Field_ranges
	return nil
}

// Represents a range of information from a Git blame.
type BlameRange struct {
	// Identifies the recency of the change, from 1 (new) to 10 (old). This is calculated as a 2-quantile and determines the length of distance between the median age of all the changes in the file and the recency of the current range's change.
	field_age int32 `json:"age"`

	// Identifies the line author
	field_commit Commit `json:"commit"`

	// The ending line for the range
	field_endingLine int32 `json:"endingLine"`

	// The starting line for the range
	field_startingLine int32 `json:"startingLine"`
}

func (o BlameRange) Age() int32 {
	return o.field_age
}

func (o BlameRange) Commit() Commit {
	return o.field_commit
}

func (o BlameRange) EndingLine() int32 {
	return o.field_endingLine
}

func (o BlameRange) StartingLine() int32 {
	return o.field_startingLine
}

func (o *BlameRange) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_age          int32  `json:"age"`
		Field_commit       Commit `json:"commit"`
		Field_endingLine   int32  `json:"endingLine"`
		Field_startingLine int32  `json:"startingLine"`
	}{
		Field_age:          o.field_age,
		Field_commit:       o.field_commit,
		Field_endingLine:   o.field_endingLine,
		Field_startingLine: o.field_startingLine,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_age = v.Field_age
	o.field_commit = v.Field_commit
	o.field_endingLine = v.Field_endingLine
	o.field_startingLine = v.Field_startingLine
	return nil
}

// Represents a Git blob.
type Blob struct {
	// An abbreviated version of the Git object ID
	field_abbreviatedOid string `json:"abbreviatedOid"`

	// Byte size of Blob object
	field_byteSize int32 `json:"byteSize"`

	// The HTTP path for this Git object
	field_commitResourcePath URI `json:"commitResourcePath"`

	// The HTTP URL for this Git object
	field_commitUrl URI `json:"commitUrl"`

	field_id ID `json:"id"`

	// Indicates whether the Blob is binary or text
	field_isBinary bool `json:"isBinary"`

	// Indicates whether the contents is truncated
	field_isTruncated bool `json:"isTruncated"`

	// The Git object ID
	field_oid GitObjectID `json:"oid"`

	// The Repository the Git object belongs to
	field_repository Repository `json:"repository"`

	// UTF8 text data or null if the Blob is binary
	field_text *string `json:"text"`
}

func (o Blob) AbbreviatedOid() string {
	return o.field_abbreviatedOid
}

func (o Blob) ByteSize() int32 {
	return o.field_byteSize
}

func (o Blob) CommitResourcePath() URI {
	return o.field_commitResourcePath
}

func (o Blob) CommitUrl() URI {
	return o.field_commitUrl
}

func (o Blob) Id() ID {
	return o.field_id
}

func (o Blob) IsBinary() bool {
	return o.field_isBinary
}

func (o Blob) IsTruncated() bool {
	return o.field_isTruncated
}

func (o Blob) Oid() GitObjectID {
	return o.field_oid
}

func (o Blob) Repository() Repository {
	return o.field_repository
}

func (o Blob) Text() *string {
	return o.field_text
}

func (o *Blob) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_abbreviatedOid     string      `json:"abbreviatedOid"`
		Field_byteSize           int32       `json:"byteSize"`
		Field_commitResourcePath URI         `json:"commitResourcePath"`
		Field_commitUrl          URI         `json:"commitUrl"`
		Field_id                 ID          `json:"id"`
		Field_isBinary           bool        `json:"isBinary"`
		Field_isTruncated        bool        `json:"isTruncated"`
		Field_oid                GitObjectID `json:"oid"`
		Field_repository         Repository  `json:"repository"`
		Field_text               *string     `json:"text"`
	}{
		Field_abbreviatedOid:     o.field_abbreviatedOid,
		Field_byteSize:           o.field_byteSize,
		Field_commitResourcePath: o.field_commitResourcePath,
		Field_commitUrl:          o.field_commitUrl,
		Field_id:                 o.field_id,
		Field_isBinary:           o.field_isBinary,
		Field_isTruncated:        o.field_isTruncated,
		Field_oid:                o.field_oid,
		Field_repository:         o.field_repository,
		Field_text:               o.field_text,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_abbreviatedOid = v.Field_abbreviatedOid
	o.field_byteSize = v.Field_byteSize
	o.field_commitResourcePath = v.Field_commitResourcePath
	o.field_commitUrl = v.Field_commitUrl
	o.field_id = v.Field_id
	o.field_isBinary = v.Field_isBinary
	o.field_isTruncated = v.Field_isTruncated
	o.field_oid = v.Field_oid
	o.field_repository = v.Field_repository
	o.field_text = v.Field_text
	return nil
}

// Represents a given language found in repositories.
type Language struct {
	// The color defined for the current language.
	field_color *string `json:"color"`

	field_id ID `json:"id"`

	// The name of the current language.
	field_name string `json:"name"`
}

func (o Language) Color() *string {
	return o.field_color
}

func (o Language) Id() ID {
	return o.field_id
}

func (o Language) Name() string {
	return o.field_name
}

func (o *Language) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_color *string `json:"color"`
		Field_id    ID      `json:"id"`
		Field_name  string  `json:"name"`
	}{
		Field_color: o.field_color,
		Field_id:    o.field_id,
		Field_name:  o.field_name,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_color = v.Field_color
	o.field_id = v.Field_id
	o.field_name = v.Field_name
	return nil
}

// The connection type for PullRequest.
type PullRequestConnection struct {
	// A list of edges.
	field_edges []*PullRequestEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*PullRequest `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o PullRequestConnection) Edges() []*PullRequestEdge {
	return o.field_edges
}

func (o PullRequestConnection) Nodes() []*PullRequest {
	return o.field_nodes
}

func (o PullRequestConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o PullRequestConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *PullRequestConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*PullRequestEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*PullRequest, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type PullRequestEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *PullRequest `json:"node"`
}

func (o PullRequestEdge) Cursor() string {
	return o.field_cursor
}

func (o PullRequestEdge) Node() *PullRequest {
	return o.field_node
}

func (o *PullRequestEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string       `json:"cursor"`
		Field_node   *PullRequest `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// A repository pull request.
type PullRequest struct {
	// A list of Users assigned to this object.
	field_assignees UserConnection `json:"assignees"`

	// The actor who authored the comment.
	field_author Actor `json:"author"`

	// Author's association with the subject of the comment.
	field_authorAssociation CommentAuthorAssociation `json:"authorAssociation"`

	// Identifies the base Ref associated with the pull request.
	field_baseRef *Ref `json:"baseRef"`

	// Identifies the name of the base Ref associated with the pull request, even if the ref has been deleted.
	field_baseRefName string `json:"baseRefName"`

	// Identifies the body of the pull request.
	field_body string `json:"body"`

	// Identifies the body of the pull request rendered to HTML.
	field_bodyHTML HTML `json:"bodyHTML"`

	// Identifies the body of the pull request rendered to text.
	field_bodyText string `json:"bodyText"`

	// `true` if the pull request is closed
	field_closed bool `json:"closed"`

	// A list of comments associated with the pull request.
	field_comments IssueCommentConnection `json:"comments"`

	// A list of commits present in this pull request's head branch not present in the base branch.
	field_commits PullRequestCommitConnection `json:"commits"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// Check if this comment was created via an email reply.
	field_createdViaEmail bool `json:"createdViaEmail"`

	// Identifies the primary key from the database.
	field_databaseId *int32 `json:"databaseId"`

	// The actor who edited this pull request's body.
	field_editor Actor `json:"editor"`

	// Identifies the head Ref associated with the pull request.
	field_headRef *Ref `json:"headRef"`

	// Identifies the name of the head Ref associated with the pull request, even if the ref has been deleted.
	field_headRefName string `json:"headRefName"`

	// The repository associated with this pull request's head Ref.
	field_headRepository *Repository `json:"headRepository"`

	// The owner of the repository associated with this pull request's head Ref.
	field_headRepositoryOwner RepositoryOwner `json:"headRepositoryOwner"`

	field_id ID `json:"id"`

	// The head and base repositories are different.
	field_isCrossRepository bool `json:"isCrossRepository"`

	// A list of labels associated with the object.
	field_labels *LabelConnection `json:"labels"`

	// The moment the editor made the last edit
	field_lastEditedAt *DateTime `json:"lastEditedAt"`

	// `true` if the pull request is locked
	field_locked bool `json:"locked"`

	// The commit that was created when this pull request was merged.
	field_mergeCommit *Commit `json:"mergeCommit"`

	// Whether or not the pull request can be merged based on the existence of merge conflicts.
	field_mergeable MergeableState `json:"mergeable"`

	// Whether or not the pull request was merged.
	field_merged bool `json:"merged"`

	// The date and time that the pull request was merged.
	field_mergedAt *DateTime `json:"mergedAt"`

	// Identifies the pull request number.
	field_number int32 `json:"number"`

	// A list of Users that are participating in the Pull Request conversation.
	field_participants UserConnection `json:"participants"`

	// The commit that GitHub automatically generated to test if this pull request could be merged. This field will not return a value if the pull request is merged, or if the test merge commit is still being generated. See the `mergeable` field for more details on the mergeability of the pull request.
	field_potentialMergeCommit *Commit `json:"potentialMergeCommit"`

	// Identifies when the comment was published at.
	field_publishedAt *DateTime `json:"publishedAt"`

	// A list of reactions grouped by content left on the subject.
	field_reactionGroups []ReactionGroup `json:"reactionGroups"`

	// A list of Reactions left on the Issue.
	field_reactions ReactionConnection `json:"reactions"`

	// The repository associated with this pull request.
	field_repository Repository `json:"repository"`

	// The HTTP path for this pull request.
	field_resourcePath URI `json:"resourcePath"`

	// A list of review requests associated with the pull request.
	field_reviewRequests *ReviewRequestConnection `json:"reviewRequests"`

	// A list of reviews associated with the pull request.
	field_reviews *PullRequestReviewConnection `json:"reviews"`

	// Identifies the state of the pull request.
	field_state PullRequestState `json:"state"`

	// A list of reviewer suggestions based on commit history and past review comments.
	field_suggestedReviewers []*SuggestedReviewer `json:"suggestedReviewers"`

	// A list of events associated with a PullRequest.
	field_timeline PullRequestTimelineConnection `json:"timeline"`

	// Identifies the pull request title.
	field_title string `json:"title"`

	// Identifies the date and time when the object was last updated.
	field_updatedAt DateTime `json:"updatedAt"`

	// The HTTP URL for this pull request.
	field_url URI `json:"url"`

	// Can user react to this subject
	field_viewerCanReact bool `json:"viewerCanReact"`

	// Check if the viewer is able to change their subscription status for the repository.
	field_viewerCanSubscribe bool `json:"viewerCanSubscribe"`

	// Check if the current viewer can update this object.
	field_viewerCanUpdate bool `json:"viewerCanUpdate"`

	// Reasons why the current viewer can not update this comment.
	field_viewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`

	// Did the viewer author this comment.
	field_viewerDidAuthor bool `json:"viewerDidAuthor"`

	// Identifies if the viewer is watching, not watching, or ignoring the repository.
	field_viewerSubscription SubscriptionState `json:"viewerSubscription"`
}

func (o PullRequest) Assignees() UserConnection {
	return o.field_assignees
}

func (o PullRequest) Author() Actor {
	return o.field_author
}

func (o PullRequest) AuthorAssociation() CommentAuthorAssociation {
	return o.field_authorAssociation
}

func (o PullRequest) BaseRef() *Ref {
	return o.field_baseRef
}

func (o PullRequest) BaseRefName() string {
	return o.field_baseRefName
}

func (o PullRequest) Body() string {
	return o.field_body
}

func (o PullRequest) BodyHTML() HTML {
	return o.field_bodyHTML
}

func (o PullRequest) BodyText() string {
	return o.field_bodyText
}

func (o PullRequest) Closed() bool {
	return o.field_closed
}

func (o PullRequest) Comments() IssueCommentConnection {
	return o.field_comments
}

func (o PullRequest) Commits() PullRequestCommitConnection {
	return o.field_commits
}

func (o PullRequest) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o PullRequest) CreatedViaEmail() bool {
	return o.field_createdViaEmail
}

func (o PullRequest) DatabaseId() *int32 {
	return o.field_databaseId
}

func (o PullRequest) Editor() Actor {
	return o.field_editor
}

func (o PullRequest) HeadRef() *Ref {
	return o.field_headRef
}

func (o PullRequest) HeadRefName() string {
	return o.field_headRefName
}

func (o PullRequest) HeadRepository() *Repository {
	return o.field_headRepository
}

func (o PullRequest) HeadRepositoryOwner() RepositoryOwner {
	return o.field_headRepositoryOwner
}

func (o PullRequest) Id() ID {
	return o.field_id
}

func (o PullRequest) IsCrossRepository() bool {
	return o.field_isCrossRepository
}

func (o PullRequest) Labels() *LabelConnection {
	return o.field_labels
}

func (o PullRequest) LastEditedAt() *DateTime {
	return o.field_lastEditedAt
}

func (o PullRequest) Locked() bool {
	return o.field_locked
}

func (o PullRequest) MergeCommit() *Commit {
	return o.field_mergeCommit
}

func (o PullRequest) Mergeable() MergeableState {
	return o.field_mergeable
}

func (o PullRequest) Merged() bool {
	return o.field_merged
}

func (o PullRequest) MergedAt() *DateTime {
	return o.field_mergedAt
}

func (o PullRequest) Number() int32 {
	return o.field_number
}

func (o PullRequest) Participants() UserConnection {
	return o.field_participants
}

func (o PullRequest) PotentialMergeCommit() *Commit {
	return o.field_potentialMergeCommit
}

func (o PullRequest) PublishedAt() *DateTime {
	return o.field_publishedAt
}

func (o PullRequest) ReactionGroups() []ReactionGroup {
	return o.field_reactionGroups
}

func (o PullRequest) Reactions() ReactionConnection {
	return o.field_reactions
}

func (o PullRequest) Repository() Repository {
	return o.field_repository
}

func (o PullRequest) ResourcePath() URI {
	return o.field_resourcePath
}

func (o PullRequest) ReviewRequests() *ReviewRequestConnection {
	return o.field_reviewRequests
}

func (o PullRequest) Reviews() *PullRequestReviewConnection {
	return o.field_reviews
}

func (o PullRequest) State() PullRequestState {
	return o.field_state
}

func (o PullRequest) SuggestedReviewers() []*SuggestedReviewer {
	return o.field_suggestedReviewers
}

func (o PullRequest) Timeline() PullRequestTimelineConnection {
	return o.field_timeline
}

func (o PullRequest) Title() string {
	return o.field_title
}

func (o PullRequest) UpdatedAt() DateTime {
	return o.field_updatedAt
}

func (o PullRequest) Url() URI {
	return o.field_url
}

func (o PullRequest) ViewerCanReact() bool {
	return o.field_viewerCanReact
}

func (o PullRequest) ViewerCanSubscribe() bool {
	return o.field_viewerCanSubscribe
}

func (o PullRequest) ViewerCanUpdate() bool {
	return o.field_viewerCanUpdate
}

func (o PullRequest) ViewerCannotUpdateReasons() []CommentCannotUpdateReason {
	return o.field_viewerCannotUpdateReasons
}

func (o PullRequest) ViewerDidAuthor() bool {
	return o.field_viewerDidAuthor
}

func (o PullRequest) ViewerSubscription() SubscriptionState {
	return o.field_viewerSubscription
}

func (o *PullRequest) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_assignees                 UserConnection                `json:"assignees"`
		Field_author                    json.RawMessage               `json:"author"`
		Field_authorAssociation         CommentAuthorAssociation      `json:"authorAssociation"`
		Field_baseRef                   *Ref                          `json:"baseRef"`
		Field_baseRefName               string                        `json:"baseRefName"`
		Field_body                      string                        `json:"body"`
		Field_bodyHTML                  HTML                          `json:"bodyHTML"`
		Field_bodyText                  string                        `json:"bodyText"`
		Field_closed                    bool                          `json:"closed"`
		Field_comments                  IssueCommentConnection        `json:"comments"`
		Field_commits                   PullRequestCommitConnection   `json:"commits"`
		Field_createdAt                 DateTime                      `json:"createdAt"`
		Field_createdViaEmail           bool                          `json:"createdViaEmail"`
		Field_databaseId                *int32                        `json:"databaseId"`
		Field_editor                    json.RawMessage               `json:"editor"`
		Field_headRef                   *Ref                          `json:"headRef"`
		Field_headRefName               string                        `json:"headRefName"`
		Field_headRepository            *Repository                   `json:"headRepository"`
		Field_headRepositoryOwner       json.RawMessage               `json:"headRepositoryOwner"`
		Field_id                        ID                            `json:"id"`
		Field_isCrossRepository         bool                          `json:"isCrossRepository"`
		Field_labels                    *LabelConnection              `json:"labels"`
		Field_lastEditedAt              *DateTime                     `json:"lastEditedAt"`
		Field_locked                    bool                          `json:"locked"`
		Field_mergeCommit               *Commit                       `json:"mergeCommit"`
		Field_mergeable                 MergeableState                `json:"mergeable"`
		Field_merged                    bool                          `json:"merged"`
		Field_mergedAt                  *DateTime                     `json:"mergedAt"`
		Field_number                    int32                         `json:"number"`
		Field_participants              UserConnection                `json:"participants"`
		Field_potentialMergeCommit      *Commit                       `json:"potentialMergeCommit"`
		Field_publishedAt               *DateTime                     `json:"publishedAt"`
		Field_reactionGroups            json.RawMessage               `json:"reactionGroups"`
		Field_reactions                 ReactionConnection            `json:"reactions"`
		Field_repository                Repository                    `json:"repository"`
		Field_resourcePath              URI                           `json:"resourcePath"`
		Field_reviewRequests            *ReviewRequestConnection      `json:"reviewRequests"`
		Field_reviews                   *PullRequestReviewConnection  `json:"reviews"`
		Field_state                     PullRequestState              `json:"state"`
		Field_suggestedReviewers        []*SuggestedReviewer          `json:"suggestedReviewers"`
		Field_timeline                  PullRequestTimelineConnection `json:"timeline"`
		Field_title                     string                        `json:"title"`
		Field_updatedAt                 DateTime                      `json:"updatedAt"`
		Field_url                       URI                           `json:"url"`
		Field_viewerCanReact            bool                          `json:"viewerCanReact"`
		Field_viewerCanSubscribe        bool                          `json:"viewerCanSubscribe"`
		Field_viewerCanUpdate           bool                          `json:"viewerCanUpdate"`
		Field_viewerCannotUpdateReasons []CommentCannotUpdateReason   `json:"viewerCannotUpdateReasons"`
		Field_viewerDidAuthor           bool                          `json:"viewerDidAuthor"`
		Field_viewerSubscription        SubscriptionState             `json:"viewerSubscription"`
	}{
		Field_assignees:                 o.field_assignees,
		Field_authorAssociation:         o.field_authorAssociation,
		Field_baseRef:                   o.field_baseRef,
		Field_baseRefName:               o.field_baseRefName,
		Field_body:                      o.field_body,
		Field_bodyHTML:                  o.field_bodyHTML,
		Field_bodyText:                  o.field_bodyText,
		Field_closed:                    o.field_closed,
		Field_comments:                  o.field_comments,
		Field_commits:                   o.field_commits,
		Field_createdAt:                 o.field_createdAt,
		Field_createdViaEmail:           o.field_createdViaEmail,
		Field_databaseId:                o.field_databaseId,
		Field_headRef:                   o.field_headRef,
		Field_headRefName:               o.field_headRefName,
		Field_headRepository:            o.field_headRepository,
		Field_id:                        o.field_id,
		Field_isCrossRepository:         o.field_isCrossRepository,
		Field_labels:                    o.field_labels,
		Field_lastEditedAt:              o.field_lastEditedAt,
		Field_locked:                    o.field_locked,
		Field_mergeCommit:               o.field_mergeCommit,
		Field_mergeable:                 o.field_mergeable,
		Field_merged:                    o.field_merged,
		Field_mergedAt:                  o.field_mergedAt,
		Field_number:                    o.field_number,
		Field_participants:              o.field_participants,
		Field_potentialMergeCommit:      o.field_potentialMergeCommit,
		Field_publishedAt:               o.field_publishedAt,
		Field_reactions:                 o.field_reactions,
		Field_repository:                o.field_repository,
		Field_resourcePath:              o.field_resourcePath,
		Field_reviewRequests:            o.field_reviewRequests,
		Field_reviews:                   o.field_reviews,
		Field_state:                     o.field_state,
		Field_suggestedReviewers:        o.field_suggestedReviewers,
		Field_timeline:                  o.field_timeline,
		Field_title:                     o.field_title,
		Field_updatedAt:                 o.field_updatedAt,
		Field_url:                       o.field_url,
		Field_viewerCanReact:            o.field_viewerCanReact,
		Field_viewerCanSubscribe:        o.field_viewerCanSubscribe,
		Field_viewerCanUpdate:           o.field_viewerCanUpdate,
		Field_viewerCannotUpdateReasons: o.field_viewerCannotUpdateReasons,
		Field_viewerDidAuthor:           o.field_viewerDidAuthor,
		Field_viewerSubscription:        o.field_viewerSubscription,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_assignees = v.Field_assignees
	if len(v.Field_author) > 0 {
		o.field_author, err = Actor_UnmarshalJSON(v.Field_author)
		if err != nil {
			return err
		}
	}
	o.field_authorAssociation = v.Field_authorAssociation
	o.field_baseRef = v.Field_baseRef
	o.field_baseRefName = v.Field_baseRefName
	o.field_body = v.Field_body
	o.field_bodyHTML = v.Field_bodyHTML
	o.field_bodyText = v.Field_bodyText
	o.field_closed = v.Field_closed
	o.field_comments = v.Field_comments
	o.field_commits = v.Field_commits
	o.field_createdAt = v.Field_createdAt
	o.field_createdViaEmail = v.Field_createdViaEmail
	o.field_databaseId = v.Field_databaseId
	if len(v.Field_editor) > 0 {
		o.field_editor, err = Actor_UnmarshalJSON(v.Field_editor)
		if err != nil {
			return err
		}
	}
	o.field_headRef = v.Field_headRef
	o.field_headRefName = v.Field_headRefName
	o.field_headRepository = v.Field_headRepository
	if len(v.Field_headRepositoryOwner) > 0 {
		o.field_headRepositoryOwner, err = RepositoryOwner_UnmarshalJSON(v.Field_headRepositoryOwner)
		if err != nil {
			return err
		}
	}
	o.field_id = v.Field_id
	o.field_isCrossRepository = v.Field_isCrossRepository
	o.field_labels = v.Field_labels
	o.field_lastEditedAt = v.Field_lastEditedAt
	o.field_locked = v.Field_locked
	o.field_mergeCommit = v.Field_mergeCommit
	o.field_mergeable = v.Field_mergeable
	o.field_merged = v.Field_merged
	o.field_mergedAt = v.Field_mergedAt
	o.field_number = v.Field_number
	o.field_participants = v.Field_participants
	o.field_potentialMergeCommit = v.Field_potentialMergeCommit
	o.field_publishedAt = v.Field_publishedAt
	if len(v.Field_reactionGroups) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_reactionGroups, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_reactionGroups = nil
		} else {
			o.field_reactionGroups = make([]ReactionGroup, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_reactionGroups[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_reactions = v.Field_reactions
	o.field_repository = v.Field_repository
	o.field_resourcePath = v.Field_resourcePath
	o.field_reviewRequests = v.Field_reviewRequests
	o.field_reviews = v.Field_reviews
	o.field_state = v.Field_state
	o.field_suggestedReviewers = v.Field_suggestedReviewers
	o.field_timeline = v.Field_timeline
	o.field_title = v.Field_title
	o.field_updatedAt = v.Field_updatedAt
	o.field_url = v.Field_url
	o.field_viewerCanReact = v.Field_viewerCanReact
	o.field_viewerCanSubscribe = v.Field_viewerCanSubscribe
	o.field_viewerCanUpdate = v.Field_viewerCanUpdate
	o.field_viewerCannotUpdateReasons = v.Field_viewerCannotUpdateReasons
	o.field_viewerDidAuthor = v.Field_viewerDidAuthor
	o.field_viewerSubscription = v.Field_viewerSubscription
	return nil
}

// The connection type for Label.
type LabelConnection struct {
	// A list of edges.
	field_edges []*LabelEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*Label `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o LabelConnection) Edges() []*LabelEdge {
	return o.field_edges
}

func (o LabelConnection) Nodes() []*Label {
	return o.field_nodes
}

func (o LabelConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o LabelConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *LabelConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*LabelEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*Label, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type LabelEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *Label `json:"node"`
}

func (o LabelEdge) Cursor() string {
	return o.field_cursor
}

func (o LabelEdge) Node() *Label {
	return o.field_node
}

func (o *LabelEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string `json:"cursor"`
		Field_node   *Label `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// A label for categorizing Issues or Milestones with a given Repository.
type Label struct {
	// Identifies the label color.
	field_color string `json:"color"`

	field_id ID `json:"id"`

	// A list of issues associated with this label.
	field_issues IssueConnection `json:"issues"`

	// Identifies the label name.
	field_name string `json:"name"`

	// A list of pull requests associated with this label.
	field_pullRequests *PullRequestConnection `json:"pullRequests"`

	// The repository associated with this label.
	field_repository Repository `json:"repository"`
}

func (o Label) Color() string {
	return o.field_color
}

func (o Label) Id() ID {
	return o.field_id
}

func (o Label) Issues() IssueConnection {
	return o.field_issues
}

func (o Label) Name() string {
	return o.field_name
}

func (o Label) PullRequests() *PullRequestConnection {
	return o.field_pullRequests
}

func (o Label) Repository() Repository {
	return o.field_repository
}

func (o *Label) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_color        string                 `json:"color"`
		Field_id           ID                     `json:"id"`
		Field_issues       IssueConnection        `json:"issues"`
		Field_name         string                 `json:"name"`
		Field_pullRequests *PullRequestConnection `json:"pullRequests"`
		Field_repository   Repository             `json:"repository"`
	}{
		Field_color:        o.field_color,
		Field_id:           o.field_id,
		Field_issues:       o.field_issues,
		Field_name:         o.field_name,
		Field_pullRequests: o.field_pullRequests,
		Field_repository:   o.field_repository,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_color = v.Field_color
	o.field_id = v.Field_id
	o.field_issues = v.Field_issues
	o.field_name = v.Field_name
	o.field_pullRequests = v.Field_pullRequests
	o.field_repository = v.Field_repository
	return nil
}

// The connection type for Issue.
type IssueConnection struct {
	// A list of edges.
	field_edges []*IssueEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*Issue `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o IssueConnection) Edges() []*IssueEdge {
	return o.field_edges
}

func (o IssueConnection) Nodes() []*Issue {
	return o.field_nodes
}

func (o IssueConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o IssueConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *IssueConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*IssueEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*Issue, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type IssueEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *Issue `json:"node"`
}

func (o IssueEdge) Cursor() string {
	return o.field_cursor
}

func (o IssueEdge) Node() *Issue {
	return o.field_node
}

func (o *IssueEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string `json:"cursor"`
		Field_node   *Issue `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// A list of repositories owned by the subject.
type RepositoryConnection struct {
	// A list of edges.
	field_edges []*RepositoryEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*Repository `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`

	// The total size in kilobytes of all repositories in the connection.
	field_totalDiskUsage int32 `json:"totalDiskUsage"`
}

func (o RepositoryConnection) Edges() []*RepositoryEdge {
	return o.field_edges
}

func (o RepositoryConnection) Nodes() []*Repository {
	return o.field_nodes
}

func (o RepositoryConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o RepositoryConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o RepositoryConnection) TotalDiskUsage() int32 {
	return o.field_totalDiskUsage
}

func (o *RepositoryConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges          json.RawMessage `json:"edges"`
		Field_nodes          json.RawMessage `json:"nodes"`
		Field_pageInfo       PageInfo        `json:"pageInfo"`
		Field_totalCount     int32           `json:"totalCount"`
		Field_totalDiskUsage int32           `json:"totalDiskUsage"`
	}{
		Field_pageInfo:       o.field_pageInfo,
		Field_totalCount:     o.field_totalCount,
		Field_totalDiskUsage: o.field_totalDiskUsage,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*RepositoryEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*Repository, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	o.field_totalDiskUsage = v.Field_totalDiskUsage
	return nil
}

// An edge in a connection.
type RepositoryEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *Repository `json:"node"`
}

func (o RepositoryEdge) Cursor() string {
	return o.field_cursor
}

func (o RepositoryEdge) Node() *Repository {
	return o.field_node
}

func (o *RepositoryEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string      `json:"cursor"`
		Field_node   *Repository `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// The connection type for IssueComment.
type IssueCommentConnection struct {
	// A list of edges.
	field_edges []*IssueCommentEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*IssueComment `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o IssueCommentConnection) Edges() []*IssueCommentEdge {
	return o.field_edges
}

func (o IssueCommentConnection) Nodes() []*IssueComment {
	return o.field_nodes
}

func (o IssueCommentConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o IssueCommentConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *IssueCommentConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*IssueCommentEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*IssueComment, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type IssueCommentEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *IssueComment `json:"node"`
}

func (o IssueCommentEdge) Cursor() string {
	return o.field_cursor
}

func (o IssueCommentEdge) Node() *IssueComment {
	return o.field_node
}

func (o *IssueCommentEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string        `json:"cursor"`
		Field_node   *IssueComment `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// Represents a comment on an Issue.
type IssueComment struct {
	// The actor who authored the comment.
	field_author Actor `json:"author"`

	// Author's association with the subject of the comment.
	field_authorAssociation CommentAuthorAssociation `json:"authorAssociation"`

	// Identifies the comment body.
	field_body string `json:"body"`

	// The comment body rendered to HTML.
	field_bodyHTML HTML `json:"bodyHTML"`

	// Identifies the body of the issue rendered to text.
	field_bodyText string `json:"bodyText"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// Check if this comment was created via an email reply.
	field_createdViaEmail bool `json:"createdViaEmail"`

	// Identifies the primary key from the database.
	field_databaseId *int32 `json:"databaseId"`

	// The actor who edited the comment.
	field_editor Actor `json:"editor"`

	field_id ID `json:"id"`

	// Identifies the issue associated with the comment.
	field_issue Issue `json:"issue"`

	// The moment the editor made the last edit
	field_lastEditedAt *DateTime `json:"lastEditedAt"`

	// Identifies when the comment was published at.
	field_publishedAt *DateTime `json:"publishedAt"`

	// A list of reactions grouped by content left on the subject.
	field_reactionGroups []ReactionGroup `json:"reactionGroups"`

	// A list of Reactions left on the Issue.
	field_reactions ReactionConnection `json:"reactions"`

	// The repository associated with this node.
	field_repository Repository `json:"repository"`

	// Identifies the date and time when the object was last updated.
	field_updatedAt DateTime `json:"updatedAt"`

	// Check if the current viewer can delete this object.
	field_viewerCanDelete bool `json:"viewerCanDelete"`

	// Can user react to this subject
	field_viewerCanReact bool `json:"viewerCanReact"`

	// Check if the current viewer can update this object.
	field_viewerCanUpdate bool `json:"viewerCanUpdate"`

	// Reasons why the current viewer can not update this comment.
	field_viewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`

	// Did the viewer author this comment.
	field_viewerDidAuthor bool `json:"viewerDidAuthor"`
}

func (o IssueComment) Author() Actor {
	return o.field_author
}

func (o IssueComment) AuthorAssociation() CommentAuthorAssociation {
	return o.field_authorAssociation
}

func (o IssueComment) Body() string {
	return o.field_body
}

func (o IssueComment) BodyHTML() HTML {
	return o.field_bodyHTML
}

func (o IssueComment) BodyText() string {
	return o.field_bodyText
}

func (o IssueComment) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o IssueComment) CreatedViaEmail() bool {
	return o.field_createdViaEmail
}

func (o IssueComment) DatabaseId() *int32 {
	return o.field_databaseId
}

func (o IssueComment) Editor() Actor {
	return o.field_editor
}

func (o IssueComment) Id() ID {
	return o.field_id
}

func (o IssueComment) Issue() Issue {
	return o.field_issue
}

func (o IssueComment) LastEditedAt() *DateTime {
	return o.field_lastEditedAt
}

func (o IssueComment) PublishedAt() *DateTime {
	return o.field_publishedAt
}

func (o IssueComment) ReactionGroups() []ReactionGroup {
	return o.field_reactionGroups
}

func (o IssueComment) Reactions() ReactionConnection {
	return o.field_reactions
}

func (o IssueComment) Repository() Repository {
	return o.field_repository
}

func (o IssueComment) UpdatedAt() DateTime {
	return o.field_updatedAt
}

func (o IssueComment) ViewerCanDelete() bool {
	return o.field_viewerCanDelete
}

func (o IssueComment) ViewerCanReact() bool {
	return o.field_viewerCanReact
}

func (o IssueComment) ViewerCanUpdate() bool {
	return o.field_viewerCanUpdate
}

func (o IssueComment) ViewerCannotUpdateReasons() []CommentCannotUpdateReason {
	return o.field_viewerCannotUpdateReasons
}

func (o IssueComment) ViewerDidAuthor() bool {
	return o.field_viewerDidAuthor
}

func (o *IssueComment) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_author                    json.RawMessage             `json:"author"`
		Field_authorAssociation         CommentAuthorAssociation    `json:"authorAssociation"`
		Field_body                      string                      `json:"body"`
		Field_bodyHTML                  HTML                        `json:"bodyHTML"`
		Field_bodyText                  string                      `json:"bodyText"`
		Field_createdAt                 DateTime                    `json:"createdAt"`
		Field_createdViaEmail           bool                        `json:"createdViaEmail"`
		Field_databaseId                *int32                      `json:"databaseId"`
		Field_editor                    json.RawMessage             `json:"editor"`
		Field_id                        ID                          `json:"id"`
		Field_issue                     Issue                       `json:"issue"`
		Field_lastEditedAt              *DateTime                   `json:"lastEditedAt"`
		Field_publishedAt               *DateTime                   `json:"publishedAt"`
		Field_reactionGroups            json.RawMessage             `json:"reactionGroups"`
		Field_reactions                 ReactionConnection          `json:"reactions"`
		Field_repository                Repository                  `json:"repository"`
		Field_updatedAt                 DateTime                    `json:"updatedAt"`
		Field_viewerCanDelete           bool                        `json:"viewerCanDelete"`
		Field_viewerCanReact            bool                        `json:"viewerCanReact"`
		Field_viewerCanUpdate           bool                        `json:"viewerCanUpdate"`
		Field_viewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`
		Field_viewerDidAuthor           bool                        `json:"viewerDidAuthor"`
	}{
		Field_authorAssociation:         o.field_authorAssociation,
		Field_body:                      o.field_body,
		Field_bodyHTML:                  o.field_bodyHTML,
		Field_bodyText:                  o.field_bodyText,
		Field_createdAt:                 o.field_createdAt,
		Field_createdViaEmail:           o.field_createdViaEmail,
		Field_databaseId:                o.field_databaseId,
		Field_id:                        o.field_id,
		Field_issue:                     o.field_issue,
		Field_lastEditedAt:              o.field_lastEditedAt,
		Field_publishedAt:               o.field_publishedAt,
		Field_reactions:                 o.field_reactions,
		Field_repository:                o.field_repository,
		Field_updatedAt:                 o.field_updatedAt,
		Field_viewerCanDelete:           o.field_viewerCanDelete,
		Field_viewerCanReact:            o.field_viewerCanReact,
		Field_viewerCanUpdate:           o.field_viewerCanUpdate,
		Field_viewerCannotUpdateReasons: o.field_viewerCannotUpdateReasons,
		Field_viewerDidAuthor:           o.field_viewerDidAuthor,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_author) > 0 {
		o.field_author, err = Actor_UnmarshalJSON(v.Field_author)
		if err != nil {
			return err
		}
	}
	o.field_authorAssociation = v.Field_authorAssociation
	o.field_body = v.Field_body
	o.field_bodyHTML = v.Field_bodyHTML
	o.field_bodyText = v.Field_bodyText
	o.field_createdAt = v.Field_createdAt
	o.field_createdViaEmail = v.Field_createdViaEmail
	o.field_databaseId = v.Field_databaseId
	if len(v.Field_editor) > 0 {
		o.field_editor, err = Actor_UnmarshalJSON(v.Field_editor)
		if err != nil {
			return err
		}
	}
	o.field_id = v.Field_id
	o.field_issue = v.Field_issue
	o.field_lastEditedAt = v.Field_lastEditedAt
	o.field_publishedAt = v.Field_publishedAt
	if len(v.Field_reactionGroups) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_reactionGroups, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_reactionGroups = nil
		} else {
			o.field_reactionGroups = make([]ReactionGroup, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_reactionGroups[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_reactions = v.Field_reactions
	o.field_repository = v.Field_repository
	o.field_updatedAt = v.Field_updatedAt
	o.field_viewerCanDelete = v.Field_viewerCanDelete
	o.field_viewerCanReact = v.Field_viewerCanReact
	o.field_viewerCanUpdate = v.Field_viewerCanUpdate
	o.field_viewerCannotUpdateReasons = v.Field_viewerCannotUpdateReasons
	o.field_viewerDidAuthor = v.Field_viewerDidAuthor
	return nil
}

// The connection type for PullRequestReview.
type PullRequestReviewConnection struct {
	// A list of edges.
	field_edges []*PullRequestReviewEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*PullRequestReview `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o PullRequestReviewConnection) Edges() []*PullRequestReviewEdge {
	return o.field_edges
}

func (o PullRequestReviewConnection) Nodes() []*PullRequestReview {
	return o.field_nodes
}

func (o PullRequestReviewConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o PullRequestReviewConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *PullRequestReviewConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*PullRequestReviewEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*PullRequestReview, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type PullRequestReviewEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *PullRequestReview `json:"node"`
}

func (o PullRequestReviewEdge) Cursor() string {
	return o.field_cursor
}

func (o PullRequestReviewEdge) Node() *PullRequestReview {
	return o.field_node
}

func (o *PullRequestReviewEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string             `json:"cursor"`
		Field_node   *PullRequestReview `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// A review object for a given pull request.
type PullRequestReview struct {
	// The actor who authored the comment.
	field_author Actor `json:"author"`

	// Author's association with the subject of the comment.
	field_authorAssociation CommentAuthorAssociation `json:"authorAssociation"`

	// Identifies the pull request review body.
	field_body string `json:"body"`

	// The body of this review rendered to HTML.
	field_bodyHTML HTML `json:"bodyHTML"`

	// The body of this review rendered as plain text.
	field_bodyText string `json:"bodyText"`

	// A list of review comments for the current pull request review.
	field_comments PullRequestReviewCommentConnection `json:"comments"`

	// Identifies the commit associated with this pull request review.
	field_commit *Commit `json:"commit"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// Check if this comment was created via an email reply.
	field_createdViaEmail bool `json:"createdViaEmail"`

	// Identifies the primary key from the database.
	field_databaseId *int32 `json:"databaseId"`

	// The actor who edited the comment.
	field_editor Actor `json:"editor"`

	field_id ID `json:"id"`

	// The moment the editor made the last edit
	field_lastEditedAt *DateTime `json:"lastEditedAt"`

	// Identifies when the comment was published at.
	field_publishedAt *DateTime `json:"publishedAt"`

	// Identifies the pull request associated with this pull request review.
	field_pullRequest PullRequest `json:"pullRequest"`

	// The repository associated with this node.
	field_repository Repository `json:"repository"`

	// The HTTP path permalink for this PullRequestReview.
	field_resourcePath URI `json:"resourcePath"`

	// Identifies the current state of the pull request review.
	field_state PullRequestReviewState `json:"state"`

	// Identifies when the Pull Request Review was submitted
	field_submittedAt *DateTime `json:"submittedAt"`

	// Identifies the date and time when the object was last updated.
	field_updatedAt DateTime `json:"updatedAt"`

	// The HTTP URL permalink for this PullRequestReview.
	field_url URI `json:"url"`

	// Check if the current viewer can delete this object.
	field_viewerCanDelete bool `json:"viewerCanDelete"`

	// Check if the current viewer can update this object.
	field_viewerCanUpdate bool `json:"viewerCanUpdate"`

	// Reasons why the current viewer can not update this comment.
	field_viewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`

	// Did the viewer author this comment.
	field_viewerDidAuthor bool `json:"viewerDidAuthor"`
}

func (o PullRequestReview) Author() Actor {
	return o.field_author
}

func (o PullRequestReview) AuthorAssociation() CommentAuthorAssociation {
	return o.field_authorAssociation
}

func (o PullRequestReview) Body() string {
	return o.field_body
}

func (o PullRequestReview) BodyHTML() HTML {
	return o.field_bodyHTML
}

func (o PullRequestReview) BodyText() string {
	return o.field_bodyText
}

func (o PullRequestReview) Comments() PullRequestReviewCommentConnection {
	return o.field_comments
}

func (o PullRequestReview) Commit() *Commit {
	return o.field_commit
}

func (o PullRequestReview) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o PullRequestReview) CreatedViaEmail() bool {
	return o.field_createdViaEmail
}

func (o PullRequestReview) DatabaseId() *int32 {
	return o.field_databaseId
}

func (o PullRequestReview) Editor() Actor {
	return o.field_editor
}

func (o PullRequestReview) Id() ID {
	return o.field_id
}

func (o PullRequestReview) LastEditedAt() *DateTime {
	return o.field_lastEditedAt
}

func (o PullRequestReview) PublishedAt() *DateTime {
	return o.field_publishedAt
}

func (o PullRequestReview) PullRequest() PullRequest {
	return o.field_pullRequest
}

func (o PullRequestReview) Repository() Repository {
	return o.field_repository
}

func (o PullRequestReview) ResourcePath() URI {
	return o.field_resourcePath
}

func (o PullRequestReview) State() PullRequestReviewState {
	return o.field_state
}

func (o PullRequestReview) SubmittedAt() *DateTime {
	return o.field_submittedAt
}

func (o PullRequestReview) UpdatedAt() DateTime {
	return o.field_updatedAt
}

func (o PullRequestReview) Url() URI {
	return o.field_url
}

func (o PullRequestReview) ViewerCanDelete() bool {
	return o.field_viewerCanDelete
}

func (o PullRequestReview) ViewerCanUpdate() bool {
	return o.field_viewerCanUpdate
}

func (o PullRequestReview) ViewerCannotUpdateReasons() []CommentCannotUpdateReason {
	return o.field_viewerCannotUpdateReasons
}

func (o PullRequestReview) ViewerDidAuthor() bool {
	return o.field_viewerDidAuthor
}

func (o *PullRequestReview) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_author                    json.RawMessage                    `json:"author"`
		Field_authorAssociation         CommentAuthorAssociation           `json:"authorAssociation"`
		Field_body                      string                             `json:"body"`
		Field_bodyHTML                  HTML                               `json:"bodyHTML"`
		Field_bodyText                  string                             `json:"bodyText"`
		Field_comments                  PullRequestReviewCommentConnection `json:"comments"`
		Field_commit                    *Commit                            `json:"commit"`
		Field_createdAt                 DateTime                           `json:"createdAt"`
		Field_createdViaEmail           bool                               `json:"createdViaEmail"`
		Field_databaseId                *int32                             `json:"databaseId"`
		Field_editor                    json.RawMessage                    `json:"editor"`
		Field_id                        ID                                 `json:"id"`
		Field_lastEditedAt              *DateTime                          `json:"lastEditedAt"`
		Field_publishedAt               *DateTime                          `json:"publishedAt"`
		Field_pullRequest               PullRequest                        `json:"pullRequest"`
		Field_repository                Repository                         `json:"repository"`
		Field_resourcePath              URI                                `json:"resourcePath"`
		Field_state                     PullRequestReviewState             `json:"state"`
		Field_submittedAt               *DateTime                          `json:"submittedAt"`
		Field_updatedAt                 DateTime                           `json:"updatedAt"`
		Field_url                       URI                                `json:"url"`
		Field_viewerCanDelete           bool                               `json:"viewerCanDelete"`
		Field_viewerCanUpdate           bool                               `json:"viewerCanUpdate"`
		Field_viewerCannotUpdateReasons []CommentCannotUpdateReason        `json:"viewerCannotUpdateReasons"`
		Field_viewerDidAuthor           bool                               `json:"viewerDidAuthor"`
	}{
		Field_authorAssociation:         o.field_authorAssociation,
		Field_body:                      o.field_body,
		Field_bodyHTML:                  o.field_bodyHTML,
		Field_bodyText:                  o.field_bodyText,
		Field_comments:                  o.field_comments,
		Field_commit:                    o.field_commit,
		Field_createdAt:                 o.field_createdAt,
		Field_createdViaEmail:           o.field_createdViaEmail,
		Field_databaseId:                o.field_databaseId,
		Field_id:                        o.field_id,
		Field_lastEditedAt:              o.field_lastEditedAt,
		Field_publishedAt:               o.field_publishedAt,
		Field_pullRequest:               o.field_pullRequest,
		Field_repository:                o.field_repository,
		Field_resourcePath:              o.field_resourcePath,
		Field_state:                     o.field_state,
		Field_submittedAt:               o.field_submittedAt,
		Field_updatedAt:                 o.field_updatedAt,
		Field_url:                       o.field_url,
		Field_viewerCanDelete:           o.field_viewerCanDelete,
		Field_viewerCanUpdate:           o.field_viewerCanUpdate,
		Field_viewerCannotUpdateReasons: o.field_viewerCannotUpdateReasons,
		Field_viewerDidAuthor:           o.field_viewerDidAuthor,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_author) > 0 {
		o.field_author, err = Actor_UnmarshalJSON(v.Field_author)
		if err != nil {
			return err
		}
	}
	o.field_authorAssociation = v.Field_authorAssociation
	o.field_body = v.Field_body
	o.field_bodyHTML = v.Field_bodyHTML
	o.field_bodyText = v.Field_bodyText
	o.field_comments = v.Field_comments
	o.field_commit = v.Field_commit
	o.field_createdAt = v.Field_createdAt
	o.field_createdViaEmail = v.Field_createdViaEmail
	o.field_databaseId = v.Field_databaseId
	if len(v.Field_editor) > 0 {
		o.field_editor, err = Actor_UnmarshalJSON(v.Field_editor)
		if err != nil {
			return err
		}
	}
	o.field_id = v.Field_id
	o.field_lastEditedAt = v.Field_lastEditedAt
	o.field_publishedAt = v.Field_publishedAt
	o.field_pullRequest = v.Field_pullRequest
	o.field_repository = v.Field_repository
	o.field_resourcePath = v.Field_resourcePath
	o.field_state = v.Field_state
	o.field_submittedAt = v.Field_submittedAt
	o.field_updatedAt = v.Field_updatedAt
	o.field_url = v.Field_url
	o.field_viewerCanDelete = v.Field_viewerCanDelete
	o.field_viewerCanUpdate = v.Field_viewerCanUpdate
	o.field_viewerCannotUpdateReasons = v.Field_viewerCannotUpdateReasons
	o.field_viewerDidAuthor = v.Field_viewerDidAuthor
	return nil
}

// The connection type for PullRequestReviewComment.
type PullRequestReviewCommentConnection struct {
	// A list of edges.
	field_edges []*PullRequestReviewCommentEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*PullRequestReviewComment `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o PullRequestReviewCommentConnection) Edges() []*PullRequestReviewCommentEdge {
	return o.field_edges
}

func (o PullRequestReviewCommentConnection) Nodes() []*PullRequestReviewComment {
	return o.field_nodes
}

func (o PullRequestReviewCommentConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o PullRequestReviewCommentConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *PullRequestReviewCommentConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*PullRequestReviewCommentEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*PullRequestReviewComment, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type PullRequestReviewCommentEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *PullRequestReviewComment `json:"node"`
}

func (o PullRequestReviewCommentEdge) Cursor() string {
	return o.field_cursor
}

func (o PullRequestReviewCommentEdge) Node() *PullRequestReviewComment {
	return o.field_node
}

func (o *PullRequestReviewCommentEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string                    `json:"cursor"`
		Field_node   *PullRequestReviewComment `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// A review comment associated with a given repository pull request.
type PullRequestReviewComment struct {
	// The actor who authored the comment.
	field_author Actor `json:"author"`

	// Author's association with the subject of the comment.
	field_authorAssociation CommentAuthorAssociation `json:"authorAssociation"`

	// The comment body of this review comment.
	field_body string `json:"body"`

	// The comment body of this review comment rendered to HTML.
	field_bodyHTML HTML `json:"bodyHTML"`

	// The comment body of this review comment rendered as plain text.
	field_bodyText string `json:"bodyText"`

	// Identifies the commit associated with the comment.
	field_commit Commit `json:"commit"`

	// Identifies when the comment was created.
	field_createdAt DateTime `json:"createdAt"`

	// Check if this comment was created via an email reply.
	field_createdViaEmail bool `json:"createdViaEmail"`

	// Identifies the primary key from the database.
	field_databaseId *int32 `json:"databaseId"`

	// The diff hunk to which the comment applies.
	field_diffHunk string `json:"diffHunk"`

	// Identifies when the comment was created in a draft state.
	field_draftedAt DateTime `json:"draftedAt"`

	// The actor who edited the comment.
	field_editor Actor `json:"editor"`

	field_id ID `json:"id"`

	// The moment the editor made the last edit
	field_lastEditedAt *DateTime `json:"lastEditedAt"`

	// Identifies the original commit associated with the comment.
	field_originalCommit Commit `json:"originalCommit"`

	// The original line index in the diff to which the comment applies.
	field_originalPosition int32 `json:"originalPosition"`

	// The path to which the comment applies.
	field_path string `json:"path"`

	// The line index in the diff to which the comment applies.
	field_position *int32 `json:"position"`

	// Identifies when the comment was published at.
	field_publishedAt *DateTime `json:"publishedAt"`

	// The pull request associated with this review comment.
	field_pullRequest PullRequest `json:"pullRequest"`

	// The pull request review associated with this review comment.
	field_pullRequestReview *PullRequestReview `json:"pullRequestReview"`

	// A list of reactions grouped by content left on the subject.
	field_reactionGroups []ReactionGroup `json:"reactionGroups"`

	// A list of Reactions left on the Issue.
	field_reactions ReactionConnection `json:"reactions"`

	// The repository associated with this review comment.
	field_repository Repository `json:"repository"`

	// The HTTP path permalink for this review comment.
	field_resourcePath URI `json:"resourcePath"`

	// Identifies when the comment was last updated.
	field_updatedAt DateTime `json:"updatedAt"`

	// The HTTP URL permalink for this review comment.
	field_url URI `json:"url"`

	// Check if the current viewer can delete this object.
	field_viewerCanDelete bool `json:"viewerCanDelete"`

	// Can user react to this subject
	field_viewerCanReact bool `json:"viewerCanReact"`

	// Check if the current viewer can update this object.
	field_viewerCanUpdate bool `json:"viewerCanUpdate"`

	// Reasons why the current viewer can not update this comment.
	field_viewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`

	// Did the viewer author this comment.
	field_viewerDidAuthor bool `json:"viewerDidAuthor"`
}

func (o PullRequestReviewComment) Author() Actor {
	return o.field_author
}

func (o PullRequestReviewComment) AuthorAssociation() CommentAuthorAssociation {
	return o.field_authorAssociation
}

func (o PullRequestReviewComment) Body() string {
	return o.field_body
}

func (o PullRequestReviewComment) BodyHTML() HTML {
	return o.field_bodyHTML
}

func (o PullRequestReviewComment) BodyText() string {
	return o.field_bodyText
}

func (o PullRequestReviewComment) Commit() Commit {
	return o.field_commit
}

func (o PullRequestReviewComment) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o PullRequestReviewComment) CreatedViaEmail() bool {
	return o.field_createdViaEmail
}

func (o PullRequestReviewComment) DatabaseId() *int32 {
	return o.field_databaseId
}

func (o PullRequestReviewComment) DiffHunk() string {
	return o.field_diffHunk
}

func (o PullRequestReviewComment) DraftedAt() DateTime {
	return o.field_draftedAt
}

func (o PullRequestReviewComment) Editor() Actor {
	return o.field_editor
}

func (o PullRequestReviewComment) Id() ID {
	return o.field_id
}

func (o PullRequestReviewComment) LastEditedAt() *DateTime {
	return o.field_lastEditedAt
}

func (o PullRequestReviewComment) OriginalCommit() Commit {
	return o.field_originalCommit
}

func (o PullRequestReviewComment) OriginalPosition() int32 {
	return o.field_originalPosition
}

func (o PullRequestReviewComment) Path() string {
	return o.field_path
}

func (o PullRequestReviewComment) Position() *int32 {
	return o.field_position
}

func (o PullRequestReviewComment) PublishedAt() *DateTime {
	return o.field_publishedAt
}

func (o PullRequestReviewComment) PullRequest() PullRequest {
	return o.field_pullRequest
}

func (o PullRequestReviewComment) PullRequestReview() *PullRequestReview {
	return o.field_pullRequestReview
}

func (o PullRequestReviewComment) ReactionGroups() []ReactionGroup {
	return o.field_reactionGroups
}

func (o PullRequestReviewComment) Reactions() ReactionConnection {
	return o.field_reactions
}

func (o PullRequestReviewComment) Repository() Repository {
	return o.field_repository
}

func (o PullRequestReviewComment) ResourcePath() URI {
	return o.field_resourcePath
}

func (o PullRequestReviewComment) UpdatedAt() DateTime {
	return o.field_updatedAt
}

func (o PullRequestReviewComment) Url() URI {
	return o.field_url
}

func (o PullRequestReviewComment) ViewerCanDelete() bool {
	return o.field_viewerCanDelete
}

func (o PullRequestReviewComment) ViewerCanReact() bool {
	return o.field_viewerCanReact
}

func (o PullRequestReviewComment) ViewerCanUpdate() bool {
	return o.field_viewerCanUpdate
}

func (o PullRequestReviewComment) ViewerCannotUpdateReasons() []CommentCannotUpdateReason {
	return o.field_viewerCannotUpdateReasons
}

func (o PullRequestReviewComment) ViewerDidAuthor() bool {
	return o.field_viewerDidAuthor
}

func (o *PullRequestReviewComment) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_author                    json.RawMessage             `json:"author"`
		Field_authorAssociation         CommentAuthorAssociation    `json:"authorAssociation"`
		Field_body                      string                      `json:"body"`
		Field_bodyHTML                  HTML                        `json:"bodyHTML"`
		Field_bodyText                  string                      `json:"bodyText"`
		Field_commit                    Commit                      `json:"commit"`
		Field_createdAt                 DateTime                    `json:"createdAt"`
		Field_createdViaEmail           bool                        `json:"createdViaEmail"`
		Field_databaseId                *int32                      `json:"databaseId"`
		Field_diffHunk                  string                      `json:"diffHunk"`
		Field_draftedAt                 DateTime                    `json:"draftedAt"`
		Field_editor                    json.RawMessage             `json:"editor"`
		Field_id                        ID                          `json:"id"`
		Field_lastEditedAt              *DateTime                   `json:"lastEditedAt"`
		Field_originalCommit            Commit                      `json:"originalCommit"`
		Field_originalPosition          int32                       `json:"originalPosition"`
		Field_path                      string                      `json:"path"`
		Field_position                  *int32                      `json:"position"`
		Field_publishedAt               *DateTime                   `json:"publishedAt"`
		Field_pullRequest               PullRequest                 `json:"pullRequest"`
		Field_pullRequestReview         *PullRequestReview          `json:"pullRequestReview"`
		Field_reactionGroups            json.RawMessage             `json:"reactionGroups"`
		Field_reactions                 ReactionConnection          `json:"reactions"`
		Field_repository                Repository                  `json:"repository"`
		Field_resourcePath              URI                         `json:"resourcePath"`
		Field_updatedAt                 DateTime                    `json:"updatedAt"`
		Field_url                       URI                         `json:"url"`
		Field_viewerCanDelete           bool                        `json:"viewerCanDelete"`
		Field_viewerCanReact            bool                        `json:"viewerCanReact"`
		Field_viewerCanUpdate           bool                        `json:"viewerCanUpdate"`
		Field_viewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`
		Field_viewerDidAuthor           bool                        `json:"viewerDidAuthor"`
	}{
		Field_authorAssociation:         o.field_authorAssociation,
		Field_body:                      o.field_body,
		Field_bodyHTML:                  o.field_bodyHTML,
		Field_bodyText:                  o.field_bodyText,
		Field_commit:                    o.field_commit,
		Field_createdAt:                 o.field_createdAt,
		Field_createdViaEmail:           o.field_createdViaEmail,
		Field_databaseId:                o.field_databaseId,
		Field_diffHunk:                  o.field_diffHunk,
		Field_draftedAt:                 o.field_draftedAt,
		Field_id:                        o.field_id,
		Field_lastEditedAt:              o.field_lastEditedAt,
		Field_originalCommit:            o.field_originalCommit,
		Field_originalPosition:          o.field_originalPosition,
		Field_path:                      o.field_path,
		Field_position:                  o.field_position,
		Field_publishedAt:               o.field_publishedAt,
		Field_pullRequest:               o.field_pullRequest,
		Field_pullRequestReview:         o.field_pullRequestReview,
		Field_reactions:                 o.field_reactions,
		Field_repository:                o.field_repository,
		Field_resourcePath:              o.field_resourcePath,
		Field_updatedAt:                 o.field_updatedAt,
		Field_url:                       o.field_url,
		Field_viewerCanDelete:           o.field_viewerCanDelete,
		Field_viewerCanReact:            o.field_viewerCanReact,
		Field_viewerCanUpdate:           o.field_viewerCanUpdate,
		Field_viewerCannotUpdateReasons: o.field_viewerCannotUpdateReasons,
		Field_viewerDidAuthor:           o.field_viewerDidAuthor,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_author) > 0 {
		o.field_author, err = Actor_UnmarshalJSON(v.Field_author)
		if err != nil {
			return err
		}
	}
	o.field_authorAssociation = v.Field_authorAssociation
	o.field_body = v.Field_body
	o.field_bodyHTML = v.Field_bodyHTML
	o.field_bodyText = v.Field_bodyText
	o.field_commit = v.Field_commit
	o.field_createdAt = v.Field_createdAt
	o.field_createdViaEmail = v.Field_createdViaEmail
	o.field_databaseId = v.Field_databaseId
	o.field_diffHunk = v.Field_diffHunk
	o.field_draftedAt = v.Field_draftedAt
	if len(v.Field_editor) > 0 {
		o.field_editor, err = Actor_UnmarshalJSON(v.Field_editor)
		if err != nil {
			return err
		}
	}
	o.field_id = v.Field_id
	o.field_lastEditedAt = v.Field_lastEditedAt
	o.field_originalCommit = v.Field_originalCommit
	o.field_originalPosition = v.Field_originalPosition
	o.field_path = v.Field_path
	o.field_position = v.Field_position
	o.field_publishedAt = v.Field_publishedAt
	o.field_pullRequest = v.Field_pullRequest
	o.field_pullRequestReview = v.Field_pullRequestReview
	if len(v.Field_reactionGroups) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_reactionGroups, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_reactionGroups = nil
		} else {
			o.field_reactionGroups = make([]ReactionGroup, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_reactionGroups[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_reactions = v.Field_reactions
	o.field_repository = v.Field_repository
	o.field_resourcePath = v.Field_resourcePath
	o.field_updatedAt = v.Field_updatedAt
	o.field_url = v.Field_url
	o.field_viewerCanDelete = v.Field_viewerCanDelete
	o.field_viewerCanReact = v.Field_viewerCanReact
	o.field_viewerCanUpdate = v.Field_viewerCanUpdate
	o.field_viewerCannotUpdateReasons = v.Field_viewerCannotUpdateReasons
	o.field_viewerDidAuthor = v.Field_viewerDidAuthor
	return nil
}

// The connection type for PullRequestCommit.
type PullRequestCommitConnection struct {
	// A list of edges.
	field_edges []*PullRequestCommitEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*PullRequestCommit `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o PullRequestCommitConnection) Edges() []*PullRequestCommitEdge {
	return o.field_edges
}

func (o PullRequestCommitConnection) Nodes() []*PullRequestCommit {
	return o.field_nodes
}

func (o PullRequestCommitConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o PullRequestCommitConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *PullRequestCommitConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*PullRequestCommitEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*PullRequestCommit, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type PullRequestCommitEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *PullRequestCommit `json:"node"`
}

func (o PullRequestCommitEdge) Cursor() string {
	return o.field_cursor
}

func (o PullRequestCommitEdge) Node() *PullRequestCommit {
	return o.field_node
}

func (o *PullRequestCommitEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string             `json:"cursor"`
		Field_node   *PullRequestCommit `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// Represents a Git commit part of a pull request.
type PullRequestCommit struct {
	// The Git commit object
	field_commit Commit `json:"commit"`

	field_id ID `json:"id"`

	// The pull request this commit belongs to
	field_pullRequest PullRequest `json:"pullRequest"`

	// The HTTP path for this pull request commit
	field_resourcePath URI `json:"resourcePath"`

	// The HTTP URL for this pull request commit
	field_url URI `json:"url"`
}

func (o PullRequestCommit) Commit() Commit {
	return o.field_commit
}

func (o PullRequestCommit) Id() ID {
	return o.field_id
}

func (o PullRequestCommit) PullRequest() PullRequest {
	return o.field_pullRequest
}

func (o PullRequestCommit) ResourcePath() URI {
	return o.field_resourcePath
}

func (o PullRequestCommit) Url() URI {
	return o.field_url
}

func (o *PullRequestCommit) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_commit       Commit      `json:"commit"`
		Field_id           ID          `json:"id"`
		Field_pullRequest  PullRequest `json:"pullRequest"`
		Field_resourcePath URI         `json:"resourcePath"`
		Field_url          URI         `json:"url"`
	}{
		Field_commit:       o.field_commit,
		Field_id:           o.field_id,
		Field_pullRequest:  o.field_pullRequest,
		Field_resourcePath: o.field_resourcePath,
		Field_url:          o.field_url,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_commit = v.Field_commit
	o.field_id = v.Field_id
	o.field_pullRequest = v.Field_pullRequest
	o.field_resourcePath = v.Field_resourcePath
	o.field_url = v.Field_url
	return nil
}

// The connection type for ReviewRequest.
type ReviewRequestConnection struct {
	// A list of edges.
	field_edges []*ReviewRequestEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*ReviewRequest `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o ReviewRequestConnection) Edges() []*ReviewRequestEdge {
	return o.field_edges
}

func (o ReviewRequestConnection) Nodes() []*ReviewRequest {
	return o.field_nodes
}

func (o ReviewRequestConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o ReviewRequestConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *ReviewRequestConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*ReviewRequestEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*ReviewRequest, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type ReviewRequestEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *ReviewRequest `json:"node"`
}

func (o ReviewRequestEdge) Cursor() string {
	return o.field_cursor
}

func (o ReviewRequestEdge) Node() *ReviewRequest {
	return o.field_node
}

func (o *ReviewRequestEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string         `json:"cursor"`
		Field_node   *ReviewRequest `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// A request for a user to review a pull request.
type ReviewRequest struct {
	// Identifies the primary key from the database.
	field_databaseId *int32 `json:"databaseId"`

	field_id ID `json:"id"`

	// Identifies the pull request associated with this review request.
	field_pullRequest PullRequest `json:"pullRequest"`

	// Identifies the author associated with this review request.
	field_reviewer *User `json:"reviewer"`
}

func (o ReviewRequest) DatabaseId() *int32 {
	return o.field_databaseId
}

func (o ReviewRequest) Id() ID {
	return o.field_id
}

func (o ReviewRequest) PullRequest() PullRequest {
	return o.field_pullRequest
}

func (o ReviewRequest) Reviewer() *User {
	return o.field_reviewer
}

func (o *ReviewRequest) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_databaseId  *int32      `json:"databaseId"`
		Field_id          ID          `json:"id"`
		Field_pullRequest PullRequest `json:"pullRequest"`
		Field_reviewer    *User       `json:"reviewer"`
	}{
		Field_databaseId:  o.field_databaseId,
		Field_id:          o.field_id,
		Field_pullRequest: o.field_pullRequest,
		Field_reviewer:    o.field_reviewer,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_databaseId = v.Field_databaseId
	o.field_id = v.Field_id
	o.field_pullRequest = v.Field_pullRequest
	o.field_reviewer = v.Field_reviewer
	return nil
}

// A team of users in an organization.
type Team struct {
	// The description of the team.
	field_description *string `json:"description"`

	// The HTTP path for editing this team
	field_editTeamResourcePath URI `json:"editTeamResourcePath"`

	// The HTTP URL for editing this team
	field_editTeamUrl URI `json:"editTeamUrl"`

	field_id ID `json:"id"`

	// A list of pending invitations for users to this team
	field_invitations *OrganizationInvitationConnection `json:"invitations"`

	// The name of the team.
	field_name string `json:"name"`

	// The organization that owns this team.
	field_organization Organization `json:"organization"`

	// The level of privacy the team has.
	field_privacy TeamPrivacy `json:"privacy"`

	// The HTTP path for this team
	field_resourcePath URI `json:"resourcePath"`

	// The slug corresponding to the team.
	field_slug string `json:"slug"`

	// The HTTP URL for this team
	field_url URI `json:"url"`
}

func (o Team) Description() *string {
	return o.field_description
}

func (o Team) EditTeamResourcePath() URI {
	return o.field_editTeamResourcePath
}

func (o Team) EditTeamUrl() URI {
	return o.field_editTeamUrl
}

func (o Team) Id() ID {
	return o.field_id
}

func (o Team) Invitations() *OrganizationInvitationConnection {
	return o.field_invitations
}

func (o Team) Name() string {
	return o.field_name
}

func (o Team) Organization() Organization {
	return o.field_organization
}

func (o Team) Privacy() TeamPrivacy {
	return o.field_privacy
}

func (o Team) ResourcePath() URI {
	return o.field_resourcePath
}

func (o Team) Slug() string {
	return o.field_slug
}

func (o Team) Url() URI {
	return o.field_url
}

func (o *Team) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_description          *string                           `json:"description"`
		Field_editTeamResourcePath URI                               `json:"editTeamResourcePath"`
		Field_editTeamUrl          URI                               `json:"editTeamUrl"`
		Field_id                   ID                                `json:"id"`
		Field_invitations          *OrganizationInvitationConnection `json:"invitations"`
		Field_name                 string                            `json:"name"`
		Field_organization         Organization                      `json:"organization"`
		Field_privacy              TeamPrivacy                       `json:"privacy"`
		Field_resourcePath         URI                               `json:"resourcePath"`
		Field_slug                 string                            `json:"slug"`
		Field_url                  URI                               `json:"url"`
	}{
		Field_description:          o.field_description,
		Field_editTeamResourcePath: o.field_editTeamResourcePath,
		Field_editTeamUrl:          o.field_editTeamUrl,
		Field_id:                   o.field_id,
		Field_invitations:          o.field_invitations,
		Field_name:                 o.field_name,
		Field_organization:         o.field_organization,
		Field_privacy:              o.field_privacy,
		Field_resourcePath:         o.field_resourcePath,
		Field_slug:                 o.field_slug,
		Field_url:                  o.field_url,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_description = v.Field_description
	o.field_editTeamResourcePath = v.Field_editTeamResourcePath
	o.field_editTeamUrl = v.Field_editTeamUrl
	o.field_id = v.Field_id
	o.field_invitations = v.Field_invitations
	o.field_name = v.Field_name
	o.field_organization = v.Field_organization
	o.field_privacy = v.Field_privacy
	o.field_resourcePath = v.Field_resourcePath
	o.field_slug = v.Field_slug
	o.field_url = v.Field_url
	return nil
}

// The connection type for OrganizationInvitation.
type OrganizationInvitationConnection struct {
	// A list of edges.
	field_edges []*OrganizationInvitationEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*OrganizationInvitation `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o OrganizationInvitationConnection) Edges() []*OrganizationInvitationEdge {
	return o.field_edges
}

func (o OrganizationInvitationConnection) Nodes() []*OrganizationInvitation {
	return o.field_nodes
}

func (o OrganizationInvitationConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o OrganizationInvitationConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *OrganizationInvitationConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*OrganizationInvitationEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*OrganizationInvitation, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type OrganizationInvitationEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *OrganizationInvitation `json:"node"`
}

func (o OrganizationInvitationEdge) Cursor() string {
	return o.field_cursor
}

func (o OrganizationInvitationEdge) Node() *OrganizationInvitation {
	return o.field_node
}

func (o *OrganizationInvitationEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string                  `json:"cursor"`
		Field_node   *OrganizationInvitation `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// An Invitation for a user to an organization.
type OrganizationInvitation struct {
	// The email address of the user invited to the organization.
	field_email *string `json:"email"`

	field_id ID `json:"id"`

	// The user who was invited to the organization.
	field_invitee *User `json:"invitee"`

	// The user who created the invitation.
	field_inviter User `json:"inviter"`

	// The user's pending role in the organization (e.g. member, owner).
	field_role OrganizationInvitationRole `json:"role"`
}

func (o OrganizationInvitation) Email() *string {
	return o.field_email
}

func (o OrganizationInvitation) Id() ID {
	return o.field_id
}

func (o OrganizationInvitation) Invitee() *User {
	return o.field_invitee
}

func (o OrganizationInvitation) Inviter() User {
	return o.field_inviter
}

func (o OrganizationInvitation) Role() OrganizationInvitationRole {
	return o.field_role
}

func (o *OrganizationInvitation) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_email   *string                    `json:"email"`
		Field_id      ID                         `json:"id"`
		Field_invitee *User                      `json:"invitee"`
		Field_inviter User                       `json:"inviter"`
		Field_role    OrganizationInvitationRole `json:"role"`
	}{
		Field_email:   o.field_email,
		Field_id:      o.field_id,
		Field_invitee: o.field_invitee,
		Field_inviter: o.field_inviter,
		Field_role:    o.field_role,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_email = v.Field_email
	o.field_id = v.Field_id
	o.field_invitee = v.Field_invitee
	o.field_inviter = v.Field_inviter
	o.field_role = v.Field_role
	return nil
}

// The connection type for Team.
type TeamConnection struct {
	// A list of edges.
	field_edges []*TeamEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*Team `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o TeamConnection) Edges() []*TeamEdge {
	return o.field_edges
}

func (o TeamConnection) Nodes() []*Team {
	return o.field_nodes
}

func (o TeamConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o TeamConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *TeamConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*TeamEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*Team, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type TeamEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *Team `json:"node"`
}

func (o TeamEdge) Cursor() string {
	return o.field_cursor
}

func (o TeamEdge) Node() *Team {
	return o.field_node
}

func (o *TeamEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string `json:"cursor"`
		Field_node   *Team  `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// An account on GitHub, with one or more owners, that has repositories, members and teams.
type Organization struct {
	// A URL pointing to the organization's public avatar.
	field_avatarUrl URI `json:"avatarUrl"`

	// Identifies the primary key from the database.
	field_databaseId *int32 `json:"databaseId"`

	field_id ID `json:"id"`

	// Is the account billed through invoices?
	field_isInvoiced bool `json:"isInvoiced"`

	// The organization's login name.
	field_login string `json:"login"`

	// A list of users who are members of this organization.
	field_members UserConnection `json:"members"`

	// The organization's public profile name.
	field_name string `json:"name"`

	// The HTTP path creating a new team
	field_newTeamResourcePath URI `json:"newTeamResourcePath"`

	// The HTTP URL creating a new team
	field_newTeamUrl URI `json:"newTeamUrl"`

	// The billing email for the organization.
	field_organizationBillingEmail *string `json:"organizationBillingEmail"`

	// Find project by number.
	field_project *Project `json:"project"`

	// A list of projects under the owner.
	field_projects ProjectConnection `json:"projects"`

	// The HTTP path listing organization's projects
	field_projectsResourcePath URI `json:"projectsResourcePath"`

	// The HTTP URL listing organization's projects
	field_projectsUrl URI `json:"projectsUrl"`

	// A list of repositories that the user owns.
	field_repositories RepositoryConnection `json:"repositories"`

	// Find Repository.
	field_repository *Repository `json:"repository"`

	// The HTTP path for this user
	field_resourcePath URI `json:"resourcePath"`

	// The Organization's SAML Identity Providers
	field_samlIdentityProvider *OrganizationIdentityProvider `json:"samlIdentityProvider"`

	// Find an organization's team by its slug.
	field_team *Team `json:"team"`

	// A list of teams in this organization.
	field_teams TeamConnection `json:"teams"`

	// The HTTP path listing organization's teams
	field_teamsResourcePath URI `json:"teamsResourcePath"`

	// The HTTP URL listing organization's teams
	field_teamsUrl URI `json:"teamsUrl"`

	// The HTTP URL for this user
	field_url URI `json:"url"`

	// Organization is adminable by the viewer.
	field_viewerCanAdminister bool `json:"viewerCanAdminister"`

	// Can the current viewer create new projects on this owner.
	field_viewerCanCreateProjects bool `json:"viewerCanCreateProjects"`

	// Viewer can create repositories on this organization
	field_viewerCanCreateRepositories bool `json:"viewerCanCreateRepositories"`

	// Viewer can create teams on this organization.
	field_viewerCanCreateTeams bool `json:"viewerCanCreateTeams"`

	// Viewer is a member of this organization.
	field_viewerIsAMember bool `json:"viewerIsAMember"`
}

func (o Organization) AvatarUrl() URI {
	return o.field_avatarUrl
}

func (o Organization) DatabaseId() *int32 {
	return o.field_databaseId
}

func (o Organization) Id() ID {
	return o.field_id
}

func (o Organization) IsInvoiced() bool {
	return o.field_isInvoiced
}

func (o Organization) Login() string {
	return o.field_login
}

func (o Organization) Members() UserConnection {
	return o.field_members
}

func (o Organization) Name() string {
	return o.field_name
}

func (o Organization) NewTeamResourcePath() URI {
	return o.field_newTeamResourcePath
}

func (o Organization) NewTeamUrl() URI {
	return o.field_newTeamUrl
}

func (o Organization) OrganizationBillingEmail() *string {
	return o.field_organizationBillingEmail
}

func (o Organization) Project() *Project {
	return o.field_project
}

func (o Organization) Projects() ProjectConnection {
	return o.field_projects
}

func (o Organization) ProjectsResourcePath() URI {
	return o.field_projectsResourcePath
}

func (o Organization) ProjectsUrl() URI {
	return o.field_projectsUrl
}

func (o Organization) Repositories() RepositoryConnection {
	return o.field_repositories
}

func (o Organization) Repository() *Repository {
	return o.field_repository
}

func (o Organization) ResourcePath() URI {
	return o.field_resourcePath
}

func (o Organization) SamlIdentityProvider() *OrganizationIdentityProvider {
	return o.field_samlIdentityProvider
}

func (o Organization) Team() *Team {
	return o.field_team
}

func (o Organization) Teams() TeamConnection {
	return o.field_teams
}

func (o Organization) TeamsResourcePath() URI {
	return o.field_teamsResourcePath
}

func (o Organization) TeamsUrl() URI {
	return o.field_teamsUrl
}

func (o Organization) Url() URI {
	return o.field_url
}

func (o Organization) ViewerCanAdminister() bool {
	return o.field_viewerCanAdminister
}

func (o Organization) ViewerCanCreateProjects() bool {
	return o.field_viewerCanCreateProjects
}

func (o Organization) ViewerCanCreateRepositories() bool {
	return o.field_viewerCanCreateRepositories
}

func (o Organization) ViewerCanCreateTeams() bool {
	return o.field_viewerCanCreateTeams
}

func (o Organization) ViewerIsAMember() bool {
	return o.field_viewerIsAMember
}

func (o *Organization) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_avatarUrl                   URI                           `json:"avatarUrl"`
		Field_databaseId                  *int32                        `json:"databaseId"`
		Field_id                          ID                            `json:"id"`
		Field_isInvoiced                  bool                          `json:"isInvoiced"`
		Field_login                       string                        `json:"login"`
		Field_members                     UserConnection                `json:"members"`
		Field_name                        string                        `json:"name"`
		Field_newTeamResourcePath         URI                           `json:"newTeamResourcePath"`
		Field_newTeamUrl                  URI                           `json:"newTeamUrl"`
		Field_organizationBillingEmail    *string                       `json:"organizationBillingEmail"`
		Field_project                     *Project                      `json:"project"`
		Field_projects                    ProjectConnection             `json:"projects"`
		Field_projectsResourcePath        URI                           `json:"projectsResourcePath"`
		Field_projectsUrl                 URI                           `json:"projectsUrl"`
		Field_repositories                RepositoryConnection          `json:"repositories"`
		Field_repository                  *Repository                   `json:"repository"`
		Field_resourcePath                URI                           `json:"resourcePath"`
		Field_samlIdentityProvider        *OrganizationIdentityProvider `json:"samlIdentityProvider"`
		Field_team                        *Team                         `json:"team"`
		Field_teams                       TeamConnection                `json:"teams"`
		Field_teamsResourcePath           URI                           `json:"teamsResourcePath"`
		Field_teamsUrl                    URI                           `json:"teamsUrl"`
		Field_url                         URI                           `json:"url"`
		Field_viewerCanAdminister         bool                          `json:"viewerCanAdminister"`
		Field_viewerCanCreateProjects     bool                          `json:"viewerCanCreateProjects"`
		Field_viewerCanCreateRepositories bool                          `json:"viewerCanCreateRepositories"`
		Field_viewerCanCreateTeams        bool                          `json:"viewerCanCreateTeams"`
		Field_viewerIsAMember             bool                          `json:"viewerIsAMember"`
	}{
		Field_avatarUrl:                   o.field_avatarUrl,
		Field_databaseId:                  o.field_databaseId,
		Field_id:                          o.field_id,
		Field_isInvoiced:                  o.field_isInvoiced,
		Field_login:                       o.field_login,
		Field_members:                     o.field_members,
		Field_name:                        o.field_name,
		Field_newTeamResourcePath:         o.field_newTeamResourcePath,
		Field_newTeamUrl:                  o.field_newTeamUrl,
		Field_organizationBillingEmail:    o.field_organizationBillingEmail,
		Field_project:                     o.field_project,
		Field_projects:                    o.field_projects,
		Field_projectsResourcePath:        o.field_projectsResourcePath,
		Field_projectsUrl:                 o.field_projectsUrl,
		Field_repositories:                o.field_repositories,
		Field_repository:                  o.field_repository,
		Field_resourcePath:                o.field_resourcePath,
		Field_samlIdentityProvider:        o.field_samlIdentityProvider,
		Field_team:                        o.field_team,
		Field_teams:                       o.field_teams,
		Field_teamsResourcePath:           o.field_teamsResourcePath,
		Field_teamsUrl:                    o.field_teamsUrl,
		Field_url:                         o.field_url,
		Field_viewerCanAdminister:         o.field_viewerCanAdminister,
		Field_viewerCanCreateProjects:     o.field_viewerCanCreateProjects,
		Field_viewerCanCreateRepositories: o.field_viewerCanCreateRepositories,
		Field_viewerCanCreateTeams:        o.field_viewerCanCreateTeams,
		Field_viewerIsAMember:             o.field_viewerIsAMember,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_avatarUrl = v.Field_avatarUrl
	o.field_databaseId = v.Field_databaseId
	o.field_id = v.Field_id
	o.field_isInvoiced = v.Field_isInvoiced
	o.field_login = v.Field_login
	o.field_members = v.Field_members
	o.field_name = v.Field_name
	o.field_newTeamResourcePath = v.Field_newTeamResourcePath
	o.field_newTeamUrl = v.Field_newTeamUrl
	o.field_organizationBillingEmail = v.Field_organizationBillingEmail
	o.field_project = v.Field_project
	o.field_projects = v.Field_projects
	o.field_projectsResourcePath = v.Field_projectsResourcePath
	o.field_projectsUrl = v.Field_projectsUrl
	o.field_repositories = v.Field_repositories
	o.field_repository = v.Field_repository
	o.field_resourcePath = v.Field_resourcePath
	o.field_samlIdentityProvider = v.Field_samlIdentityProvider
	o.field_team = v.Field_team
	o.field_teams = v.Field_teams
	o.field_teamsResourcePath = v.Field_teamsResourcePath
	o.field_teamsUrl = v.Field_teamsUrl
	o.field_url = v.Field_url
	o.field_viewerCanAdminister = v.Field_viewerCanAdminister
	o.field_viewerCanCreateProjects = v.Field_viewerCanCreateProjects
	o.field_viewerCanCreateRepositories = v.Field_viewerCanCreateRepositories
	o.field_viewerCanCreateTeams = v.Field_viewerCanCreateTeams
	o.field_viewerIsAMember = v.Field_viewerIsAMember
	return nil
}

// A list of languages associated with the parent.
type LanguageConnection struct {
	// A list of edges.
	field_edges []*LanguageEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*Language `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`

	// The total size in bytes of files written in that language.
	field_totalSize int32 `json:"totalSize"`
}

func (o LanguageConnection) Edges() []*LanguageEdge {
	return o.field_edges
}

func (o LanguageConnection) Nodes() []*Language {
	return o.field_nodes
}

func (o LanguageConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o LanguageConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o LanguageConnection) TotalSize() int32 {
	return o.field_totalSize
}

func (o *LanguageConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
		Field_totalSize  int32           `json:"totalSize"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
		Field_totalSize:  o.field_totalSize,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*LanguageEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*Language, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	o.field_totalSize = v.Field_totalSize
	return nil
}

// Represents the language of a repository.
type LanguageEdge struct {
	field_cursor string `json:"cursor"`

	field_node Language `json:"node"`

	// The number of bytes of code written in the language.
	field_size int32 `json:"size"`
}

func (o LanguageEdge) Cursor() string {
	return o.field_cursor
}

func (o LanguageEdge) Node() Language {
	return o.field_node
}

func (o LanguageEdge) Size() int32 {
	return o.field_size
}

func (o *LanguageEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string   `json:"cursor"`
		Field_node   Language `json:"node"`
		Field_size   int32    `json:"size"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
		Field_size:   o.field_size,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	o.field_size = v.Field_size
	return nil
}

// A special type of user which takes actions on behalf of GitHub Apps.
type Bot struct {
	// A URL pointing to the GitHub App's public avatar.
	field_avatarUrl URI `json:"avatarUrl"`

	// Identifies the primary key from the database.
	field_databaseId *int32 `json:"databaseId"`

	field_id ID `json:"id"`

	// The username of the actor.
	field_login string `json:"login"`

	// The HTTP path for this bot
	field_resourcePath URI `json:"resourcePath"`

	// The HTTP URL for this bot
	field_url URI `json:"url"`
}

func (o Bot) AvatarUrl() URI {
	return o.field_avatarUrl
}

func (o Bot) DatabaseId() *int32 {
	return o.field_databaseId
}

func (o Bot) Id() ID {
	return o.field_id
}

func (o Bot) Login() string {
	return o.field_login
}

func (o Bot) ResourcePath() URI {
	return o.field_resourcePath
}

func (o Bot) Url() URI {
	return o.field_url
}

func (o *Bot) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_avatarUrl    URI    `json:"avatarUrl"`
		Field_databaseId   *int32 `json:"databaseId"`
		Field_id           ID     `json:"id"`
		Field_login        string `json:"login"`
		Field_resourcePath URI    `json:"resourcePath"`
		Field_url          URI    `json:"url"`
	}{
		Field_avatarUrl:    o.field_avatarUrl,
		Field_databaseId:   o.field_databaseId,
		Field_id:           o.field_id,
		Field_login:        o.field_login,
		Field_resourcePath: o.field_resourcePath,
		Field_url:          o.field_url,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_avatarUrl = v.Field_avatarUrl
	o.field_databaseId = v.Field_databaseId
	o.field_id = v.Field_id
	o.field_login = v.Field_login
	o.field_resourcePath = v.Field_resourcePath
	o.field_url = v.Field_url
	return nil
}

// An Identity Provider configured to provision SAML and SCIM identities for Organizations
type OrganizationIdentityProvider struct {
	// The digest algorithm used to sign SAML requests for the Identity Provider.
	field_digestMethod *URI `json:"digestMethod"`

	// External Identities provisioned by this Identity Provider
	field_externalIdentities ExternalIdentityConnection `json:"externalIdentities"`

	field_id ID `json:"id"`

	// The x509 certificate used by the Identity Provder to sign assertions and responses.
	field_idpCertificate *X509Certificate `json:"idpCertificate"`

	// The Issuer Entity ID for the SAML Identity Provider
	field_issuer *string `json:"issuer"`

	// Organization this Identity Provider belongs to
	field_organization *Organization `json:"organization"`

	// The signature algorithm used to sign SAML requests for the Identity Provider.
	field_signatureMethod *URI `json:"signatureMethod"`

	// The URL endpoint for the Identity Provider's SAML SSO.
	field_ssoUrl *URI `json:"ssoUrl"`
}

func (o OrganizationIdentityProvider) DigestMethod() *URI {
	return o.field_digestMethod
}

func (o OrganizationIdentityProvider) ExternalIdentities() ExternalIdentityConnection {
	return o.field_externalIdentities
}

func (o OrganizationIdentityProvider) Id() ID {
	return o.field_id
}

func (o OrganizationIdentityProvider) IdpCertificate() *X509Certificate {
	return o.field_idpCertificate
}

func (o OrganizationIdentityProvider) Issuer() *string {
	return o.field_issuer
}

func (o OrganizationIdentityProvider) Organization() *Organization {
	return o.field_organization
}

func (o OrganizationIdentityProvider) SignatureMethod() *URI {
	return o.field_signatureMethod
}

func (o OrganizationIdentityProvider) SsoUrl() *URI {
	return o.field_ssoUrl
}

func (o *OrganizationIdentityProvider) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_digestMethod       *URI                       `json:"digestMethod"`
		Field_externalIdentities ExternalIdentityConnection `json:"externalIdentities"`
		Field_id                 ID                         `json:"id"`
		Field_idpCertificate     *X509Certificate           `json:"idpCertificate"`
		Field_issuer             *string                    `json:"issuer"`
		Field_organization       *Organization              `json:"organization"`
		Field_signatureMethod    *URI                       `json:"signatureMethod"`
		Field_ssoUrl             *URI                       `json:"ssoUrl"`
	}{
		Field_digestMethod:       o.field_digestMethod,
		Field_externalIdentities: o.field_externalIdentities,
		Field_id:                 o.field_id,
		Field_idpCertificate:     o.field_idpCertificate,
		Field_issuer:             o.field_issuer,
		Field_organization:       o.field_organization,
		Field_signatureMethod:    o.field_signatureMethod,
		Field_ssoUrl:             o.field_ssoUrl,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_digestMethod = v.Field_digestMethod
	o.field_externalIdentities = v.Field_externalIdentities
	o.field_id = v.Field_id
	o.field_idpCertificate = v.Field_idpCertificate
	o.field_issuer = v.Field_issuer
	o.field_organization = v.Field_organization
	o.field_signatureMethod = v.Field_signatureMethod
	o.field_ssoUrl = v.Field_ssoUrl
	return nil
}

// The connection type for ExternalIdentity.
type ExternalIdentityConnection struct {
	// A list of edges.
	field_edges []*ExternalIdentityEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*ExternalIdentity `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o ExternalIdentityConnection) Edges() []*ExternalIdentityEdge {
	return o.field_edges
}

func (o ExternalIdentityConnection) Nodes() []*ExternalIdentity {
	return o.field_nodes
}

func (o ExternalIdentityConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o ExternalIdentityConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *ExternalIdentityConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*ExternalIdentityEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*ExternalIdentity, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type ExternalIdentityEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *ExternalIdentity `json:"node"`
}

func (o ExternalIdentityEdge) Cursor() string {
	return o.field_cursor
}

func (o ExternalIdentityEdge) Node() *ExternalIdentity {
	return o.field_node
}

func (o *ExternalIdentityEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string            `json:"cursor"`
		Field_node   *ExternalIdentity `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// An external identity provisioned by SAML SSO or SCIM.
type ExternalIdentity struct {
	// The GUID for this identity
	field_guid string `json:"guid"`

	field_id ID `json:"id"`

	// Organization invitation for this SCIM-provisioned external identity
	field_organizationInvitation *OrganizationInvitation `json:"organizationInvitation"`

	// SAML Identity attributes
	field_samlIdentity *ExternalIdentitySamlAttributes `json:"samlIdentity"`

	// SCIM Identity attributes
	field_scimIdentity *ExternalIdentityScimAttributes `json:"scimIdentity"`

	// User linked to this external identity
	field_user *User `json:"user"`
}

func (o ExternalIdentity) Guid() string {
	return o.field_guid
}

func (o ExternalIdentity) Id() ID {
	return o.field_id
}

func (o ExternalIdentity) OrganizationInvitation() *OrganizationInvitation {
	return o.field_organizationInvitation
}

func (o ExternalIdentity) SamlIdentity() *ExternalIdentitySamlAttributes {
	return o.field_samlIdentity
}

func (o ExternalIdentity) ScimIdentity() *ExternalIdentityScimAttributes {
	return o.field_scimIdentity
}

func (o ExternalIdentity) User() *User {
	return o.field_user
}

func (o *ExternalIdentity) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_guid                   string                          `json:"guid"`
		Field_id                     ID                              `json:"id"`
		Field_organizationInvitation *OrganizationInvitation         `json:"organizationInvitation"`
		Field_samlIdentity           *ExternalIdentitySamlAttributes `json:"samlIdentity"`
		Field_scimIdentity           *ExternalIdentityScimAttributes `json:"scimIdentity"`
		Field_user                   *User                           `json:"user"`
	}{
		Field_guid:                   o.field_guid,
		Field_id:                     o.field_id,
		Field_organizationInvitation: o.field_organizationInvitation,
		Field_samlIdentity:           o.field_samlIdentity,
		Field_scimIdentity:           o.field_scimIdentity,
		Field_user:                   o.field_user,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_guid = v.Field_guid
	o.field_id = v.Field_id
	o.field_organizationInvitation = v.Field_organizationInvitation
	o.field_samlIdentity = v.Field_samlIdentity
	o.field_scimIdentity = v.Field_scimIdentity
	o.field_user = v.Field_user
	return nil
}

// SAML attributes for the External Identity
type ExternalIdentitySamlAttributes struct {
	// The NameID of the SAML identity
	field_nameId *string `json:"nameId"`
}

func (o ExternalIdentitySamlAttributes) NameId() *string {
	return o.field_nameId
}

func (o *ExternalIdentitySamlAttributes) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_nameId *string `json:"nameId"`
	}{
		Field_nameId: o.field_nameId,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_nameId = v.Field_nameId
	return nil
}

// SCIM attributes for the External Identity
type ExternalIdentityScimAttributes struct {
	// The userName of the SCIM identity
	field_username *string `json:"username"`
}

func (o ExternalIdentityScimAttributes) Username() *string {
	return o.field_username
}

func (o *ExternalIdentityScimAttributes) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_username *string `json:"username"`
	}{
		Field_username: o.field_username,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_username = v.Field_username
	return nil
}

// The connection type for PullRequestTimelineItem.
type PullRequestTimelineConnection struct {
	// A list of edges.
	field_edges []*PullRequestTimelineItemEdge `json:"edges"`

	// A list of nodes.
	field_nodes []PullRequestTimelineItem `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o PullRequestTimelineConnection) Edges() []*PullRequestTimelineItemEdge {
	return o.field_edges
}

func (o PullRequestTimelineConnection) Nodes() []PullRequestTimelineItem {
	return o.field_nodes
}

func (o PullRequestTimelineConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o PullRequestTimelineConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *PullRequestTimelineConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*PullRequestTimelineItemEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]PullRequestTimelineItem, len(l))
			for li, lv := range l {
				o.field_nodes[li], err = PullRequestTimelineItem_UnmarshalJSON(lv)
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type PullRequestTimelineItemEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node PullRequestTimelineItem `json:"node"`
}

func (o PullRequestTimelineItemEdge) Cursor() string {
	return o.field_cursor
}

func (o PullRequestTimelineItemEdge) Node() PullRequestTimelineItem {
	return o.field_node
}

func (o *PullRequestTimelineItemEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string          `json:"cursor"`
		Field_node   json.RawMessage `json:"node"`
	}{
		Field_cursor: o.field_cursor,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	if len(v.Field_node) > 0 {
		o.field_node, err = PullRequestTimelineItem_UnmarshalJSON(v.Field_node)
		if err != nil {
			return err
		}
	}
	return nil
}

// A threaded list of comments for a given pull request.
type PullRequestReviewThread struct {
	// A list of pull request comments associated with the thread.
	field_comments PullRequestReviewCommentConnection `json:"comments"`

	field_id ID `json:"id"`

	// Identifies the pull request associated with this thread.
	field_pullRequest PullRequest `json:"pullRequest"`
}

func (o PullRequestReviewThread) Comments() PullRequestReviewCommentConnection {
	return o.field_comments
}

func (o PullRequestReviewThread) Id() ID {
	return o.field_id
}

func (o PullRequestReviewThread) PullRequest() PullRequest {
	return o.field_pullRequest
}

func (o *PullRequestReviewThread) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_comments    PullRequestReviewCommentConnection `json:"comments"`
		Field_id          ID                                 `json:"id"`
		Field_pullRequest PullRequest                        `json:"pullRequest"`
	}{
		Field_comments:    o.field_comments,
		Field_id:          o.field_id,
		Field_pullRequest: o.field_pullRequest,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_comments = v.Field_comments
	o.field_id = v.Field_id
	o.field_pullRequest = v.Field_pullRequest
	return nil
}

// Represents a 'closed' event on any `Closable`.
type ClosedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Object that was closed.
	field_closable Closable `json:"closable"`

	// Identifies the commit associated with the 'closed' event.
	field_commit *Commit `json:"commit"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`
}

func (o ClosedEvent) Actor() Actor {
	return o.field_actor
}

func (o ClosedEvent) Closable() Closable {
	return o.field_closable
}

func (o ClosedEvent) Commit() *Commit {
	return o.field_commit
}

func (o ClosedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o ClosedEvent) Id() ID {
	return o.field_id
}

func (o *ClosedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor     json.RawMessage `json:"actor"`
		Field_closable  Closable        `json:"closable"`
		Field_commit    *Commit         `json:"commit"`
		Field_createdAt DateTime        `json:"createdAt"`
		Field_id        ID              `json:"id"`
	}{
		Field_closable:  o.field_closable,
		Field_commit:    o.field_commit,
		Field_createdAt: o.field_createdAt,
		Field_id:        o.field_id,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_closable = v.Field_closable
	o.field_commit = v.Field_commit
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	return nil
}

// Represents a 'reopened' event on any `Closable`.
type ReopenedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Object that was reopened.
	field_closable Closable `json:"closable"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`
}

func (o ReopenedEvent) Actor() Actor {
	return o.field_actor
}

func (o ReopenedEvent) Closable() Closable {
	return o.field_closable
}

func (o ReopenedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o ReopenedEvent) Id() ID {
	return o.field_id
}

func (o *ReopenedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor     json.RawMessage `json:"actor"`
		Field_closable  Closable        `json:"closable"`
		Field_createdAt DateTime        `json:"createdAt"`
		Field_id        ID              `json:"id"`
	}{
		Field_closable:  o.field_closable,
		Field_createdAt: o.field_createdAt,
		Field_id:        o.field_id,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_closable = v.Field_closable
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	return nil
}

// Represents a 'subscribed' event on a given `Subscribable`.
type SubscribedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`

	// Object referenced by event.
	field_subscribable Subscribable `json:"subscribable"`
}

func (o SubscribedEvent) Actor() Actor {
	return o.field_actor
}

func (o SubscribedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o SubscribedEvent) Id() ID {
	return o.field_id
}

func (o SubscribedEvent) Subscribable() Subscribable {
	return o.field_subscribable
}

func (o *SubscribedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor        json.RawMessage `json:"actor"`
		Field_createdAt    DateTime        `json:"createdAt"`
		Field_id           ID              `json:"id"`
		Field_subscribable Subscribable    `json:"subscribable"`
	}{
		Field_createdAt:    o.field_createdAt,
		Field_id:           o.field_id,
		Field_subscribable: o.field_subscribable,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	o.field_subscribable = v.Field_subscribable
	return nil
}

// Represents an 'unsubscribed' event on a given `Subscribable`.
type UnsubscribedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`

	// Object referenced by event.
	field_subscribable Subscribable `json:"subscribable"`
}

func (o UnsubscribedEvent) Actor() Actor {
	return o.field_actor
}

func (o UnsubscribedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o UnsubscribedEvent) Id() ID {
	return o.field_id
}

func (o UnsubscribedEvent) Subscribable() Subscribable {
	return o.field_subscribable
}

func (o *UnsubscribedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor        json.RawMessage `json:"actor"`
		Field_createdAt    DateTime        `json:"createdAt"`
		Field_id           ID              `json:"id"`
		Field_subscribable Subscribable    `json:"subscribable"`
	}{
		Field_createdAt:    o.field_createdAt,
		Field_id:           o.field_id,
		Field_subscribable: o.field_subscribable,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	o.field_subscribable = v.Field_subscribable
	return nil
}

// Represents a 'merged' event on a given pull request.
type MergedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the commit associated with the `merge` event.
	field_commit *Commit `json:"commit"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`

	// Identifies the Ref associated with the `merge` event.
	field_mergeRef *Ref `json:"mergeRef"`

	// Identifies the name of the Ref associated with the `merge` event.
	field_mergeRefName string `json:"mergeRefName"`

	// PullRequest referenced by event.
	field_pullRequest PullRequest `json:"pullRequest"`

	// The HTTP path for this merged event.
	field_resourcePath URI `json:"resourcePath"`

	// The HTTP URL for this merged event.
	field_url URI `json:"url"`
}

func (o MergedEvent) Actor() Actor {
	return o.field_actor
}

func (o MergedEvent) Commit() *Commit {
	return o.field_commit
}

func (o MergedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o MergedEvent) Id() ID {
	return o.field_id
}

func (o MergedEvent) MergeRef() *Ref {
	return o.field_mergeRef
}

func (o MergedEvent) MergeRefName() string {
	return o.field_mergeRefName
}

func (o MergedEvent) PullRequest() PullRequest {
	return o.field_pullRequest
}

func (o MergedEvent) ResourcePath() URI {
	return o.field_resourcePath
}

func (o MergedEvent) Url() URI {
	return o.field_url
}

func (o *MergedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor        json.RawMessage `json:"actor"`
		Field_commit       *Commit         `json:"commit"`
		Field_createdAt    DateTime        `json:"createdAt"`
		Field_id           ID              `json:"id"`
		Field_mergeRef     *Ref            `json:"mergeRef"`
		Field_mergeRefName string          `json:"mergeRefName"`
		Field_pullRequest  PullRequest     `json:"pullRequest"`
		Field_resourcePath URI             `json:"resourcePath"`
		Field_url          URI             `json:"url"`
	}{
		Field_commit:       o.field_commit,
		Field_createdAt:    o.field_createdAt,
		Field_id:           o.field_id,
		Field_mergeRef:     o.field_mergeRef,
		Field_mergeRefName: o.field_mergeRefName,
		Field_pullRequest:  o.field_pullRequest,
		Field_resourcePath: o.field_resourcePath,
		Field_url:          o.field_url,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_commit = v.Field_commit
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	o.field_mergeRef = v.Field_mergeRef
	o.field_mergeRefName = v.Field_mergeRefName
	o.field_pullRequest = v.Field_pullRequest
	o.field_resourcePath = v.Field_resourcePath
	o.field_url = v.Field_url
	return nil
}

// Represents a 'referenced' event on a given `ReferencedSubject`.
type ReferencedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the commit associated with the 'referenced' event.
	field_commit *Commit `json:"commit"`

	// Identifies the repository associated with the 'referenced' event.
	field_commitRepository Repository `json:"commitRepository"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`

	// Reference originated in a different repository.
	field_isCrossReference bool `json:"isCrossReference"`

	// Checks if the commit message itself references the subject. Can be false in the case of a commit comment reference.
	field_isDirectReference bool `json:"isDirectReference"`

	// Object referenced by event.
	field_subject ReferencedSubject `json:"subject"`
}

func (o ReferencedEvent) Actor() Actor {
	return o.field_actor
}

func (o ReferencedEvent) Commit() *Commit {
	return o.field_commit
}

func (o ReferencedEvent) CommitRepository() Repository {
	return o.field_commitRepository
}

func (o ReferencedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o ReferencedEvent) Id() ID {
	return o.field_id
}

func (o ReferencedEvent) IsCrossReference() bool {
	return o.field_isCrossReference
}

func (o ReferencedEvent) IsDirectReference() bool {
	return o.field_isDirectReference
}

func (o ReferencedEvent) Subject() ReferencedSubject {
	return o.field_subject
}

func (o *ReferencedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor             json.RawMessage   `json:"actor"`
		Field_commit            *Commit           `json:"commit"`
		Field_commitRepository  Repository        `json:"commitRepository"`
		Field_createdAt         DateTime          `json:"createdAt"`
		Field_id                ID                `json:"id"`
		Field_isCrossReference  bool              `json:"isCrossReference"`
		Field_isDirectReference bool              `json:"isDirectReference"`
		Field_subject           ReferencedSubject `json:"subject"`
	}{
		Field_commit:            o.field_commit,
		Field_commitRepository:  o.field_commitRepository,
		Field_createdAt:         o.field_createdAt,
		Field_id:                o.field_id,
		Field_isCrossReference:  o.field_isCrossReference,
		Field_isDirectReference: o.field_isDirectReference,
		Field_subject:           o.field_subject,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_commit = v.Field_commit
	o.field_commitRepository = v.Field_commitRepository
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	o.field_isCrossReference = v.Field_isCrossReference
	o.field_isDirectReference = v.Field_isDirectReference
	o.field_subject = v.Field_subject
	return nil
}

// Represents an 'assigned' event on any assignable object.
type AssignedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the assignable associated with the event.
	field_assignable Assignable `json:"assignable"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`

	// Identifies the user who was assigned.
	field_user *User `json:"user"`
}

func (o AssignedEvent) Actor() Actor {
	return o.field_actor
}

func (o AssignedEvent) Assignable() Assignable {
	return o.field_assignable
}

func (o AssignedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o AssignedEvent) Id() ID {
	return o.field_id
}

func (o AssignedEvent) User() *User {
	return o.field_user
}

func (o *AssignedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor      json.RawMessage `json:"actor"`
		Field_assignable Assignable      `json:"assignable"`
		Field_createdAt  DateTime        `json:"createdAt"`
		Field_id         ID              `json:"id"`
		Field_user       *User           `json:"user"`
	}{
		Field_assignable: o.field_assignable,
		Field_createdAt:  o.field_createdAt,
		Field_id:         o.field_id,
		Field_user:       o.field_user,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_assignable = v.Field_assignable
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	o.field_user = v.Field_user
	return nil
}

// Represents an 'unassigned' event on any assignable object.
type UnassignedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the assignable associated with the event.
	field_assignable Assignable `json:"assignable"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`

	// Identifies the subject (user) who was unassigned.
	field_user *User `json:"user"`
}

func (o UnassignedEvent) Actor() Actor {
	return o.field_actor
}

func (o UnassignedEvent) Assignable() Assignable {
	return o.field_assignable
}

func (o UnassignedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o UnassignedEvent) Id() ID {
	return o.field_id
}

func (o UnassignedEvent) User() *User {
	return o.field_user
}

func (o *UnassignedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor      json.RawMessage `json:"actor"`
		Field_assignable Assignable      `json:"assignable"`
		Field_createdAt  DateTime        `json:"createdAt"`
		Field_id         ID              `json:"id"`
		Field_user       *User           `json:"user"`
	}{
		Field_assignable: o.field_assignable,
		Field_createdAt:  o.field_createdAt,
		Field_id:         o.field_id,
		Field_user:       o.field_user,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_assignable = v.Field_assignable
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	o.field_user = v.Field_user
	return nil
}

// Represents a 'labeled' event on a given issue or pull request.
type LabeledEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`

	// Identifies the label associated with the 'labeled' event.
	field_label Label `json:"label"`

	// Identifies the `Labelable` associated with the event.
	field_labelable Labelable `json:"labelable"`
}

func (o LabeledEvent) Actor() Actor {
	return o.field_actor
}

func (o LabeledEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o LabeledEvent) Id() ID {
	return o.field_id
}

func (o LabeledEvent) Label() Label {
	return o.field_label
}

func (o LabeledEvent) Labelable() Labelable {
	return o.field_labelable
}

func (o *LabeledEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor     json.RawMessage `json:"actor"`
		Field_createdAt DateTime        `json:"createdAt"`
		Field_id        ID              `json:"id"`
		Field_label     Label           `json:"label"`
		Field_labelable Labelable       `json:"labelable"`
	}{
		Field_createdAt: o.field_createdAt,
		Field_id:        o.field_id,
		Field_label:     o.field_label,
		Field_labelable: o.field_labelable,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	o.field_label = v.Field_label
	o.field_labelable = v.Field_labelable
	return nil
}

// Represents an 'unlabeled' event on a given issue or pull request.
type UnlabeledEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`

	// Identifies the label associated with the 'unlabeled' event.
	field_label Label `json:"label"`

	// Identifies the `Labelable` associated with the event.
	field_labelable Labelable `json:"labelable"`
}

func (o UnlabeledEvent) Actor() Actor {
	return o.field_actor
}

func (o UnlabeledEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o UnlabeledEvent) Id() ID {
	return o.field_id
}

func (o UnlabeledEvent) Label() Label {
	return o.field_label
}

func (o UnlabeledEvent) Labelable() Labelable {
	return o.field_labelable
}

func (o *UnlabeledEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor     json.RawMessage `json:"actor"`
		Field_createdAt DateTime        `json:"createdAt"`
		Field_id        ID              `json:"id"`
		Field_label     Label           `json:"label"`
		Field_labelable Labelable       `json:"labelable"`
	}{
		Field_createdAt: o.field_createdAt,
		Field_id:        o.field_id,
		Field_label:     o.field_label,
		Field_labelable: o.field_labelable,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	o.field_label = v.Field_label
	o.field_labelable = v.Field_labelable
	return nil
}

// Represents a 'milestoned' event on a given issue or pull request.
type MilestonedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`

	// Identifies the milestone title associated with the 'milestoned' event.
	field_milestoneTitle string `json:"milestoneTitle"`

	// Object referenced by event.
	field_subject MilestoneItem `json:"subject"`
}

func (o MilestonedEvent) Actor() Actor {
	return o.field_actor
}

func (o MilestonedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o MilestonedEvent) Id() ID {
	return o.field_id
}

func (o MilestonedEvent) MilestoneTitle() string {
	return o.field_milestoneTitle
}

func (o MilestonedEvent) Subject() MilestoneItem {
	return o.field_subject
}

func (o *MilestonedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor          json.RawMessage `json:"actor"`
		Field_createdAt      DateTime        `json:"createdAt"`
		Field_id             ID              `json:"id"`
		Field_milestoneTitle string          `json:"milestoneTitle"`
		Field_subject        MilestoneItem   `json:"subject"`
	}{
		Field_createdAt:      o.field_createdAt,
		Field_id:             o.field_id,
		Field_milestoneTitle: o.field_milestoneTitle,
		Field_subject:        o.field_subject,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	o.field_milestoneTitle = v.Field_milestoneTitle
	o.field_subject = v.Field_subject
	return nil
}

// Represents a 'demilestoned' event on a given issue or pull request.
type DemilestonedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`

	// Identifies the milestone title associated with the 'demilestoned' event.
	field_milestoneTitle string `json:"milestoneTitle"`

	// Object referenced by event.
	field_subject MilestoneItem `json:"subject"`
}

func (o DemilestonedEvent) Actor() Actor {
	return o.field_actor
}

func (o DemilestonedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o DemilestonedEvent) Id() ID {
	return o.field_id
}

func (o DemilestonedEvent) MilestoneTitle() string {
	return o.field_milestoneTitle
}

func (o DemilestonedEvent) Subject() MilestoneItem {
	return o.field_subject
}

func (o *DemilestonedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor          json.RawMessage `json:"actor"`
		Field_createdAt      DateTime        `json:"createdAt"`
		Field_id             ID              `json:"id"`
		Field_milestoneTitle string          `json:"milestoneTitle"`
		Field_subject        MilestoneItem   `json:"subject"`
	}{
		Field_createdAt:      o.field_createdAt,
		Field_id:             o.field_id,
		Field_milestoneTitle: o.field_milestoneTitle,
		Field_subject:        o.field_subject,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	o.field_milestoneTitle = v.Field_milestoneTitle
	o.field_subject = v.Field_subject
	return nil
}

// Represents a 'renamed' event on a given issue or pull request
type RenamedTitleEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// Identifies the current title of the issue or pull request.
	field_currentTitle string `json:"currentTitle"`

	field_id ID `json:"id"`

	// Identifies the previous title of the issue or pull request.
	field_previousTitle string `json:"previousTitle"`

	// Subject that was renamed.
	field_subject RenamedTitleSubject `json:"subject"`
}

func (o RenamedTitleEvent) Actor() Actor {
	return o.field_actor
}

func (o RenamedTitleEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o RenamedTitleEvent) CurrentTitle() string {
	return o.field_currentTitle
}

func (o RenamedTitleEvent) Id() ID {
	return o.field_id
}

func (o RenamedTitleEvent) PreviousTitle() string {
	return o.field_previousTitle
}

func (o RenamedTitleEvent) Subject() RenamedTitleSubject {
	return o.field_subject
}

func (o *RenamedTitleEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor         json.RawMessage     `json:"actor"`
		Field_createdAt     DateTime            `json:"createdAt"`
		Field_currentTitle  string              `json:"currentTitle"`
		Field_id            ID                  `json:"id"`
		Field_previousTitle string              `json:"previousTitle"`
		Field_subject       RenamedTitleSubject `json:"subject"`
	}{
		Field_createdAt:     o.field_createdAt,
		Field_currentTitle:  o.field_currentTitle,
		Field_id:            o.field_id,
		Field_previousTitle: o.field_previousTitle,
		Field_subject:       o.field_subject,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_createdAt = v.Field_createdAt
	o.field_currentTitle = v.Field_currentTitle
	o.field_id = v.Field_id
	o.field_previousTitle = v.Field_previousTitle
	o.field_subject = v.Field_subject
	return nil
}

// Represents a 'locked' event on a given issue or pull request.
type LockedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`

	// Object that was locked.
	field_lockable Lockable `json:"lockable"`
}

func (o LockedEvent) Actor() Actor {
	return o.field_actor
}

func (o LockedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o LockedEvent) Id() ID {
	return o.field_id
}

func (o LockedEvent) Lockable() Lockable {
	return o.field_lockable
}

func (o *LockedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor     json.RawMessage `json:"actor"`
		Field_createdAt DateTime        `json:"createdAt"`
		Field_id        ID              `json:"id"`
		Field_lockable  Lockable        `json:"lockable"`
	}{
		Field_createdAt: o.field_createdAt,
		Field_id:        o.field_id,
		Field_lockable:  o.field_lockable,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	o.field_lockable = v.Field_lockable
	return nil
}

// Represents an 'unlocked' event on a given issue or pull request.
type UnlockedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`

	// Object that was unlocked.
	field_lockable Lockable `json:"lockable"`
}

func (o UnlockedEvent) Actor() Actor {
	return o.field_actor
}

func (o UnlockedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o UnlockedEvent) Id() ID {
	return o.field_id
}

func (o UnlockedEvent) Lockable() Lockable {
	return o.field_lockable
}

func (o *UnlockedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor     json.RawMessage `json:"actor"`
		Field_createdAt DateTime        `json:"createdAt"`
		Field_id        ID              `json:"id"`
		Field_lockable  Lockable        `json:"lockable"`
	}{
		Field_createdAt: o.field_createdAt,
		Field_id:        o.field_id,
		Field_lockable:  o.field_lockable,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	o.field_lockable = v.Field_lockable
	return nil
}

// Represents a 'deployed' event on a given pull request.
type DeployedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// Identifies the primary key from the database.
	field_databaseId *int32 `json:"databaseId"`

	// The deployment associated with the 'deployed' event.
	field_deployment Deployment `json:"deployment"`

	field_id ID `json:"id"`

	// PullRequest referenced by event.
	field_pullRequest PullRequest `json:"pullRequest"`

	// The ref associated with the 'deployed' event.
	field_ref *Ref `json:"ref"`
}

func (o DeployedEvent) Actor() Actor {
	return o.field_actor
}

func (o DeployedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o DeployedEvent) DatabaseId() *int32 {
	return o.field_databaseId
}

func (o DeployedEvent) Deployment() Deployment {
	return o.field_deployment
}

func (o DeployedEvent) Id() ID {
	return o.field_id
}

func (o DeployedEvent) PullRequest() PullRequest {
	return o.field_pullRequest
}

func (o DeployedEvent) Ref() *Ref {
	return o.field_ref
}

func (o *DeployedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor       json.RawMessage `json:"actor"`
		Field_createdAt   DateTime        `json:"createdAt"`
		Field_databaseId  *int32          `json:"databaseId"`
		Field_deployment  Deployment      `json:"deployment"`
		Field_id          ID              `json:"id"`
		Field_pullRequest PullRequest     `json:"pullRequest"`
		Field_ref         *Ref            `json:"ref"`
	}{
		Field_createdAt:   o.field_createdAt,
		Field_databaseId:  o.field_databaseId,
		Field_deployment:  o.field_deployment,
		Field_id:          o.field_id,
		Field_pullRequest: o.field_pullRequest,
		Field_ref:         o.field_ref,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_createdAt = v.Field_createdAt
	o.field_databaseId = v.Field_databaseId
	o.field_deployment = v.Field_deployment
	o.field_id = v.Field_id
	o.field_pullRequest = v.Field_pullRequest
	o.field_ref = v.Field_ref
	return nil
}

// Represents triggered deployment instance.
type Deployment struct {
	// Identifies the commit sha of the deployment.
	field_commit *Commit `json:"commit"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// Identifies the actor who triggered the deployment.
	field_creator Actor `json:"creator"`

	// The environment to which this deployment was made.
	field_environment *string `json:"environment"`

	field_id ID `json:"id"`

	// The latest status of this deployment.
	field_latestStatus *DeploymentStatus `json:"latestStatus"`

	// Identifies the repository associated with the deployment.
	field_repository Repository `json:"repository"`

	// The current state of the deployment.
	field_state *DeploymentState `json:"state"`

	// A list of statuses associated with the deployment.
	field_statuses *DeploymentStatusConnection `json:"statuses"`
}

func (o Deployment) Commit() *Commit {
	return o.field_commit
}

func (o Deployment) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o Deployment) Creator() Actor {
	return o.field_creator
}

func (o Deployment) Environment() *string {
	return o.field_environment
}

func (o Deployment) Id() ID {
	return o.field_id
}

func (o Deployment) LatestStatus() *DeploymentStatus {
	return o.field_latestStatus
}

func (o Deployment) Repository() Repository {
	return o.field_repository
}

func (o Deployment) State() *DeploymentState {
	return o.field_state
}

func (o Deployment) Statuses() *DeploymentStatusConnection {
	return o.field_statuses
}

func (o *Deployment) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_commit       *Commit                     `json:"commit"`
		Field_createdAt    DateTime                    `json:"createdAt"`
		Field_creator      json.RawMessage             `json:"creator"`
		Field_environment  *string                     `json:"environment"`
		Field_id           ID                          `json:"id"`
		Field_latestStatus *DeploymentStatus           `json:"latestStatus"`
		Field_repository   Repository                  `json:"repository"`
		Field_state        *DeploymentState            `json:"state"`
		Field_statuses     *DeploymentStatusConnection `json:"statuses"`
	}{
		Field_commit:       o.field_commit,
		Field_createdAt:    o.field_createdAt,
		Field_environment:  o.field_environment,
		Field_id:           o.field_id,
		Field_latestStatus: o.field_latestStatus,
		Field_repository:   o.field_repository,
		Field_state:        o.field_state,
		Field_statuses:     o.field_statuses,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_commit = v.Field_commit
	o.field_createdAt = v.Field_createdAt
	if len(v.Field_creator) > 0 {
		o.field_creator, err = Actor_UnmarshalJSON(v.Field_creator)
		if err != nil {
			return err
		}
	}
	o.field_environment = v.Field_environment
	o.field_id = v.Field_id
	o.field_latestStatus = v.Field_latestStatus
	o.field_repository = v.Field_repository
	o.field_state = v.Field_state
	o.field_statuses = v.Field_statuses
	return nil
}

// The connection type for DeploymentStatus.
type DeploymentStatusConnection struct {
	// A list of edges.
	field_edges []*DeploymentStatusEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*DeploymentStatus `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o DeploymentStatusConnection) Edges() []*DeploymentStatusEdge {
	return o.field_edges
}

func (o DeploymentStatusConnection) Nodes() []*DeploymentStatus {
	return o.field_nodes
}

func (o DeploymentStatusConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o DeploymentStatusConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *DeploymentStatusConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*DeploymentStatusEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*DeploymentStatus, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type DeploymentStatusEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *DeploymentStatus `json:"node"`
}

func (o DeploymentStatusEdge) Cursor() string {
	return o.field_cursor
}

func (o DeploymentStatusEdge) Node() *DeploymentStatus {
	return o.field_node
}

func (o *DeploymentStatusEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string            `json:"cursor"`
		Field_node   *DeploymentStatus `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// Describes the status of a given deployment attempt.
type DeploymentStatus struct {
	// Identifies the actor who triggered the deployment.
	field_creator Actor `json:"creator"`

	// Identifies the deployment associated with status.
	field_deployment Deployment `json:"deployment"`

	// Identifies the description of the deployment.
	field_description *string `json:"description"`

	// Identifies the environment URL of the deployment.
	field_environmentUrl *URI `json:"environmentUrl"`

	field_id ID `json:"id"`

	// Identifies the log URL of the deployment.
	field_logUrl *URI `json:"logUrl"`

	// Identifies the current state of the deployment.
	field_state DeploymentStatusState `json:"state"`
}

func (o DeploymentStatus) Creator() Actor {
	return o.field_creator
}

func (o DeploymentStatus) Deployment() Deployment {
	return o.field_deployment
}

func (o DeploymentStatus) Description() *string {
	return o.field_description
}

func (o DeploymentStatus) EnvironmentUrl() *URI {
	return o.field_environmentUrl
}

func (o DeploymentStatus) Id() ID {
	return o.field_id
}

func (o DeploymentStatus) LogUrl() *URI {
	return o.field_logUrl
}

func (o DeploymentStatus) State() DeploymentStatusState {
	return o.field_state
}

func (o *DeploymentStatus) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_creator        json.RawMessage       `json:"creator"`
		Field_deployment     Deployment            `json:"deployment"`
		Field_description    *string               `json:"description"`
		Field_environmentUrl *URI                  `json:"environmentUrl"`
		Field_id             ID                    `json:"id"`
		Field_logUrl         *URI                  `json:"logUrl"`
		Field_state          DeploymentStatusState `json:"state"`
	}{
		Field_deployment:     o.field_deployment,
		Field_description:    o.field_description,
		Field_environmentUrl: o.field_environmentUrl,
		Field_id:             o.field_id,
		Field_logUrl:         o.field_logUrl,
		Field_state:          o.field_state,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_creator) > 0 {
		o.field_creator, err = Actor_UnmarshalJSON(v.Field_creator)
		if err != nil {
			return err
		}
	}
	o.field_deployment = v.Field_deployment
	o.field_description = v.Field_description
	o.field_environmentUrl = v.Field_environmentUrl
	o.field_id = v.Field_id
	o.field_logUrl = v.Field_logUrl
	o.field_state = v.Field_state
	return nil
}

// Represents a 'head_ref_deleted' event on a given pull request.
type HeadRefDeletedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// Identifies the Ref associated with the `head_ref_deleted` event.
	field_headRef *Ref `json:"headRef"`

	// Identifies the name of the Ref associated with the `head_ref_deleted` event.
	field_headRefName string `json:"headRefName"`

	field_id ID `json:"id"`

	// PullRequest referenced by event.
	field_pullRequest PullRequest `json:"pullRequest"`
}

func (o HeadRefDeletedEvent) Actor() Actor {
	return o.field_actor
}

func (o HeadRefDeletedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o HeadRefDeletedEvent) HeadRef() *Ref {
	return o.field_headRef
}

func (o HeadRefDeletedEvent) HeadRefName() string {
	return o.field_headRefName
}

func (o HeadRefDeletedEvent) Id() ID {
	return o.field_id
}

func (o HeadRefDeletedEvent) PullRequest() PullRequest {
	return o.field_pullRequest
}

func (o *HeadRefDeletedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor       json.RawMessage `json:"actor"`
		Field_createdAt   DateTime        `json:"createdAt"`
		Field_headRef     *Ref            `json:"headRef"`
		Field_headRefName string          `json:"headRefName"`
		Field_id          ID              `json:"id"`
		Field_pullRequest PullRequest     `json:"pullRequest"`
	}{
		Field_createdAt:   o.field_createdAt,
		Field_headRef:     o.field_headRef,
		Field_headRefName: o.field_headRefName,
		Field_id:          o.field_id,
		Field_pullRequest: o.field_pullRequest,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_createdAt = v.Field_createdAt
	o.field_headRef = v.Field_headRef
	o.field_headRefName = v.Field_headRefName
	o.field_id = v.Field_id
	o.field_pullRequest = v.Field_pullRequest
	return nil
}

// Represents a 'head_ref_restored' event on a given pull request.
type HeadRefRestoredEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`

	// PullRequest referenced by event.
	field_pullRequest PullRequest `json:"pullRequest"`
}

func (o HeadRefRestoredEvent) Actor() Actor {
	return o.field_actor
}

func (o HeadRefRestoredEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o HeadRefRestoredEvent) Id() ID {
	return o.field_id
}

func (o HeadRefRestoredEvent) PullRequest() PullRequest {
	return o.field_pullRequest
}

func (o *HeadRefRestoredEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor       json.RawMessage `json:"actor"`
		Field_createdAt   DateTime        `json:"createdAt"`
		Field_id          ID              `json:"id"`
		Field_pullRequest PullRequest     `json:"pullRequest"`
	}{
		Field_createdAt:   o.field_createdAt,
		Field_id:          o.field_id,
		Field_pullRequest: o.field_pullRequest,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	o.field_pullRequest = v.Field_pullRequest
	return nil
}

// Represents a 'head_ref_force_pushed' event on a given pull request.
type HeadRefForcePushedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the after commit SHA for the 'head_ref_force_pushed' event.
	field_afterCommit Commit `json:"afterCommit"`

	// Identifies the before commit SHA for the 'head_ref_force_pushed' event.
	field_beforeCommit Commit `json:"beforeCommit"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`

	// PullRequest referenced by event.
	field_pullRequest PullRequest `json:"pullRequest"`

	// Identifies the fully qualified ref name for the 'head_ref_force_pushed' event.
	field_ref *Ref `json:"ref"`
}

func (o HeadRefForcePushedEvent) Actor() Actor {
	return o.field_actor
}

func (o HeadRefForcePushedEvent) AfterCommit() Commit {
	return o.field_afterCommit
}

func (o HeadRefForcePushedEvent) BeforeCommit() Commit {
	return o.field_beforeCommit
}

func (o HeadRefForcePushedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o HeadRefForcePushedEvent) Id() ID {
	return o.field_id
}

func (o HeadRefForcePushedEvent) PullRequest() PullRequest {
	return o.field_pullRequest
}

func (o HeadRefForcePushedEvent) Ref() *Ref {
	return o.field_ref
}

func (o *HeadRefForcePushedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor        json.RawMessage `json:"actor"`
		Field_afterCommit  Commit          `json:"afterCommit"`
		Field_beforeCommit Commit          `json:"beforeCommit"`
		Field_createdAt    DateTime        `json:"createdAt"`
		Field_id           ID              `json:"id"`
		Field_pullRequest  PullRequest     `json:"pullRequest"`
		Field_ref          *Ref            `json:"ref"`
	}{
		Field_afterCommit:  o.field_afterCommit,
		Field_beforeCommit: o.field_beforeCommit,
		Field_createdAt:    o.field_createdAt,
		Field_id:           o.field_id,
		Field_pullRequest:  o.field_pullRequest,
		Field_ref:          o.field_ref,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_afterCommit = v.Field_afterCommit
	o.field_beforeCommit = v.Field_beforeCommit
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	o.field_pullRequest = v.Field_pullRequest
	o.field_ref = v.Field_ref
	return nil
}

// Represents a 'base_ref_force_pushed' event on a given pull request.
type BaseRefForcePushedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the after commit SHA for the 'base_ref_force_pushed' event.
	field_afterCommit Commit `json:"afterCommit"`

	// Identifies the before commit SHA for the 'base_ref_force_pushed' event.
	field_beforeCommit Commit `json:"beforeCommit"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`

	// PullRequest referenced by event.
	field_pullRequest PullRequest `json:"pullRequest"`

	// Identifies the fully qualified ref name for the 'base_ref_force_pushed' event.
	field_ref *Ref `json:"ref"`
}

func (o BaseRefForcePushedEvent) Actor() Actor {
	return o.field_actor
}

func (o BaseRefForcePushedEvent) AfterCommit() Commit {
	return o.field_afterCommit
}

func (o BaseRefForcePushedEvent) BeforeCommit() Commit {
	return o.field_beforeCommit
}

func (o BaseRefForcePushedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o BaseRefForcePushedEvent) Id() ID {
	return o.field_id
}

func (o BaseRefForcePushedEvent) PullRequest() PullRequest {
	return o.field_pullRequest
}

func (o BaseRefForcePushedEvent) Ref() *Ref {
	return o.field_ref
}

func (o *BaseRefForcePushedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor        json.RawMessage `json:"actor"`
		Field_afterCommit  Commit          `json:"afterCommit"`
		Field_beforeCommit Commit          `json:"beforeCommit"`
		Field_createdAt    DateTime        `json:"createdAt"`
		Field_id           ID              `json:"id"`
		Field_pullRequest  PullRequest     `json:"pullRequest"`
		Field_ref          *Ref            `json:"ref"`
	}{
		Field_afterCommit:  o.field_afterCommit,
		Field_beforeCommit: o.field_beforeCommit,
		Field_createdAt:    o.field_createdAt,
		Field_id:           o.field_id,
		Field_pullRequest:  o.field_pullRequest,
		Field_ref:          o.field_ref,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_afterCommit = v.Field_afterCommit
	o.field_beforeCommit = v.Field_beforeCommit
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	o.field_pullRequest = v.Field_pullRequest
	o.field_ref = v.Field_ref
	return nil
}

// Represents an 'review_requested' event on a given pull request.
type ReviewRequestedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`

	// PullRequest referenced by event.
	field_pullRequest PullRequest `json:"pullRequest"`

	// Identifies the user whose review was requested.
	field_subject User `json:"subject"`
}

func (o ReviewRequestedEvent) Actor() Actor {
	return o.field_actor
}

func (o ReviewRequestedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o ReviewRequestedEvent) Id() ID {
	return o.field_id
}

func (o ReviewRequestedEvent) PullRequest() PullRequest {
	return o.field_pullRequest
}

func (o ReviewRequestedEvent) Subject() User {
	return o.field_subject
}

func (o *ReviewRequestedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor       json.RawMessage `json:"actor"`
		Field_createdAt   DateTime        `json:"createdAt"`
		Field_id          ID              `json:"id"`
		Field_pullRequest PullRequest     `json:"pullRequest"`
		Field_subject     User            `json:"subject"`
	}{
		Field_createdAt:   o.field_createdAt,
		Field_id:          o.field_id,
		Field_pullRequest: o.field_pullRequest,
		Field_subject:     o.field_subject,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	o.field_pullRequest = v.Field_pullRequest
	o.field_subject = v.Field_subject
	return nil
}

// Represents an 'review_request_removed' event on a given pull request.
type ReviewRequestRemovedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	field_id ID `json:"id"`

	// PullRequest referenced by event.
	field_pullRequest PullRequest `json:"pullRequest"`

	// Identifies the user whose review request was removed.
	field_subject User `json:"subject"`
}

func (o ReviewRequestRemovedEvent) Actor() Actor {
	return o.field_actor
}

func (o ReviewRequestRemovedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o ReviewRequestRemovedEvent) Id() ID {
	return o.field_id
}

func (o ReviewRequestRemovedEvent) PullRequest() PullRequest {
	return o.field_pullRequest
}

func (o ReviewRequestRemovedEvent) Subject() User {
	return o.field_subject
}

func (o *ReviewRequestRemovedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor       json.RawMessage `json:"actor"`
		Field_createdAt   DateTime        `json:"createdAt"`
		Field_id          ID              `json:"id"`
		Field_pullRequest PullRequest     `json:"pullRequest"`
		Field_subject     User            `json:"subject"`
	}{
		Field_createdAt:   o.field_createdAt,
		Field_id:          o.field_id,
		Field_pullRequest: o.field_pullRequest,
		Field_subject:     o.field_subject,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_createdAt = v.Field_createdAt
	o.field_id = v.Field_id
	o.field_pullRequest = v.Field_pullRequest
	o.field_subject = v.Field_subject
	return nil
}

// Represents a 'review_dismissed' event on a given issue or pull request.
type ReviewDismissedEvent struct {
	// Identifies the actor who performed the event.
	field_actor Actor `json:"actor"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// Identifies the primary key from the database.
	field_databaseId *int32 `json:"databaseId"`

	field_id ID `json:"id"`

	// Identifies the message associated with the 'review_dismissed' event.
	field_message string `json:"message"`

	// The message associated with the event, rendered to HTML.
	field_messageHtml HTML `json:"messageHtml"`

	// Identifies the previous state of the review with the 'review_dismissed' event.
	field_previousReviewState PullRequestReviewState `json:"previousReviewState"`

	// PullRequest referenced by event.
	field_pullRequest PullRequest `json:"pullRequest"`

	// Identifies the commit which caused the review to become stale.
	field_pullRequestCommit *PullRequestCommit `json:"pullRequestCommit"`

	// The HTTP path for this ReviewDismissedEvent.
	field_resourcePath URI `json:"resourcePath"`

	// Identifies the review associated with the 'review_dismissed' event.
	field_review *PullRequestReview `json:"review"`

	// The HTTP URL for this ReviewDismissedEvent.
	field_url URI `json:"url"`
}

func (o ReviewDismissedEvent) Actor() Actor {
	return o.field_actor
}

func (o ReviewDismissedEvent) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o ReviewDismissedEvent) DatabaseId() *int32 {
	return o.field_databaseId
}

func (o ReviewDismissedEvent) Id() ID {
	return o.field_id
}

func (o ReviewDismissedEvent) Message() string {
	return o.field_message
}

func (o ReviewDismissedEvent) MessageHtml() HTML {
	return o.field_messageHtml
}

func (o ReviewDismissedEvent) PreviousReviewState() PullRequestReviewState {
	return o.field_previousReviewState
}

func (o ReviewDismissedEvent) PullRequest() PullRequest {
	return o.field_pullRequest
}

func (o ReviewDismissedEvent) PullRequestCommit() *PullRequestCommit {
	return o.field_pullRequestCommit
}

func (o ReviewDismissedEvent) ResourcePath() URI {
	return o.field_resourcePath
}

func (o ReviewDismissedEvent) Review() *PullRequestReview {
	return o.field_review
}

func (o ReviewDismissedEvent) Url() URI {
	return o.field_url
}

func (o *ReviewDismissedEvent) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor               json.RawMessage        `json:"actor"`
		Field_createdAt           DateTime               `json:"createdAt"`
		Field_databaseId          *int32                 `json:"databaseId"`
		Field_id                  ID                     `json:"id"`
		Field_message             string                 `json:"message"`
		Field_messageHtml         HTML                   `json:"messageHtml"`
		Field_previousReviewState PullRequestReviewState `json:"previousReviewState"`
		Field_pullRequest         PullRequest            `json:"pullRequest"`
		Field_pullRequestCommit   *PullRequestCommit     `json:"pullRequestCommit"`
		Field_resourcePath        URI                    `json:"resourcePath"`
		Field_review              *PullRequestReview     `json:"review"`
		Field_url                 URI                    `json:"url"`
	}{
		Field_createdAt:           o.field_createdAt,
		Field_databaseId:          o.field_databaseId,
		Field_id:                  o.field_id,
		Field_message:             o.field_message,
		Field_messageHtml:         o.field_messageHtml,
		Field_previousReviewState: o.field_previousReviewState,
		Field_pullRequest:         o.field_pullRequest,
		Field_pullRequestCommit:   o.field_pullRequestCommit,
		Field_resourcePath:        o.field_resourcePath,
		Field_review:              o.field_review,
		Field_url:                 o.field_url,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = Actor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_createdAt = v.Field_createdAt
	o.field_databaseId = v.Field_databaseId
	o.field_id = v.Field_id
	o.field_message = v.Field_message
	o.field_messageHtml = v.Field_messageHtml
	o.field_previousReviewState = v.Field_previousReviewState
	o.field_pullRequest = v.Field_pullRequest
	o.field_pullRequestCommit = v.Field_pullRequestCommit
	o.field_resourcePath = v.Field_resourcePath
	o.field_review = v.Field_review
	o.field_url = v.Field_url
	return nil
}

// A suggestion to review a pull request based on a user's commit history and review comments.
type SuggestedReviewer struct {
	// Is this suggestion based on past commits?
	field_isAuthor bool `json:"isAuthor"`

	// Is this suggestion based on past review comments?
	field_isCommenter bool `json:"isCommenter"`

	// Identifies the user suggested to review the pull request.
	field_reviewer User `json:"reviewer"`
}

func (o SuggestedReviewer) IsAuthor() bool {
	return o.field_isAuthor
}

func (o SuggestedReviewer) IsCommenter() bool {
	return o.field_isCommenter
}

func (o SuggestedReviewer) Reviewer() User {
	return o.field_reviewer
}

func (o *SuggestedReviewer) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_isAuthor    bool `json:"isAuthor"`
		Field_isCommenter bool `json:"isCommenter"`
		Field_reviewer    User `json:"reviewer"`
	}{
		Field_isAuthor:    o.field_isAuthor,
		Field_isCommenter: o.field_isCommenter,
		Field_reviewer:    o.field_reviewer,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_isAuthor = v.Field_isAuthor
	o.field_isCommenter = v.Field_isCommenter
	o.field_reviewer = v.Field_reviewer
	return nil
}

// The connection type for ReleaseAsset.
type ReleaseAssetConnection struct {
	// A list of edges.
	field_edges []*ReleaseAssetEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*ReleaseAsset `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o ReleaseAssetConnection) Edges() []*ReleaseAssetEdge {
	return o.field_edges
}

func (o ReleaseAssetConnection) Nodes() []*ReleaseAsset {
	return o.field_nodes
}

func (o ReleaseAssetConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o ReleaseAssetConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *ReleaseAssetConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*ReleaseAssetEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*ReleaseAsset, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type ReleaseAssetEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *ReleaseAsset `json:"node"`
}

func (o ReleaseAssetEdge) Cursor() string {
	return o.field_cursor
}

func (o ReleaseAssetEdge) Node() *ReleaseAsset {
	return o.field_node
}

func (o *ReleaseAssetEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string        `json:"cursor"`
		Field_node   *ReleaseAsset `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// A release asset contains the content for a release asset.
type ReleaseAsset struct {
	field_id ID `json:"id"`

	// Identifies the title of the release asset.
	field_name string `json:"name"`

	// release that the asset is associated with
	field_release *Release `json:"release"`

	// Identifies the URL of the release asset.
	field_url URI `json:"url"`
}

func (o ReleaseAsset) Id() ID {
	return o.field_id
}

func (o ReleaseAsset) Name() string {
	return o.field_name
}

func (o ReleaseAsset) Release() *Release {
	return o.field_release
}

func (o ReleaseAsset) Url() URI {
	return o.field_url
}

func (o *ReleaseAsset) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_id      ID       `json:"id"`
		Field_name    string   `json:"name"`
		Field_release *Release `json:"release"`
		Field_url     URI      `json:"url"`
	}{
		Field_id:      o.field_id,
		Field_name:    o.field_name,
		Field_release: o.field_release,
		Field_url:     o.field_url,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_id = v.Field_id
	o.field_name = v.Field_name
	o.field_release = v.Field_release
	o.field_url = v.Field_url
	return nil
}

// The connection type for User.
type FollowingConnection struct {
	// A list of edges.
	field_edges []*UserEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*User `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o FollowingConnection) Edges() []*UserEdge {
	return o.field_edges
}

func (o FollowingConnection) Nodes() []*User {
	return o.field_nodes
}

func (o FollowingConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o FollowingConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *FollowingConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*UserEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*User, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// The connection type for User.
type FollowerConnection struct {
	// A list of edges.
	field_edges []*UserEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*User `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o FollowerConnection) Edges() []*UserEdge {
	return o.field_edges
}

func (o FollowerConnection) Nodes() []*User {
	return o.field_nodes
}

func (o FollowerConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o FollowerConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *FollowerConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*UserEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*User, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// A Gist.
type Gist struct {
	// A list of comments associated with the gist
	field_comments GistCommentConnection `json:"comments"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// The gist description.
	field_description *string `json:"description"`

	field_id ID `json:"id"`

	// Whether the gist is public or not.
	field_isPublic bool `json:"isPublic"`

	// The gist name.
	field_name string `json:"name"`

	// The gist owner.
	field_owner RepositoryOwner `json:"owner"`

	// A list of users who have starred this starrable.
	field_stargazers StargazerConnection `json:"stargazers"`

	// Identifies the date and time when the object was last updated.
	field_updatedAt DateTime `json:"updatedAt"`

	// Returns a boolean indicating whether the viewing user has starred this starrable.
	field_viewerHasStarred bool `json:"viewerHasStarred"`
}

func (o Gist) Comments() GistCommentConnection {
	return o.field_comments
}

func (o Gist) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o Gist) Description() *string {
	return o.field_description
}

func (o Gist) Id() ID {
	return o.field_id
}

func (o Gist) IsPublic() bool {
	return o.field_isPublic
}

func (o Gist) Name() string {
	return o.field_name
}

func (o Gist) Owner() RepositoryOwner {
	return o.field_owner
}

func (o Gist) Stargazers() StargazerConnection {
	return o.field_stargazers
}

func (o Gist) UpdatedAt() DateTime {
	return o.field_updatedAt
}

func (o Gist) ViewerHasStarred() bool {
	return o.field_viewerHasStarred
}

func (o *Gist) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_comments         GistCommentConnection `json:"comments"`
		Field_createdAt        DateTime              `json:"createdAt"`
		Field_description      *string               `json:"description"`
		Field_id               ID                    `json:"id"`
		Field_isPublic         bool                  `json:"isPublic"`
		Field_name             string                `json:"name"`
		Field_owner            json.RawMessage       `json:"owner"`
		Field_stargazers       StargazerConnection   `json:"stargazers"`
		Field_updatedAt        DateTime              `json:"updatedAt"`
		Field_viewerHasStarred bool                  `json:"viewerHasStarred"`
	}{
		Field_comments:         o.field_comments,
		Field_createdAt:        o.field_createdAt,
		Field_description:      o.field_description,
		Field_id:               o.field_id,
		Field_isPublic:         o.field_isPublic,
		Field_name:             o.field_name,
		Field_stargazers:       o.field_stargazers,
		Field_updatedAt:        o.field_updatedAt,
		Field_viewerHasStarred: o.field_viewerHasStarred,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_comments = v.Field_comments
	o.field_createdAt = v.Field_createdAt
	o.field_description = v.Field_description
	o.field_id = v.Field_id
	o.field_isPublic = v.Field_isPublic
	o.field_name = v.Field_name
	if len(v.Field_owner) > 0 {
		o.field_owner, err = RepositoryOwner_UnmarshalJSON(v.Field_owner)
		if err != nil {
			return err
		}
	}
	o.field_stargazers = v.Field_stargazers
	o.field_updatedAt = v.Field_updatedAt
	o.field_viewerHasStarred = v.Field_viewerHasStarred
	return nil
}

// The connection type for User.
type StargazerConnection struct {
	// A list of edges.
	field_edges []*StargazerEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*User `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o StargazerConnection) Edges() []*StargazerEdge {
	return o.field_edges
}

func (o StargazerConnection) Nodes() []*User {
	return o.field_nodes
}

func (o StargazerConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o StargazerConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *StargazerConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*StargazerEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*User, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// Represents a user that's starred a repository.
type StargazerEdge struct {
	field_cursor string `json:"cursor"`

	field_node User `json:"node"`

	// Identifies when the item was starred.
	field_starredAt DateTime `json:"starredAt"`
}

func (o StargazerEdge) Cursor() string {
	return o.field_cursor
}

func (o StargazerEdge) Node() User {
	return o.field_node
}

func (o StargazerEdge) StarredAt() DateTime {
	return o.field_starredAt
}

func (o *StargazerEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor    string   `json:"cursor"`
		Field_node      User     `json:"node"`
		Field_starredAt DateTime `json:"starredAt"`
	}{
		Field_cursor:    o.field_cursor,
		Field_node:      o.field_node,
		Field_starredAt: o.field_starredAt,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	o.field_starredAt = v.Field_starredAt
	return nil
}

// The connection type for GistComment.
type GistCommentConnection struct {
	// A list of edges.
	field_edges []*GistCommentEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*GistComment `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o GistCommentConnection) Edges() []*GistCommentEdge {
	return o.field_edges
}

func (o GistCommentConnection) Nodes() []*GistComment {
	return o.field_nodes
}

func (o GistCommentConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o GistCommentConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *GistCommentConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*GistCommentEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*GistComment, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type GistCommentEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *GistComment `json:"node"`
}

func (o GistCommentEdge) Cursor() string {
	return o.field_cursor
}

func (o GistCommentEdge) Node() *GistComment {
	return o.field_node
}

func (o *GistCommentEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string       `json:"cursor"`
		Field_node   *GistComment `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// Represents a comment on an Gist.
type GistComment struct {
	// The actor who authored the comment.
	field_author Actor `json:"author"`

	// Author's association with the gist.
	field_authorAssociation CommentAuthorAssociation `json:"authorAssociation"`

	// Identifies the comment body.
	field_body string `json:"body"`

	// The comment body rendered to HTML.
	field_bodyHTML HTML `json:"bodyHTML"`

	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// Check if this comment was created via an email reply.
	field_createdViaEmail bool `json:"createdViaEmail"`

	// The actor who edited the comment.
	field_editor Actor `json:"editor"`

	field_id ID `json:"id"`

	// The moment the editor made the last edit
	field_lastEditedAt *DateTime `json:"lastEditedAt"`

	// Identifies when the comment was published at.
	field_publishedAt *DateTime `json:"publishedAt"`

	// Identifies the date and time when the object was last updated.
	field_updatedAt DateTime `json:"updatedAt"`

	// Check if the current viewer can delete this object.
	field_viewerCanDelete bool `json:"viewerCanDelete"`

	// Check if the current viewer can update this object.
	field_viewerCanUpdate bool `json:"viewerCanUpdate"`

	// Reasons why the current viewer can not update this comment.
	field_viewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`

	// Did the viewer author this comment.
	field_viewerDidAuthor bool `json:"viewerDidAuthor"`
}

func (o GistComment) Author() Actor {
	return o.field_author
}

func (o GistComment) AuthorAssociation() CommentAuthorAssociation {
	return o.field_authorAssociation
}

func (o GistComment) Body() string {
	return o.field_body
}

func (o GistComment) BodyHTML() HTML {
	return o.field_bodyHTML
}

func (o GistComment) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o GistComment) CreatedViaEmail() bool {
	return o.field_createdViaEmail
}

func (o GistComment) Editor() Actor {
	return o.field_editor
}

func (o GistComment) Id() ID {
	return o.field_id
}

func (o GistComment) LastEditedAt() *DateTime {
	return o.field_lastEditedAt
}

func (o GistComment) PublishedAt() *DateTime {
	return o.field_publishedAt
}

func (o GistComment) UpdatedAt() DateTime {
	return o.field_updatedAt
}

func (o GistComment) ViewerCanDelete() bool {
	return o.field_viewerCanDelete
}

func (o GistComment) ViewerCanUpdate() bool {
	return o.field_viewerCanUpdate
}

func (o GistComment) ViewerCannotUpdateReasons() []CommentCannotUpdateReason {
	return o.field_viewerCannotUpdateReasons
}

func (o GistComment) ViewerDidAuthor() bool {
	return o.field_viewerDidAuthor
}

func (o *GistComment) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_author                    json.RawMessage             `json:"author"`
		Field_authorAssociation         CommentAuthorAssociation    `json:"authorAssociation"`
		Field_body                      string                      `json:"body"`
		Field_bodyHTML                  HTML                        `json:"bodyHTML"`
		Field_createdAt                 DateTime                    `json:"createdAt"`
		Field_createdViaEmail           bool                        `json:"createdViaEmail"`
		Field_editor                    json.RawMessage             `json:"editor"`
		Field_id                        ID                          `json:"id"`
		Field_lastEditedAt              *DateTime                   `json:"lastEditedAt"`
		Field_publishedAt               *DateTime                   `json:"publishedAt"`
		Field_updatedAt                 DateTime                    `json:"updatedAt"`
		Field_viewerCanDelete           bool                        `json:"viewerCanDelete"`
		Field_viewerCanUpdate           bool                        `json:"viewerCanUpdate"`
		Field_viewerCannotUpdateReasons []CommentCannotUpdateReason `json:"viewerCannotUpdateReasons"`
		Field_viewerDidAuthor           bool                        `json:"viewerDidAuthor"`
	}{
		Field_authorAssociation:         o.field_authorAssociation,
		Field_body:                      o.field_body,
		Field_bodyHTML:                  o.field_bodyHTML,
		Field_createdAt:                 o.field_createdAt,
		Field_createdViaEmail:           o.field_createdViaEmail,
		Field_id:                        o.field_id,
		Field_lastEditedAt:              o.field_lastEditedAt,
		Field_publishedAt:               o.field_publishedAt,
		Field_updatedAt:                 o.field_updatedAt,
		Field_viewerCanDelete:           o.field_viewerCanDelete,
		Field_viewerCanUpdate:           o.field_viewerCanUpdate,
		Field_viewerCannotUpdateReasons: o.field_viewerCannotUpdateReasons,
		Field_viewerDidAuthor:           o.field_viewerDidAuthor,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_author) > 0 {
		o.field_author, err = Actor_UnmarshalJSON(v.Field_author)
		if err != nil {
			return err
		}
	}
	o.field_authorAssociation = v.Field_authorAssociation
	o.field_body = v.Field_body
	o.field_bodyHTML = v.Field_bodyHTML
	o.field_createdAt = v.Field_createdAt
	o.field_createdViaEmail = v.Field_createdViaEmail
	if len(v.Field_editor) > 0 {
		o.field_editor, err = Actor_UnmarshalJSON(v.Field_editor)
		if err != nil {
			return err
		}
	}
	o.field_id = v.Field_id
	o.field_lastEditedAt = v.Field_lastEditedAt
	o.field_publishedAt = v.Field_publishedAt
	o.field_updatedAt = v.Field_updatedAt
	o.field_viewerCanDelete = v.Field_viewerCanDelete
	o.field_viewerCanUpdate = v.Field_viewerCanUpdate
	o.field_viewerCannotUpdateReasons = v.Field_viewerCannotUpdateReasons
	o.field_viewerDidAuthor = v.Field_viewerDidAuthor
	return nil
}

// The connection type for Gist.
type GistConnection struct {
	// A list of edges.
	field_edges []*GistEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*Gist `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o GistConnection) Edges() []*GistEdge {
	return o.field_edges
}

func (o GistConnection) Nodes() []*Gist {
	return o.field_nodes
}

func (o GistConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o GistConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *GistConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*GistEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*Gist, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type GistEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *Gist `json:"node"`
}

func (o GistEdge) Cursor() string {
	return o.field_cursor
}

func (o GistEdge) Node() *Gist {
	return o.field_node
}

func (o *GistEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string `json:"cursor"`
		Field_node   *Gist  `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// The connection type for Organization.
type OrganizationConnection struct {
	// A list of edges.
	field_edges []*OrganizationEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*Organization `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o OrganizationConnection) Edges() []*OrganizationEdge {
	return o.field_edges
}

func (o OrganizationConnection) Nodes() []*Organization {
	return o.field_nodes
}

func (o OrganizationConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o OrganizationConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *OrganizationConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*OrganizationEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*Organization, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type OrganizationEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *Organization `json:"node"`
}

func (o OrganizationEdge) Cursor() string {
	return o.field_cursor
}

func (o OrganizationEdge) Node() *Organization {
	return o.field_node
}

func (o *OrganizationEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string        `json:"cursor"`
		Field_node   *Organization `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// The connection type for Repository.
type StarredRepositoryConnection struct {
	// A list of edges.
	field_edges []*StarredRepositoryEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*Repository `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o StarredRepositoryConnection) Edges() []*StarredRepositoryEdge {
	return o.field_edges
}

func (o StarredRepositoryConnection) Nodes() []*Repository {
	return o.field_nodes
}

func (o StarredRepositoryConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o StarredRepositoryConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *StarredRepositoryConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*StarredRepositoryEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*Repository, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// Represents a starred repository.
type StarredRepositoryEdge struct {
	field_cursor string `json:"cursor"`

	field_node Repository `json:"node"`

	// Identifies when the item was starred.
	field_starredAt DateTime `json:"starredAt"`
}

func (o StarredRepositoryEdge) Cursor() string {
	return o.field_cursor
}

func (o StarredRepositoryEdge) Node() Repository {
	return o.field_node
}

func (o StarredRepositoryEdge) StarredAt() DateTime {
	return o.field_starredAt
}

func (o *StarredRepositoryEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor    string     `json:"cursor"`
		Field_node      Repository `json:"node"`
		Field_starredAt DateTime   `json:"starredAt"`
	}{
		Field_cursor:    o.field_cursor,
		Field_node:      o.field_node,
		Field_starredAt: o.field_starredAt,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	o.field_starredAt = v.Field_starredAt
	return nil
}

// Represents a Milestone object on a given repository.
type Milestone struct {
	// Identifies the actor who created the milestone.
	field_creator Actor `json:"creator"`

	// Identifies the description of the milestone.
	field_description *string `json:"description"`

	// Identifies the due date of the milestone.
	field_dueOn *DateTime `json:"dueOn"`

	field_id ID `json:"id"`

	// Identifies the number of the milestone.
	field_number int32 `json:"number"`

	// The repository associated with this milestone.
	field_repository Repository `json:"repository"`

	// The HTTP path for this milestone
	field_resourcePath URI `json:"resourcePath"`

	// Identifies the state of the milestone.
	field_state MilestoneState `json:"state"`

	// Identifies the title of the milestone.
	field_title string `json:"title"`

	// The HTTP URL for this milestone
	field_url URI `json:"url"`
}

func (o Milestone) Creator() Actor {
	return o.field_creator
}

func (o Milestone) Description() *string {
	return o.field_description
}

func (o Milestone) DueOn() *DateTime {
	return o.field_dueOn
}

func (o Milestone) Id() ID {
	return o.field_id
}

func (o Milestone) Number() int32 {
	return o.field_number
}

func (o Milestone) Repository() Repository {
	return o.field_repository
}

func (o Milestone) ResourcePath() URI {
	return o.field_resourcePath
}

func (o Milestone) State() MilestoneState {
	return o.field_state
}

func (o Milestone) Title() string {
	return o.field_title
}

func (o Milestone) Url() URI {
	return o.field_url
}

func (o *Milestone) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_creator      json.RawMessage `json:"creator"`
		Field_description  *string         `json:"description"`
		Field_dueOn        *DateTime       `json:"dueOn"`
		Field_id           ID              `json:"id"`
		Field_number       int32           `json:"number"`
		Field_repository   Repository      `json:"repository"`
		Field_resourcePath URI             `json:"resourcePath"`
		Field_state        MilestoneState  `json:"state"`
		Field_title        string          `json:"title"`
		Field_url          URI             `json:"url"`
	}{
		Field_description:  o.field_description,
		Field_dueOn:        o.field_dueOn,
		Field_id:           o.field_id,
		Field_number:       o.field_number,
		Field_repository:   o.field_repository,
		Field_resourcePath: o.field_resourcePath,
		Field_state:        o.field_state,
		Field_title:        o.field_title,
		Field_url:          o.field_url,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_creator) > 0 {
		o.field_creator, err = Actor_UnmarshalJSON(v.Field_creator)
		if err != nil {
			return err
		}
	}
	o.field_description = v.Field_description
	o.field_dueOn = v.Field_dueOn
	o.field_id = v.Field_id
	o.field_number = v.Field_number
	o.field_repository = v.Field_repository
	o.field_resourcePath = v.Field_resourcePath
	o.field_state = v.Field_state
	o.field_title = v.Field_title
	o.field_url = v.Field_url
	return nil
}

// The connection type for IssueTimelineItem.
type IssueTimelineConnection struct {
	// A list of edges.
	field_edges []*IssueTimelineItemEdge `json:"edges"`

	// A list of nodes.
	field_nodes []IssueTimelineItem `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o IssueTimelineConnection) Edges() []*IssueTimelineItemEdge {
	return o.field_edges
}

func (o IssueTimelineConnection) Nodes() []IssueTimelineItem {
	return o.field_nodes
}

func (o IssueTimelineConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o IssueTimelineConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *IssueTimelineConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*IssueTimelineItemEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]IssueTimelineItem, len(l))
			for li, lv := range l {
				o.field_nodes[li], err = IssueTimelineItem_UnmarshalJSON(lv)
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type IssueTimelineItemEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node IssueTimelineItem `json:"node"`
}

func (o IssueTimelineItemEdge) Cursor() string {
	return o.field_cursor
}

func (o IssueTimelineItemEdge) Node() IssueTimelineItem {
	return o.field_node
}

func (o *IssueTimelineItemEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string          `json:"cursor"`
		Field_node   json.RawMessage `json:"node"`
	}{
		Field_cursor: o.field_cursor,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	if len(v.Field_node) > 0 {
		o.field_node, err = IssueTimelineItem_UnmarshalJSON(v.Field_node)
		if err != nil {
			return err
		}
	}
	return nil
}

// The connection type for RepositoryTopic.
type RepositoryTopicConnection struct {
	// A list of edges.
	field_edges []*RepositoryTopicEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*RepositoryTopic `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o RepositoryTopicConnection) Edges() []*RepositoryTopicEdge {
	return o.field_edges
}

func (o RepositoryTopicConnection) Nodes() []*RepositoryTopic {
	return o.field_nodes
}

func (o RepositoryTopicConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o RepositoryTopicConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *RepositoryTopicConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*RepositoryTopicEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*RepositoryTopic, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type RepositoryTopicEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *RepositoryTopic `json:"node"`
}

func (o RepositoryTopicEdge) Cursor() string {
	return o.field_cursor
}

func (o RepositoryTopicEdge) Node() *RepositoryTopic {
	return o.field_node
}

func (o *RepositoryTopicEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string           `json:"cursor"`
		Field_node   *RepositoryTopic `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// A repository-topic connects a repository to a topic.
type RepositoryTopic struct {
	field_id ID `json:"id"`

	// The HTTP path for this repository-topic.
	field_resourcePath URI `json:"resourcePath"`

	// The topic.
	field_topic Topic `json:"topic"`

	// The HTTP URL for this repository-topic.
	field_url URI `json:"url"`
}

func (o RepositoryTopic) Id() ID {
	return o.field_id
}

func (o RepositoryTopic) ResourcePath() URI {
	return o.field_resourcePath
}

func (o RepositoryTopic) Topic() Topic {
	return o.field_topic
}

func (o RepositoryTopic) Url() URI {
	return o.field_url
}

func (o *RepositoryTopic) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_id           ID    `json:"id"`
		Field_resourcePath URI   `json:"resourcePath"`
		Field_topic        Topic `json:"topic"`
		Field_url          URI   `json:"url"`
	}{
		Field_id:           o.field_id,
		Field_resourcePath: o.field_resourcePath,
		Field_topic:        o.field_topic,
		Field_url:          o.field_url,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_id = v.Field_id
	o.field_resourcePath = v.Field_resourcePath
	o.field_topic = v.Field_topic
	o.field_url = v.Field_url
	return nil
}

// A topic aggregates entities that are related to a subject.
type Topic struct {
	field_id ID `json:"id"`

	// The topic's name.
	field_name string `json:"name"`

	// A list of related topics sorted with the most relevant first.
	field_relatedTopics []Topic `json:"relatedTopics"`
}

func (o Topic) Id() ID {
	return o.field_id
}

func (o Topic) Name() string {
	return o.field_name
}

func (o Topic) RelatedTopics() []Topic {
	return o.field_relatedTopics
}

func (o *Topic) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_id            ID      `json:"id"`
		Field_name          string  `json:"name"`
		Field_relatedTopics []Topic `json:"relatedTopics"`
	}{
		Field_id:            o.field_id,
		Field_name:          o.field_name,
		Field_relatedTopics: o.field_relatedTopics,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_id = v.Field_id
	o.field_name = v.Field_name
	o.field_relatedTopics = v.Field_relatedTopics
	return nil
}

// The connection type for ProtectedBranch.
type ProtectedBranchConnection struct {
	// A list of edges.
	field_edges []*ProtectedBranchEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*ProtectedBranch `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o ProtectedBranchConnection) Edges() []*ProtectedBranchEdge {
	return o.field_edges
}

func (o ProtectedBranchConnection) Nodes() []*ProtectedBranch {
	return o.field_nodes
}

func (o ProtectedBranchConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o ProtectedBranchConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *ProtectedBranchConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*ProtectedBranchEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*ProtectedBranch, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type ProtectedBranchEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *ProtectedBranch `json:"node"`
}

func (o ProtectedBranchEdge) Cursor() string {
	return o.field_cursor
}

func (o ProtectedBranchEdge) Node() *ProtectedBranch {
	return o.field_node
}

func (o *ProtectedBranchEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string           `json:"cursor"`
		Field_node   *ProtectedBranch `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// A repository protected branch.
type ProtectedBranch struct {
	// The actor who created this protected branch.
	field_creator Actor `json:"creator"`

	// Will new commits pushed to this branch dismiss pull request review approvals.
	field_hasDismissableStaleReviews bool `json:"hasDismissableStaleReviews"`

	// Are reviews required to update this branch.
	field_hasRequiredReviews bool `json:"hasRequiredReviews"`

	// Are status checks required to update this branch.
	field_hasRequiredStatusChecks bool `json:"hasRequiredStatusChecks"`

	// Is pushing to this branch restricted.
	field_hasRestrictedPushes bool `json:"hasRestrictedPushes"`

	// Is dismissal of pull request reviews restricted.
	field_hasRestrictedReviewDismissals bool `json:"hasRestrictedReviewDismissals"`

	// Are branches required to be up to date before merging.
	field_hasStrictRequiredStatusChecks bool `json:"hasStrictRequiredStatusChecks"`

	field_id ID `json:"id"`

	// Can admins overwrite branch protection.
	field_isAdminEnforced bool `json:"isAdminEnforced"`

	// Identifies the name of the protected branch.
	field_name string `json:"name"`

	// A list push allowances for this protected branch.
	field_pushAllowances PushAllowanceConnection `json:"pushAllowances"`

	// The repository associated with this protected branch.
	field_repository Repository `json:"repository"`

	// List of required status check contexts that must pass for commits to be accepted to this branch.
	field_requiredStatusCheckContexts []*string `json:"requiredStatusCheckContexts"`

	// A list review dismissal allowances for this protected branch.
	field_reviewDismissalAllowances ReviewDismissalAllowanceConnection `json:"reviewDismissalAllowances"`
}

func (o ProtectedBranch) Creator() Actor {
	return o.field_creator
}

func (o ProtectedBranch) HasDismissableStaleReviews() bool {
	return o.field_hasDismissableStaleReviews
}

func (o ProtectedBranch) HasRequiredReviews() bool {
	return o.field_hasRequiredReviews
}

func (o ProtectedBranch) HasRequiredStatusChecks() bool {
	return o.field_hasRequiredStatusChecks
}

func (o ProtectedBranch) HasRestrictedPushes() bool {
	return o.field_hasRestrictedPushes
}

func (o ProtectedBranch) HasRestrictedReviewDismissals() bool {
	return o.field_hasRestrictedReviewDismissals
}

func (o ProtectedBranch) HasStrictRequiredStatusChecks() bool {
	return o.field_hasStrictRequiredStatusChecks
}

func (o ProtectedBranch) Id() ID {
	return o.field_id
}

func (o ProtectedBranch) IsAdminEnforced() bool {
	return o.field_isAdminEnforced
}

func (o ProtectedBranch) Name() string {
	return o.field_name
}

func (o ProtectedBranch) PushAllowances() PushAllowanceConnection {
	return o.field_pushAllowances
}

func (o ProtectedBranch) Repository() Repository {
	return o.field_repository
}

func (o ProtectedBranch) RequiredStatusCheckContexts() []*string {
	return o.field_requiredStatusCheckContexts
}

func (o ProtectedBranch) ReviewDismissalAllowances() ReviewDismissalAllowanceConnection {
	return o.field_reviewDismissalAllowances
}

func (o *ProtectedBranch) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_creator                       json.RawMessage                    `json:"creator"`
		Field_hasDismissableStaleReviews    bool                               `json:"hasDismissableStaleReviews"`
		Field_hasRequiredReviews            bool                               `json:"hasRequiredReviews"`
		Field_hasRequiredStatusChecks       bool                               `json:"hasRequiredStatusChecks"`
		Field_hasRestrictedPushes           bool                               `json:"hasRestrictedPushes"`
		Field_hasRestrictedReviewDismissals bool                               `json:"hasRestrictedReviewDismissals"`
		Field_hasStrictRequiredStatusChecks bool                               `json:"hasStrictRequiredStatusChecks"`
		Field_id                            ID                                 `json:"id"`
		Field_isAdminEnforced               bool                               `json:"isAdminEnforced"`
		Field_name                          string                             `json:"name"`
		Field_pushAllowances                PushAllowanceConnection            `json:"pushAllowances"`
		Field_repository                    Repository                         `json:"repository"`
		Field_requiredStatusCheckContexts   json.RawMessage                    `json:"requiredStatusCheckContexts"`
		Field_reviewDismissalAllowances     ReviewDismissalAllowanceConnection `json:"reviewDismissalAllowances"`
	}{
		Field_hasDismissableStaleReviews:    o.field_hasDismissableStaleReviews,
		Field_hasRequiredReviews:            o.field_hasRequiredReviews,
		Field_hasRequiredStatusChecks:       o.field_hasRequiredStatusChecks,
		Field_hasRestrictedPushes:           o.field_hasRestrictedPushes,
		Field_hasRestrictedReviewDismissals: o.field_hasRestrictedReviewDismissals,
		Field_hasStrictRequiredStatusChecks: o.field_hasStrictRequiredStatusChecks,
		Field_id:                        o.field_id,
		Field_isAdminEnforced:           o.field_isAdminEnforced,
		Field_name:                      o.field_name,
		Field_pushAllowances:            o.field_pushAllowances,
		Field_repository:                o.field_repository,
		Field_reviewDismissalAllowances: o.field_reviewDismissalAllowances,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_creator) > 0 {
		o.field_creator, err = Actor_UnmarshalJSON(v.Field_creator)
		if err != nil {
			return err
		}
	}
	o.field_hasDismissableStaleReviews = v.Field_hasDismissableStaleReviews
	o.field_hasRequiredReviews = v.Field_hasRequiredReviews
	o.field_hasRequiredStatusChecks = v.Field_hasRequiredStatusChecks
	o.field_hasRestrictedPushes = v.Field_hasRestrictedPushes
	o.field_hasRestrictedReviewDismissals = v.Field_hasRestrictedReviewDismissals
	o.field_hasStrictRequiredStatusChecks = v.Field_hasStrictRequiredStatusChecks
	o.field_id = v.Field_id
	o.field_isAdminEnforced = v.Field_isAdminEnforced
	o.field_name = v.Field_name
	o.field_pushAllowances = v.Field_pushAllowances
	o.field_repository = v.Field_repository
	if len(v.Field_requiredStatusCheckContexts) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_requiredStatusCheckContexts, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_requiredStatusCheckContexts = nil
		} else {
			o.field_requiredStatusCheckContexts = make([]*string, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_requiredStatusCheckContexts[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_reviewDismissalAllowances = v.Field_reviewDismissalAllowances
	return nil
}

// The connection type for ReviewDismissalAllowance.
type ReviewDismissalAllowanceConnection struct {
	// A list of edges.
	field_edges []*ReviewDismissalAllowanceEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*ReviewDismissalAllowance `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o ReviewDismissalAllowanceConnection) Edges() []*ReviewDismissalAllowanceEdge {
	return o.field_edges
}

func (o ReviewDismissalAllowanceConnection) Nodes() []*ReviewDismissalAllowance {
	return o.field_nodes
}

func (o ReviewDismissalAllowanceConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o ReviewDismissalAllowanceConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *ReviewDismissalAllowanceConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*ReviewDismissalAllowanceEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*ReviewDismissalAllowance, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type ReviewDismissalAllowanceEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *ReviewDismissalAllowance `json:"node"`
}

func (o ReviewDismissalAllowanceEdge) Cursor() string {
	return o.field_cursor
}

func (o ReviewDismissalAllowanceEdge) Node() *ReviewDismissalAllowance {
	return o.field_node
}

func (o *ReviewDismissalAllowanceEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string                    `json:"cursor"`
		Field_node   *ReviewDismissalAllowance `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// A team or user who has the ability to dismiss a review on a protected branch.
type ReviewDismissalAllowance struct {
	// The actor that can dismiss.
	field_actor ReviewDismissalAllowanceActor `json:"actor"`

	field_id ID `json:"id"`

	// Identifies the protected branch associated with the allowed user or team.
	field_protectedBranch ProtectedBranch `json:"protectedBranch"`
}

func (o ReviewDismissalAllowance) Actor() ReviewDismissalAllowanceActor {
	return o.field_actor
}

func (o ReviewDismissalAllowance) Id() ID {
	return o.field_id
}

func (o ReviewDismissalAllowance) ProtectedBranch() ProtectedBranch {
	return o.field_protectedBranch
}

func (o *ReviewDismissalAllowance) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor           json.RawMessage `json:"actor"`
		Field_id              ID              `json:"id"`
		Field_protectedBranch ProtectedBranch `json:"protectedBranch"`
	}{
		Field_id:              o.field_id,
		Field_protectedBranch: o.field_protectedBranch,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = ReviewDismissalAllowanceActor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_id = v.Field_id
	o.field_protectedBranch = v.Field_protectedBranch
	return nil
}

// The connection type for PushAllowance.
type PushAllowanceConnection struct {
	// A list of edges.
	field_edges []*PushAllowanceEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*PushAllowance `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o PushAllowanceConnection) Edges() []*PushAllowanceEdge {
	return o.field_edges
}

func (o PushAllowanceConnection) Nodes() []*PushAllowance {
	return o.field_nodes
}

func (o PushAllowanceConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o PushAllowanceConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *PushAllowanceConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*PushAllowanceEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*PushAllowance, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type PushAllowanceEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *PushAllowance `json:"node"`
}

func (o PushAllowanceEdge) Cursor() string {
	return o.field_cursor
}

func (o PushAllowanceEdge) Node() *PushAllowance {
	return o.field_node
}

func (o *PushAllowanceEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string         `json:"cursor"`
		Field_node   *PushAllowance `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// A team or user who has the ability to push to a protected branch.
type PushAllowance struct {
	// The actor that can push.
	field_actor PushAllowanceActor `json:"actor"`

	field_id ID `json:"id"`

	// Identifies the protected branch associated with the allowed user or team.
	field_protectedBranch ProtectedBranch `json:"protectedBranch"`
}

func (o PushAllowance) Actor() PushAllowanceActor {
	return o.field_actor
}

func (o PushAllowance) Id() ID {
	return o.field_id
}

func (o PushAllowance) ProtectedBranch() ProtectedBranch {
	return o.field_protectedBranch
}

func (o *PushAllowance) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_actor           json.RawMessage `json:"actor"`
		Field_id              ID              `json:"id"`
		Field_protectedBranch ProtectedBranch `json:"protectedBranch"`
	}{
		Field_id:              o.field_id,
		Field_protectedBranch: o.field_protectedBranch,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_actor) > 0 {
		o.field_actor, err = PushAllowanceActor_UnmarshalJSON(v.Field_actor)
		if err != nil {
			return err
		}
	}
	o.field_id = v.Field_id
	o.field_protectedBranch = v.Field_protectedBranch
	return nil
}

// The connection type for Milestone.
type MilestoneConnection struct {
	// A list of edges.
	field_edges []*MilestoneEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*Milestone `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o MilestoneConnection) Edges() []*MilestoneEdge {
	return o.field_edges
}

func (o MilestoneConnection) Nodes() []*Milestone {
	return o.field_nodes
}

func (o MilestoneConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o MilestoneConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *MilestoneConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*MilestoneEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*Milestone, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type MilestoneEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *Milestone `json:"node"`
}

func (o MilestoneEdge) Cursor() string {
	return o.field_cursor
}

func (o MilestoneEdge) Node() *Milestone {
	return o.field_node
}

func (o *MilestoneEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string     `json:"cursor"`
		Field_node   *Milestone `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// The Code of Conduct for a repository
type CodeOfConduct struct {
	// The body of the CoC
	field_body *string `json:"body"`

	// The key for the CoC
	field_key string `json:"key"`

	// The formal name of the CoC
	field_name string `json:"name"`

	// The path to the CoC
	field_url *URI `json:"url"`
}

func (o CodeOfConduct) Body() *string {
	return o.field_body
}

func (o CodeOfConduct) Key() string {
	return o.field_key
}

func (o CodeOfConduct) Name() string {
	return o.field_name
}

func (o CodeOfConduct) Url() *URI {
	return o.field_url
}

func (o *CodeOfConduct) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_body *string `json:"body"`
		Field_key  string  `json:"key"`
		Field_name string  `json:"name"`
		Field_url  *URI    `json:"url"`
	}{
		Field_body: o.field_body,
		Field_key:  o.field_key,
		Field_name: o.field_name,
		Field_url:  o.field_url,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_body = v.Field_body
	o.field_key = v.Field_key
	o.field_name = v.Field_name
	o.field_url = v.Field_url
	return nil
}

// The connection type for Ref.
type RefConnection struct {
	// A list of edges.
	field_edges []*RefEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*Ref `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o RefConnection) Edges() []*RefEdge {
	return o.field_edges
}

func (o RefConnection) Nodes() []*Ref {
	return o.field_nodes
}

func (o RefConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o RefConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *RefConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*RefEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*Ref, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type RefEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *Ref `json:"node"`
}

func (o RefEdge) Cursor() string {
	return o.field_cursor
}

func (o RefEdge) Node() *Ref {
	return o.field_node
}

func (o *RefEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string `json:"cursor"`
		Field_node   *Ref   `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// The connection type for Release.
type ReleaseConnection struct {
	// A list of edges.
	field_edges []*ReleaseEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*Release `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o ReleaseConnection) Edges() []*ReleaseEdge {
	return o.field_edges
}

func (o ReleaseConnection) Nodes() []*Release {
	return o.field_nodes
}

func (o ReleaseConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o ReleaseConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *ReleaseConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*ReleaseEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*Release, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type ReleaseEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *Release `json:"node"`
}

func (o ReleaseEdge) Cursor() string {
	return o.field_cursor
}

func (o ReleaseEdge) Node() *Release {
	return o.field_node
}

func (o *ReleaseEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string   `json:"cursor"`
		Field_node   *Release `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// The connection type for Deployment.
type DeploymentConnection struct {
	// A list of edges.
	field_edges []*DeploymentEdge `json:"edges"`

	// A list of nodes.
	field_nodes []*Deployment `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// Identifies the total count of items in the connection.
	field_totalCount int32 `json:"totalCount"`
}

func (o DeploymentConnection) Edges() []*DeploymentEdge {
	return o.field_edges
}

func (o DeploymentConnection) Nodes() []*Deployment {
	return o.field_nodes
}

func (o DeploymentConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o DeploymentConnection) TotalCount() int32 {
	return o.field_totalCount
}

func (o *DeploymentConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_edges      json.RawMessage `json:"edges"`
		Field_nodes      json.RawMessage `json:"nodes"`
		Field_pageInfo   PageInfo        `json:"pageInfo"`
		Field_totalCount int32           `json:"totalCount"`
	}{
		Field_pageInfo:   o.field_pageInfo,
		Field_totalCount: o.field_totalCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*DeploymentEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]*Deployment, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_nodes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_totalCount = v.Field_totalCount
	return nil
}

// An edge in a connection.
type DeploymentEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node *Deployment `json:"node"`
}

func (o DeploymentEdge) Cursor() string {
	return o.field_cursor
}

func (o DeploymentEdge) Node() *Deployment {
	return o.field_node
}

func (o *DeploymentEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string      `json:"cursor"`
		Field_node   *Deployment `json:"node"`
	}{
		Field_cursor: o.field_cursor,
		Field_node:   o.field_node,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	o.field_node = v.Field_node
	return nil
}

// Represents a GPG signature on a Commit or Tag.
type GpgSignature struct {
	// Email used to sign this object.
	field_email string `json:"email"`

	// True if the signature is valid and verified by GitHub.
	field_isValid bool `json:"isValid"`

	// Hex-encoded ID of the key that signed this object.
	field_keyId *string `json:"keyId"`

	// Payload for GPG signing object. Raw ODB object without the signature header.
	field_payload string `json:"payload"`

	// ASCII-armored signature header from object.
	field_signature string `json:"signature"`

	// GitHub user corresponding to the email signing this commit.
	field_signer *User `json:"signer"`

	// The state of this signature. `VALID` if signature is valid and verified by GitHub, otherwise represents reason why signature is considered invalid.
	field_state GitSignatureState `json:"state"`
}

func (o GpgSignature) Email() string {
	return o.field_email
}

func (o GpgSignature) IsValid() bool {
	return o.field_isValid
}

func (o GpgSignature) KeyId() *string {
	return o.field_keyId
}

func (o GpgSignature) Payload() string {
	return o.field_payload
}

func (o GpgSignature) Signature() string {
	return o.field_signature
}

func (o GpgSignature) Signer() *User {
	return o.field_signer
}

func (o GpgSignature) State() GitSignatureState {
	return o.field_state
}

func (o *GpgSignature) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_email     string            `json:"email"`
		Field_isValid   bool              `json:"isValid"`
		Field_keyId     *string           `json:"keyId"`
		Field_payload   string            `json:"payload"`
		Field_signature string            `json:"signature"`
		Field_signer    *User             `json:"signer"`
		Field_state     GitSignatureState `json:"state"`
	}{
		Field_email:     o.field_email,
		Field_isValid:   o.field_isValid,
		Field_keyId:     o.field_keyId,
		Field_payload:   o.field_payload,
		Field_signature: o.field_signature,
		Field_signer:    o.field_signer,
		Field_state:     o.field_state,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_email = v.Field_email
	o.field_isValid = v.Field_isValid
	o.field_keyId = v.Field_keyId
	o.field_payload = v.Field_payload
	o.field_signature = v.Field_signature
	o.field_signer = v.Field_signer
	o.field_state = v.Field_state
	return nil
}

// An invitation for a user to be added to a repository.
type RepositoryInvitation struct {
	field_id ID `json:"id"`

	// The user who received the invitation.
	field_invitee User `json:"invitee"`

	// The user who created the invitation.
	field_inviter User `json:"inviter"`

	// The Repository the user is invited to.
	field_repository *RepositoryInvitationRepository `json:"repository"`
}

func (o RepositoryInvitation) Id() ID {
	return o.field_id
}

func (o RepositoryInvitation) Invitee() User {
	return o.field_invitee
}

func (o RepositoryInvitation) Inviter() User {
	return o.field_inviter
}

func (o RepositoryInvitation) Repository() *RepositoryInvitationRepository {
	return o.field_repository
}

func (o *RepositoryInvitation) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_id         ID                              `json:"id"`
		Field_invitee    User                            `json:"invitee"`
		Field_inviter    User                            `json:"inviter"`
		Field_repository *RepositoryInvitationRepository `json:"repository"`
	}{
		Field_id:         o.field_id,
		Field_invitee:    o.field_invitee,
		Field_inviter:    o.field_inviter,
		Field_repository: o.field_repository,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_id = v.Field_id
	o.field_invitee = v.Field_invitee
	o.field_inviter = v.Field_inviter
	o.field_repository = v.Field_repository
	return nil
}

// A subset of repository info shared with potential collaborators.
type RepositoryInvitationRepository struct {
	// Identifies the date and time when the object was created.
	field_createdAt DateTime `json:"createdAt"`

	// The description of the repository.
	field_description *string `json:"description"`

	// The description of the repository rendered to HTML.
	field_descriptionHTML HTML `json:"descriptionHTML"`

	// Indicates if the repository has issues feature enabled.
	field_hasIssuesEnabled bool `json:"hasIssuesEnabled"`

	// Indicates if the repository has wiki feature enabled.
	field_hasWikiEnabled bool `json:"hasWikiEnabled"`

	// The repository's URL.
	field_homepageUrl *URI `json:"homepageUrl"`

	// Identifies if the repository is a fork.
	field_isFork bool `json:"isFork"`

	// Indicates if the repository has been locked or not.
	field_isLocked bool `json:"isLocked"`

	// Identifies if the repository is a mirror.
	field_isMirror bool `json:"isMirror"`

	// Identifies if the repository is private.
	field_isPrivate bool `json:"isPrivate"`

	// The license associated with the repository
	field_license *string `json:"license"`

	// The reason the repository has been locked.
	field_lockReason *RepositoryLockReason `json:"lockReason"`

	// The repository's original mirror URL.
	field_mirrorUrl *URI `json:"mirrorUrl"`

	// The name of the repository.
	field_name string `json:"name"`

	// The repository's name with owner.
	field_nameWithOwner string `json:"nameWithOwner"`

	// The User owner of the repository.
	field_owner RepositoryOwner `json:"owner"`

	// Identifies when the repository was last pushed to.
	field_pushedAt *DateTime `json:"pushedAt"`

	// The HTTP path for this repository
	field_resourcePath URI `json:"resourcePath"`

	// Identifies the date and time when the object was last updated.
	field_updatedAt DateTime `json:"updatedAt"`

	// The HTTP URL for this repository
	field_url URI `json:"url"`
}

func (o RepositoryInvitationRepository) CreatedAt() DateTime {
	return o.field_createdAt
}

func (o RepositoryInvitationRepository) Description() *string {
	return o.field_description
}

func (o RepositoryInvitationRepository) DescriptionHTML() HTML {
	return o.field_descriptionHTML
}

func (o RepositoryInvitationRepository) HasIssuesEnabled() bool {
	return o.field_hasIssuesEnabled
}

func (o RepositoryInvitationRepository) HasWikiEnabled() bool {
	return o.field_hasWikiEnabled
}

func (o RepositoryInvitationRepository) HomepageUrl() *URI {
	return o.field_homepageUrl
}

func (o RepositoryInvitationRepository) IsFork() bool {
	return o.field_isFork
}

func (o RepositoryInvitationRepository) IsLocked() bool {
	return o.field_isLocked
}

func (o RepositoryInvitationRepository) IsMirror() bool {
	return o.field_isMirror
}

func (o RepositoryInvitationRepository) IsPrivate() bool {
	return o.field_isPrivate
}

func (o RepositoryInvitationRepository) License() *string {
	return o.field_license
}

func (o RepositoryInvitationRepository) LockReason() *RepositoryLockReason {
	return o.field_lockReason
}

func (o RepositoryInvitationRepository) MirrorUrl() *URI {
	return o.field_mirrorUrl
}

func (o RepositoryInvitationRepository) Name() string {
	return o.field_name
}

func (o RepositoryInvitationRepository) NameWithOwner() string {
	return o.field_nameWithOwner
}

func (o RepositoryInvitationRepository) Owner() RepositoryOwner {
	return o.field_owner
}

func (o RepositoryInvitationRepository) PushedAt() *DateTime {
	return o.field_pushedAt
}

func (o RepositoryInvitationRepository) ResourcePath() URI {
	return o.field_resourcePath
}

func (o RepositoryInvitationRepository) UpdatedAt() DateTime {
	return o.field_updatedAt
}

func (o RepositoryInvitationRepository) Url() URI {
	return o.field_url
}

func (o *RepositoryInvitationRepository) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_createdAt        DateTime              `json:"createdAt"`
		Field_description      *string               `json:"description"`
		Field_descriptionHTML  HTML                  `json:"descriptionHTML"`
		Field_hasIssuesEnabled bool                  `json:"hasIssuesEnabled"`
		Field_hasWikiEnabled   bool                  `json:"hasWikiEnabled"`
		Field_homepageUrl      *URI                  `json:"homepageUrl"`
		Field_isFork           bool                  `json:"isFork"`
		Field_isLocked         bool                  `json:"isLocked"`
		Field_isMirror         bool                  `json:"isMirror"`
		Field_isPrivate        bool                  `json:"isPrivate"`
		Field_license          *string               `json:"license"`
		Field_lockReason       *RepositoryLockReason `json:"lockReason"`
		Field_mirrorUrl        *URI                  `json:"mirrorUrl"`
		Field_name             string                `json:"name"`
		Field_nameWithOwner    string                `json:"nameWithOwner"`
		Field_owner            RepositoryOwner       `json:"owner"`
		Field_pushedAt         *DateTime             `json:"pushedAt"`
		Field_resourcePath     URI                   `json:"resourcePath"`
		Field_updatedAt        DateTime              `json:"updatedAt"`
		Field_url              URI                   `json:"url"`
	}{
		Field_createdAt:        o.field_createdAt,
		Field_description:      o.field_description,
		Field_descriptionHTML:  o.field_descriptionHTML,
		Field_hasIssuesEnabled: o.field_hasIssuesEnabled,
		Field_hasWikiEnabled:   o.field_hasWikiEnabled,
		Field_homepageUrl:      o.field_homepageUrl,
		Field_isFork:           o.field_isFork,
		Field_isLocked:         o.field_isLocked,
		Field_isMirror:         o.field_isMirror,
		Field_isPrivate:        o.field_isPrivate,
		Field_license:          o.field_license,
		Field_lockReason:       o.field_lockReason,
		Field_mirrorUrl:        o.field_mirrorUrl,
		Field_name:             o.field_name,
		Field_nameWithOwner:    o.field_nameWithOwner,
		Field_owner:            o.field_owner,
		Field_pushedAt:         o.field_pushedAt,
		Field_resourcePath:     o.field_resourcePath,
		Field_updatedAt:        o.field_updatedAt,
		Field_url:              o.field_url,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_createdAt = v.Field_createdAt
	o.field_description = v.Field_description
	o.field_descriptionHTML = v.Field_descriptionHTML
	o.field_hasIssuesEnabled = v.Field_hasIssuesEnabled
	o.field_hasWikiEnabled = v.Field_hasWikiEnabled
	o.field_homepageUrl = v.Field_homepageUrl
	o.field_isFork = v.Field_isFork
	o.field_isLocked = v.Field_isLocked
	o.field_isMirror = v.Field_isMirror
	o.field_isPrivate = v.Field_isPrivate
	o.field_license = v.Field_license
	o.field_lockReason = v.Field_lockReason
	o.field_mirrorUrl = v.Field_mirrorUrl
	o.field_name = v.Field_name
	o.field_nameWithOwner = v.Field_nameWithOwner
	o.field_owner = v.Field_owner
	o.field_pushedAt = v.Field_pushedAt
	o.field_resourcePath = v.Field_resourcePath
	o.field_updatedAt = v.Field_updatedAt
	o.field_url = v.Field_url
	return nil
}

// Represents an S/MIME signature on a Commit or Tag.
type SmimeSignature struct {
	// Email used to sign this object.
	field_email string `json:"email"`

	// True if the signature is valid and verified by GitHub.
	field_isValid bool `json:"isValid"`

	// Payload for GPG signing object. Raw ODB object without the signature header.
	field_payload string `json:"payload"`

	// ASCII-armored signature header from object.
	field_signature string `json:"signature"`

	// GitHub user corresponding to the email signing this commit.
	field_signer *User `json:"signer"`

	// The state of this signature. `VALID` if signature is valid and verified by GitHub, otherwise represents reason why signature is considered invalid.
	field_state GitSignatureState `json:"state"`
}

func (o SmimeSignature) Email() string {
	return o.field_email
}

func (o SmimeSignature) IsValid() bool {
	return o.field_isValid
}

func (o SmimeSignature) Payload() string {
	return o.field_payload
}

func (o SmimeSignature) Signature() string {
	return o.field_signature
}

func (o SmimeSignature) Signer() *User {
	return o.field_signer
}

func (o SmimeSignature) State() GitSignatureState {
	return o.field_state
}

func (o *SmimeSignature) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_email     string            `json:"email"`
		Field_isValid   bool              `json:"isValid"`
		Field_payload   string            `json:"payload"`
		Field_signature string            `json:"signature"`
		Field_signer    *User             `json:"signer"`
		Field_state     GitSignatureState `json:"state"`
	}{
		Field_email:     o.field_email,
		Field_isValid:   o.field_isValid,
		Field_payload:   o.field_payload,
		Field_signature: o.field_signature,
		Field_signer:    o.field_signer,
		Field_state:     o.field_state,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_email = v.Field_email
	o.field_isValid = v.Field_isValid
	o.field_payload = v.Field_payload
	o.field_signature = v.Field_signature
	o.field_signer = v.Field_signer
	o.field_state = v.Field_state
	return nil
}

// Represents a Git tag.
type Tag struct {
	// An abbreviated version of the Git object ID
	field_abbreviatedOid string `json:"abbreviatedOid"`

	// The HTTP path for this Git object
	field_commitResourcePath URI `json:"commitResourcePath"`

	// The HTTP URL for this Git object
	field_commitUrl URI `json:"commitUrl"`

	field_id ID `json:"id"`

	// The Git tag message.
	field_message *string `json:"message"`

	// The Git tag name.
	field_name string `json:"name"`

	// The Git object ID
	field_oid GitObjectID `json:"oid"`

	// The Repository the Git object belongs to
	field_repository Repository `json:"repository"`

	// Details about the tag author.
	field_tagger *GitActor `json:"tagger"`

	// The Git object the tag points to.
	field_target GitObject `json:"target"`
}

func (o Tag) AbbreviatedOid() string {
	return o.field_abbreviatedOid
}

func (o Tag) CommitResourcePath() URI {
	return o.field_commitResourcePath
}

func (o Tag) CommitUrl() URI {
	return o.field_commitUrl
}

func (o Tag) Id() ID {
	return o.field_id
}

func (o Tag) Message() *string {
	return o.field_message
}

func (o Tag) Name() string {
	return o.field_name
}

func (o Tag) Oid() GitObjectID {
	return o.field_oid
}

func (o Tag) Repository() Repository {
	return o.field_repository
}

func (o Tag) Tagger() *GitActor {
	return o.field_tagger
}

func (o Tag) Target() GitObject {
	return o.field_target
}

func (o *Tag) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_abbreviatedOid     string      `json:"abbreviatedOid"`
		Field_commitResourcePath URI         `json:"commitResourcePath"`
		Field_commitUrl          URI         `json:"commitUrl"`
		Field_id                 ID          `json:"id"`
		Field_message            *string     `json:"message"`
		Field_name               string      `json:"name"`
		Field_oid                GitObjectID `json:"oid"`
		Field_repository         Repository  `json:"repository"`
		Field_tagger             *GitActor   `json:"tagger"`
		Field_target             GitObject   `json:"target"`
	}{
		Field_abbreviatedOid:     o.field_abbreviatedOid,
		Field_commitResourcePath: o.field_commitResourcePath,
		Field_commitUrl:          o.field_commitUrl,
		Field_id:                 o.field_id,
		Field_message:            o.field_message,
		Field_name:               o.field_name,
		Field_oid:                o.field_oid,
		Field_repository:         o.field_repository,
		Field_tagger:             o.field_tagger,
		Field_target:             o.field_target,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_abbreviatedOid = v.Field_abbreviatedOid
	o.field_commitResourcePath = v.Field_commitResourcePath
	o.field_commitUrl = v.Field_commitUrl
	o.field_id = v.Field_id
	o.field_message = v.Field_message
	o.field_name = v.Field_name
	o.field_oid = v.Field_oid
	o.field_repository = v.Field_repository
	o.field_tagger = v.Field_tagger
	o.field_target = v.Field_target
	return nil
}

// Represents an unknown signature on a Commit or Tag.
type UnknownSignature struct {
	// Email used to sign this object.
	field_email string `json:"email"`

	// True if the signature is valid and verified by GitHub.
	field_isValid bool `json:"isValid"`

	// Payload for GPG signing object. Raw ODB object without the signature header.
	field_payload string `json:"payload"`

	// ASCII-armored signature header from object.
	field_signature string `json:"signature"`

	// GitHub user corresponding to the email signing this commit.
	field_signer *User `json:"signer"`

	// The state of this signature. `VALID` if signature is valid and verified by GitHub, otherwise represents reason why signature is considered invalid.
	field_state GitSignatureState `json:"state"`
}

func (o UnknownSignature) Email() string {
	return o.field_email
}

func (o UnknownSignature) IsValid() bool {
	return o.field_isValid
}

func (o UnknownSignature) Payload() string {
	return o.field_payload
}

func (o UnknownSignature) Signature() string {
	return o.field_signature
}

func (o UnknownSignature) Signer() *User {
	return o.field_signer
}

func (o UnknownSignature) State() GitSignatureState {
	return o.field_state
}

func (o *UnknownSignature) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_email     string            `json:"email"`
		Field_isValid   bool              `json:"isValid"`
		Field_payload   string            `json:"payload"`
		Field_signature string            `json:"signature"`
		Field_signer    *User             `json:"signer"`
		Field_state     GitSignatureState `json:"state"`
	}{
		Field_email:     o.field_email,
		Field_isValid:   o.field_isValid,
		Field_payload:   o.field_payload,
		Field_signature: o.field_signature,
		Field_signer:    o.field_signer,
		Field_state:     o.field_state,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_email = v.Field_email
	o.field_isValid = v.Field_isValid
	o.field_payload = v.Field_payload
	o.field_signature = v.Field_signature
	o.field_signer = v.Field_signer
	o.field_state = v.Field_state
	return nil
}

// The query root of GitHub's GraphQL interface.
type Query struct {
	// Look up a code of conduct by its key
	field_codeOfConduct *CodeOfConduct `json:"codeOfConduct"`

	// Look up a code of conduct by its key
	field_codesOfConduct []*CodeOfConduct `json:"codesOfConduct"`

	// Fetches an object given its ID.
	field_node Node `json:"node"`

	// Lookup nodes by a list of IDs.
	field_nodes []Node `json:"nodes"`

	// Lookup a organization by login.
	field_organization *Organization `json:"organization"`

	// The client's rate limit information.
	field_rateLimit *RateLimit `json:"rateLimit"`

	// Hack to workaround https://github.com/facebook/relay/issues/112 re-exposing the root query object
	field_relay *Query `json:"relay"`

	// Lookup a given repository by the owner and repository name.
	field_repository *Repository `json:"repository"`

	// Lookup a repository owner (ie. either a User or an Organization) by login.
	field_repositoryOwner RepositoryOwner `json:"repositoryOwner"`

	// Lookup resource by a URL.
	field_resource UniformResourceLocatable `json:"resource"`

	// Perform a search across resources.
	field_search SearchResultItemConnection `json:"search"`

	// Look up a topic by name.
	field_topic *Topic `json:"topic"`

	// Lookup a user by login.
	field_user *User `json:"user"`

	// The currently authenticated user.
	field_viewer User `json:"viewer"`
}

func (o Query) CodeOfConduct() *CodeOfConduct {
	return o.field_codeOfConduct
}

func (o Query) CodesOfConduct() []*CodeOfConduct {
	return o.field_codesOfConduct
}

func (o Query) Node() Node {
	return o.field_node
}

func (o Query) Nodes() []Node {
	return o.field_nodes
}

func (o Query) Organization() *Organization {
	return o.field_organization
}

func (o Query) RateLimit() *RateLimit {
	return o.field_rateLimit
}

func (o Query) Relay() *Query {
	return o.field_relay
}

func (o Query) Repository() *Repository {
	return o.field_repository
}

func (o Query) RepositoryOwner() RepositoryOwner {
	return o.field_repositoryOwner
}

func (o Query) Resource() UniformResourceLocatable {
	return o.field_resource
}

func (o Query) Search() SearchResultItemConnection {
	return o.field_search
}

func (o Query) Topic() *Topic {
	return o.field_topic
}

func (o Query) User() *User {
	return o.field_user
}

func (o Query) Viewer() User {
	return o.field_viewer
}

func (o *Query) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_codeOfConduct   *CodeOfConduct             `json:"codeOfConduct"`
		Field_codesOfConduct  json.RawMessage            `json:"codesOfConduct"`
		Field_node            json.RawMessage            `json:"node"`
		Field_nodes           []Node                     `json:"nodes"`
		Field_organization    *Organization              `json:"organization"`
		Field_rateLimit       *RateLimit                 `json:"rateLimit"`
		Field_relay           *Query                     `json:"relay"`
		Field_repository      *Repository                `json:"repository"`
		Field_repositoryOwner json.RawMessage            `json:"repositoryOwner"`
		Field_resource        json.RawMessage            `json:"resource"`
		Field_search          SearchResultItemConnection `json:"search"`
		Field_topic           *Topic                     `json:"topic"`
		Field_user            *User                      `json:"user"`
		Field_viewer          User                       `json:"viewer"`
	}{
		Field_codeOfConduct: o.field_codeOfConduct,
		Field_nodes:         o.field_nodes,
		Field_organization:  o.field_organization,
		Field_rateLimit:     o.field_rateLimit,
		Field_relay:         o.field_relay,
		Field_repository:    o.field_repository,
		Field_search:        o.field_search,
		Field_topic:         o.field_topic,
		Field_user:          o.field_user,
		Field_viewer:        o.field_viewer,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_codeOfConduct = v.Field_codeOfConduct
	if len(v.Field_codesOfConduct) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_codesOfConduct, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_codesOfConduct = nil
		} else {
			o.field_codesOfConduct = make([]*CodeOfConduct, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_codesOfConduct[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_node) > 0 {
		o.field_node, err = Node_UnmarshalJSON(v.Field_node)
		if err != nil {
			return err
		}
	}
	o.field_nodes = v.Field_nodes
	o.field_organization = v.Field_organization
	o.field_rateLimit = v.Field_rateLimit
	o.field_relay = v.Field_relay
	o.field_repository = v.Field_repository
	if len(v.Field_repositoryOwner) > 0 {
		o.field_repositoryOwner, err = RepositoryOwner_UnmarshalJSON(v.Field_repositoryOwner)
		if err != nil {
			return err
		}
	}
	if len(v.Field_resource) > 0 {
		o.field_resource, err = UniformResourceLocatable_UnmarshalJSON(v.Field_resource)
		if err != nil {
			return err
		}
	}
	o.field_search = v.Field_search
	o.field_topic = v.Field_topic
	o.field_user = v.Field_user
	o.field_viewer = v.Field_viewer
	return nil
}

// Represents the client's rate limit.
type RateLimit struct {
	// The point cost for the current query counting against the rate limit.
	field_cost int32 `json:"cost"`

	// The maximum number of points the client is permitted to consume in a 60 minute window.
	field_limit int32 `json:"limit"`

	// The number of points remaining in the current rate limit window.
	field_remaining int32 `json:"remaining"`

	// The time at which the current rate limit window resets in UTC epoch seconds.
	field_resetAt DateTime `json:"resetAt"`
}

func (o RateLimit) Cost() int32 {
	return o.field_cost
}

func (o RateLimit) Limit() int32 {
	return o.field_limit
}

func (o RateLimit) Remaining() int32 {
	return o.field_remaining
}

func (o RateLimit) ResetAt() DateTime {
	return o.field_resetAt
}

func (o *RateLimit) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cost      int32    `json:"cost"`
		Field_limit     int32    `json:"limit"`
		Field_remaining int32    `json:"remaining"`
		Field_resetAt   DateTime `json:"resetAt"`
	}{
		Field_cost:      o.field_cost,
		Field_limit:     o.field_limit,
		Field_remaining: o.field_remaining,
		Field_resetAt:   o.field_resetAt,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cost = v.Field_cost
	o.field_limit = v.Field_limit
	o.field_remaining = v.Field_remaining
	o.field_resetAt = v.Field_resetAt
	return nil
}

// A list of results that matched against a search query.
type SearchResultItemConnection struct {
	// The number of pieces of code that matched the search query.
	field_codeCount int32 `json:"codeCount"`

	// A list of edges.
	field_edges []*SearchResultItemEdge `json:"edges"`

	// The number of issues that matched the search query.
	field_issueCount int32 `json:"issueCount"`

	// A list of nodes.
	field_nodes []SearchResultItem `json:"nodes"`

	// Information to aid in pagination.
	field_pageInfo PageInfo `json:"pageInfo"`

	// The number of repositories that matched the search query.
	field_repositoryCount int32 `json:"repositoryCount"`

	// The number of users that matched the search query.
	field_userCount int32 `json:"userCount"`

	// The number of wiki pages that matched the search query.
	field_wikiCount int32 `json:"wikiCount"`
}

func (o SearchResultItemConnection) CodeCount() int32 {
	return o.field_codeCount
}

func (o SearchResultItemConnection) Edges() []*SearchResultItemEdge {
	return o.field_edges
}

func (o SearchResultItemConnection) IssueCount() int32 {
	return o.field_issueCount
}

func (o SearchResultItemConnection) Nodes() []SearchResultItem {
	return o.field_nodes
}

func (o SearchResultItemConnection) PageInfo() PageInfo {
	return o.field_pageInfo
}

func (o SearchResultItemConnection) RepositoryCount() int32 {
	return o.field_repositoryCount
}

func (o SearchResultItemConnection) UserCount() int32 {
	return o.field_userCount
}

func (o SearchResultItemConnection) WikiCount() int32 {
	return o.field_wikiCount
}

func (o *SearchResultItemConnection) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_codeCount       int32           `json:"codeCount"`
		Field_edges           json.RawMessage `json:"edges"`
		Field_issueCount      int32           `json:"issueCount"`
		Field_nodes           json.RawMessage `json:"nodes"`
		Field_pageInfo        PageInfo        `json:"pageInfo"`
		Field_repositoryCount int32           `json:"repositoryCount"`
		Field_userCount       int32           `json:"userCount"`
		Field_wikiCount       int32           `json:"wikiCount"`
	}{
		Field_codeCount:       o.field_codeCount,
		Field_issueCount:      o.field_issueCount,
		Field_pageInfo:        o.field_pageInfo,
		Field_repositoryCount: o.field_repositoryCount,
		Field_userCount:       o.field_userCount,
		Field_wikiCount:       o.field_wikiCount,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_codeCount = v.Field_codeCount
	if len(v.Field_edges) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_edges, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_edges = nil
		} else {
			o.field_edges = make([]*SearchResultItemEdge, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_edges[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_issueCount = v.Field_issueCount
	if len(v.Field_nodes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_nodes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_nodes = nil
		} else {
			o.field_nodes = make([]SearchResultItem, len(l))
			for li, lv := range l {
				o.field_nodes[li], err = SearchResultItem_UnmarshalJSON(lv)
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_pageInfo = v.Field_pageInfo
	o.field_repositoryCount = v.Field_repositoryCount
	o.field_userCount = v.Field_userCount
	o.field_wikiCount = v.Field_wikiCount
	return nil
}

// An edge in a connection.
type SearchResultItemEdge struct {
	// A cursor for use in pagination.
	field_cursor string `json:"cursor"`

	// The item at the end of the edge.
	field_node SearchResultItem `json:"node"`
}

func (o SearchResultItemEdge) Cursor() string {
	return o.field_cursor
}

func (o SearchResultItemEdge) Node() SearchResultItem {
	return o.field_node
}

func (o *SearchResultItemEdge) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cursor string          `json:"cursor"`
		Field_node   json.RawMessage `json:"node"`
	}{
		Field_cursor: o.field_cursor,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cursor = v.Field_cursor
	if len(v.Field_node) > 0 {
		o.field_node, err = SearchResultItem_UnmarshalJSON(v.Field_node)
		if err != nil {
			return err
		}
	}
	return nil
}

// The root query for implementing GraphQL mutations.
type Mutation struct {
	// Applies a suggested topic to the repository.
	field_acceptTopicSuggestion *AcceptTopicSuggestionPayload `json:"acceptTopicSuggestion"`

	// Adds a comment to an Issue or Pull Request.
	field_addComment *AddCommentPayload `json:"addComment"`

	// Adds a card to a ProjectColumn. Either `contentId` or `note` must be provided but **not** both.
	field_addProjectCard *AddProjectCardPayload `json:"addProjectCard"`

	// Adds a column to a Project.
	field_addProjectColumn *AddProjectColumnPayload `json:"addProjectColumn"`

	// Adds a review to a Pull Request.
	field_addPullRequestReview *AddPullRequestReviewPayload `json:"addPullRequestReview"`

	// Adds a comment to a review.
	field_addPullRequestReviewComment *AddPullRequestReviewCommentPayload `json:"addPullRequestReviewComment"`

	// Adds a reaction to a subject.
	field_addReaction *AddReactionPayload `json:"addReaction"`

	// Adds a star to a Starrable.
	field_addStar *AddStarPayload `json:"addStar"`

	// Creates a new project.
	field_createProject *CreateProjectPayload `json:"createProject"`

	// Rejects a suggested topic for the repository.
	field_declineTopicSuggestion *DeclineTopicSuggestionPayload `json:"declineTopicSuggestion"`

	// Deletes a project.
	field_deleteProject *DeleteProjectPayload `json:"deleteProject"`

	// Deletes a project card.
	field_deleteProjectCard *DeleteProjectCardPayload `json:"deleteProjectCard"`

	// Deletes a project column.
	field_deleteProjectColumn *DeleteProjectColumnPayload `json:"deleteProjectColumn"`

	// Deletes a pull request review.
	field_deletePullRequestReview *DeletePullRequestReviewPayload `json:"deletePullRequestReview"`

	// Dismisses an approved or rejected pull request review.
	field_dismissPullRequestReview *DismissPullRequestReviewPayload `json:"dismissPullRequestReview"`

	// Moves a project card to another place.
	field_moveProjectCard *MoveProjectCardPayload `json:"moveProjectCard"`

	// Moves a project column to another place.
	field_moveProjectColumn *MoveProjectColumnPayload `json:"moveProjectColumn"`

	// Removes outside collaborator from all repositories in an organization.
	field_removeOutsideCollaborator *RemoveOutsideCollaboratorPayload `json:"removeOutsideCollaborator"`

	// Removes a reaction from a subject.
	field_removeReaction *RemoveReactionPayload `json:"removeReaction"`

	// Removes a star from a Starrable.
	field_removeStar *RemoveStarPayload `json:"removeStar"`

	// Set review requests on a pull request.
	field_requestReviews *RequestReviewsPayload `json:"requestReviews"`

	// Submits a pending pull request review.
	field_submitPullRequestReview *SubmitPullRequestReviewPayload `json:"submitPullRequestReview"`

	// Updates an existing project.
	field_updateProject *UpdateProjectPayload `json:"updateProject"`

	// Updates an existing project card.
	field_updateProjectCard *UpdateProjectCardPayload `json:"updateProjectCard"`

	// Updates an existing project column.
	field_updateProjectColumn *UpdateProjectColumnPayload `json:"updateProjectColumn"`

	// Updates the body of a pull request review.
	field_updatePullRequestReview *UpdatePullRequestReviewPayload `json:"updatePullRequestReview"`

	// Updates a pull request review comment.
	field_updatePullRequestReviewComment *UpdatePullRequestReviewCommentPayload `json:"updatePullRequestReviewComment"`

	// Updates viewers repository subscription state.
	field_updateSubscription *UpdateSubscriptionPayload `json:"updateSubscription"`

	// Replaces the repository's topics with the given topics.
	field_updateTopics *UpdateTopicsPayload `json:"updateTopics"`
}

func (o Mutation) AcceptTopicSuggestion() *AcceptTopicSuggestionPayload {
	return o.field_acceptTopicSuggestion
}

func (o Mutation) AddComment() *AddCommentPayload {
	return o.field_addComment
}

func (o Mutation) AddProjectCard() *AddProjectCardPayload {
	return o.field_addProjectCard
}

func (o Mutation) AddProjectColumn() *AddProjectColumnPayload {
	return o.field_addProjectColumn
}

func (o Mutation) AddPullRequestReview() *AddPullRequestReviewPayload {
	return o.field_addPullRequestReview
}

func (o Mutation) AddPullRequestReviewComment() *AddPullRequestReviewCommentPayload {
	return o.field_addPullRequestReviewComment
}

func (o Mutation) AddReaction() *AddReactionPayload {
	return o.field_addReaction
}

func (o Mutation) AddStar() *AddStarPayload {
	return o.field_addStar
}

func (o Mutation) CreateProject() *CreateProjectPayload {
	return o.field_createProject
}

func (o Mutation) DeclineTopicSuggestion() *DeclineTopicSuggestionPayload {
	return o.field_declineTopicSuggestion
}

func (o Mutation) DeleteProject() *DeleteProjectPayload {
	return o.field_deleteProject
}

func (o Mutation) DeleteProjectCard() *DeleteProjectCardPayload {
	return o.field_deleteProjectCard
}

func (o Mutation) DeleteProjectColumn() *DeleteProjectColumnPayload {
	return o.field_deleteProjectColumn
}

func (o Mutation) DeletePullRequestReview() *DeletePullRequestReviewPayload {
	return o.field_deletePullRequestReview
}

func (o Mutation) DismissPullRequestReview() *DismissPullRequestReviewPayload {
	return o.field_dismissPullRequestReview
}

func (o Mutation) MoveProjectCard() *MoveProjectCardPayload {
	return o.field_moveProjectCard
}

func (o Mutation) MoveProjectColumn() *MoveProjectColumnPayload {
	return o.field_moveProjectColumn
}

func (o Mutation) RemoveOutsideCollaborator() *RemoveOutsideCollaboratorPayload {
	return o.field_removeOutsideCollaborator
}

func (o Mutation) RemoveReaction() *RemoveReactionPayload {
	return o.field_removeReaction
}

func (o Mutation) RemoveStar() *RemoveStarPayload {
	return o.field_removeStar
}

func (o Mutation) RequestReviews() *RequestReviewsPayload {
	return o.field_requestReviews
}

func (o Mutation) SubmitPullRequestReview() *SubmitPullRequestReviewPayload {
	return o.field_submitPullRequestReview
}

func (o Mutation) UpdateProject() *UpdateProjectPayload {
	return o.field_updateProject
}

func (o Mutation) UpdateProjectCard() *UpdateProjectCardPayload {
	return o.field_updateProjectCard
}

func (o Mutation) UpdateProjectColumn() *UpdateProjectColumnPayload {
	return o.field_updateProjectColumn
}

func (o Mutation) UpdatePullRequestReview() *UpdatePullRequestReviewPayload {
	return o.field_updatePullRequestReview
}

func (o Mutation) UpdatePullRequestReviewComment() *UpdatePullRequestReviewCommentPayload {
	return o.field_updatePullRequestReviewComment
}

func (o Mutation) UpdateSubscription() *UpdateSubscriptionPayload {
	return o.field_updateSubscription
}

func (o Mutation) UpdateTopics() *UpdateTopicsPayload {
	return o.field_updateTopics
}

func (o *Mutation) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_acceptTopicSuggestion          *AcceptTopicSuggestionPayload          `json:"acceptTopicSuggestion"`
		Field_addComment                     *AddCommentPayload                     `json:"addComment"`
		Field_addProjectCard                 *AddProjectCardPayload                 `json:"addProjectCard"`
		Field_addProjectColumn               *AddProjectColumnPayload               `json:"addProjectColumn"`
		Field_addPullRequestReview           *AddPullRequestReviewPayload           `json:"addPullRequestReview"`
		Field_addPullRequestReviewComment    *AddPullRequestReviewCommentPayload    `json:"addPullRequestReviewComment"`
		Field_addReaction                    *AddReactionPayload                    `json:"addReaction"`
		Field_addStar                        *AddStarPayload                        `json:"addStar"`
		Field_createProject                  *CreateProjectPayload                  `json:"createProject"`
		Field_declineTopicSuggestion         *DeclineTopicSuggestionPayload         `json:"declineTopicSuggestion"`
		Field_deleteProject                  *DeleteProjectPayload                  `json:"deleteProject"`
		Field_deleteProjectCard              *DeleteProjectCardPayload              `json:"deleteProjectCard"`
		Field_deleteProjectColumn            *DeleteProjectColumnPayload            `json:"deleteProjectColumn"`
		Field_deletePullRequestReview        *DeletePullRequestReviewPayload        `json:"deletePullRequestReview"`
		Field_dismissPullRequestReview       *DismissPullRequestReviewPayload       `json:"dismissPullRequestReview"`
		Field_moveProjectCard                *MoveProjectCardPayload                `json:"moveProjectCard"`
		Field_moveProjectColumn              *MoveProjectColumnPayload              `json:"moveProjectColumn"`
		Field_removeOutsideCollaborator      *RemoveOutsideCollaboratorPayload      `json:"removeOutsideCollaborator"`
		Field_removeReaction                 *RemoveReactionPayload                 `json:"removeReaction"`
		Field_removeStar                     *RemoveStarPayload                     `json:"removeStar"`
		Field_requestReviews                 *RequestReviewsPayload                 `json:"requestReviews"`
		Field_submitPullRequestReview        *SubmitPullRequestReviewPayload        `json:"submitPullRequestReview"`
		Field_updateProject                  *UpdateProjectPayload                  `json:"updateProject"`
		Field_updateProjectCard              *UpdateProjectCardPayload              `json:"updateProjectCard"`
		Field_updateProjectColumn            *UpdateProjectColumnPayload            `json:"updateProjectColumn"`
		Field_updatePullRequestReview        *UpdatePullRequestReviewPayload        `json:"updatePullRequestReview"`
		Field_updatePullRequestReviewComment *UpdatePullRequestReviewCommentPayload `json:"updatePullRequestReviewComment"`
		Field_updateSubscription             *UpdateSubscriptionPayload             `json:"updateSubscription"`
		Field_updateTopics                   *UpdateTopicsPayload                   `json:"updateTopics"`
	}{
		Field_acceptTopicSuggestion:          o.field_acceptTopicSuggestion,
		Field_addComment:                     o.field_addComment,
		Field_addProjectCard:                 o.field_addProjectCard,
		Field_addProjectColumn:               o.field_addProjectColumn,
		Field_addPullRequestReview:           o.field_addPullRequestReview,
		Field_addPullRequestReviewComment:    o.field_addPullRequestReviewComment,
		Field_addReaction:                    o.field_addReaction,
		Field_addStar:                        o.field_addStar,
		Field_createProject:                  o.field_createProject,
		Field_declineTopicSuggestion:         o.field_declineTopicSuggestion,
		Field_deleteProject:                  o.field_deleteProject,
		Field_deleteProjectCard:              o.field_deleteProjectCard,
		Field_deleteProjectColumn:            o.field_deleteProjectColumn,
		Field_deletePullRequestReview:        o.field_deletePullRequestReview,
		Field_dismissPullRequestReview:       o.field_dismissPullRequestReview,
		Field_moveProjectCard:                o.field_moveProjectCard,
		Field_moveProjectColumn:              o.field_moveProjectColumn,
		Field_removeOutsideCollaborator:      o.field_removeOutsideCollaborator,
		Field_removeReaction:                 o.field_removeReaction,
		Field_removeStar:                     o.field_removeStar,
		Field_requestReviews:                 o.field_requestReviews,
		Field_submitPullRequestReview:        o.field_submitPullRequestReview,
		Field_updateProject:                  o.field_updateProject,
		Field_updateProjectCard:              o.field_updateProjectCard,
		Field_updateProjectColumn:            o.field_updateProjectColumn,
		Field_updatePullRequestReview:        o.field_updatePullRequestReview,
		Field_updatePullRequestReviewComment: o.field_updatePullRequestReviewComment,
		Field_updateSubscription:             o.field_updateSubscription,
		Field_updateTopics:                   o.field_updateTopics,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_acceptTopicSuggestion = v.Field_acceptTopicSuggestion
	o.field_addComment = v.Field_addComment
	o.field_addProjectCard = v.Field_addProjectCard
	o.field_addProjectColumn = v.Field_addProjectColumn
	o.field_addPullRequestReview = v.Field_addPullRequestReview
	o.field_addPullRequestReviewComment = v.Field_addPullRequestReviewComment
	o.field_addReaction = v.Field_addReaction
	o.field_addStar = v.Field_addStar
	o.field_createProject = v.Field_createProject
	o.field_declineTopicSuggestion = v.Field_declineTopicSuggestion
	o.field_deleteProject = v.Field_deleteProject
	o.field_deleteProjectCard = v.Field_deleteProjectCard
	o.field_deleteProjectColumn = v.Field_deleteProjectColumn
	o.field_deletePullRequestReview = v.Field_deletePullRequestReview
	o.field_dismissPullRequestReview = v.Field_dismissPullRequestReview
	o.field_moveProjectCard = v.Field_moveProjectCard
	o.field_moveProjectColumn = v.Field_moveProjectColumn
	o.field_removeOutsideCollaborator = v.Field_removeOutsideCollaborator
	o.field_removeReaction = v.Field_removeReaction
	o.field_removeStar = v.Field_removeStar
	o.field_requestReviews = v.Field_requestReviews
	o.field_submitPullRequestReview = v.Field_submitPullRequestReview
	o.field_updateProject = v.Field_updateProject
	o.field_updateProjectCard = v.Field_updateProjectCard
	o.field_updateProjectColumn = v.Field_updateProjectColumn
	o.field_updatePullRequestReview = v.Field_updatePullRequestReview
	o.field_updatePullRequestReviewComment = v.Field_updatePullRequestReviewComment
	o.field_updateSubscription = v.Field_updateSubscription
	o.field_updateTopics = v.Field_updateTopics
	return nil
}

// Autogenerated return type of AddReaction
type AddReactionPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The reaction object.
	field_reaction Reaction `json:"reaction"`

	// The reactable subject.
	field_subject Reactable `json:"subject"`
}

func (o AddReactionPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o AddReactionPayload) Reaction() Reaction {
	return o.field_reaction
}

func (o AddReactionPayload) Subject() Reactable {
	return o.field_subject
}

func (o *AddReactionPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string   `json:"clientMutationId"`
		Field_reaction         Reaction  `json:"reaction"`
		Field_subject          Reactable `json:"subject"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_reaction:         o.field_reaction,
		Field_subject:          o.field_subject,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_reaction = v.Field_reaction
	o.field_subject = v.Field_subject
	return nil
}

// Autogenerated return type of RemoveReaction
type RemoveReactionPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The reaction object.
	field_reaction Reaction `json:"reaction"`

	// The reactable subject.
	field_subject Reactable `json:"subject"`
}

func (o RemoveReactionPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o RemoveReactionPayload) Reaction() Reaction {
	return o.field_reaction
}

func (o RemoveReactionPayload) Subject() Reactable {
	return o.field_subject
}

func (o *RemoveReactionPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string   `json:"clientMutationId"`
		Field_reaction         Reaction  `json:"reaction"`
		Field_subject          Reactable `json:"subject"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_reaction:         o.field_reaction,
		Field_subject:          o.field_subject,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_reaction = v.Field_reaction
	o.field_subject = v.Field_subject
	return nil
}

// Autogenerated return type of AddComment
type AddCommentPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The edge from the subject's comment connection.
	field_commentEdge IssueCommentEdge `json:"commentEdge"`

	// The subject
	field_subject Node `json:"subject"`

	// The edge from the subject's timeline connection.
	field_timelineEdge IssueTimelineItemEdge `json:"timelineEdge"`
}

func (o AddCommentPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o AddCommentPayload) CommentEdge() IssueCommentEdge {
	return o.field_commentEdge
}

func (o AddCommentPayload) Subject() Node {
	return o.field_subject
}

func (o AddCommentPayload) TimelineEdge() IssueTimelineItemEdge {
	return o.field_timelineEdge
}

func (o *AddCommentPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string               `json:"clientMutationId"`
		Field_commentEdge      IssueCommentEdge      `json:"commentEdge"`
		Field_subject          Node                  `json:"subject"`
		Field_timelineEdge     IssueTimelineItemEdge `json:"timelineEdge"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_commentEdge:      o.field_commentEdge,
		Field_subject:          o.field_subject,
		Field_timelineEdge:     o.field_timelineEdge,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_commentEdge = v.Field_commentEdge
	o.field_subject = v.Field_subject
	o.field_timelineEdge = v.Field_timelineEdge
	return nil
}

// Autogenerated return type of UpdateSubscription
type UpdateSubscriptionPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The input subscribable entity.
	field_subscribable Subscribable `json:"subscribable"`
}

func (o UpdateSubscriptionPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o UpdateSubscriptionPayload) Subscribable() Subscribable {
	return o.field_subscribable
}

func (o *UpdateSubscriptionPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string      `json:"clientMutationId"`
		Field_subscribable     Subscribable `json:"subscribable"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_subscribable:     o.field_subscribable,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_subscribable = v.Field_subscribable
	return nil
}

// Autogenerated return type of CreateProject
type CreateProjectPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The new project.
	field_project Project `json:"project"`
}

func (o CreateProjectPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o CreateProjectPayload) Project() Project {
	return o.field_project
}

func (o *CreateProjectPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string `json:"clientMutationId"`
		Field_project          Project `json:"project"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_project:          o.field_project,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_project = v.Field_project
	return nil
}

// Autogenerated return type of UpdateProject
type UpdateProjectPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The updated project.
	field_project Project `json:"project"`
}

func (o UpdateProjectPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o UpdateProjectPayload) Project() Project {
	return o.field_project
}

func (o *UpdateProjectPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string `json:"clientMutationId"`
		Field_project          Project `json:"project"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_project:          o.field_project,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_project = v.Field_project
	return nil
}

// Autogenerated return type of DeleteProject
type DeleteProjectPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The repository or organization the project was removed from.
	field_owner ProjectOwner `json:"owner"`
}

func (o DeleteProjectPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o DeleteProjectPayload) Owner() ProjectOwner {
	return o.field_owner
}

func (o *DeleteProjectPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string      `json:"clientMutationId"`
		Field_owner            ProjectOwner `json:"owner"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_owner:            o.field_owner,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_owner = v.Field_owner
	return nil
}

// Autogenerated return type of AddProjectColumn
type AddProjectColumnPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The edge from the project's column connection.
	field_columnEdge ProjectColumnEdge `json:"columnEdge"`

	// The project
	field_project Project `json:"project"`
}

func (o AddProjectColumnPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o AddProjectColumnPayload) ColumnEdge() ProjectColumnEdge {
	return o.field_columnEdge
}

func (o AddProjectColumnPayload) Project() Project {
	return o.field_project
}

func (o *AddProjectColumnPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string           `json:"clientMutationId"`
		Field_columnEdge       ProjectColumnEdge `json:"columnEdge"`
		Field_project          Project           `json:"project"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_columnEdge:       o.field_columnEdge,
		Field_project:          o.field_project,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_columnEdge = v.Field_columnEdge
	o.field_project = v.Field_project
	return nil
}

// Autogenerated return type of MoveProjectColumn
type MoveProjectColumnPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The new edge of the moved column.
	field_columnEdge ProjectColumnEdge `json:"columnEdge"`
}

func (o MoveProjectColumnPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o MoveProjectColumnPayload) ColumnEdge() ProjectColumnEdge {
	return o.field_columnEdge
}

func (o *MoveProjectColumnPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string           `json:"clientMutationId"`
		Field_columnEdge       ProjectColumnEdge `json:"columnEdge"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_columnEdge:       o.field_columnEdge,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_columnEdge = v.Field_columnEdge
	return nil
}

// Autogenerated return type of UpdateProjectColumn
type UpdateProjectColumnPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The updated project column.
	field_projectColumn ProjectColumn `json:"projectColumn"`
}

func (o UpdateProjectColumnPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o UpdateProjectColumnPayload) ProjectColumn() ProjectColumn {
	return o.field_projectColumn
}

func (o *UpdateProjectColumnPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string       `json:"clientMutationId"`
		Field_projectColumn    ProjectColumn `json:"projectColumn"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_projectColumn:    o.field_projectColumn,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_projectColumn = v.Field_projectColumn
	return nil
}

// Autogenerated return type of DeleteProjectColumn
type DeleteProjectColumnPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The deleted column ID.
	field_deletedColumnId ID `json:"deletedColumnId"`

	// The project the deleted column was in.
	field_project Project `json:"project"`
}

func (o DeleteProjectColumnPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o DeleteProjectColumnPayload) DeletedColumnId() ID {
	return o.field_deletedColumnId
}

func (o DeleteProjectColumnPayload) Project() Project {
	return o.field_project
}

func (o *DeleteProjectColumnPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string `json:"clientMutationId"`
		Field_deletedColumnId  ID      `json:"deletedColumnId"`
		Field_project          Project `json:"project"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_deletedColumnId:  o.field_deletedColumnId,
		Field_project:          o.field_project,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_deletedColumnId = v.Field_deletedColumnId
	o.field_project = v.Field_project
	return nil
}

// Autogenerated return type of AddProjectCard
type AddProjectCardPayload struct {
	// The edge from the ProjectColumn's card connection.
	field_cardEdge ProjectCardEdge `json:"cardEdge"`

	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The ProjectColumn
	field_projectColumn Project `json:"projectColumn"`
}

func (o AddProjectCardPayload) CardEdge() ProjectCardEdge {
	return o.field_cardEdge
}

func (o AddProjectCardPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o AddProjectCardPayload) ProjectColumn() Project {
	return o.field_projectColumn
}

func (o *AddProjectCardPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cardEdge         ProjectCardEdge `json:"cardEdge"`
		Field_clientMutationId *string         `json:"clientMutationId"`
		Field_projectColumn    Project         `json:"projectColumn"`
	}{
		Field_cardEdge:         o.field_cardEdge,
		Field_clientMutationId: o.field_clientMutationId,
		Field_projectColumn:    o.field_projectColumn,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cardEdge = v.Field_cardEdge
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_projectColumn = v.Field_projectColumn
	return nil
}

// Autogenerated return type of UpdateProjectCard
type UpdateProjectCardPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The updated ProjectCard.
	field_projectCard ProjectCard `json:"projectCard"`
}

func (o UpdateProjectCardPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o UpdateProjectCardPayload) ProjectCard() ProjectCard {
	return o.field_projectCard
}

func (o *UpdateProjectCardPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string     `json:"clientMutationId"`
		Field_projectCard      ProjectCard `json:"projectCard"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_projectCard:      o.field_projectCard,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_projectCard = v.Field_projectCard
	return nil
}

// Autogenerated return type of MoveProjectCard
type MoveProjectCardPayload struct {
	// The new edge of the moved card.
	field_cardEdge ProjectCardEdge `json:"cardEdge"`

	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`
}

func (o MoveProjectCardPayload) CardEdge() ProjectCardEdge {
	return o.field_cardEdge
}

func (o MoveProjectCardPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o *MoveProjectCardPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_cardEdge         ProjectCardEdge `json:"cardEdge"`
		Field_clientMutationId *string         `json:"clientMutationId"`
	}{
		Field_cardEdge:         o.field_cardEdge,
		Field_clientMutationId: o.field_clientMutationId,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_cardEdge = v.Field_cardEdge
	o.field_clientMutationId = v.Field_clientMutationId
	return nil
}

// Autogenerated return type of DeleteProjectCard
type DeleteProjectCardPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The column the deleted card was in.
	field_column ProjectColumn `json:"column"`

	// The deleted card ID.
	field_deletedCardId ID `json:"deletedCardId"`
}

func (o DeleteProjectCardPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o DeleteProjectCardPayload) Column() ProjectColumn {
	return o.field_column
}

func (o DeleteProjectCardPayload) DeletedCardId() ID {
	return o.field_deletedCardId
}

func (o *DeleteProjectCardPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string       `json:"clientMutationId"`
		Field_column           ProjectColumn `json:"column"`
		Field_deletedCardId    ID            `json:"deletedCardId"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_column:           o.field_column,
		Field_deletedCardId:    o.field_deletedCardId,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_column = v.Field_column
	o.field_deletedCardId = v.Field_deletedCardId
	return nil
}

// Autogenerated return type of AddPullRequestReview
type AddPullRequestReviewPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The newly created pull request review.
	field_pullRequestReview PullRequestReview `json:"pullRequestReview"`

	// The edge from the pull request's review connection.
	field_reviewEdge PullRequestReviewEdge `json:"reviewEdge"`
}

func (o AddPullRequestReviewPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o AddPullRequestReviewPayload) PullRequestReview() PullRequestReview {
	return o.field_pullRequestReview
}

func (o AddPullRequestReviewPayload) ReviewEdge() PullRequestReviewEdge {
	return o.field_reviewEdge
}

func (o *AddPullRequestReviewPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId  *string               `json:"clientMutationId"`
		Field_pullRequestReview PullRequestReview     `json:"pullRequestReview"`
		Field_reviewEdge        PullRequestReviewEdge `json:"reviewEdge"`
	}{
		Field_clientMutationId:  o.field_clientMutationId,
		Field_pullRequestReview: o.field_pullRequestReview,
		Field_reviewEdge:        o.field_reviewEdge,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_pullRequestReview = v.Field_pullRequestReview
	o.field_reviewEdge = v.Field_reviewEdge
	return nil
}

// Autogenerated return type of SubmitPullRequestReview
type SubmitPullRequestReviewPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The submitted pull request review.
	field_pullRequestReview PullRequestReview `json:"pullRequestReview"`
}

func (o SubmitPullRequestReviewPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o SubmitPullRequestReviewPayload) PullRequestReview() PullRequestReview {
	return o.field_pullRequestReview
}

func (o *SubmitPullRequestReviewPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId  *string           `json:"clientMutationId"`
		Field_pullRequestReview PullRequestReview `json:"pullRequestReview"`
	}{
		Field_clientMutationId:  o.field_clientMutationId,
		Field_pullRequestReview: o.field_pullRequestReview,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_pullRequestReview = v.Field_pullRequestReview
	return nil
}

// Autogenerated return type of UpdatePullRequestReview
type UpdatePullRequestReviewPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The updated pull request review.
	field_pullRequestReview PullRequestReview `json:"pullRequestReview"`
}

func (o UpdatePullRequestReviewPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o UpdatePullRequestReviewPayload) PullRequestReview() PullRequestReview {
	return o.field_pullRequestReview
}

func (o *UpdatePullRequestReviewPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId  *string           `json:"clientMutationId"`
		Field_pullRequestReview PullRequestReview `json:"pullRequestReview"`
	}{
		Field_clientMutationId:  o.field_clientMutationId,
		Field_pullRequestReview: o.field_pullRequestReview,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_pullRequestReview = v.Field_pullRequestReview
	return nil
}

// Autogenerated return type of DismissPullRequestReview
type DismissPullRequestReviewPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The dismissed pull request review.
	field_pullRequestReview PullRequestReview `json:"pullRequestReview"`
}

func (o DismissPullRequestReviewPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o DismissPullRequestReviewPayload) PullRequestReview() PullRequestReview {
	return o.field_pullRequestReview
}

func (o *DismissPullRequestReviewPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId  *string           `json:"clientMutationId"`
		Field_pullRequestReview PullRequestReview `json:"pullRequestReview"`
	}{
		Field_clientMutationId:  o.field_clientMutationId,
		Field_pullRequestReview: o.field_pullRequestReview,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_pullRequestReview = v.Field_pullRequestReview
	return nil
}

// Autogenerated return type of DeletePullRequestReview
type DeletePullRequestReviewPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The deleted pull request review.
	field_pullRequestReview PullRequestReview `json:"pullRequestReview"`
}

func (o DeletePullRequestReviewPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o DeletePullRequestReviewPayload) PullRequestReview() PullRequestReview {
	return o.field_pullRequestReview
}

func (o *DeletePullRequestReviewPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId  *string           `json:"clientMutationId"`
		Field_pullRequestReview PullRequestReview `json:"pullRequestReview"`
	}{
		Field_clientMutationId:  o.field_clientMutationId,
		Field_pullRequestReview: o.field_pullRequestReview,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_pullRequestReview = v.Field_pullRequestReview
	return nil
}

// Autogenerated return type of AddPullRequestReviewComment
type AddPullRequestReviewCommentPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The newly created comment.
	field_comment PullRequestReviewComment `json:"comment"`

	// The edge from the review's comment connection.
	field_commentEdge PullRequestReviewCommentEdge `json:"commentEdge"`
}

func (o AddPullRequestReviewCommentPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o AddPullRequestReviewCommentPayload) Comment() PullRequestReviewComment {
	return o.field_comment
}

func (o AddPullRequestReviewCommentPayload) CommentEdge() PullRequestReviewCommentEdge {
	return o.field_commentEdge
}

func (o *AddPullRequestReviewCommentPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string                      `json:"clientMutationId"`
		Field_comment          PullRequestReviewComment     `json:"comment"`
		Field_commentEdge      PullRequestReviewCommentEdge `json:"commentEdge"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_comment:          o.field_comment,
		Field_commentEdge:      o.field_commentEdge,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_comment = v.Field_comment
	o.field_commentEdge = v.Field_commentEdge
	return nil
}

// Autogenerated return type of UpdatePullRequestReviewComment
type UpdatePullRequestReviewCommentPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The updated comment.
	field_pullRequestReviewComment PullRequestReviewComment `json:"pullRequestReviewComment"`
}

func (o UpdatePullRequestReviewCommentPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o UpdatePullRequestReviewCommentPayload) PullRequestReviewComment() PullRequestReviewComment {
	return o.field_pullRequestReviewComment
}

func (o *UpdatePullRequestReviewCommentPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId         *string                  `json:"clientMutationId"`
		Field_pullRequestReviewComment PullRequestReviewComment `json:"pullRequestReviewComment"`
	}{
		Field_clientMutationId:         o.field_clientMutationId,
		Field_pullRequestReviewComment: o.field_pullRequestReviewComment,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_pullRequestReviewComment = v.Field_pullRequestReviewComment
	return nil
}

// Autogenerated return type of RemoveOutsideCollaborator
type RemoveOutsideCollaboratorPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The user that was removed as an outside collaborator.
	field_removedUser User `json:"removedUser"`
}

func (o RemoveOutsideCollaboratorPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o RemoveOutsideCollaboratorPayload) RemovedUser() User {
	return o.field_removedUser
}

func (o *RemoveOutsideCollaboratorPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string `json:"clientMutationId"`
		Field_removedUser      User    `json:"removedUser"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_removedUser:      o.field_removedUser,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_removedUser = v.Field_removedUser
	return nil
}

// Autogenerated return type of RequestReviews
type RequestReviewsPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The pull request that is getting requests.
	field_pullRequest PullRequest `json:"pullRequest"`

	// The edge from the pull request to the requested reviewers.
	field_requestedReviewersEdge UserEdge `json:"requestedReviewersEdge"`
}

func (o RequestReviewsPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o RequestReviewsPayload) PullRequest() PullRequest {
	return o.field_pullRequest
}

func (o RequestReviewsPayload) RequestedReviewersEdge() UserEdge {
	return o.field_requestedReviewersEdge
}

func (o *RequestReviewsPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId       *string     `json:"clientMutationId"`
		Field_pullRequest            PullRequest `json:"pullRequest"`
		Field_requestedReviewersEdge UserEdge    `json:"requestedReviewersEdge"`
	}{
		Field_clientMutationId:       o.field_clientMutationId,
		Field_pullRequest:            o.field_pullRequest,
		Field_requestedReviewersEdge: o.field_requestedReviewersEdge,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_pullRequest = v.Field_pullRequest
	o.field_requestedReviewersEdge = v.Field_requestedReviewersEdge
	return nil
}

// Autogenerated return type of AddStar
type AddStarPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The starrable.
	field_starrable Starrable `json:"starrable"`
}

func (o AddStarPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o AddStarPayload) Starrable() Starrable {
	return o.field_starrable
}

func (o *AddStarPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string   `json:"clientMutationId"`
		Field_starrable        Starrable `json:"starrable"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_starrable:        o.field_starrable,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_starrable = v.Field_starrable
	return nil
}

// Autogenerated return type of RemoveStar
type RemoveStarPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The starrable.
	field_starrable Starrable `json:"starrable"`
}

func (o RemoveStarPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o RemoveStarPayload) Starrable() Starrable {
	return o.field_starrable
}

func (o *RemoveStarPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string   `json:"clientMutationId"`
		Field_starrable        Starrable `json:"starrable"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_starrable:        o.field_starrable,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_starrable = v.Field_starrable
	return nil
}

// Autogenerated return type of AcceptTopicSuggestion
type AcceptTopicSuggestionPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The accepted topic.
	field_topic Topic `json:"topic"`
}

func (o AcceptTopicSuggestionPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o AcceptTopicSuggestionPayload) Topic() Topic {
	return o.field_topic
}

func (o *AcceptTopicSuggestionPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string `json:"clientMutationId"`
		Field_topic            Topic   `json:"topic"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_topic:            o.field_topic,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_topic = v.Field_topic
	return nil
}

// Autogenerated return type of DeclineTopicSuggestion
type DeclineTopicSuggestionPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// The declined topic.
	field_topic Topic `json:"topic"`
}

func (o DeclineTopicSuggestionPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o DeclineTopicSuggestionPayload) Topic() Topic {
	return o.field_topic
}

func (o *DeclineTopicSuggestionPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId *string `json:"clientMutationId"`
		Field_topic            Topic   `json:"topic"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_topic:            o.field_topic,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	o.field_topic = v.Field_topic
	return nil
}

// Autogenerated return type of UpdateTopics
type UpdateTopicsPayload struct {
	// A unique identifier for the client performing the mutation.
	field_clientMutationId *string `json:"clientMutationId"`

	// Names of the provided topics that are not valid.
	field_invalidTopicNames []string `json:"invalidTopicNames"`

	// The updated repository.
	field_repository Repository `json:"repository"`
}

func (o UpdateTopicsPayload) ClientMutationId() *string {
	return o.field_clientMutationId
}

func (o UpdateTopicsPayload) InvalidTopicNames() []string {
	return o.field_invalidTopicNames
}

func (o UpdateTopicsPayload) Repository() Repository {
	return o.field_repository
}

func (o *UpdateTopicsPayload) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_clientMutationId  *string         `json:"clientMutationId"`
		Field_invalidTopicNames json.RawMessage `json:"invalidTopicNames"`
		Field_repository        Repository      `json:"repository"`
	}{
		Field_clientMutationId: o.field_clientMutationId,
		Field_repository:       o.field_repository,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_clientMutationId = v.Field_clientMutationId
	if len(v.Field_invalidTopicNames) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_invalidTopicNames, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_invalidTopicNames = nil
		} else {
			o.field_invalidTopicNames = make([]string, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_invalidTopicNames[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_repository = v.Field_repository
	return nil
}

// A GraphQL Schema defines the capabilities of a GraphQL server. It exposes all available types and directives on the server, as well as the entry points for query, mutation, and subscription operations.
type __Schema struct {
	// A list of all directives supported by this server.
	field_directives []__Directive `json:"directives"`

	// If this server supports mutation, the type that mutation operations will be rooted at.
	field_mutationType *__Type `json:"mutationType"`

	// The type that query operations will be rooted at.
	field_queryType __Type `json:"queryType"`

	// If this server support subscription, the type that subscription operations will be rooted at.
	field_subscriptionType *__Type `json:"subscriptionType"`

	// A list of all types supported by this server.
	field_types []__Type `json:"types"`
}

func (o __Schema) Directives() []__Directive {
	return o.field_directives
}

func (o __Schema) MutationType() *__Type {
	return o.field_mutationType
}

func (o __Schema) QueryType() __Type {
	return o.field_queryType
}

func (o __Schema) SubscriptionType() *__Type {
	return o.field_subscriptionType
}

func (o __Schema) Types() []__Type {
	return o.field_types
}

func (o *__Schema) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_directives       []__Directive `json:"directives"`
		Field_mutationType     *__Type       `json:"mutationType"`
		Field_queryType        __Type        `json:"queryType"`
		Field_subscriptionType *__Type       `json:"subscriptionType"`
		Field_types            []__Type      `json:"types"`
	}{
		Field_directives:       o.field_directives,
		Field_mutationType:     o.field_mutationType,
		Field_queryType:        o.field_queryType,
		Field_subscriptionType: o.field_subscriptionType,
		Field_types:            o.field_types,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_directives = v.Field_directives
	o.field_mutationType = v.Field_mutationType
	o.field_queryType = v.Field_queryType
	o.field_subscriptionType = v.Field_subscriptionType
	o.field_types = v.Field_types
	return nil
}

// The fundamental unit of any GraphQL Schema is the type. There are many kinds of types in GraphQL as represented by the `__TypeKind` enum.
//
// Depending on the kind of a type, certain fields describe information about that type. Scalar types provide no information beyond a name and description, while Enum types provide their values. Object and Interface types provide the fields they describe. Abstract types, Union and Interface, provide the Object types possible at runtime. List and NonNull types compose other types.
type __Type struct {
	field_description *string `json:"description"`

	field_enumValues []__EnumValue `json:"enumValues"`

	field_fields []__Field `json:"fields"`

	field_inputFields []__InputValue `json:"inputFields"`

	field_interfaces []__Type `json:"interfaces"`

	field_kind __TypeKind `json:"kind"`

	field_name *string `json:"name"`

	field_ofType *__Type `json:"ofType"`

	field_possibleTypes []__Type `json:"possibleTypes"`
}

func (o __Type) Description() *string {
	return o.field_description
}

func (o __Type) EnumValues() []__EnumValue {
	return o.field_enumValues
}

func (o __Type) Fields() []__Field {
	return o.field_fields
}

func (o __Type) InputFields() []__InputValue {
	return o.field_inputFields
}

func (o __Type) Interfaces() []__Type {
	return o.field_interfaces
}

func (o __Type) Kind() __TypeKind {
	return o.field_kind
}

func (o __Type) Name() *string {
	return o.field_name
}

func (o __Type) OfType() *__Type {
	return o.field_ofType
}

func (o __Type) PossibleTypes() []__Type {
	return o.field_possibleTypes
}

func (o *__Type) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_description   *string         `json:"description"`
		Field_enumValues    json.RawMessage `json:"enumValues"`
		Field_fields        json.RawMessage `json:"fields"`
		Field_inputFields   json.RawMessage `json:"inputFields"`
		Field_interfaces    json.RawMessage `json:"interfaces"`
		Field_kind          __TypeKind      `json:"kind"`
		Field_name          *string         `json:"name"`
		Field_ofType        *__Type         `json:"ofType"`
		Field_possibleTypes json.RawMessage `json:"possibleTypes"`
	}{
		Field_description: o.field_description,
		Field_kind:        o.field_kind,
		Field_name:        o.field_name,
		Field_ofType:      o.field_ofType,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_description = v.Field_description
	if len(v.Field_enumValues) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_enumValues, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_enumValues = nil
		} else {
			o.field_enumValues = make([]__EnumValue, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_enumValues[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_fields) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_fields, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_fields = nil
		} else {
			o.field_fields = make([]__Field, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_fields[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_inputFields) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_inputFields, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_inputFields = nil
		} else {
			o.field_inputFields = make([]__InputValue, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_inputFields[li])
				if err != nil {
					return err
				}
			}
		}
	}
	if len(v.Field_interfaces) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_interfaces, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_interfaces = nil
		} else {
			o.field_interfaces = make([]__Type, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_interfaces[li])
				if err != nil {
					return err
				}
			}
		}
	}
	o.field_kind = v.Field_kind
	o.field_name = v.Field_name
	o.field_ofType = v.Field_ofType
	if len(v.Field_possibleTypes) > 0 {
		var l []json.RawMessage
		err = json.Unmarshal(v.Field_possibleTypes, &l)
		if err != nil {
			return err
		}
		if l == nil {
			o.field_possibleTypes = nil
		} else {
			o.field_possibleTypes = make([]__Type, len(l))
			for li, lv := range l {
				err = json.Unmarshal(lv, &o.field_possibleTypes[li])
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Object and Interface types are described by a list of Fields, each of which has a name, potentially a list of arguments, and a return type.
type __Field struct {
	field_args []__InputValue `json:"args"`

	field_deprecationReason *string `json:"deprecationReason"`

	field_description *string `json:"description"`

	field_isDeprecated bool `json:"isDeprecated"`

	field_name string `json:"name"`

	field_type __Type `json:"type"`
}

func (o __Field) Args() []__InputValue {
	return o.field_args
}

func (o __Field) DeprecationReason() *string {
	return o.field_deprecationReason
}

func (o __Field) Description() *string {
	return o.field_description
}

func (o __Field) IsDeprecated() bool {
	return o.field_isDeprecated
}

func (o __Field) Name() string {
	return o.field_name
}

func (o __Field) Type() __Type {
	return o.field_type
}

func (o *__Field) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_args              []__InputValue `json:"args"`
		Field_deprecationReason *string        `json:"deprecationReason"`
		Field_description       *string        `json:"description"`
		Field_isDeprecated      bool           `json:"isDeprecated"`
		Field_name              string         `json:"name"`
		Field_type              __Type         `json:"type"`
	}{
		Field_args:              o.field_args,
		Field_deprecationReason: o.field_deprecationReason,
		Field_description:       o.field_description,
		Field_isDeprecated:      o.field_isDeprecated,
		Field_name:              o.field_name,
		Field_type:              o.field_type,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_args = v.Field_args
	o.field_deprecationReason = v.Field_deprecationReason
	o.field_description = v.Field_description
	o.field_isDeprecated = v.Field_isDeprecated
	o.field_name = v.Field_name
	o.field_type = v.Field_type
	return nil
}

// Arguments provided to Fields or Directives and the input fields of an InputObject are represented as Input Values which describe their type and optionally a default value.
type __InputValue struct {
	// A GraphQL-formatted string representing the default value for this input value.
	field_defaultValue *string `json:"defaultValue"`

	field_description *string `json:"description"`

	field_name string `json:"name"`

	field_type __Type `json:"type"`
}

func (o __InputValue) DefaultValue() *string {
	return o.field_defaultValue
}

func (o __InputValue) Description() *string {
	return o.field_description
}

func (o __InputValue) Name() string {
	return o.field_name
}

func (o __InputValue) Type() __Type {
	return o.field_type
}

func (o *__InputValue) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_defaultValue *string `json:"defaultValue"`
		Field_description  *string `json:"description"`
		Field_name         string  `json:"name"`
		Field_type         __Type  `json:"type"`
	}{
		Field_defaultValue: o.field_defaultValue,
		Field_description:  o.field_description,
		Field_name:         o.field_name,
		Field_type:         o.field_type,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_defaultValue = v.Field_defaultValue
	o.field_description = v.Field_description
	o.field_name = v.Field_name
	o.field_type = v.Field_type
	return nil
}

// One possible value for a given Enum. Enum values are unique values, not a placeholder for a string or numeric value. However an Enum value is returned in a JSON response as a string.
type __EnumValue struct {
	field_deprecationReason *string `json:"deprecationReason"`

	field_description *string `json:"description"`

	field_isDeprecated bool `json:"isDeprecated"`

	field_name string `json:"name"`
}

func (o __EnumValue) DeprecationReason() *string {
	return o.field_deprecationReason
}

func (o __EnumValue) Description() *string {
	return o.field_description
}

func (o __EnumValue) IsDeprecated() bool {
	return o.field_isDeprecated
}

func (o __EnumValue) Name() string {
	return o.field_name
}

func (o *__EnumValue) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_deprecationReason *string `json:"deprecationReason"`
		Field_description       *string `json:"description"`
		Field_isDeprecated      bool    `json:"isDeprecated"`
		Field_name              string  `json:"name"`
	}{
		Field_deprecationReason: o.field_deprecationReason,
		Field_description:       o.field_description,
		Field_isDeprecated:      o.field_isDeprecated,
		Field_name:              o.field_name,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_deprecationReason = v.Field_deprecationReason
	o.field_description = v.Field_description
	o.field_isDeprecated = v.Field_isDeprecated
	o.field_name = v.Field_name
	return nil
}

// A Directive provides a way to describe alternate runtime execution and type validation behavior in a GraphQL document.
//
// In some cases, you need to provide options to alter GraphQL's execution behavior in ways field arguments will not suffice, such as conditionally including or skipping a field. Directives provide this by describing additional information to the executor.
type __Directive struct {
	field_args []__InputValue `json:"args"`

	field_description *string `json:"description"`

	field_locations []__DirectiveLocation `json:"locations"`

	field_name string `json:"name"`

	field_onField bool `json:"onField"`

	field_onFragment bool `json:"onFragment"`

	field_onOperation bool `json:"onOperation"`
}

func (o __Directive) Args() []__InputValue {
	return o.field_args
}

func (o __Directive) Description() *string {
	return o.field_description
}

func (o __Directive) Locations() []__DirectiveLocation {
	return o.field_locations
}

func (o __Directive) Name() string {
	return o.field_name
}

func (o __Directive) OnField() bool {
	return o.field_onField
}

func (o __Directive) OnFragment() bool {
	return o.field_onFragment
}

func (o __Directive) OnOperation() bool {
	return o.field_onOperation
}

func (o *__Directive) UnmarshalJSON(data []byte) error {
	v := struct {
		Field_args        []__InputValue        `json:"args"`
		Field_description *string               `json:"description"`
		Field_locations   []__DirectiveLocation `json:"locations"`
		Field_name        string                `json:"name"`
		Field_onField     bool                  `json:"onField"`
		Field_onFragment  bool                  `json:"onFragment"`
		Field_onOperation bool                  `json:"onOperation"`
	}{
		Field_args:        o.field_args,
		Field_description: o.field_description,
		Field_locations:   o.field_locations,
		Field_name:        o.field_name,
		Field_onField:     o.field_onField,
		Field_onFragment:  o.field_onFragment,
		Field_onOperation: o.field_onOperation,
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	o.field_args = v.Field_args
	o.field_description = v.Field_description
	o.field_locations = v.Field_locations
	o.field_name = v.Field_name
	o.field_onField = v.Field_onField
	o.field_onFragment = v.Field_onFragment
	o.field_onOperation = v.Field_onOperation
	return nil
}
