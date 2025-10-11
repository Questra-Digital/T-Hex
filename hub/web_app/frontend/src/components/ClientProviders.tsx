"use client";

import { SnackbarProvider } from "@/contexts/SnackbarContext";
import { Provider } from "react-redux";
import { store } from "@/store/store";
import { ReactNode } from "react";
import PipeLineEventsListener from "./PipeLineEventsListener";

interface ClientProvidersProps {
    children: ReactNode;
}

export default function ClientProviders({ children }: ClientProvidersProps) {
    return (
        <SnackbarProvider>
            <Provider store={store}>
                {/* Listen to pipeline events */}
                <PipeLineEventsListener />
                {children}
            </Provider>
        </SnackbarProvider>
    );
}
