package models

import (
	"time"
)

// GitHubWebhookPayload is the top-level webhook payload
// GitHubWebhookPayload represents the complete webhook payload
type GitHubWebhookPayload struct {
	Action      string              `json:"action"`
	WorkflowRun *WorkflowRunWebhook `json:"workflow_run"`
	Repository  Repository          `json:"repository"`
	Workflow    Workflow            `json:"workflow"`
	Sender      User                `json:"sender"`
}

// WorkflowRunWebhook represents the workflow_run object from GitHub
type WorkflowRunWebhook struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	HeadBranch      string    `json:"head_branch"`
	HeadSHA         string    `json:"head_sha"`
	Path            string    `json:"path"`
	DisplayTitle    string    `json:"display_title"`
	RunNumber       int       `json:"run_number"`
	Event           string    `json:"event"`
	Status          string    `json:"status"`
	Conclusion      string    `json:"conclusion"`
	WorkflowID      int64     `json:"workflow_id"`
	CheckSuiteID    int64     `json:"check_suite_id"`
	Actor           User      `json:"actor"`
	RunAttempt      int       `json:"run_attempt"`
	RunStartedAt    time.Time `json:"run_started_at"`
	TriggeringActor User      `json:"triggering_actor"`
	HTMLURL         string    `json:"html_url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	HeadCommit      Commit    `json:"head_commit"`
}

type Repository struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
}

type Workflow struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

type User struct {
	Login string `json:"login"`
}

type Commit struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Author    Author    `json:"author"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
