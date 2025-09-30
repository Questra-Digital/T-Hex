import styles from "./Summary.module.scss";
import { Eye, EyeOff, Play, GitBranch, Tag, FileText } from "lucide-react";
import { useState, forwardRef, useImperativeHandle } from "react";
import { PipelineFormData } from "../types";

interface SummaryProps {
  data: PipelineFormData;
}

export interface SummaryRef {
  validate(): boolean;
  getData(): PipelineFormData;
  clearErrors(): void;
}

const Summary = forwardRef<SummaryRef, SummaryProps>(
  ({ data }, ref) => {
    const [showToken, setShowToken] = useState(false);

    // Expose methods to parent via ref
    useImperativeHandle(ref, () => ({
      validate(): boolean {
        // Summary step is always valid (it's just a review)
        return true;
      },
      getData(): PipelineFormData {
        return data;
      },
      clearErrors(): void {
        // No errors to clear in summary
      },
    }));

    const maskToken = (token: string) => {
      if (!token) return "";
      if (token.length <= 8) return token;
      return `${token.substring(0, 8)}${"*".repeat(token.length - 8)}`;
    };

    return (
      <div className={styles.summaryContainer}>
        {/* Pipeline Info Section */}
        <div className={styles.summarySection}>
          <div className={styles.sectionHeader}>
            <FileText size={20} className={styles.sectionIcon} />
            <h3 className={styles.sectionTitle}>Pipeline Information</h3>
          </div>
          <div className={styles.summaryCard}>
            <div className={styles.summaryItem}>
              <span className={styles.itemLabel}>Pipeline Name:</span>
              <span className={styles.itemValue}>{data.pipelineName || "Not specified"}</span>
            </div>
            <div className={styles.summaryItem}>
              <span className={styles.itemLabel}>Description:</span>
              <span className={styles.itemValue}>{data.description || "No description provided"}</span>
            </div>
          </div>
        </div>

        {/* Labels Section */}
        <div className={styles.summarySection}>
          <div className={styles.sectionHeader}>
            <Tag size={20} className={styles.sectionIcon} />
            <h3 className={styles.sectionTitle}>Labels</h3>
          </div>
          <div className={styles.summaryCard}>
            {data.labels.length > 0 ? (
              <div className={styles.labelsContainer}>
                {data.labels.map((label, index) => (
                  <span key={index} className={styles.labelTag}>
                    {label}
                  </span>
                ))}
              </div>
            ) : (
              <span className={styles.emptyState}>No labels added</span>
            )}
          </div>
        </div>

        {/* Triggers Section */}
        <div className={styles.summarySection}>
          <div className={styles.sectionHeader}>
            {data.triggerType === "manual" ? (
              <Play size={20} className={styles.sectionIcon} />
            ) : (
              <GitBranch size={20} className={styles.sectionIcon} />
            )}
            <h3 className={styles.sectionTitle}>Trigger Configuration</h3>
          </div>
          <div className={styles.summaryCard}>
            <div className={styles.summaryItem}>
              <span className={styles.itemLabel}>Type:</span>
              <span className={styles.itemValue}>
                {data.triggerType === "manual" ? "Manual Trigger" : "Commit to Branch"}
              </span>
            </div>
            {data.triggerType === "commit" && data.branchName && (
              <div className={styles.summaryItem}>
                <span className={styles.itemLabel}>Branch:</span>
                <span className={styles.itemValue}>{data.branchName}</span>
              </div>
            )}
          </div>
        </div>

        {/* GitHub Configuration Section */}
        <div className={styles.summarySection}>
          <div className={styles.sectionHeader}>
            <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor" className={styles.sectionIcon}>
              <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z" />
            </svg>
            <h3 className={styles.sectionTitle}>GitHub Configuration</h3>
          </div>
          <div className={styles.summaryCard}>
            <div className={styles.summaryItem}>
              <span className={styles.itemLabel}>Token Type:</span>
              <span className={styles.itemValue}>
                {data.tokenType ? "Fine-Grained Token" : "Personal Access Token"}
              </span>
            </div>
            <div className={styles.summaryItem}>
              <span className={styles.itemLabel}>GitHub Token:</span>
              <div className={styles.tokenDisplay}>
                <span className={styles.itemValue}>
                  {showToken ? data.githubToken || "Not provided" : maskToken(data.githubToken)}
                </span>
                {data.githubToken && (
                  <button
                    type="button"
                    className={styles.toggleTokenBtn}
                    onClick={() => setShowToken(!showToken)}
                    title={showToken ? "Hide token" : "Show token"}
                  >
                    {showToken ? <EyeOff size={16} /> : <Eye size={16} />}
                  </button>
                )}
              </div>
            </div>
            <div className={styles.summaryItem}>
              <span className={styles.itemLabel}>Repository Path:</span>
              <span className={styles.itemValue}>{data.repoPath || "Not specified"}</span>
            </div>
          </div>
        </div>

      </div>
    );
  });

Summary.displayName = "Summary";

export default Summary;