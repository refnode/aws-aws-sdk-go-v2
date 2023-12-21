// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ApprovalState string

// Enum values for ApprovalState
const (
	ApprovalStateApprove ApprovalState = "APPROVE"
	ApprovalStateRevoke  ApprovalState = "REVOKE"
)

// Values returns all known values for ApprovalState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ApprovalState) Values() []ApprovalState {
	return []ApprovalState{
		"APPROVE",
		"REVOKE",
	}
}

type BatchGetRepositoriesErrorCodeEnum string

// Enum values for BatchGetRepositoriesErrorCodeEnum
const (
	BatchGetRepositoriesErrorCodeEnumEncryptionIntegrityChecksFailedException BatchGetRepositoriesErrorCodeEnum = "EncryptionIntegrityChecksFailedException"
	BatchGetRepositoriesErrorCodeEnumEncryptionKeyAccessDeniedException       BatchGetRepositoriesErrorCodeEnum = "EncryptionKeyAccessDeniedException"
	BatchGetRepositoriesErrorCodeEnumEncryptionKeyDisabledException           BatchGetRepositoriesErrorCodeEnum = "EncryptionKeyDisabledException"
	BatchGetRepositoriesErrorCodeEnumEncryptionKeyNotFoundException           BatchGetRepositoriesErrorCodeEnum = "EncryptionKeyNotFoundException"
	BatchGetRepositoriesErrorCodeEnumEncryptionKeyUnavailableException        BatchGetRepositoriesErrorCodeEnum = "EncryptionKeyUnavailableException"
	BatchGetRepositoriesErrorCodeEnumRepositoryDoesNotExistException          BatchGetRepositoriesErrorCodeEnum = "RepositoryDoesNotExistException"
)

// Values returns all known values for BatchGetRepositoriesErrorCodeEnum. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (BatchGetRepositoriesErrorCodeEnum) Values() []BatchGetRepositoriesErrorCodeEnum {
	return []BatchGetRepositoriesErrorCodeEnum{
		"EncryptionIntegrityChecksFailedException",
		"EncryptionKeyAccessDeniedException",
		"EncryptionKeyDisabledException",
		"EncryptionKeyNotFoundException",
		"EncryptionKeyUnavailableException",
		"RepositoryDoesNotExistException",
	}
}

type ChangeTypeEnum string

// Enum values for ChangeTypeEnum
const (
	ChangeTypeEnumAdded    ChangeTypeEnum = "A"
	ChangeTypeEnumModified ChangeTypeEnum = "M"
	ChangeTypeEnumDeleted  ChangeTypeEnum = "D"
)

// Values returns all known values for ChangeTypeEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ChangeTypeEnum) Values() []ChangeTypeEnum {
	return []ChangeTypeEnum{
		"A",
		"M",
		"D",
	}
}

type ConflictDetailLevelTypeEnum string

// Enum values for ConflictDetailLevelTypeEnum
const (
	ConflictDetailLevelTypeEnumFileLevel ConflictDetailLevelTypeEnum = "FILE_LEVEL"
	ConflictDetailLevelTypeEnumLineLevel ConflictDetailLevelTypeEnum = "LINE_LEVEL"
)

// Values returns all known values for ConflictDetailLevelTypeEnum. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConflictDetailLevelTypeEnum) Values() []ConflictDetailLevelTypeEnum {
	return []ConflictDetailLevelTypeEnum{
		"FILE_LEVEL",
		"LINE_LEVEL",
	}
}

type ConflictResolutionStrategyTypeEnum string

// Enum values for ConflictResolutionStrategyTypeEnum
const (
	ConflictResolutionStrategyTypeEnumNone              ConflictResolutionStrategyTypeEnum = "NONE"
	ConflictResolutionStrategyTypeEnumAcceptSource      ConflictResolutionStrategyTypeEnum = "ACCEPT_SOURCE"
	ConflictResolutionStrategyTypeEnumAcceptDestination ConflictResolutionStrategyTypeEnum = "ACCEPT_DESTINATION"
	ConflictResolutionStrategyTypeEnumAutomerge         ConflictResolutionStrategyTypeEnum = "AUTOMERGE"
)

