import { WorkflowStatus, WorkflowConclusion } from "@/types/workflow";
import { Badge } from "@/components/ui/badge";
import { CheckCircle2, XCircle, Clock, Loader2, Ban } from "lucide-react";

interface StatusBadgeProps {
  status: WorkflowStatus;
  conclusion: WorkflowConclusion;
}

export const StatusBadge = ({ status, conclusion }: StatusBadgeProps) => {
  const getStatusInfo = () => {
    if (status === "completed") {
      switch (conclusion) {
        case "success":
          return {
            label: "Success",
            variant: "success" as const,
            icon: CheckCircle2,
          };
        case "failure":
          return {
            label: "Failed",
            variant: "failure" as const,
            icon: XCircle,
          };
        case "cancelled":
          return {
            label: "Cancelled",
            variant: "cancelled" as const,
            icon: Ban,
          };
        default:
          return {
            label: "Completed",
            variant: "success" as const,
            icon: CheckCircle2,
          };
      }
    } else if (status === "in_progress") {
      return {
        label: "Running",
        variant: "running" as const,
        icon: Loader2,
      };
    } else {
      return {
        label: "Queued",
        variant: "queued" as const,
        icon: Clock,
      };
    }
  };

  const { label, variant, icon: Icon } = getStatusInfo();

  return (
    <Badge variant={variant} className="flex items-center gap-1.5">
      <Icon className={`h-3.5 w-3.5 ${variant === "running" ? "animate-spin" : ""}`} />
      {label}
    </Badge>
  );
};
