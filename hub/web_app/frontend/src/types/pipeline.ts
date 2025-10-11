import { APIResponse } from "./API";

// Shared types for pipeline management across API routes and context
 export type PipelineStatus = "running" | "success" | "failed";
 export type TriggerType = "manual" | "commit";
 export type EventType = "test";
 export type EventStatus = "success" | "failed" | "running";

export interface PipelineEvent {
  id: number;
  pipeline_id: number;
  status: EventStatus;
  timestamp: string;
  duration: number;
  details: string;
  type: EventType;
}

export interface Pipeline {
  id?: number;
  name: string;
  description: string;
  status: PipelineStatus;
  last_run: string;
  trigger_type: TriggerType;
  branch_name: string;
  repository_path: string;
  labels: string[];
  events: PipelineEvent[];
}

export interface CreatePipelineRequest {
  pipeline: Omit<Pipeline, "id" | "events">;
  access_token: string;
}

export interface CreatePipelineResponse extends APIResponse<{"pipeline_id": number}> {}

export interface GetPipelinesResponse extends APIResponse<{"pipelines": Pipeline[]}> {}

export interface PipeLineEventRequest extends APIResponse<{"pipeline_id": number, "event": PipelineEvent}> {}