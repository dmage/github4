package github4

// State of the project; either 'open' or 'closed'
type ProjectState string

const (
	// The project is open.
	ProjectState_OPEN ProjectState = "OPEN"

	// The project is closed.
	ProjectState_CLOSED ProjectState = "CLOSED"
)

// Properties by which project connections can be ordered.
type ProjectOrderField string

const (
	// Order projects by creation time
	ProjectOrderField_CREATED_AT ProjectOrderField = "CREATED_AT"

	// Order projects by update time
	ProjectOrderField_UPDATED_AT ProjectOrderField = "UPDATED_AT"

	// Order projects by name
	ProjectOrderField_NAME ProjectOrderField = "NAME"
)

// Possible directions in which to order a list of items when provided an `orderBy` argument.
type OrderDirection string

const (
	// Specifies an ascending order for a given `orderBy` argument.
	OrderDirection_ASC OrderDirection = "ASC"

	// Specifies a descending order for a given `orderBy` argument.
	OrderDirection_DESC OrderDirection = "DESC"
)

// Various content states of a ProjectCard
type ProjectCardState string

const (
	// The card has content only.
	ProjectCardState_CONTENT_ONLY ProjectCardState = "CONTENT_ONLY"

	// The card has a note only.
	ProjectCardState_NOTE_ONLY ProjectCardState = "NOTE_ONLY"

	// The card is redacted.
	ProjectCardState_REDACTED ProjectCardState = "REDACTED"
)

// The possible states of a subscription.
type SubscriptionState string

const (
	// The User is only notified when particpating or @mentioned.
	SubscriptionState_UNSUBSCRIBED SubscriptionState = "UNSUBSCRIBED"

	// The User is notified of all conversations.
	SubscriptionState_SUBSCRIBED SubscriptionState = "SUBSCRIBED"

	// The User is never notified.
	SubscriptionState_IGNORED SubscriptionState = "IGNORED"
)

// A comment author association with repository.
type CommentAuthorAssociation string

const (
	// Author is a member of the organization that owns the repository.
	CommentAuthorAssociation_MEMBER CommentAuthorAssociation = "MEMBER"

	// Author is the owner of the repository.
	CommentAuthorAssociation_OWNER CommentAuthorAssociation = "OWNER"

	// Author has been invited to collaborate on the repository.
	CommentAuthorAssociation_COLLABORATOR CommentAuthorAssociation = "COLLABORATOR"

	// Author has previously committed to the repository.
	CommentAuthorAssociation_CONTRIBUTOR CommentAuthorAssociation = "CONTRIBUTOR"

	// Author has not previously committed to the repository.
	CommentAuthorAssociation_FIRST_TIME_CONTRIBUTOR CommentAuthorAssociation = "FIRST_TIME_CONTRIBUTOR"

	// Author has no association with the repository.
	CommentAuthorAssociation_NONE CommentAuthorAssociation = "NONE"
)

// The possible errors that will prevent a user from updating a comment.
type CommentCannotUpdateReason string

const (
	// You must be the author or have write access to this repository to update this comment.
	CommentCannotUpdateReason_INSUFFICIENT_ACCESS CommentCannotUpdateReason = "INSUFFICIENT_ACCESS"

	// Unable to create comment because issue is locked.
	CommentCannotUpdateReason_LOCKED CommentCannotUpdateReason = "LOCKED"

	// You must be logged in to update this comment.
	CommentCannotUpdateReason_LOGIN_REQUIRED CommentCannotUpdateReason = "LOGIN_REQUIRED"

	// Repository is under maintenance.
	CommentCannotUpdateReason_MAINTENANCE CommentCannotUpdateReason = "MAINTENANCE"

	// At least one email address must be verified to update this comment.
	CommentCannotUpdateReason_VERIFIED_EMAIL_REQUIRED CommentCannotUpdateReason = "VERIFIED_EMAIL_REQUIRED"
)

