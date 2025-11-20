import styles from "./RepoConfig.module.scss";
import { useState, forwardRef, useImperativeHandle, useCallback } from "react";
import { Eye, EyeOff, Loader2 } from "lucide-react";
import React from "react";
import InputField from "@/components/InputField/InputField";
import { validateTokenAndRepo } from "@/services/githubApi";
import { useSnackbar } from "@/contexts/SnackbarContext";

interface RepoConfigProps {
  initialTokenType?: boolean;
  initialGithubToken?: string;
  initialRepoPath?: string;
}

export interface RepoConfigRef {
  validate(): Promise<boolean>;
  getData(): { tokenType: boolean; githubToken: string; repoPath: string };
  clearErrors(): void;
}

const RepoConfig = forwardRef<RepoConfigRef, RepoConfigProps>(
  ({ initialTokenType = false, initialGithubToken = "", initialRepoPath = "" }, ref) => {
    const [tokenType, setTokenType] = useState(initialTokenType);
    const [githubToken, setGithubToken] = useState(initialGithubToken);
    const [repoPath, setRepoPath] = useState(initialRepoPath);
    const [showToken, setShowToken] = useState(false);
    const [errors, setErrors] = useState<{ githubToken?: string; repoPath?: string }>({});
    const [showErrors, setShowErrors] = useState(false);
    const [isValidated, setIsValidated] = useState(false);

    const { showSnackbar } = useSnackbar();

    // Validation functions
    const validateGitHubToken = (token: string): string => {
      if (!token.trim()) return "GitHub token is required";
      if (tokenType === false && !token.startsWith("ghp_")) return "Personal access token must start with 'ghp_'";
      if (tokenType === true && !token.startsWith("github_pat_")) return "Fine-grained token must start with 'github_pat_'";
      if (token.length < 20) return "Token appears to be too short";
      return "";
    };

    const validateRepoPath = (path: string): string => {
      if (!path.trim()) return "Repository path is required";
      if (!/^[a-zA-Z0-9._-]+\/[a-zA-Z0-9._-]+$/.test(path.trim())) return "Repository path must be in format 'owner/repository'";
      if (path.trim().length > 100) return "Repository path must be less than 100 characters";
      return "";
    };

    // API validation function
    const validateWithApi = (async (): Promise<boolean> => {
      // First, do basic validation
      const githubTokenError = validateGitHubToken(githubToken);
      const repoPathError = validateRepoPath(repoPath);

      if (githubTokenError || repoPathError) {
        const newErrors = {
          githubToken: githubTokenError || undefined,
          repoPath: repoPathError || undefined,
        };
        setErrors(newErrors);
        setShowErrors(true);
        return false;
      }

      setErrors({});
      setShowErrors(false);

      try {
        const response = await validateTokenAndRepo(githubToken, repoPath);

        if (!response.success) {
          setErrors({ githubToken: response.error || response.message });
          setShowErrors(true);
          showSnackbar(response.error || response.message || "GitHub validation failed", "error");
          return false;
        }

        setIsValidated(true);
        showSnackbar(
          `Successfully validated access to ${repoPath}`,
          "success"
        );
        return true;
      } catch (error) {
        console.error("API validation error:", error);
        showSnackbar("Failed to validate GitHub credentials", "error");
        return false;
      }
    });

    // Expose methods to parent via ref
    useImperativeHandle(ref, () => ({
      async validate(): Promise<boolean> {
        // First, do basic validation
        const githubTokenError = validateGitHubToken(githubToken);
        const repoPathError = validateRepoPath(repoPath);

        if (githubTokenError || repoPathError) {
          const newErrors = {
            githubToken: githubTokenError || undefined,
            repoPath: repoPathError || undefined,
          };
          setErrors(newErrors);
          setShowErrors(true);
          return false;
        }

        // If basic validation passes, do API validation
        return await validateWithApi();
      },
      getData(): { tokenType: boolean; githubToken: string; repoPath: string } {
        return { tokenType, githubToken, repoPath };
      },
      clearErrors(): void {
        setErrors({});
        setShowErrors(false);
        setIsValidated(false);
      },
    }));

    return (
      <div className={styles.repoConfig}>
        {/* Token Type Selection */}
        <div className={styles.tokenTypeContainer}>
          <p className={styles.sectionLabel}>Select Token Type</p>
          <div className={styles.tokenTypeButtons}>
            <button
              type="button"
              className={`${styles.tokenTypeButton} ${tokenType === false ? styles.active : ""
                }`}
              onClick={() => setTokenType(false)}
            >
              <div className={styles.tokenTypeContent}>
                <p className={styles.tokenTypeTitle}>Personal Access Token</p>
                <p className={styles.tokenTypeSubtitle}>
                  Classic token with full access
                </p>
              </div>
            </button>
            <button
              type="button"
              className={`${styles.tokenTypeButton} ${tokenType === true ? styles.active : ""
                }`}
              onClick={() => setTokenType(true)}
            >
              <div className={styles.tokenTypeContent}>
                <p className={styles.tokenTypeTitle}>Fine-Grained Token</p>
                <p className={styles.tokenTypeSubtitle}>
                  Enhanced security & permissions
                </p>
              </div>
            </button>
          </div>
        </div>

        <div className={styles.helpSection}>
          <div className={styles.helpCard}>
            <h4 className={styles.helpTitle}>
              {tokenType === false
                ? "Personal Access Token"
                : "Fine-Grained Token"}{" "}
              Setup
            </h4>
            <p className={styles.helpText}>
              {tokenType === false
                ? "Go to GitHub Settings → Developer settings → Personal access tokens → Tokens (classic). Generate a new token with 'repo' scope for full repository access."
                : "Go to GitHub Settings → Developer settings → Personal access tokens → Fine-grained tokens. Create a token with specific repository access and required permissions."}
            </p>
            <ul className={styles.helpSteps}>
              <li className={styles.helpStep}>
                <p className={styles.stepNumber}>1</p>
                <p>Navigate to GitHub Settings</p>
              </li>
              <li className={styles.helpStep}>
                <p className={styles.stepNumber}>2</p>
                <p>Select Developer settings → Personal access tokens</p>
              </li>
              <li className={styles.helpStep}>
                <p className={styles.stepNumber}>3</p>
                <p>Generate new token with repository permissions</p>
              </li>
            </ul>
          </div>
        </div>

        {/* GitHub Token Input */}
        <div className={styles.inputGroup}>
          <label className={styles.inputLabel}>
            <p className={styles.labelText}>GitHub Token</p>
            <p className={styles.labelRequired}>*</p>
          </label>
          <InputField
            value={githubToken}
            onChange={(value) => {
              setGithubToken(value);
              // Clear error when user starts typing
              if (showErrors && errors.githubToken) {
                setErrors(prev => ({ ...prev, githubToken: undefined }));
              }
            }}
            placeholder={
              tokenType === false
                ? "ghp_xxxxxxxxxxxxxxxxxxxxxxxxxxx"
                : "github_pat_xxxxxxxxxxxxxxxxxxxx"
            }
            icon={showToken ? <EyeOff /> : <Eye />}
            onIconClick={() => setShowToken(!showToken)}
            error={showErrors ? errors.githubToken : undefined}
          />
        </div>

        {/* Repository Path Input */}
        <div className={styles.inputGroup}>
          <label className={styles.inputLabel}>
            <span className={styles.labelText}>Repository Path</span>
            <span className={styles.labelRequired}>*</span>
          </label>
          <InputField
            value={repoPath}
            onChange={(value) => {
              setRepoPath(value);
              // Clear error when user starts typing
              if (showErrors && errors.repoPath) {
                setErrors(prev => ({ ...prev, repoPath: undefined }));
              }
            }}
            placeholder="e.g., octocat/Hello-World"
            error={showErrors ? errors.repoPath : undefined}
          />
          <div className={styles.inputHint}>
            Format: username/repository-name or organization/repository-name
          </div>
        </div>
      </div>
    );
  });

export default RepoConfig;