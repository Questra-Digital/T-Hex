// Pipeline service for API interactions

import { HttpRequestOptions, httpUtils } from "./httpUtils";
import { CreatePipelineRequest, CreatePipelineResponse, GetPipelinesResponse } from "@/types/pipeline";

export async function createPipeline(newPipelineRequest: CreatePipelineRequest): Promise<CreatePipelineResponse> {

  //The data type of the response is { pipeline_id: number }
  const response = await httpUtils.post<CreatePipelineResponse>("/create_pipeline", {
    body: newPipelineRequest,
  });

  return response;
}

export async function getPipelines(user_id: number): Promise<GetPipelinesResponse> {
  const options: HttpRequestOptions = {
    body: { user_id: user_id },
  };
  const response = await httpUtils.post<GetPipelinesResponse>("/pipelines", options);
  console.log("response from getPipelines", response);
  return response;
}

export default {
  createPipeline,
  getPipelines,
}