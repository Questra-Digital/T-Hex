// components/PipelineEventsListener.tsx
"use client";

import { useEffect } from "react";
import { useDispatch } from "react-redux";
import { addEventToPipeline } from "@/store/pipelinesSlice";
import { PipelineEvent } from "@/types/pipeline";

export default function PipeLineEventsListener() {
    const dispatch = useDispatch();

    useEffect(() => {
        const eventSource = new EventSource("/api/pipeline-status");

        eventSource.onmessage = (event) => {
            try {
                const data = JSON.parse(event.data) as {
                    pipeline_id: number;
                    event: PipelineEvent;
                };

                console.log("SSE event received:", data);
                dispatch(addEventToPipeline({ pipelineId: data.pipeline_id, event: data.event }));
            } catch (err) {
                console.error("Error parsing SSE event:", err);
            }
        };

        eventSource.onerror = (err) => {
            console.error("SSE connection error:", err);
            eventSource.close();
        };

        return () => {
            eventSource.close();
        };
    }, [dispatch]);

    return null; // No UI, just background listener
}
