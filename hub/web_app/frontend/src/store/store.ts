import { configureStore } from '@reduxjs/toolkit';
import pipelinesReducer from './pipelinesSlice';

export const store = configureStore({
    reducer: {
        pipelines: pipelinesReducer,
    },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;