// Emojis that can be attached to Issues, Pull Requests and Comments.
type ReactionContent string

const (
	// Represents the 👍 emoji.
	ReactionContent_THUMBS_UP ReactionContent = "THUMBS_UP"

	// Represents the 👎 emoji.
	ReactionContent_THUMBS_DOWN ReactionContent = "THUMBS_DOWN"

	// Represents the 😄 emoji.
	ReactionContent_LAUGH ReactionContent = "LAUGH"

	// Represents the 🎉 emoji.
	ReactionContent_HOORAY ReactionContent = "HOORAY"

	// Represents the 😕 emoji.
	ReactionContent_CONFUSED ReactionContent = "CONFUSED"

	// Represents the ❤️ emoji.
	ReactionContent_HEART ReactionContent = "HEART"
)

// A list of fields that reactions can be ordered by.
type ReactionOrderField string

const (
	// Allows ordering a list of reactions by when they were created.
	ReactionOrderField_CREATED_AT ReactionOrderField = "CREATED_AT"
)

// The state of a Git signature.
type GitSignatureState string

const (
	// Valid signature and verified by GitHub.
	GitSignatureState_VALID GitSignatureState = "VALID"

	// Invalid signature.
	GitSignatureState_INVALID GitSignatureState = "INVALID"

	// Malformed signature.
	GitSignatureState_MALFORMED_SIG GitSignatureState = "MALFORMED_SIG"

	// Key used for signing not known to GitHub.
	GitSignatureState_UNKNOWN_KEY GitSignatureState = "UNKNOWN_KEY"

	// Invalid email used for signing.
	GitSignatureState_BAD_EMAIL GitSignatureState = "BAD_EMAIL"

	// Email used for signing unverified on GitHub.
	GitSignatureState_UNVERIFIED_EMAIL GitSignatureState = "UNVERIFIED_EMAIL"

	// Email used for signing not known to GitHub.
	GitSignatureState_NO_USER GitSignatureState = "NO_USER"

	// Unknown signature type.
	GitSignatureState_UNKNOWN_SIG_TYPE GitSignatureState = "UNKNOWN_SIG_TYPE"

	// Unsigned.
	GitSignatureState_UNSIGNED GitSignatureState = "UNSIGNED"

	// Internal error - the GPG verification service is unavailable at the moment.
	GitSignatureState_GPGVERIFY_UNAVAILABLE GitSignatureState = "GPGVERIFY_UNAVAILABLE"

	// Internal error - the GPG verification service misbehaved.
	GitSignatureState_GPGVERIFY_ERROR GitSignatureState = "GPGVERIFY_ERROR"

	// The usage flags for the key that signed this don't allow signing.
	GitSignatureState_NOT_SIGNING_KEY GitSignatureState = "NOT_SIGNING_KEY"

	// Signing key expired.
	GitSignatureState_EXPIRED_KEY GitSignatureState = "EXPIRED_KEY"
)

// The possible commit status states.
type StatusState string

const (
	// Status is expected.
	StatusState_EXPECTED StatusState = "EXPECTED"

	// Status is errored.
	StatusState_ERROR StatusState = "ERROR"

	// Status is failing.
	StatusState_FAILURE StatusState = "FAILURE"

	// Status is pending.
	StatusState_PENDING StatusState = "PENDING"

	// Status is successful.
	StatusState_SUCCESS StatusState = "SUCCESS"
)

// The possible states of an issue.
type IssueState string

const (
	// An issue that is still open
	IssueState_OPEN IssueState = "OPEN"

	// An issue that has been closed
	IssueState_CLOSED IssueState = "CLOSED"
)

// Properties by which issue connections can be ordered.
type IssueOrderField string

const (
	// Order issues by creation time
	IssueOrderField_CREATED_AT IssueOrderField = "CREATED_AT"

	// Order issues by update time
	IssueOrderField_UPDATED_AT IssueOrderField = "UPDATED_AT"

	// Order issues by comment count
	IssueOrderField_COMMENTS IssueOrderField = "COMMENTS"
)

