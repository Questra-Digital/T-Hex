"use client";
import styles from "@/styles/pages/GettingStarted.module.scss";
import { outfit } from "@/fonts/fonts";
import React, { useState } from "react";
import Image from "next/image";
import PipelineForm from "@/components/PipelineForm";

const steps = [
  { number: 1, name: "Enter Pipeline Name" },
  { number: 2, name: "Enter Project Description" },
  { number: 3, name: "Add Labels" },
  {
    number: 4, name: "GitHub Configuration",
    children: [
      { name: "Token setup" },
      { name: "Repository path" },
    ],
  },
  {
    number: 5,
    name: "Configure Triggers",
    children: [
      { name: "Choose trigger type" },
      { name: "Set branch name" },
    ],
  },

  { number: 6, name: "Review & Create Pipeline" },
];

const pipelineSteps = [
  {
    icon: "Icons/GettingStarted/pipeline.svg",
    title: "Pipeline",
    description:
      "Pipelines orchestrate executable commands and scripts for your CI/CD process",
  },
  {
    icon: "/Icons/GettingStarted/code.svg",
    title: "Code",
    description: "Pipelines can checkout code from VCS repositories",
  },
  {
    icon: "/Icons/GettingStarted/config.svg",
    title: "Config",
    description: "Config files use YAML to define what runs your pipeline",
  },
  {
    icon: "Icons/GettingStarted/trigger.svg",
    title: "Triggers",
    description:
      "Triggers automatically run your pipelines when an event occurs",
  },
];


export default function Page() {
  const [currentStep, setCurrentStep] = useState(0);
  const [completedSteps, setCompletedSteps] = useState<number[]>([]);
  return (
    <section className={`${styles.gettingStarted} ${outfit.variable}`}>
      <div className={styles.sideBar}>
        <ul>
          {steps.map((step) => {
            const isCurrentStep = currentStep === step.number - 1;
            const isCompleted = completedSteps.includes(step.number - 1);

            return (
              <React.Fragment key={step.number}>
                <li className={`${styles.sidebarItem} ${isCurrentStep ? styles.currentStep : ''}`}>
                  <span className={`${styles.circle} ${isCompleted ? styles.completed : ''} ${isCurrentStep ? styles.current : ''}`}>
                    {isCompleted ? (
                      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41L9 16.17z" fill="currentColor" />
                      </svg>
                    ) : (
                      step.number
                    )}
                  </span>
                  <span className={styles.itemName}>{step.name}</span>
                </li>

                {step.children &&
                  step.children.map((child, index) => (
                    <li
                      key={`${step.number}-${index}`}
                      className={`${styles.sidebarItem} ${styles.subListItem}`}
                    >
                      <span className={styles.subConnector}></span>
                      <span className={styles.itemName}>{child.name}</span>
                    </li>
                  ))}
              </React.Fragment>
            );
          })}
        </ul>
      </div>
      <div className={styles.mainContent}>
        <div className={styles.headingContainer}>
          <p className={styles.heading}>Lets get a pipeline setup</p>
          <p className={styles.subHeading}>
            First, a few details on how pipelines work. We’ll help you set it
            all up.
          </p>
        </div>
        <div className={styles.pipelineInfo}>
          <div className={styles.note}>
            <div className={styles.exclamation}>!</div>
            <p>
              T-hex projects gets work done by running <span>pipelines</span>
            </p>
          </div>
          <div className={styles.pipelineSteps}>
            <ul className={styles.stepList}>
              {pipelineSteps.map((step, index) => (
                <li key={index} className={styles.step}>
                  <div className={styles.icon}>
                    <Image
                      src={step.icon}
                      alt={`${step.title} icon`}
                      width={24}
                      height={24}
                    />
                  </div>
                  <div className={styles.text}>
                    <p className={styles.title}>{step.title}</p>
                    <p className={styles.description}>{step.description}</p>
                  </div>
                </li>
              ))}
            </ul>
          </div>
        </div>


        <PipelineForm
          onStepChange={setCurrentStep}
          onStepComplete={(stepIndex) => {
            setCompletedSteps(prev => [...prev, stepIndex]);
          }}
          onStepIncomplete={(stepIndex) => {
            setCompletedSteps(prev => prev.filter(step => step !== stepIndex));
          }}
        />

      </div>
    </section>
  );
}
