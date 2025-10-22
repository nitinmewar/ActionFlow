export interface WorkflowRun {
  ID: number;
  RunID: number;
  WorkflowName: string;
  WorkflowID: number;
  Repository: string;
  HeadBranch: string;
  HeadSHA: string;
  HeadSHAShort: string;
  DisplayTitle: string;
  WorkflowPath: string;
  CheckSuiteID: number;
  RunNumber: number;
  RunAttempt: number;
  Event: string;
  Status: "completed" | "in_progress" | "queued" | "waiting";
  Conclusion: "success" | "failure" | "cancelled" | "skipped" | null;
  ActorLogin: string;
  TriggeringActorLogin: string;
  GitHubURL: string;
  RunStartedAt: string;
  CreatedAt: string;
  UpdatedAt: string;
  CompletedAt: string | null;
  Duration: number;
  CommitMessage: string;
  CommitTimestamp: string;
  CommitAuthorName: string;
  CommitAuthorEmail: string;
  CreatedAtDB: string;
  UpdatedAtDB: string;
}

export interface WorkflowMetrics {
  totalRuns: number;
  successRate: number;
  failures: number;
  avgDuration: number;
}

export interface ApiResponse<T> {
  data: T;
  pagination?: {
    page: number;
    page_size: number;
    pages: number;
    total: number;
  };
}