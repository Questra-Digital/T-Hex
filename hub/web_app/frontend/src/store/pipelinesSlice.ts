import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Pipeline, PipelineEvent } from '@/types/pipeline';

interface PipelinesState {
    pipelines: Pipeline[];
}

const initialState: PipelinesState = {
    pipelines: [],
}

const pipelinesSlice = createSlice({
    name: 'pipelines',
    initialState,
    reducers: {
        setPipelines: (state, action: PayloadAction<Pipeline[]>) => {
            state.pipelines = action.payload;
        },
        addPipeline: (state, action: PayloadAction<Pipeline>) => {
            state.pipelines.push(action.payload);
        },
        deletePipeline: (state, action: PayloadAction<number>) => {
            state.pipelines = state.pipelines.filter(pipeline => pipeline.id !== action.payload);
        },
        updatePipeline: (state, action: PayloadAction<Pipeline>) => {
            const index = state.pipelines.findIndex(pipeline => pipeline.id === action.payload.id);
            if (index !== -1) {
                state.pipelines[index] = action.payload;
            }
        },
        addEventToPipeline: (state, action: PayloadAction<{ pipelineId: number, event: PipelineEvent }>) => {
            const index = state.pipelines.findIndex(pipeline => pipeline.id === action.payload.pipelineId);
            if (index !== -1) {
                state.pipelines[index].events.push(action.payload.event);
            }
        },
        resetPipelines: (state) => {
            state.pipelines = initialState.pipelines;
        }
    },
});

export const { setPipelines, addPipeline, deletePipeline, updatePipeline, addEventToPipeline, resetPipelines } = pipelinesSlice.actions;
export default pipelinesSlice.reducer;

// Selectors
export const selectPipelines = (state: { pipelines: PipelinesState }) => state.pipelines.pipelines;
export const selectNewPipelineId = (state: { pipelines: PipelinesState }) => state.pipelines.pipelines.length + 1;
