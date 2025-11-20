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

export interface TestSession{
  test_id: number;
  pipeline_event_id: number;
  time: number;
  key: string;
  proj: string;
  current: boolean;
}

export interface SessionSelenium{
  session_id: string;
  test_id: number;
  time: number;
  valid: boolean;
  status: boolean;
  message: string;
}

export interface Event{
  id: number;
  session_id: string;
  time: number;
  method: string;
  path: string;
  req_body: string;
  status: number;
  res: string;
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
  user_id: number;
}

export interface CreatePipelineResponse extends APIResponse<{"pipeline_id": number}> {}

export interface GetPipelinesResponse extends APIResponse<{"pipelines": Pipeline[]}> {}

export interface PipeLineEventRequest extends APIResponse<{"pipeline_id": number, "event": PipelineEvent}> {}

export interface GetTestSessionsResponse extends APIResponse<{"test_sessions": TestSession[]}> {}

export interface GetSeleniumSessionsResponse extends APIResponse<{"selenium_sessions": SessionSelenium[]}> {}

export interface GetSeleniumEventsResponse extends APIResponse<{"selenium_events": Event[]}> {}