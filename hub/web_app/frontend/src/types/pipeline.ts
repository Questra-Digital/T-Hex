// Shared types for pipeline management across API routes and context
export interface PipelineEvent {
  id: string;
  type: "Test";
  status: "success" | "failed" | "running" | "pending";
  timestamp: string;
  duration: string;
  details: string;
}

export interface Pipeline {
  id: string;
  name: string;
  description: string;
  status: "running" | "success" | "failed" | "pending";
  lastRun: string;
  triggerType: "manual" | "commit";
  branchName: string;
  repositoryPath: string;
  labels: string[];
  events: PipelineEvent[];
}

// Initial pipeline data (shared between context and API)
export const initialPipelines: Pipeline[] = [
  {
    id: "1",
    name: "build-and-test",
    description: "Automated build and test pipeline for the main application",
    status: "running",
    lastRun: "2024-01-15T10:30:00Z",
    triggerType: "commit",
    branchName: "main",
    repositoryPath: "github.com/company/frontend-app",
    labels: ["frontend", "testing", "automation"],
    events: [
      {
        id: "1",
        type: "Test",
        status: "success",
        timestamp: "2024-01-15T10:30:00Z",
        duration: "2m 34s",
        details: "Unit tests completed successfully with 0 errors"
      },
      {
        id: "2",
        type: "Test",
        status: "success",
        timestamp: "2024-01-15T10:32:34Z",
        duration: "1m 12s",
        details: "Integration tests passed - All 156 tests passed"
      },
      {
        id: "3",
        type: "Test",
        status: "running",
        timestamp: "2024-01-15T10:33:46Z",
        duration: "0m 45s",
        details: "E2E tests running..."
      }
    ]
  },
  {
    id: "2",
    name: "security-scan",
    description: "Security vulnerability scanning and code analysis",
    status: "success",
    lastRun: "2024-01-15T09:15:00Z",
    triggerType: "manual",
    branchName: "develop",
    repositoryPath: "github.com/company/security-tools",
    labels: ["security", "analysis"],
    events: [
      {
        id: "4",
        type: "Test",
        status: "success",
        timestamp: "2024-01-15T09:15:00Z",
        duration: "5m 23s",
        details: "Security tests completed - no critical vulnerabilities found"
      }
    ]
  },
  {
    id: "3",
    name: "performance-test",
    description: "Performance testing and load testing pipeline",
    status: "failed",
    lastRun: "2024-01-15T08:45:00Z",
    triggerType: "commit",
    branchName: "feature/performance",
    repositoryPath: "github.com/company/performance-suite",
    labels: ["performance", "testing"],
    events: [
      {
        id: "5",
        type: "Test",
        status: "failed",
        timestamp: "2024-01-15T08:45:00Z",
        duration: "3m 12s",
        details: "Performance tests failed - response time exceeded threshold"
      }
    ]
  }
];