// Values returns all known values for ConflictResolutionStrategyTypeEnum. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ConflictResolutionStrategyTypeEnum) Values() []ConflictResolutionStrategyTypeEnum {
	return []ConflictResolutionStrategyTypeEnum{
		"NONE",
		"ACCEPT_SOURCE",
		"ACCEPT_DESTINATION",
		"AUTOMERGE",
	}
}

type FileModeTypeEnum string

// Enum values for FileModeTypeEnum
const (
	FileModeTypeEnumExecutable FileModeTypeEnum = "EXECUTABLE"
	FileModeTypeEnumNormal     FileModeTypeEnum = "NORMAL"
	FileModeTypeEnumSymlink    FileModeTypeEnum = "SYMLINK"
)

// Values returns all known values for FileModeTypeEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FileModeTypeEnum) Values() []FileModeTypeEnum {
	return []FileModeTypeEnum{
		"EXECUTABLE",
		"NORMAL",
		"SYMLINK",
	}
}

type MergeOptionTypeEnum string

// Enum values for MergeOptionTypeEnum
const (
	MergeOptionTypeEnumFastForwardMerge MergeOptionTypeEnum = "FAST_FORWARD_MERGE"
	MergeOptionTypeEnumSquashMerge      MergeOptionTypeEnum = "SQUASH_MERGE"
	MergeOptionTypeEnumThreeWayMerge    MergeOptionTypeEnum = "THREE_WAY_MERGE"
)

// Values returns all known values for MergeOptionTypeEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MergeOptionTypeEnum) Values() []MergeOptionTypeEnum {
	return []MergeOptionTypeEnum{
		"FAST_FORWARD_MERGE",
		"SQUASH_MERGE",
		"THREE_WAY_MERGE",
	}
}

type ObjectTypeEnum string

// Enum values for ObjectTypeEnum
const (
	ObjectTypeEnumFile         ObjectTypeEnum = "FILE"
	ObjectTypeEnumDirectory    ObjectTypeEnum = "DIRECTORY"
	ObjectTypeEnumGitLink      ObjectTypeEnum = "GIT_LINK"
	ObjectTypeEnumSymbolicLink ObjectTypeEnum = "SYMBOLIC_LINK"
)

// Values returns all known values for ObjectTypeEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ObjectTypeEnum) Values() []ObjectTypeEnum {
	return []ObjectTypeEnum{
		"FILE",
		"DIRECTORY",
		"GIT_LINK",
		"SYMBOLIC_LINK",
	}
}

type OrderEnum string

// Enum values for OrderEnum
const (
	OrderEnumAscending  OrderEnum = "ascending"
	OrderEnumDescending OrderEnum = "descending"
)

// Values returns all known values for OrderEnum. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OrderEnum) Values() []OrderEnum {
	return []OrderEnum{
		"ascending",
		"descending",
	}
}

type OverrideStatus string

// Enum values for OverrideStatus
const (
	OverrideStatusOverride OverrideStatus = "OVERRIDE"
	OverrideStatusRevoke   OverrideStatus = "REVOKE"
)

// Values returns all known values for OverrideStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OverrideStatus) Values() []OverrideStatus {
	return []OverrideStatus{
		"OVERRIDE",
		"REVOKE",
	}
}

type PullRequestEventType string

// Enum values for PullRequestEventType
const (
	PullRequestEventTypePullRequestCreated                PullRequestEventType = "PULL_REQUEST_CREATED"
	PullRequestEventTypePullRequestStatusChanged          PullRequestEventType = "PULL_REQUEST_STATUS_CHANGED"
	PullRequestEventTypePullRequestSourceReferenceUpdated PullRequestEventType = "PULL_REQUEST_SOURCE_REFERENCE_UPDATED"
	PullRequestEventTypePullRequestMergeStateChanged      PullRequestEventType = "PULL_REQUEST_MERGE_STATE_CHANGED"
	PullRequestEventTypePullRequestApprovalRuleCreated    PullRequestEventType = "PULL_REQUEST_APPROVAL_RULE_CREATED"
	PullRequestEventTypePullRequestApprovalRuleUpdated    PullRequestEventType = "PULL_REQUEST_APPROVAL_RULE_UPDATED"
	PullRequestEventTypePullRequestApprovalRuleDeleted    PullRequestEventType = "PULL_REQUEST_APPROVAL_RULE_DELETED"
	PullRequestEventTypePullRequestApprovalRuleOverridden PullRequestEventType = "PULL_REQUEST_APPROVAL_RULE_OVERRIDDEN"
	PullRequestEventTypePullRequestApprovalStateChanged   PullRequestEventType = "PULL_REQUEST_APPROVAL_STATE_CHANGED"
)

