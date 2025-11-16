import { fetchSeleniumSessions} from "@/services/test_cases";
import SeleniumSessions from "@/components/TestCases/SeleniumSessions";
import { GetSeleniumSessionsResponse } from "@/types/pipeline";
import styles from "@/styles/pages/TestCasesPage.module.scss";

export default async function TestSessionsPage({
  params,
}: {
  params: Promise<{ user_id: string; pipeline_id: string; pipeline_event_id: string }>;
}) {
  const { user_id, pipeline_id, pipeline_event_id } = await params; // ✅ await the promise

  const response: GetSeleniumSessionsResponse = await fetchSeleniumSessions(
    pipeline_id,
    pipeline_event_id,
    parseInt(user_id)
  );

  console.log("API Response:", response);

  if (!response.success) {
    return (
      <div className={styles.testCasesPage}>
        <div className={styles.container}>
          <div className={styles.header}>
            <h1 className={styles.title}>Test <span>Cases</span></h1>
            <p className={styles.subtitle}>Monitor and track your selenium test sessions</p>
          </div>
          <div className={styles.dashboardContent}>
            <div className={styles.contentWrapper}>
              <div>Error: {response.message}</div>
            </div>
          </div>
        </div>
      </div>
    );
  }

  // Handle case where no test sessions exist
  if (!response.data || !response.data.selenium_sessions || response.data.selenium_sessions.length === 0) {
    return (
      <div className={styles.testCasesPage}>
        <div className={styles.container}>
          <div className={styles.header}>
            <h1 className={styles.title}>Test <span>Cases</span></h1>
            <p className={styles.subtitle}>Monitor and track your selenium test sessions</p>
          </div>
          <div className={styles.dashboardContent}>
            <div className={styles.contentWrapper}>
              <div>No test sessions found for this pipeline event.</div>
            </div>
          </div>
        </div>
      </div>
    );
  }

  console.log("Test Session Data:", response.data.selenium_sessions[0]);

  return (
    <div className={styles.testCasesPage}>
      <div className={styles.container}>
        <div className={styles.header}>
          <h1 className={styles.title}>Test <span>Cases</span></h1>
          <p className={styles.subtitle}>Monitor and track your selenium test sessions</p>
        </div>
        <div className={styles.dashboardContent}>
          <div className={styles.contentWrapper}>
            <SeleniumSessions 
              seleniumSessions={response.data.selenium_sessions!} 
              userId={user_id}
              pipelineId={pipeline_id}
              pipelineEventId={pipeline_event_id}
            />
          </div>
        </div>
      </div>
    </div>
  );
}
