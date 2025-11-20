"use client";

import { SnackbarProvider } from "@/contexts/SnackbarContext";
import { Provider } from "react-redux";
import { PersistGate } from 'redux-persist/integration/react';
import { store, persistor } from "@/store/store";
import { ReactNode } from "react";
import SessionInitializer from "./SessionInitializer";

interface ClientProvidersProps {
    children: ReactNode;
}

export default function ClientProviders({ children }: ClientProvidersProps) {
    return (
        <SnackbarProvider>
            <Provider store={store}>
                <PersistGate loading={null} persistor={persistor}>
                    <SessionInitializer />
                    {children}
                </PersistGate>
            </Provider>
        </SnackbarProvider>
    );
}