// The privacy of a repository
type RepositoryPrivacy string

const (
	// Public
	RepositoryPrivacy_PUBLIC RepositoryPrivacy = "PUBLIC"

	// Private
	RepositoryPrivacy_PRIVATE RepositoryPrivacy = "PRIVATE"
)

// Properties by which repository connections can be ordered.
type RepositoryOrderField string

const (
	// Order repositories by creation time
	RepositoryOrderField_CREATED_AT RepositoryOrderField = "CREATED_AT"

	// Order repositories by update time
	RepositoryOrderField_UPDATED_AT RepositoryOrderField = "UPDATED_AT"

	// Order repositories by push time
	RepositoryOrderField_PUSHED_AT RepositoryOrderField = "PUSHED_AT"

	// Order repositories by name
	RepositoryOrderField_NAME RepositoryOrderField = "NAME"

	// Order repositories by number of stargazers
	RepositoryOrderField_STARGAZERS RepositoryOrderField = "STARGAZERS"
)

// The affiliation of a user to a repository
type RepositoryAffiliation string

const (
	// Repositories that are owned by the authenticated user.
	RepositoryAffiliation_OWNER RepositoryAffiliation = "OWNER"

	// Repositories that the user has been added to as a collaborator.
	RepositoryAffiliation_COLLABORATOR RepositoryAffiliation = "COLLABORATOR"

	// Repositories that the user has access to through being a member of an organization. This includes every repository on every team that the user is on.
	RepositoryAffiliation_ORGANIZATION_MEMBER RepositoryAffiliation = "ORGANIZATION_MEMBER"
)

// The possible states of a pull request.
type PullRequestState string

const (
	// A pull request that is still open.
	PullRequestState_OPEN PullRequestState = "OPEN"

	// A pull request that has been closed without being merged.
	PullRequestState_CLOSED PullRequestState = "CLOSED"

	// A pull request that has been closed by being merged.
	PullRequestState_MERGED PullRequestState = "MERGED"
)

// Whether or not a PullRequest can be merged.
type MergeableState string

const (
	// The pull request can be merged.
	MergeableState_MERGEABLE MergeableState = "MERGEABLE"

	// The pull request cannot be merged due to merge conflicts.
	MergeableState_CONFLICTING MergeableState = "CONFLICTING"

	// The mergeability of the pull request is still being calculated.
	MergeableState_UNKNOWN MergeableState = "UNKNOWN"
)

// The possible PubSub channels for an issue.
type IssuePubSubTopic string

const (
	// The channel ID for observing issue updates.
	IssuePubSubTopic_UPDATED IssuePubSubTopic = "UPDATED"

	// The channel ID for marking an issue as read.
	IssuePubSubTopic_MARKASREAD IssuePubSubTopic = "MARKASREAD"
)

// The possible states of a pull request review.
type PullRequestReviewState string

const (
	// A review that has not yet been submitted.
	PullRequestReviewState_PENDING PullRequestReviewState = "PENDING"

	// An informational review.
	PullRequestReviewState_COMMENTED PullRequestReviewState = "COMMENTED"

	// A review allowing the pull request to merge.
	PullRequestReviewState_APPROVED PullRequestReviewState = "APPROVED"

	// A review blocking the pull request from merging.
	PullRequestReviewState_CHANGES_REQUESTED PullRequestReviewState = "CHANGES_REQUESTED"

	// A review that has been dismissed.
	PullRequestReviewState_DISMISSED PullRequestReviewState = "DISMISSED"
)

// The possible PubSub channels for a pull request.
type PullRequestPubSubTopic string

const (
	// The channel ID for observing pull request updates.
	PullRequestPubSubTopic_UPDATED PullRequestPubSubTopic = "UPDATED"

	// The channel ID for marking an pull request as read.
	PullRequestPubSubTopic_MARKASREAD PullRequestPubSubTopic = "MARKASREAD"

	// The channel ID for observing head ref updates.
	PullRequestPubSubTopic_HEAD_REF PullRequestPubSubTopic = "HEAD_REF"
)

