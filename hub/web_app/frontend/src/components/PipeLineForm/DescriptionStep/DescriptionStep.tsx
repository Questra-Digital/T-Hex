import React, { useState, forwardRef, useImperativeHandle } from "react";
import styles from "@/styles/components/PipelineForm.module.scss";

interface DescriptionStepProps {
  initialValue?: string;
}

export interface DescriptionStepRef {
  validate(): boolean;
  getData(): { description: string };
  clearErrors(): void;
}

const DescriptionStep = forwardRef<DescriptionStepRef, DescriptionStepProps>(
  ({ initialValue = "" }, ref) => {
    const [description, setDescription] = useState(initialValue);
    const [error, setError] = useState<string>("");
    const [showError, setShowError] = useState(false);

    // Validation function
    const validateDescription = (desc: string): string => {
      if (!desc.trim()) return "Description is required";
      if (desc.trim().length < 10) return "Description must be at least 10 characters";
      if (desc.trim().length > 500) return "Description must be less than 500 characters";
      return "";
    };

    // Expose methods to parent via ref
    useImperativeHandle(ref, () => ({
      validate(): boolean {
        const errorMessage = validateDescription(description);
        setError(errorMessage);
        setShowError(!!errorMessage);
        return !errorMessage;
      },
      getData(): { description: string } {
        return { description };
      },
      clearErrors(): void {
        setError("");
        setShowError(false);
      },
    }));

    const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
      setDescription(e.target.value);
      // Clear error when user starts typing
      if (showError) {
        setShowError(false);
        setError("");
      }
    };

    return (
      <div className={styles.textareaContainer}>
        <textarea
          className={`${styles.descriptionInput} ${showError && error ? styles.error : ""}`}
          value={description}
          onChange={handleChange}
          placeholder="Enter a project description"
        />
        {showError && error && (
          <div className={styles.errorMessage}>{error}</div>
        )}
      </div>
    );
  }
);

export default DescriptionStep;