import { fetchSeleniumEvents } from "@/services/test_cases";
import SeleniumEvents from "@/components/TestCases/SeleniumEvents";
import { GetSeleniumEventsResponse } from "@/types/pipeline";
import styles from "@/styles/pages/TestCasesPage.module.scss";

export default async function SeleniumEventsPage({
  params,
}: {
  params: Promise<{ 
    user_id: string; 
    pipeline_id: string; 
    pipeline_event_id: string; 
    session_id: string;
  }>;
}) {
  const { user_id, pipeline_id, pipeline_event_id, session_id } = await params;

  const response: GetSeleniumEventsResponse = await fetchSeleniumEvents(
    pipeline_id,
    pipeline_event_id,
    parseInt(user_id),
    session_id
  );

  console.log("Selenium Events API Response:", response);

  if (!response.success) {
    return (
      <div className={styles.testCasesPage}>
        <div className={styles.container}>
          <div className={styles.header}>
            <h1 className={styles.title}>Session <span>Events</span></h1>
            <p className={styles.subtitle}>Monitor and track selenium session events</p>
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

  // Handle case where no events exist
  if (!response.data || !response.data.selenium_events || response.data.selenium_events.length === 0) {
    return (
      <div className={styles.testCasesPage}>
        <div className={styles.container}>
          <div className={styles.header}>
            <h1 className={styles.title}>Session <span>Events</span></h1>
            <p className={styles.subtitle}>Monitor and track selenium session events</p>
          </div>
          <div className={styles.dashboardContent}>
            <div className={styles.contentWrapper}>
              <div>No events found for this session.</div>
            </div>
          </div>
        </div>
      </div>
    );
  }

  console.log("Selenium Events Data:", response.data.selenium_events);

  return (
    <div className={styles.testCasesPage}>
      <div className={styles.container}>
        <div className={styles.header}>
          <h1 className={styles.title}>Session <span>Events</span></h1>
          <p className={styles.subtitle}>Monitor and track selenium session events</p>
        </div>
        <div className={styles.dashboardContent}>
          <div className={styles.contentWrapper}>
            <SeleniumEvents 
              events={response.data.selenium_events!} 
              sessionId={session_id}
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
