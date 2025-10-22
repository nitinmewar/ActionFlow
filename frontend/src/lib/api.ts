import { WorkflowRun, ApiResponse } from "@/types/workflow";

const API_BASE_URL = "http://localhost:8080/api";

export const workflowAPI = {
  // Get all workflow runs
  async getRuns(): Promise<WorkflowRun[]> {
    const response = await fetch(`${API_BASE_URL}/runs`);
    if (!response.ok) {
      throw new Error(`Failed to fetch runs: ${response.statusText}`);
    }
    const data: ApiResponse<WorkflowRun[]> = await response.json();
    return data.data;
  },

  // Get a specific run by ID
  async getRun(runId: number): Promise<WorkflowRun> {
    const response = await fetch(`${API_BASE_URL}/runs/${runId}`);
    if (!response.ok) {
      throw new Error(`Failed to fetch run ${runId}: ${response.statusText}`);
    }
    const data: ApiResponse<WorkflowRun> = await response.json();
    return data.data;
  },
};