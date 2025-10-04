"use client";

import { useState, useRef, useCallback, useEffect } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { useRouter } from "next/navigation";
import styles from "@/styles/components/PipelineForm.module.scss";
import { ArrowLeft, ArrowRight, Loader2 } from "lucide-react";
import { useSnackbar } from "@/contexts/SnackbarContext";
import { pipelineService } from "@/services/pipelineService";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { addPipeline, selectNewPipelineId } from "@/store/pipelinesSlice";
import { Pipeline } from "@/types/pipeline";
// Step components
import PipelineNameStep, { PipelineNameStepRef } from "@/components/PipeLineForm/PipelineNameStep/PipelineNameStep";
import DescriptionStep, { DescriptionStepRef } from "@/components/PipeLineForm/DescriptionStep/DescriptionStep";
import LabelsInput, { LabelsInputRef } from "@/components/PipeLineForm/LabelsInput/LabelsInput";
import Triggers, { TriggersRef } from "@/components/PipeLineForm/Triggers/Triggers";
import RepoConfig, { RepoConfigRef } from "@/components/PipeLineForm/RepoConfig/RepoConfig";
import Summary, { SummaryRef } from "@/components/PipeLineForm/Summary/Summary";

interface PipelineFormProps {
  onStepChange?: (stepIndex: number) => void;
  onStepComplete?: (stepIndex: number) => void;
  onStepIncomplete?: (stepIndex: number) => void;
}

interface PipelineFormData {
  pipelineName: string;
  description: string;
  labels: string[];
  triggerType: "manual" | "commit";
  branchName: string;
  tokenType: boolean;
  githubToken: string;
  repoPath: string;
}