// The possible team privacy values.
type TeamPrivacy string

const (
	// A secret team can only be seen by its members.
	TeamPrivacy_SECRET TeamPrivacy = "SECRET"

	// A visible team can be seen and @mentioned by every member of the organization.
	TeamPrivacy_VISIBLE TeamPrivacy = "VISIBLE"
)

// Properties by which user connections can be ordered.
type UserOrderField string

const (
	// Allows ordering a list of users by their login.
	UserOrderField_LOGIN UserOrderField = "LOGIN"

	// Allows ordering a list of users by their ability action
	UserOrderField_ACTION UserOrderField = "ACTION"
)

// The possible organization invitation roles.
type OrganizationInvitationRole string

const (
	// The user is invited to be a direct member of the organization.
	OrganizationInvitationRole_DIRECT_MEMBER OrganizationInvitationRole = "DIRECT_MEMBER"

	// The user is invited to be an admin of the organization.
	OrganizationInvitationRole_ADMIN OrganizationInvitationRole = "ADMIN"

	// The user is invited to be a billing manager of the organization.
	OrganizationInvitationRole_BILLING_MANAGER OrganizationInvitationRole = "BILLING_MANAGER"

	// The user's previous role will be reinstated.
	OrganizationInvitationRole_REINSTATE OrganizationInvitationRole = "REINSTATE"
)

// Properties by which team connections can be ordered.
type TeamOrderField string

const (
	// Allows ordering a list of teams by name.
	TeamOrderField_NAME TeamOrderField = "NAME"
)

// The possible default permissions for organization-owned repositories.
type DefaultRepositoryPermissionField string

const (
	// Members have read access to org repos by default
	DefaultRepositoryPermissionField_READ DefaultRepositoryPermissionField = "READ"

	// Members have read and write access to org repos by default
	DefaultRepositoryPermissionField_WRITE DefaultRepositoryPermissionField = "WRITE"

	// Members have read, write, and admin access to org repos by default
	DefaultRepositoryPermissionField_ADMIN DefaultRepositoryPermissionField = "ADMIN"
)

// The role of a user on a team.
type TeamRole string

const (
	// User has admin rights on the team.
	TeamRole_ADMIN TeamRole = "ADMIN"

	// User is a member of the team.
	TeamRole_MEMBER TeamRole = "MEMBER"
)

// The possible states for a deployment status.
type DeploymentStatusState string

const (
	// The deployment is pending.
	DeploymentStatusState_PENDING DeploymentStatusState = "PENDING"

	// The deployment was successful.
	DeploymentStatusState_SUCCESS DeploymentStatusState = "SUCCESS"

	// The deployment has failed.
	DeploymentStatusState_FAILURE DeploymentStatusState = "FAILURE"

	// The deployment is inactive.
	DeploymentStatusState_INACTIVE DeploymentStatusState = "INACTIVE"

	// The deployment experienced an error.
	DeploymentStatusState_ERROR DeploymentStatusState = "ERROR"
)

// The possible states in which a deployment can be.
type DeploymentState string

const (
	// The pending deployment was not updated after 30 minutes.
	DeploymentState_ABANDONED DeploymentState = "ABANDONED"

	// The deployment is currently active.
	DeploymentState_ACTIVE DeploymentState = "ACTIVE"

	// An inactive transient deployment.
	DeploymentState_DESTROYED DeploymentState = "DESTROYED"

	// The deployment experienced an error.
	DeploymentState_ERROR DeploymentState = "ERROR"

	// The deployment has failed.
	DeploymentState_FAILURE DeploymentState = "FAILURE"

	// The deployment is inactive.
	DeploymentState_INACTIVE DeploymentState = "INACTIVE"

	// The deployment is pending.
	DeploymentState_PENDING DeploymentState = "PENDING"
)

