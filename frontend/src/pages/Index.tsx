import { useState, useEffect, useMemo } from "react";
import { WorkflowRun, WorkflowMetrics } from "@/types/workflow";
import { workflowAPI } from "@/lib/api";
import { MetricsCard } from "@/components/MetricsCard";
import { FilterBar } from "@/components/FilterBar";
import { WorkflowRunsTable } from "@/components/WorkflowRunsTable";
import { RunDetailModal } from "@/components/RunDetailModal";
import { Activity, CheckCircle2, XCircle, Clock, RefreshCw } from "lucide-react";
import { toast } from "sonner";

const Index = () => {
  const [runs, setRuns] = useState<WorkflowRun[]>([]);
  const [selectedRun, setSelectedRun] = useState<WorkflowRun | null>(null);
  const [repositoryFilter, setRepositoryFilter] = useState("");
  const [statusFilter, setStatusFilter] = useState("all");
  const [branchFilter, setBranchFilter] = useState("all");
  const [loading, setLoading] = useState(true);
  const [refreshing, setRefreshing] = useState(false);

  // Fetch runs from backend
  const fetchRuns = async (showRefresh = false) => {
    try {
      if (showRefresh) {
        setRefreshing(true);
      } else {
        setLoading(true);
      }
      
      const data = await workflowAPI.getRuns();
      setRuns(data);
    } catch (error) {
      console.error("Error fetching runs:", error);
      toast.error("Failed to fetch workflow runs");
    } finally {
      setLoading(false);
      setRefreshing(false);
    }
  };

  // Initial data fetch
  useEffect(() => {
    fetchRuns();
  }, []);

  // Simulate real-time updates for running workflows
  useEffect(() => {
    const interval = setInterval(() => {
      // Only refresh if there are running workflows
      const hasRunningWorkflows = runs.some(run => 
        run.Status === "in_progress" || run.Status === "queued"
      );
      
      if (hasRunningWorkflows) {
        fetchRuns(true);
      }
    }, 10000); // Update every 10 seconds

    return () => clearInterval(interval);
  }, [runs]);

  // Calculate metrics
  const metrics = useMemo<WorkflowMetrics>(() => {
    const completedRuns = runs.filter((run) => run.Status === "completed");
    const successfulRuns = completedRuns.filter((run) => run.Conclusion === "success");
    const failedRuns = completedRuns.filter((run) => run.Conclusion === "failure");
    const totalDuration = completedRuns.reduce((sum, run) => sum + run.Duration, 0);

    return {
      totalRuns: runs.length,
      successRate: completedRuns.length > 0 ? (successfulRuns.length / completedRuns.length) * 100 : 0,
      failures: failedRuns.length,
      avgDuration: completedRuns.length > 0 ? Math.round(totalDuration / completedRuns.length) : 0,
    };
  }, [runs]);

  // Get unique repositories and branches for filters
  const repositories = useMemo(() => [...new Set(runs.map((run) => run.Repository))], [runs]);
  const branches = useMemo(() => [...new Set(runs.map((run) => run.HeadBranch))], [runs]);

  // Filter runs
  const filteredRuns = useMemo(() => {
    return runs.filter((run) => {
      const matchesRepository = repositoryFilter === "" || 
        run.Repository.toLowerCase().includes(repositoryFilter.toLowerCase());
      const matchesStatus = statusFilter === "all" || 
        (statusFilter === "success" && run.Conclusion === "success") ||
        (statusFilter === "failure" && run.Conclusion === "failure") ||
        (statusFilter === "cancelled" && run.Conclusion === "cancelled") ||
        (statusFilter === "in_progress" && run.Status === "in_progress") ||
        (statusFilter === "queued" && run.Status === "queued");
      const matchesBranch = branchFilter === "all" || run.HeadBranch === branchFilter;

      return matchesRepository && matchesStatus && matchesBranch;
    });
  }, [runs, repositoryFilter, statusFilter, branchFilter]);

  // Handle run click - fetch detailed run data
  const handleRunClick = async (run: WorkflowRun) => {
    try {
      // You can either use the existing data or fetch fresh data
      // For now, using existing data. Uncomment below if you want fresh data:
      // const detailedRun = await workflowAPI.getRun(run.RunID);
      // setSelectedRun(detailedRun);
      setSelectedRun(run);
    } catch (error) {
      console.error("Error fetching run details:", error);
      toast.error("Failed to fetch run details");
    }
  };

  // Manual refresh
  const handleRefresh = () => {
    fetchRuns(true);
  };

  if (loading) {
    return (
      <div className="min-h-screen bg-background flex items-center justify-center">
        <div className="text-center">
          <RefreshCw className="h-8 w-8 animate-spin mx-auto mb-4" />
          <p className="text-muted-foreground">Loading workflow runs...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-background">
      <div className="container mx-auto py-8 px-4">
        <div className="flex justify-between items-center mb-8">
          <div>
            <h1 className="text-4xl font-bold mb-2">GitHub Actions Dashboard</h1>
            <p className="text-muted-foreground">Monitor and track your workflow runs in real-time</p>
          </div>
          <button
            onClick={handleRefresh}
            disabled={refreshing}
            className="flex items-center gap-2 px-4 py-2 bg-primary text-primary-foreground rounded-md hover:bg-primary/90 disabled:opacity-50"
          >
            <RefreshCw className={`h-4 w-4 ${refreshing ? 'animate-spin' : ''}`} />
            {refreshing ? 'Refreshing...' : 'Refresh'}
          </button>
        </div>

        {/* Metrics Cards */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
          <MetricsCard
            title="Total Runs"
            value={metrics.totalRuns}
            icon={Activity}
            description="All workflow runs"
          />
          <MetricsCard
            title="Success Rate"
            value={`${metrics.successRate.toFixed(1)}%`}
            icon={CheckCircle2}
            description="Completed successfully"
            trend={{ value: 5.2, isPositive: true }}
          />
          <MetricsCard
            title="Failures"
            value={metrics.failures}
            icon={XCircle}
            description="Failed runs"
            trend={{ value: 2.1, isPositive: false }}
          />
          <MetricsCard
            title="Avg Duration"
            value={`${Math.floor(metrics.avgDuration / 60)}m ${metrics.avgDuration % 60}s`}
            icon={Clock}
            description="Average completion time"
          />
        </div>

        {/* Filters */}
        <FilterBar
          onRepositoryChange={setRepositoryFilter}
          onStatusChange={setStatusFilter}
          onBranchChange={setBranchFilter}
          repositories={repositories}
          branches={branches}
        />

        {/* Workflow Runs Table */}
        {filteredRuns.length === 0 ? (
          <div className="text-center py-12">
            <p className="text-muted-foreground">No workflow runs found</p>
          </div>
        ) : (
          <WorkflowRunsTable runs={filteredRuns} onRunClick={handleRunClick} />
        )}

        {/* Run Detail Modal */}
        <RunDetailModal
          run={selectedRun}
          open={selectedRun !== null}
          onOpenChange={(open) => !open && setSelectedRun(null)}
        />
      </div>
    </div>
  );
};

export default Index;