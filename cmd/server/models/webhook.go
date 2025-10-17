package models

type WorkflowRunPayload struct {
	Action      string      `json:"action"`
	WorkFlowRun WorkFlowRun `json:"workflow_run"`
	Repository  Repository  `json:"repository"`
}

type WorkFlowRun struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	Status     string     `json:"status"`
	Conclusion string     `json:"conclusion"`
	HeadBranch string     `json:"head_branch"`
	HeadSha    string     `json:"head_sha"`
	HtmlURL    string     `json:"html_url"`
	Event      string     `json:"event"`
	HeadCommit HeadCommit `json:"head_commit"`
	CreatedAt  string     `json:"created_at"`
	UpdatedAt  string     `json:"updated_at"`
}

type HeadCommit struct {
	Message string `json:"message"`
	Author  Author `json:"author"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Repository struct {
	FullName string `json:"full_name"`
}