// Properties by which star connections can be ordered.
type StarOrderField string

const (
	// Allows ordering a list of stars by when they were created.
	StarOrderField_STARRED_AT StarOrderField = "STARRED_AT"
)

// The privacy of a Gist
type GistPrivacy string

const (
	// Public
	GistPrivacy_PUBLIC GistPrivacy = "PUBLIC"

	// Secret
	GistPrivacy_SECRET GistPrivacy = "SECRET"

	// Gists that are public and secret
	GistPrivacy_ALL GistPrivacy = "ALL"
)

// The possible states of a milestone.
type MilestoneState string

const (
	// A milestone that is still open.
	MilestoneState_OPEN MilestoneState = "OPEN"

	// A milestone that has been closed.
	MilestoneState_CLOSED MilestoneState = "CLOSED"
)

// The possible reasons a given repsitory could be in a locked state.
type RepositoryLockReason string

const (
	// The repository is locked due to a move.
	RepositoryLockReason_MOVING RepositoryLockReason = "MOVING"

	// The repository is locked due to a billing related reason.
	RepositoryLockReason_BILLING RepositoryLockReason = "BILLING"

	// The repository is locked due to a rename.
	RepositoryLockReason_RENAME RepositoryLockReason = "RENAME"

	// The repository is locked due to a migration.
	RepositoryLockReason_MIGRATING RepositoryLockReason = "MIGRATING"
)

// The affiliation type between collaborator and repository.
type RepositoryCollaboratorAffiliation string

const (
	// All collaborators of the repository.
	RepositoryCollaboratorAffiliation_ALL RepositoryCollaboratorAffiliation = "ALL"

	// All outside collaborators of an organization-owned repository.
	RepositoryCollaboratorAffiliation_OUTSIDE RepositoryCollaboratorAffiliation = "OUTSIDE"
)

// Properties by which language connections can be ordered.
type LanguageOrderField string

const (
	// Order languages by the size of all files containing the language
	LanguageOrderField_SIZE LanguageOrderField = "SIZE"
)

// Represents the individual results of a search.
type SearchType string

const (
	// Returns results matching issues in repositories.
	SearchType_ISSUE SearchType = "ISSUE"

	// Returns results matching repositories.
	SearchType_REPOSITORY SearchType = "REPOSITORY"

	// Returns results matching users on GitHub.
	SearchType_USER SearchType = "USER"
)

// The possible events to perform on a pull request review.
type PullRequestReviewEvent string

const (
	// Submit general feedback without explicit approval.
	PullRequestReviewEvent_COMMENT PullRequestReviewEvent = "COMMENT"

	// Submit feedback and approve merging these changes.
	PullRequestReviewEvent_APPROVE PullRequestReviewEvent = "APPROVE"

	// Submit feedback that must be addressed before merging.
	PullRequestReviewEvent_REQUEST_CHANGES PullRequestReviewEvent = "REQUEST_CHANGES"

	// Dismiss review so it now longer effects merging.
	PullRequestReviewEvent_DISMISS PullRequestReviewEvent = "DISMISS"
)

// Reason that the suggested topic is declined.
type TopicSuggestionDeclineReason string

const (
	// The suggested topic is not relevant to the repository.
	TopicSuggestionDeclineReason_NOT_RELEVANT TopicSuggestionDeclineReason = "NOT_RELEVANT"

	// The suggested topic is too specific for the repository (e.g. #ruby-on-rails-version-4-2-1).
	TopicSuggestionDeclineReason_TOO_SPECIFIC TopicSuggestionDeclineReason = "TOO_SPECIFIC"

	// The viewer does not like the suggested topic.
	TopicSuggestionDeclineReason_PERSONAL_PREFERENCE TopicSuggestionDeclineReason = "PERSONAL_PREFERENCE"

	// The suggested topic is too general for the repository.
	TopicSuggestionDeclineReason_TOO_GENERAL TopicSuggestionDeclineReason = "TOO_GENERAL"
)

