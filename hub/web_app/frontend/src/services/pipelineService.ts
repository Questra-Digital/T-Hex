// Pipeline service for API interactions

import { httpUtils } from "./httpUtils";
import { CreatePipelineRequest, CreatePipelineResponse, GetPipelinesResponse } from "@/types/pipeline";

export async function createPipeline(newPipelineRequest: CreatePipelineRequest): Promise<CreatePipelineResponse> {

  //The data type of the response is { pipeline_id: number }
  const response = await httpUtils.post<CreatePipelineResponse>("/create_pipeline", {
    body: newPipelineRequest,
  });

  return response;
}

export async function getPipelines(): Promise<GetPipelinesResponse> {
  const response = await httpUtils.get<GetPipelinesResponse>("/pipelines");
  return response;
}

export default {
  createPipeline,
  getPipelines,
}