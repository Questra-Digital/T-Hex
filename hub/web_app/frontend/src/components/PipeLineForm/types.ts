// Pipeline form data structure
export interface PipelineFormData {
  pipelineName: string;
  description: string;
  labels: string[];
  triggerType: "manual" | "commit";
  branchName: string;
  tokenType: boolean;
  githubToken: string;
  repoPath: string;
}
