import React, { useState, forwardRef, useImperativeHandle } from "react";
import InputField from "@/components/InputField/InputField";

interface PipelineNameStepProps {
  initialValue?: string;
}

export interface PipelineNameStepRef {
  validate(): boolean;
  getData(): { pipelineName: string };
  clearErrors(): void;
}

const PipelineNameStep = forwardRef<PipelineNameStepRef, PipelineNameStepProps>(
  ({ initialValue = "" }, ref) => {
    const [pipelineName, setPipelineName] = useState(initialValue);
    const [error, setError] = useState<string>("");
    const [showError, setShowError] = useState(false);

    // Validation function
    const validatePipelineName = (name: string): string => {
      if (!name.trim()) return "Pipeline name is required";
      if (name.trim().length < 3) return "Pipeline name must be at least 3 characters";
      if (name.trim().length > 50) return "Pipeline name must be less than 50 characters";
      if (!/^[a-zA-Z0-9-_]+$/.test(name.trim())) return "Pipeline name can only contain letters, numbers, hyphens, and underscores";
      return "";
    };

    // Expose methods to parent via ref
    useImperativeHandle(ref, () => ({
      validate(): boolean {
        const errorMessage = validatePipelineName(pipelineName);
        setError(errorMessage);
        setShowError(!!errorMessage);
        return !errorMessage;
      },
      getData(): { pipelineName: string } {
        return { pipelineName };
      },
      clearErrors(): void {
        setError("");
        setShowError(false);
      },
    }));

    const handleChange = (value: string) => {
      setPipelineName(value);
      // Clear error when user starts typing
      if (showError) {
        setShowError(false);
        setError("");
      }
    };

    return (
      <div>
        <InputField
          value={pipelineName}
          onChange={handleChange}
          placeholder="Enter Pipeline Name"
          error={showError ? error : undefined}
        />
      </div>
    );
  }
);

export default PipelineNameStep;
