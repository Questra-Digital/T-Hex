import PipelinesList from "@/components/PipelinesDashboard/PipelinesList/PipelinesList";
import AnalyticsPanel from "@/components/PipelinesDashboard/AnalyticsPanel/AnalyticsPanel";
import styles from "@/styles/pages/PipelinesPage.module.scss";
import { GetPipelinesResponse, Pipeline } from "@/types/pipeline";
import {getPipelines} from "@/services/pipelineService";


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

export default async function PipelinesPage() {
  const response: GetPipelinesResponse = await getPipelines();
  const pipelines: Pipeline[] = response.success ? response.data?.pipelines ?? [] : [];
  const error = response.success ? null : response.error || response.message || "Failed to fetch pipelines";

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
