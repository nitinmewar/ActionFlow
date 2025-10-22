package models

import "time"

type WorkflowRunSummary struct {
	ID               uint       `json:"id"`
	RunID            int64      `json:"run_id"`
	WorkflowName     string     `json:"workflow_name"`
	Repository       string     `json:"repository"`
	HeadBranch       string     `json:"head_branch"`
	HeadSHAShort     string     `json:"head_sha_short"`
	DisplayTitle     string     `json:"display_title"`
	Event            string     `json:"event"`
	Status           string     `json:"status"`
	Conclusion       string     `json:"conclusion"`
	ActorLogin       string     `json:"actor_login"`
	RunNumber        int        `json:"run_number"`
	RunStartedAt     time.Time  `json:"run_started_at"`
	CompletedAt      *time.Time `json:"completed_at"`
	Duration         *int64     `json:"duration"`
	CommitMessage    string     `json:"commit_message"`
	CommitAuthorName string     `json:"commit_author_name"`
}
