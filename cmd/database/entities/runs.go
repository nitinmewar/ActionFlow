package entities

import (
	"time"
)

const (
	StatusQueued             = "queued"
	StatusInProgress         = "in_progress"
	StatusCompleted          = "completed"
	ConclusionSuccess        = "success"
	ConclusionFailure        = "failure"
	ConclusionNeutral        = "neutral"
	ConclusionCancelled      = "cancelled"
	ConclusionSkipped        = "skipped"
	ConclusionTimedout       = "timed_out"
	ConclusionActionRequired = "action_required"
	ConclusionStale          = "stale"
)

// WorkflowRun represents a GitHub Actions workflow run
type WorkflowRun struct {
	ID                   uint       `gorm:"primaryKey;autoIncrement"`
	RunID                int64      `gorm:"column:run_id;uniqueIndex;not null"` // GitHub's unique run ID
	WorkflowName         string     `gorm:"column:workflow_name;not null;index"`
	WorkflowID           int64      `gorm:"column:workflow_id;not null"`
	Repository           string     `gorm:"column:repository;not null;index"` // format: owner/repo
	HeadBranch           string     `gorm:"column:head_branch;index"`
	HeadSHA              string     `gorm:"column:head_sha;not null"`
	HeadSHAShort         string     `gorm:"column:head_sha_short"`
	DisplayTitle         string     `gorm:"column:display_title"`
	WorkflowPath         string     `gorm:"column:workflow_path"`
	CheckSuiteID         int64      `gorm:"column:check_suite_id"`
	RunNumber            int        `gorm:"column:run_number;not null"`
	RunAttempt           int        `gorm:"column:run_attempt;default:1"`
	Event                string     `gorm:"column:event;not null"`        // push, pull_request, etc.
	Status               string     `gorm:"column:status;not null;index"` // queued, in_progress, completed
	Conclusion           string     `gorm:"column:conclusion"`            // success, failure, etc.
	ActorLogin           string     `gorm:"column:actor_login"`
	TriggeringActorLogin string     `gorm:"column:triggering_actor_login"`
	GitHubURL            string     `gorm:"column:github_url"`
	RunStartedAt         time.Time  `gorm:"column:run_started_at;not null;index"`
	CreatedAt            time.Time  `gorm:"column:created_at;not null"`
	UpdatedAt            time.Time  `gorm:"column:updated_at;not null"`
	CompletedAt          *time.Time `gorm:"column:completed_at"`
	Duration             *int64     `gorm:"column:duration"` // in seconds

	// Commit info (flattened)
	CommitMessage     string    `gorm:"column:commit_message"`
	CommitTimestamp   time.Time `gorm:"column:commit_timestamp"`
	CommitAuthorName  string    `gorm:"column:commit_author_name"`
	CommitAuthorEmail string    `gorm:"column:commit_author_email"`

	CreatedAtDB time.Time `gorm:"column:created_at_db;autoCreateTime"`
	UpdatedAtDB time.Time `gorm:"column:updated_at_db;autoUpdateTime"`
}

func (WorkflowRun) TableName() string {
	return "workflow_runs"
}
