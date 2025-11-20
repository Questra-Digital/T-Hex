import styles from "./Triggers.module.scss";
import { motion } from "framer-motion";
import { Play, GitBranch, Loader2 } from "lucide-react";
import InputField from "@/components/InputField/InputField";
import { useState, forwardRef, useImperativeHandle, useCallback } from "react";
import { validateBranch } from "@/services/githubApi";
import { useSnackbar } from "@/contexts/SnackbarContext";

type TriggerType = "manual" | "commit";

interface TriggersProps {
  initialTriggerType?: TriggerType;
  initialBranchName?: string;
  githubToken?: string;
  repoPath?: string;
}

export interface TriggersRef {
  validate(): Promise<boolean>;
  getData(): { triggerType: TriggerType; branchName: string };
  clearErrors(): void;
}

const Triggers = forwardRef<TriggersRef, TriggersProps>(
  ({ initialTriggerType = "manual", initialBranchName = "", githubToken, repoPath }, ref) => {
    const [triggerType, setTriggerType] = useState<TriggerType>(initialTriggerType);
    const [branchName, setBranchName] = useState(initialBranchName);
    const [error, setError] = useState<string>("");
    const [showError, setShowError] = useState(false);
    const [isBranchValidated, setIsBranchValidated] = useState(false);

    const { showSnackbar } = useSnackbar();

    // Validation function
    const validateBranchName = (branch: string, trigger: "manual" | "commit"): string => {
      if (trigger === "commit") {
        if (!branch.trim()) return "Branch name is required for commit triggers";
        if (!/^[a-zA-Z0-9/._-]+$/.test(branch.trim())) return "Branch name contains invalid characters";
        if (branch.trim().length > 100) return "Branch name must be less than 100 characters";
      }
      return "";
    };

    // API validation function for branch
    const validateBranchWithApi = (async (): Promise<boolean> => {
      if (triggerType !== "commit") {
        return true; // No validation needed for manual triggers
      }

      // First, do basic validation
      const branchError = validateBranchName(branchName, triggerType);
      if (branchError) {
        setError(branchError);
        setShowError(true);
        return false;
      }

      if (!githubToken || !repoPath) {
        showSnackbar("GitHub credentials are required to validate branch", "error");
        return false;
      }

      setError("");
      setShowError(false);

      try {
        const response = await validateBranch(githubToken, repoPath, branchName);

        if (!response.success) {
          setError(response.error || response.message || "Branch validation failed");
          setShowError(true);
          showSnackbar(response.error || response.message || "Branch validation failed", "error");
          return false;
        }

        setIsBranchValidated(true);
        showSnackbar(`Branch '${branchName}' validated successfully`, "success");
        return true;
      } catch (error) {
        console.error("Branch validation error:", error);
        showSnackbar("Failed to validate branch", "error");
        return false;
      }
    });

    // Expose methods to parent via ref
    useImperativeHandle(ref, () => ({
      async validate(): Promise<boolean> {
        // First, do basic validation
        const errorMessage = validateBranchName(branchName, triggerType);
        if (errorMessage) {
          setError(errorMessage);
          setShowError(true);
          return false;
        }

        // If commit trigger is selected and GitHub credentials are available, validate branch
        if (triggerType === "commit" && githubToken && repoPath) {
          return await validateBranchWithApi();
        }

        return true;
      },
      getData(): { triggerType: TriggerType; branchName: string } {
        return { triggerType, branchName };
      },
      clearErrors(): void {
        setError("");
        setShowError(false);
        setIsBranchValidated(false);
      },
    }));
    const triggerOptions = [
      {
        type: "manual" as TriggerType,
        title: "Manual",
        description: "Trigger the pipeline manually when needed",
        icon: <Play size={18} />,
      },
      {
        type: "commit" as TriggerType,
        title: "Commit to Branch",
        description: "Auto-trigger on commits to branch",
        icon: <GitBranch size={18} />,
      },
    ];

    return (
      <div className={styles.triggersContainer}>
        {/* Trigger Type Selection */}
        <div className={styles.triggerTypeContainer}>
          <p className={styles.sectionLabel}>Select Trigger Type</p>
          <div className={styles.triggerOptions}>
            {triggerOptions.map((option) => (
              <button
                key={option.type}
                type="button"
                className={`${styles.triggerOption} ${triggerType === option.type ? styles.selected : ""
                  }`}
                onClick={() => setTriggerType(option.type)}
              >
                <div className={styles.triggerIcon}>{option.icon}</div>
                <div className={styles.triggerContent}>
                  <div className={styles.triggerTitle}>{option.title}</div>
                  <div className={styles.triggerDescription}>
                    {option.description}
                  </div>
                </div>
              </button>
            ))}
          </div>
        </div>

        {triggerType === "commit" && (
          <motion.div
            className={styles.branchConfig}
            initial={{ opacity: 0, height: 0 }}
            animate={{ opacity: 1, height: "auto" }}
            exit={{ opacity: 0, height: 0 }}
            transition={{ duration: 0.3 }}
          >
            <div className={styles.inputGroup}>
              <label className={styles.inputLabel}>
                <span className={styles.labelText}>Branch Name</span>
                <span className={styles.labelRequired}>*</span>
              </label>
              <InputField
                value={branchName}
                onChange={(val) => {
                  setBranchName(val);
                  // Clear error when user starts typing
                  if (showError) {
                    setShowError(false);
                    setError("");
                  }
                }}
                placeholder="Enter branch name (e.g., main, develop)"
                error={showError ? error : undefined}
              />
              <p className={styles.inputHint}>
                The pipeline will trigger automatically when commits are pushed to
                this branch
              </p>
            </div>
          </motion.div>
        )}
      </div>
    );
  });

export default Triggers;