// Values returns all known values for PullRequestEventType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PullRequestEventType) Values() []PullRequestEventType {
	return []PullRequestEventType{
		"PULL_REQUEST_CREATED",
		"PULL_REQUEST_STATUS_CHANGED",
		"PULL_REQUEST_SOURCE_REFERENCE_UPDATED",
		"PULL_REQUEST_MERGE_STATE_CHANGED",
		"PULL_REQUEST_APPROVAL_RULE_CREATED",
		"PULL_REQUEST_APPROVAL_RULE_UPDATED",
		"PULL_REQUEST_APPROVAL_RULE_DELETED",
		"PULL_REQUEST_APPROVAL_RULE_OVERRIDDEN",
		"PULL_REQUEST_APPROVAL_STATE_CHANGED",
	}
}

type PullRequestStatusEnum string

// Enum values for PullRequestStatusEnum
const (
	PullRequestStatusEnumOpen   PullRequestStatusEnum = "OPEN"
	PullRequestStatusEnumClosed PullRequestStatusEnum = "CLOSED"
)

// Values returns all known values for PullRequestStatusEnum. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PullRequestStatusEnum) Values() []PullRequestStatusEnum {
	return []PullRequestStatusEnum{
		"OPEN",
		"CLOSED",
	}
}

type RelativeFileVersionEnum string

// Enum values for RelativeFileVersionEnum
const (
	RelativeFileVersionEnumBefore RelativeFileVersionEnum = "BEFORE"
	RelativeFileVersionEnumAfter  RelativeFileVersionEnum = "AFTER"
)

// Values returns all known values for RelativeFileVersionEnum. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RelativeFileVersionEnum) Values() []RelativeFileVersionEnum {
	return []RelativeFileVersionEnum{
		"BEFORE",
		"AFTER",
	}
}

type ReplacementTypeEnum string

// Enum values for ReplacementTypeEnum
const (
	ReplacementTypeEnumKeepBase        ReplacementTypeEnum = "KEEP_BASE"
	ReplacementTypeEnumKeepSource      ReplacementTypeEnum = "KEEP_SOURCE"
	ReplacementTypeEnumKeepDestination ReplacementTypeEnum = "KEEP_DESTINATION"
	ReplacementTypeEnumUseNewContent   ReplacementTypeEnum = "USE_NEW_CONTENT"
)

// Values returns all known values for ReplacementTypeEnum. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReplacementTypeEnum) Values() []ReplacementTypeEnum {
	return []ReplacementTypeEnum{
		"KEEP_BASE",
		"KEEP_SOURCE",
		"KEEP_DESTINATION",
		"USE_NEW_CONTENT",
	}
}

type RepositoryTriggerEventEnum string

// Enum values for RepositoryTriggerEventEnum
const (
	RepositoryTriggerEventEnumAll             RepositoryTriggerEventEnum = "all"
	RepositoryTriggerEventEnumUpdateReference RepositoryTriggerEventEnum = "updateReference"
	RepositoryTriggerEventEnumCreateReference RepositoryTriggerEventEnum = "createReference"
	RepositoryTriggerEventEnumDeleteReference RepositoryTriggerEventEnum = "deleteReference"
)

// Values returns all known values for RepositoryTriggerEventEnum. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (RepositoryTriggerEventEnum) Values() []RepositoryTriggerEventEnum {
	return []RepositoryTriggerEventEnum{
		"all",
		"updateReference",
		"createReference",
		"deleteReference",
	}
}

type SortByEnum string

// Enum values for SortByEnum
const (
	SortByEnumRepositoryName SortByEnum = "repositoryName"
	SortByEnumModifiedDate   SortByEnum = "lastModifiedDate"
)

// Values returns all known values for SortByEnum. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SortByEnum) Values() []SortByEnum {
	return []SortByEnum{
		"repositoryName",
		"lastModifiedDate",
	}
}
