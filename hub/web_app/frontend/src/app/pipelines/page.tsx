'use client';
import PipelinesList from "@/components/PipelinesDashboard/PipelinesList/PipelinesList";
import AnalyticsPanel from "@/components/PipelinesDashboard/AnalyticsPanel/AnalyticsPanel";
import styles from "@/styles/pages/PipelinesPage.module.scss";
import { Pipeline } from "@/types/pipeline";

import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { selectPipelines, setPipelines } from "@/store/pipelinesSlice";
import { useEffect } from "react";

//Stub Pipelines data for the pipelines list
export const initialPipelines: Pipeline[] = [
  {
    id: "1",
    name: "build-and-test",
    description: "Automated build and test pipeline for the main application",
    status: "running",
    lastRun: "2024-01-15T10:30:00Z",
    triggerType: "commit",
    branchName: "main",
    repositoryPath: "github.com/company/frontend-app",
    labels: ["frontend", "testing", "automation"],
    events: [
      {
        id: "1",
        type: "Test",
        status: "success",
        timestamp: "2024-01-15T10:30:00Z",
        duration: "2m 34s",
        details: "Unit tests completed successfully with 0 errors"
      },
      {
        id: "2",
        type: "Test",
        status: "success",
        timestamp: "2024-01-15T10:32:34Z",
        duration: "1m 12s",
        details: "Integration tests passed - All 156 tests passed"
      },
      {
        id: "3",
        type: "Test",
        status: "running",
        timestamp: "2024-01-15T10:33:46Z",
        duration: "0m 45s",
        details: "E2E tests running..."
      }
    ]
  },
  {
    id: "2",
    name: "security-scan",
    description: "Security vulnerability scanning and code analysis",
    status: "success",
    lastRun: "2024-01-15T09:15:00Z",
    triggerType: "manual",
    branchName: "develop",
    repositoryPath: "github.com/company/security-tools",
    labels: ["security", "analysis"],
    events: [
      {
        id: "4",
        type: "Test",
        status: "success",
        timestamp: "2024-01-15T09:15:00Z",
        duration: "5m 23s",
        details: "Security tests completed - no critical vulnerabilities found"
      }
    ]
  },
  {
    id: "3",
    name: "performance-test",
    description: "Performance testing and load testing pipeline",
    status: "failed",
    lastRun: "2024-01-15T08:45:00Z",
    triggerType: "commit",
    branchName: "feature/performance",
    repositoryPath: "github.com/company/performance-suite",
    labels: ["performance", "testing"],
    events: [
      {
        id: "5",
        type: "Test",
        status: "failed",
        timestamp: "2024-01-15T08:45:00Z",
        duration: "3m 12s",
        details: "Performance tests failed - response time exceeded threshold"
      }
    ]
  }
];

// Simulated analytics data from the backend server
const stubAnalytics = {
  successRate: 87.5,
  totalRuns: 1247,
  avgDuration: 4.2,
  weeklyRuns: [
    { day: "Mon", runs: 45 },
    { day: "Tue", runs: 52 },
    { day: "Wed", runs: 38 },
    { day: "Thu", runs: 61 },
    { day: "Fri", runs: 48 },
    { day: "Sat", runs: 23 },
    { day: "Sun", runs: 19 }
  ],
  pipelineStats: [
    { name: "build-and-test", success: 89, failed: 11 },
    { name: "security-scan", success: 95, failed: 5 },
    { name: "performance-test", success: 78, failed: 22 }
  ]
};

export default function PipelinesPage() {
  const dispatch = useAppDispatch();
  //Hydrate store with initial pipelines
  useEffect(() => {
    if (initialPipelines && initialPipelines.length > 0) {
      dispatch(setPipelines(initialPipelines));
    }
  }, [dispatch, initialPipelines]);

  return (
    <div className={`${styles.pipelinesPage}`}>
      <div className={styles.container}>
        <div className={styles.header}>
          <h1 className={styles.title}>Pipelines <span>Dashboard</span></h1>
          <p className={styles.subtitle}>Manage and monitor your CI/CD pipelines</p>
        </div>

        <div className={styles.dashboardContent}>
          <PipelinesList />
          <AnalyticsPanel analytics={stubAnalytics} />
        </div>
      </div>
    </div>
  );
}
