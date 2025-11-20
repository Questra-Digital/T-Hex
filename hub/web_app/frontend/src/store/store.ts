import { configureStore } from '@reduxjs/toolkit';
import { persistStore, persistReducer } from 'redux-persist';
import storage from 'redux-persist/lib/storage';
import { combineReducers } from '@reduxjs/toolkit';
import pipelinesReducer from './pipelinesSlice';
import userDataReducer from './userData';
import testCasesReducer from './testCasesSlice';

const persistConfig = {
    key: 'root',
    storage,
    whitelist: ['userData'], // Only persist userData slice
};

const rootReducer = combineReducers({
    pipelines: pipelinesReducer,
    userData: userDataReducer,
    testCases: testCasesReducer,
});

const persistedReducer = persistReducer(persistConfig, rootReducer);

export const store = configureStore({
    reducer: persistedReducer,
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware({
            serializableCheck: {
                ignoredActions: ['persist/PERSIST', 'persist/REHYDRATE'],
            },
        }),
});

export const persistor = persistStore(store);

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;