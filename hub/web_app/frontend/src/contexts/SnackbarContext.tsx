"use client";

import React, { createContext, useContext, useState, useCallback, ReactNode } from "react";
import Snackbar, { SnackbarProps } from "@/components/Snackbar/Snackbar";

interface SnackbarContextType {
  showSnackbar: (message: string, type: "success" | "error" | "info", duration?: number) => void;
  hideSnackbar: () => void;
}

const SnackbarContext = createContext<SnackbarContextType | undefined>(undefined);

interface SnackbarProviderProps {
  children: ReactNode;
}

export function SnackbarProvider({ children }: SnackbarProviderProps) {
  const [snackbar, setSnackbar] = useState<{
    message: string;
    type: "success" | "error" | "info";
    isVisible: boolean;
    duration?: number;
  }>({
    message: "",
    type: "info",
    isVisible: false,
    duration: 5000,
  });

  const showSnackbar = useCallback((
    message: string,
    type: "success" | "error" | "info",
    duration = 5000
  ) => {
    setSnackbar({
      message,
      type,
      isVisible: true,
      duration,
    });
  }, []);

  const hideSnackbar = useCallback(() => {
    setSnackbar(prev => ({
      ...prev,
      isVisible: false,
    }));
  }, []);

  return (
    <SnackbarContext.Provider value={{ showSnackbar, hideSnackbar }}>
      {children}
      <Snackbar
        message={snackbar.message}
        type={snackbar.type}
        isVisible={snackbar.isVisible}
        onClose={hideSnackbar}
        duration={snackbar.duration}
      />
    </SnackbarContext.Provider>
  );
}

export function useSnackbar() {
  const context = useContext(SnackbarContext);
  if (context === undefined) {
    throw new Error("useSnackbar must be used within a SnackbarProvider");
  }
  return context;
}
