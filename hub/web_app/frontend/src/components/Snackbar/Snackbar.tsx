"use client";

import React, { useEffect, useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { CheckCircle, AlertCircle, Info, X } from "lucide-react";
import styles from "./Snackbar.module.scss";

export interface SnackbarProps {
  message: string;
  type: "success" | "error" | "info";
  isVisible: boolean;
  onClose: () => void;
  duration?: number; // Auto-close duration in milliseconds
}

export default function Snackbar({
  message,
  type,
  isVisible,
  onClose,
  duration = 5000,
}: SnackbarProps) {
  const [isClosing, setIsClosing] = useState(false);

  useEffect(() => {
    if (isVisible && duration > 0) {
      const timer = setTimeout(() => {
        handleClose();
      }, duration);

      return () => clearTimeout(timer);
    }
  }, [isVisible, duration]);

  const handleClose = () => {
    setIsClosing(true);
    setTimeout(() => {
      onClose();
      setIsClosing(false);
    }, 300); // Match animation duration
  };

  const getIcon = () => {
    switch (type) {
      case "success":
        return <CheckCircle size={20} />;
      case "error":
        return <AlertCircle size={20} />;
      case "info":
        return <Info size={20} />;
      default:
        return <Info size={20} />;
    }
  };

  return (
    <AnimatePresence>
      {isVisible && (
        <motion.div
          initial={{ opacity: 0, y: -100, scale: 0.95 }}
          animate={{ 
            opacity: 1, 
            y: 0, 
            scale: 1,
            transition: { type: "spring", stiffness: 300, damping: 30 }
          }}
          exit={{ 
            opacity: 0, 
            y: -100, 
            scale: 0.95,
            transition: { duration: 0.3 }
          }}
          className={`${styles.snackbar} ${styles[type]} ${isClosing ? styles.closing : ''}`}
        >
          <div className={styles.content}>
            <div className={styles.icon}>
              {getIcon()}
            </div>
            <div className={styles.message}>
              {message}
            </div>
            <button
              className={styles.closeButton}
              onClick={handleClose}
              aria-label="Close notification"
            >
              <X size={16} />
            </button>
          </div>
        </motion.div>
      )}
    </AnimatePresence>
  );
}
