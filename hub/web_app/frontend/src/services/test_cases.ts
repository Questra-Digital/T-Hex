import { GetSeleniumEventsResponse, GetSeleniumSessionsResponse, GetTestSessionsResponse } from "@/types/pipeline";
import { httpUtils, HttpRequestOptions } from "./httpUtils";

export async function fetchTestSessions(pipelineId: string, pipelineEventId: string, user_id: number) {

    const options: HttpRequestOptions = {
        body: {
            user_id: user_id,
            pipeline_id: pipelineId,
            pipeline_event_id: pipelineEventId,
        },
    };
    const response = await httpUtils.post<GetTestSessionsResponse>('/test-session', options);
    console.log("response from fetchTestSessions", response);

    return response;
}

export async function fetchSeleniumSessions(pipelineId: string, pipelineEventId: string, user_id: number) {
    const options: HttpRequestOptions = {
        body: {
            user_id: user_id,
            pipeline_id: pipelineId,
            pipeline_event_id: pipelineEventId,
        },
    };
    const response = await httpUtils.post<GetSeleniumSessionsResponse>('/selenium-sessions', options);
    console.log("response from fetchSeleniumSessions", response);
    return response;
}

export async function fetchSeleniumEvents(pipelineId: string, pipelineEventId: string, user_id: number, session_id: string) {
    const options: HttpRequestOptions = {
        body: {
            user_id: user_id,
            pipeline_id: pipelineId,
            pipeline_event_id: pipelineEventId,
            session_id: session_id,
        },
    };
    const response = await httpUtils.post<GetSeleniumEventsResponse>('/selenium-events', options);
    console.log("response from fetchSeleniumEvents", response);
    return response;
}