function PipelineFormContent({ onStepChange, onStepComplete, onStepIncomplete }: PipelineFormProps) {
  const router = useRouter();
  const { showSnackbar } = useSnackbar();
  const dispatch = useAppDispatch();
  const newPipelineId = useAppSelector(selectNewPipelineId);


  // Current step state
  const [currentStep, setCurrentStep] = useState(0);
  const [isValidating, setIsValidating] = useState(false);

  // Form data state
  const [formData, setFormData] = useState<PipelineFormData>({
    pipelineName: "",
    description: "",
    labels: [],
    triggerType: "manual",
    branchName: "",
    tokenType: false,
    githubToken: "",
    repoPath: "",
  });

  // Refs for step components
  const stepRefs = useRef<(PipelineNameStepRef | DescriptionStepRef | LabelsInputRef | TriggersRef | RepoConfigRef | SummaryRef)[]>([]);

  // Notify parent of initial step
  useEffect(() => {
    onStepChange?.(currentStep);
  }, [currentStep, onStepChange]);

  // Step configuration
  const steps = [
    {
      id: 0,
      title: "Enter Pipeline Name",
      subtitle: "We suggest naming pipelines after the work they will perform i.e. build-and-test",
      children: (
        <PipelineNameStep
          ref={(ref) => {
            if (ref) stepRefs.current[0] = ref;
          }}
          initialValue={formData.pipelineName}
        />
      ),
    },
    {
      id: 1,
      title: "Enter Project Description",
      subtitle: "This helps others understand what the project is about.",
      children: (
        <DescriptionStep
          ref={(ref) => {
            if (ref) stepRefs.current[1] = ref;
          }}
          initialValue={formData.description}
        />
      ),
    },
    {
      id: 2,
      title: "Add Labels",
      subtitle: "You can organize your project by adding custom labels. Selected labels appear in the input field above.",
      children: (
        <LabelsInput
          ref={(ref) => {
            if (ref) stepRefs.current[2] = ref;
          }}
          initialLabels={formData.labels}
        />
      ),
    },
    {
      id: 3,
      title: "GitHub Configuration",
      children: (
        <RepoConfig
          ref={(ref) => {
            if (ref) stepRefs.current[3] = ref;
          }}
          initialTokenType={formData.tokenType}
          initialGithubToken={formData.githubToken}
          initialRepoPath={formData.repoPath}
        />
      ),
    },
    {
      id: 4,
      title: "Configure Triggers",
      subtitle: "Choose how your pipeline will be triggered. You can run it manually or automatically when commits are pushed to a branch.",
      children: (
        <Triggers
          ref={(ref) => {
            if (ref) stepRefs.current[4] = ref;
          }}
          initialTriggerType={formData.triggerType}
          initialBranchName={formData.branchName}
          githubToken={formData.githubToken}
          repoPath={formData.repoPath}
        />
      ),
    },
    {
      id: 5,
      title: "Review & Create Pipeline",
      subtitle: "Review your pipeline configuration before creating it.",
      children: (
        <Summary
          ref={(ref) => {
            if (ref) stepRefs.current[5] = ref;
          }}
          data={formData}
        />
      ),
    },
  ];

  // Handle step navigation
  const handleNextStep = (async () => {
    const currentStepRef = stepRefs.current[currentStep];
    if (!currentStepRef) return;

    setIsValidating(true);
    try {
      // Validate current step's inputs (now async)
      const isValid = await currentStepRef.validate();
      if (isValid) {
        // Collect data from current step
        const stepData = currentStepRef.getData();
        setFormData(prev => ({ ...prev, ...stepData }));

        // Mark current step as completed
        onStepComplete?.(currentStep);

        if (currentStep < steps.length - 1) {
          const nextStep = currentStep + 1;
          setCurrentStep(nextStep);
          onStepChange?.(nextStep);
        }
      }
    } catch (error) {
      console.error("Validation error:", error);
    } finally {
      setIsValidating(false);
    }
  });

  const handlePrevStep = (() => {
    if (currentStep > 0) {
      // Mark current step as incomplete when going back
      onStepIncomplete?.(currentStep);

      const prevStep = currentStep - 1;
      setCurrentStep(prevStep);
      onStepChange?.(prevStep);
    }
  });


  const step = steps[currentStep];

  return (
    <div className={styles.wrapper}>
      <AnimatePresence mode="wait">
        <motion.div
          key={currentStep}
          initial={{ opacity: 0, x: 100 }}
          animate={{ opacity: 1, x: 0 }}
          exit={{ opacity: 0, x: -100 }}
          transition={{ duration: 0.4 }}
          className={styles.infoContainer}
        >
          <p className={styles.infoTitle}>{step.title}</p>

          {/* Step component being rendered */}
          {step.children}

          {step.subtitle && (
            <p className={styles.infoSubtitle}>{step.subtitle}</p>
          )}

          {/* Navigation Buttons */}
          <div className={styles.navigation}>
            {currentStep > 0 && (
              <button
                className={styles.button}
                onClick={handlePrevStep}
              >
                <ArrowLeft size={26} />
                {steps[currentStep - 1].title}
              </button>
            )}

            {currentStep < steps.length - 1 && (
              <button
                className={styles.button}
                onClick={handleNextStep}
                disabled={isValidating}
              >
                {isValidating ? (
                  <>
                    <Loader2 className={styles.spinner} size={26} />
                    Validating...
                  </>
                ) : (
                  <>
                    {steps[currentStep + 1].title} <ArrowRight size={26} />
                  </>
                )}
              </button>
            )}

            {currentStep === steps.length - 1 && (
              <button
                className={styles.button}
                onClick={async () => {
                  try {

                    const newPipeLine: Pipeline = {
                      id: (newPipelineId).toString(),
                      name: formData.pipelineName,
                      description: formData.description,
                      status: "pending",
                      lastRun: new Date().toISOString(),
                      triggerType: formData.triggerType,
                      branchName: formData.branchName,
                      repositoryPath: formData.repoPath,
                      labels: formData.labels,
                      events: [],
                    };

                    // Create pipeline via API service using the ID from context
                    const result = await pipelineService.createPipeline({
                      pipeline_id: newPipeLine.id,
                      repo_path: formData.repoPath,
                      access_token: formData.githubToken,
                    });

                    if (result.success) {
                      // Show success snackbar
                      showSnackbar(result.message, "success", 3000);

                      // Add the pipeline to the context
                      dispatch(addPipeline(newPipeLine));

                      // Redirect to pipelines page after a short delay
                      setTimeout(() => {
                        router.push("/pipelines");
                      }, 1000);
                    } else {
                      // Show error snackbar
                      showSnackbar(result.error || result.message, "error", 5000);
                    }
                  } catch (error) {
                    console.error("Pipeline creation error:", error);
                    showSnackbar("Failed to create pipeline. Please try again.", "error", 5000);
                  }
                }}
              >
                Create Pipeline
              </button>
            )}
          </div>
        </motion.div>
      </AnimatePresence>
    </div>
  );
}

export default function PipelineForm({ onStepChange, onStepComplete, onStepIncomplete }: PipelineFormProps) {
  return (
    <PipelineFormContent
      onStepChange={onStepChange}
      onStepComplete={onStepComplete}
      onStepIncomplete={onStepIncomplete}
    />
  );
}