// An enum describing what kind of type a given `__Type` is.
type __TypeKind string

const (
	// Indicates this type is a scalar.
	__TypeKind_SCALAR __TypeKind = "SCALAR"

	// Indicates this type is an object. `fields` and `interfaces` are valid fields.
	__TypeKind_OBJECT __TypeKind = "OBJECT"

	// Indicates this type is an interface. `fields` and `possibleTypes` are valid fields.
	__TypeKind_INTERFACE __TypeKind = "INTERFACE"

	// Indicates this type is a union. `possibleTypes` is a valid field.
	__TypeKind_UNION __TypeKind = "UNION"

	// Indicates this type is an enum. `enumValues` is a valid field.
	__TypeKind_ENUM __TypeKind = "ENUM"

	// Indicates this type is an input object. `inputFields` is a valid field.
	__TypeKind_INPUT_OBJECT __TypeKind = "INPUT_OBJECT"

	// Indicates this type is a list. `ofType` is a valid field.
	__TypeKind_LIST __TypeKind = "LIST"

	// Indicates this type is a non-null. `ofType` is a valid field.
	__TypeKind_NON_NULL __TypeKind = "NON_NULL"
)

// A Directive can be adjacent to many parts of the GraphQL language, a __DirectiveLocation describes one such possible adjacencies.
type __DirectiveLocation string

const (
	// Location adjacent to a query operation.
	__DirectiveLocation_QUERY __DirectiveLocation = "QUERY"

	// Location adjacent to a mutation operation.
	__DirectiveLocation_MUTATION __DirectiveLocation = "MUTATION"

	// Location adjacent to a subscription operation.
	__DirectiveLocation_SUBSCRIPTION __DirectiveLocation = "SUBSCRIPTION"

	// Location adjacent to a field.
	__DirectiveLocation_FIELD __DirectiveLocation = "FIELD"

	// Location adjacent to a fragment definition.
	__DirectiveLocation_FRAGMENT_DEFINITION __DirectiveLocation = "FRAGMENT_DEFINITION"

	// Location adjacent to a fragment spread.
	__DirectiveLocation_FRAGMENT_SPREAD __DirectiveLocation = "FRAGMENT_SPREAD"

	// Location adjacent to an inline fragment.
	__DirectiveLocation_INLINE_FRAGMENT __DirectiveLocation = "INLINE_FRAGMENT"

	// Location adjacent to a schema definition.
	__DirectiveLocation_SCHEMA __DirectiveLocation = "SCHEMA"

	// Location adjacent to a scalar definition.
	__DirectiveLocation_SCALAR __DirectiveLocation = "SCALAR"

	// Location adjacent to an object type definition.
	__DirectiveLocation_OBJECT __DirectiveLocation = "OBJECT"

	// Location adjacent to a field definition.
	__DirectiveLocation_FIELD_DEFINITION __DirectiveLocation = "FIELD_DEFINITION"

	// Location adjacent to an argument definition.
	__DirectiveLocation_ARGUMENT_DEFINITION __DirectiveLocation = "ARGUMENT_DEFINITION"

	// Location adjacent to an interface definition.
	__DirectiveLocation_INTERFACE __DirectiveLocation = "INTERFACE"

	// Location adjacent to a union definition.
	__DirectiveLocation_UNION __DirectiveLocation = "UNION"

	// Location adjacent to an enum definition.
	__DirectiveLocation_ENUM __DirectiveLocation = "ENUM"

	// Location adjacent to an enum value definition.
	__DirectiveLocation_ENUM_VALUE __DirectiveLocation = "ENUM_VALUE"

	// Location adjacent to an input object type definition.
	__DirectiveLocation_INPUT_OBJECT __DirectiveLocation = "INPUT_OBJECT"

	// Location adjacent to an input object field definition.
	__DirectiveLocation_INPUT_FIELD_DEFINITION __DirectiveLocation = "INPUT_FIELD_DEFINITION"
)
