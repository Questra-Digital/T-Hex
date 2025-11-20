import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { SessionSelenium, Event } from '@/types/pipeline';
import { fetchSeleniumSessions, fetchSeleniumEvents } from '@/services/test_cases';

// Async thunks for API calls
export const fetchSeleniumSessionsAsync = createAsyncThunk(
  'testCases/fetchSeleniumSessions',
  async ({ pipelineId, pipelineEventId, userId }: { 
    pipelineId: string; 
    pipelineEventId: string; 
    userId: number; 
  }) => {
    const response = await fetchSeleniumSessions(pipelineId, pipelineEventId, userId);
    return { response, pipelineId, pipelineEventId, userId };
  }
);

export const fetchSeleniumEventsAsync = createAsyncThunk(
  'testCases/fetchSeleniumEvents',
  async ({ pipelineId, pipelineEventId, userId, sessionId }: { 
    pipelineId: string; 
    pipelineEventId: string; 
    userId: number; 
    sessionId: string; 
  }) => {
    const response = await fetchSeleniumEvents(pipelineId, pipelineEventId, userId, sessionId);
    return { response, pipelineId, pipelineEventId, userId, sessionId };
  }
);

interface TestCasesState {
  // Cache for selenium sessions by pipeline event
  sessionsCache: {
    [key: string]: {
      data: SessionSelenium[];
      timestamp: number;
      loading: boolean;
      error: string | null;
    };
  };
  
  // Cache for selenium events by session
  eventsCache: {
    [key: string]: {
      data: Event[];
      timestamp: number;
      loading: boolean;
      error: string | null;
    };
  };
  
  // Cache expiration time (5 minutes)
  cacheExpiration: number;
}

const initialState: TestCasesState = {
  sessionsCache: {},
  eventsCache: {},
  cacheExpiration: 5 * 60 * 1000, // 5 minutes in milliseconds
};

const testCasesSlice = createSlice({
  name: 'testCases',
  initialState,
  reducers: {
    clearCache: (state) => {
      state.sessionsCache = {};
      state.eventsCache = {};
    },
    clearSessionsCache: (state, action: PayloadAction<string>) => {
      delete state.sessionsCache[action.payload];
    },
    clearEventsCache: (state, action: PayloadAction<string>) => {
      delete state.eventsCache[action.payload];
    },
  },
  extraReducers: (builder) => {
    // Handle selenium sessions
    builder
      .addCase(fetchSeleniumSessionsAsync.pending, (state, action) => {
        const key = `${action.meta.arg.pipelineId}-${action.meta.arg.pipelineEventId}-${action.meta.arg.userId}`;
        state.sessionsCache[key] = {
          data: [],
          timestamp: Date.now(),
          loading: true,
          error: null,
        };
      })
      .addCase(fetchSeleniumSessionsAsync.fulfilled, (state, action) => {
        const key = `${action.meta.arg.pipelineId}-${action.meta.arg.pipelineEventId}-${action.meta.arg.userId}`;
        if (action.payload.response.success && action.payload.response.data) {
          state.sessionsCache[key] = {
            data: action.payload.response.data.selenium_sessions || [],
            timestamp: Date.now(),
            loading: false,
            error: null,
          };
        } else {
          state.sessionsCache[key] = {
            data: [],
            timestamp: Date.now(),
            loading: false,
            error: action.payload.response.message || 'Failed to fetch sessions',
          };
        }
      })
      .addCase(fetchSeleniumSessionsAsync.rejected, (state, action) => {
        const key = `${action.meta.arg.pipelineId}-${action.meta.arg.pipelineEventId}-${action.meta.arg.userId}`;
        state.sessionsCache[key] = {
          data: [],
          timestamp: Date.now(),
          loading: false,
          error: action.error.message || 'Failed to fetch sessions',
        };
      });

    // Handle selenium events
    builder
      .addCase(fetchSeleniumEventsAsync.pending, (state, action) => {
        const key = `${action.meta.arg.pipelineId}-${action.meta.arg.pipelineEventId}-${action.meta.arg.userId}-${action.meta.arg.sessionId}`;
        state.eventsCache[key] = {
          data: [],
          timestamp: Date.now(),
          loading: true,
          error: null,
        };
      })
      .addCase(fetchSeleniumEventsAsync.fulfilled, (state, action) => {
        const key = `${action.meta.arg.pipelineId}-${action.meta.arg.pipelineEventId}-${action.meta.arg.userId}-${action.meta.arg.sessionId}`;
        if (action.payload.response.success && action.payload.response.data) {
          state.eventsCache[key] = {
            data: action.payload.response.data.selenium_events || [],
            timestamp: Date.now(),
            loading: false,
            error: null,
          };
        } else {
          state.eventsCache[key] = {
            data: [],
            timestamp: Date.now(),
            loading: false,
            error: action.payload.response.message || 'Failed to fetch events',
          };
        }
      })
      .addCase(fetchSeleniumEventsAsync.rejected, (state, action) => {
        const key = `${action.meta.arg.pipelineId}-${action.meta.arg.pipelineEventId}-${action.meta.arg.userId}-${action.meta.arg.sessionId}`;
        state.eventsCache[key] = {
          data: [],
          timestamp: Date.now(),
          loading: false,
          error: action.error.message || 'Failed to fetch events',
        };
      });
  },
});

export const { clearCache, clearSessionsCache, clearEventsCache } = testCasesSlice.actions;

// Selectors
export const selectSessionsCache = (state: { testCases: TestCasesState }, key: string) => {
  const cache = state.testCases.sessionsCache[key];
  if (!cache) return null;
  
  // Check if cache is expired
  if (Date.now() - cache.timestamp > state.testCases.cacheExpiration) {
    return null;
  }
  
  return cache;
};

export const selectEventsCache = (state: { testCases: TestCasesState }, key: string) => {
  const cache = state.testCases.eventsCache[key];
  if (!cache) return null;
  
  // Check if cache is expired
  if (Date.now() - cache.timestamp > state.testCases.cacheExpiration) {
    return null;
  }
  
  return cache;
};

export default testCasesSlice.reducer;
