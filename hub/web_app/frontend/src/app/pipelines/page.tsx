"use client";

import PipelinesList from "@/components/PipelinesDashboard/PipelinesList/PipelinesList";
import AnalyticsPanel from "@/components/PipelinesDashboard/AnalyticsPanel/AnalyticsPanel";
import styles from "@/styles/pages/PipelinesPage.module.scss";
import { GetPipelinesResponse, Pipeline } from "@/types/pipeline";
import {getPipelines} from "@/services/pipelineService";
import { useAppSelector } from "@/store/hooks";
import { selectUserData } from "@/store/userData";
import { useEffect, useState } from "react";

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
  const userData = useAppSelector(selectUserData);
  const [pipelines, setPipelines] = useState<Pipeline[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchPipelines = async () => {
      if (!userData.user_id || userData.user_id === 0) {
        setError("User not authenticated");
        setLoading(false);
        return;
      }

      try {
        const response: GetPipelinesResponse = await getPipelines(userData.user_id);
        if (response.success) {
          setPipelines(response.data?.pipelines ?? []);
          setError(null);
        } else {
          setError(response.error || response.message || "Failed to fetch pipelines");
        }
      } catch (err) {
        console.log("Error fetching pipelines:", err);
        setError("Failed to fetch pipelines");
      } finally {
        setLoading(false);
      }
    };

    fetchPipelines();
  }, [userData.user_id]);
  
  

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div className={`${styles.pipelinesPage}`}>
      <div className={styles.container}>
        <div className={styles.header}>
          <h1 className={styles.title}>Pipelines <span>Dashboard</span></h1>
          <p className={styles.subtitle}>Manage and monitor your CI/CD pipelines</p>
        </div>

        <div className={styles.dashboardContent}>
          <PipelinesList initialPipelines={pipelines} error={error} />
          <AnalyticsPanel analytics={stubAnalytics} />
        </div>
      </div>
    </div>
  );
}
