import { useState } from "react";
import { WorkflowRun } from "@/types/workflow";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { StatusBadge } from "./StatusBadge";
import { Button } from "@/components/ui/button";
import { ExternalLink, ArrowUpDown, GitBranch, User } from "lucide-react";
import { formatDistanceToNow } from "date-fns";

interface WorkflowRunsTableProps {
  runs: WorkflowRun[];
  onRunClick: (run: WorkflowRun) => void;
}

export const WorkflowRunsTable = ({ runs, onRunClick }: WorkflowRunsTableProps) => {
  const [sortBy, setSortBy] = useState<"time" | "status">("time");
  const [sortOrder, setSortOrder] = useState<"asc" | "desc">("desc");

  const handleSort = (column: "time" | "status") => {
    if (sortBy === column) {
      setSortOrder(sortOrder === "asc" ? "desc" : "asc");
    } else {
      setSortBy(column);
      setSortOrder("desc");
    }
  };

  const sortedRuns = [...runs].sort((a, b) => {
    if (sortBy === "time") {
      const timeA = new Date(a.run_started_at).getTime();
      const timeB = new Date(b.run_started_at).getTime();
      return sortOrder === "asc" ? timeA - timeB : timeB - timeA;
    } else {
      const statusOrder = { queued: 0, in_progress: 1, completed: 2 };
      const statusA = statusOrder[a.status];
      const statusB = statusOrder[b.status];
      return sortOrder === "asc" ? statusA - statusB : statusB - statusA;
    }
  });

  const formatDuration = (seconds: number) => {
    if (seconds === 0) return "-";
    const minutes = Math.floor(seconds / 60);
    const secs = seconds % 60;
    return `${minutes}m ${secs}s`;
  };

  return (
    <div className="border rounded-lg overflow-hidden">
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[200px]">Workflow</TableHead>
            <TableHead className="w-[180px]">Repository</TableHead>
            <TableHead className="w-[120px]">
              <Button
                variant="ghost"
                size="sm"
                className="h-8 px-2 flex items-center gap-1"
                onClick={() => handleSort("status")}
              >
                Status
                <ArrowUpDown className="h-3.5 w-3.5" />
              </Button>
            </TableHead>
            <TableHead className="w-[120px]">Branch</TableHead>
            <TableHead className="w-[140px]">Actor</TableHead>
            <TableHead className="w-[140px]">
              <Button
                variant="ghost"
                size="sm"
                className="h-8 px-2 flex items-center gap-1"
                onClick={() => handleSort("time")}
              >
                Started
                <ArrowUpDown className="h-3.5 w-3.5" />
              </Button>
            </TableHead>
            <TableHead className="w-[100px]">Duration</TableHead>
            <TableHead className="w-[50px]"></TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {sortedRuns.map((run) => (
            <TableRow
              key={run.run_id}
              className="cursor-pointer hover:bg-muted/50"
              onClick={() => onRunClick(run)}
            >
              <TableCell className="font-medium">{run.workflow_name}</TableCell>
              <TableCell className="text-muted-foreground">{run.repository}</TableCell>
              <TableCell>
                <StatusBadge status={run.status} conclusion={run.conclusion} />
              </TableCell>
              <TableCell>
                <div className="flex items-center gap-1.5 text-muted-foreground">
                  <GitBranch className="h-3.5 w-3.5" />
                  <span className="text-sm">{run.head_branch}</span>
                </div>
              </TableCell>
              <TableCell>
                <div className="flex items-center gap-1.5 text-muted-foreground">
                  <User className="h-3.5 w-3.5" />
                  <span className="text-sm">{run.actor_login}</span>
                </div>
              </TableCell>
              <TableCell className="text-sm text-muted-foreground">
                {formatDistanceToNow(new Date(run.run_started_at), { addSuffix: true })}
              </TableCell>
              <TableCell className="text-sm text-muted-foreground">
                {formatDuration(run.duration)}
              </TableCell>
              <TableCell>
                <Button
                  variant="ghost"
                  size="sm"
                  className="h-8 w-8 p-0"
                  onClick={(e) => {
                    e.stopPropagation();
                    window.open(run.github_url, "_blank");
                  }}
                >
                  <ExternalLink className="h-4 w-4" />
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
};
