import { WorkflowRun } from "@/types/workflow";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { StatusBadge } from "./StatusBadge";
import { Button } from "@/components/ui/button";
import { ExternalLink, GitBranch, User, Clock, GitCommit } from "lucide-react";
import { format } from "date-fns";

interface RunDetailModalProps {
  run: WorkflowRun | null;
  open: boolean;
  onOpenChange: (open: boolean) => void;
}

export const RunDetailModal = ({ run, open, onOpenChange }: RunDetailModalProps) => {
  if (!run) return null;

  const formatDuration = (seconds: number) => {
    if (seconds === 0) return "Not started";
    const minutes = Math.floor(seconds / 60);
    const secs = seconds % 60;
    return `${minutes}m ${secs}s`;
  };

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="max-w-2xl">
        <DialogHeader>
          <DialogTitle className="flex items-center justify-between">
            <span>{run.workflow_name}</span>
            <StatusBadge status={run.status} conclusion={run.conclusion} />
          </DialogTitle>
          <DialogDescription>Run #{run.run_id}</DialogDescription>
        </DialogHeader>

        <div className="space-y-6 py-4">
          <div className="grid grid-cols-2 gap-4">
            <div>
              <div className="text-sm font-medium text-muted-foreground mb-1">Repository</div>
              <div className="text-sm">{run.repository}</div>
            </div>
            <div>
              <div className="text-sm font-medium text-muted-foreground mb-1">Branch</div>
              <div className="flex items-center gap-1.5 text-sm">
                <GitBranch className="h-4 w-4" />
                {run.head_branch}
              </div>
            </div>
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div>
              <div className="text-sm font-medium text-muted-foreground mb-1">Triggered by</div>
              <div className="flex items-center gap-1.5 text-sm">
                <User className="h-4 w-4" />
                {run.actor_login}
              </div>
            </div>
            <div>
              <div className="text-sm font-medium text-muted-foreground mb-1">Duration</div>
              <div className="flex items-center gap-1.5 text-sm">
                <Clock className="h-4 w-4" />
                {formatDuration(run.duration)}
              </div>
            </div>
          </div>

          <div>
            <div className="text-sm font-medium text-muted-foreground mb-1">Started at</div>
            <div className="text-sm">
              {format(new Date(run.run_started_at), "PPpp")}
            </div>
          </div>

          <div>
            <div className="text-sm font-medium text-muted-foreground mb-1">Commit</div>
            <div className="flex items-center gap-2 text-sm bg-muted p-3 rounded-md">
              <GitCommit className="h-4 w-4 flex-shrink-0" />
              <span className="flex-1">{run.commit_message}</span>
            </div>
          </div>

          <Button
            className="w-full"
            onClick={() => window.open(run.github_url, "_blank")}
          >
            <ExternalLink className="h-4 w-4 mr-2" />
            View on GitHub
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  );
};
