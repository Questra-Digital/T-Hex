"use client";

import { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { ChevronDown, ChevronUp, Play, Settings, Trash2, Calendar, GitBranch, Clock, Check, X, GitCommit, Eye } from "lucide-react";
import styles from "./PipelinesList.module.scss";
import { Pipeline, PipelineEvent } from "@/types/pipeline";

import { useSnackbar } from "@/contexts/SnackbarContext";
import { useEffect } from "react";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { addEventToPipeline, selectPipelines, setPipelines } from "@/store/pipelinesSlice";
import { selectUserData } from "@/store/userData";
import { useRouter } from "next/navigation";

export default function PipelinesList({ initialPipelines, error }: { initialPipelines: Pipeline[], error: string | null }) {
  const [expandedPipeline, setExpandedPipeline] = useState<string | null>(null);

  const { showSnackbar } = useSnackbar();
  const dispatch = useAppDispatch();
  const pipelines: Pipeline[] = useAppSelector((state: any) => selectPipelines(state));
  const userData = useAppSelector((state: any) => selectUserData(state));
  const router = useRouter();

  //SSE Connection to get pipeline events
  useEffect(() => {
    let eventSource: EventSource | null = null;
    let reconnectTimeout: NodeJS.Timeout | null = null;
    let reconnectAttempts = 0;
    const maxReconnectAttempts = 5;
    const reconnectDelay = 1000; // Start with 1 second

    const connect = () => {
      if (eventSource) {
        eventSource.close();
      }

      eventSource = new EventSource("/api/pipeline-status");

      eventSource.onopen = () => {
        console.log("SSE connection opened");
        reconnectAttempts = 0; // Reset attempts on successful connection
      };

      eventSource.onmessage = (event) => {
        try {
          const data = JSON.parse(event.data) as {
            pipeline_id: number;
            event: PipelineEvent;
          };
          
          console.log("SSE event received:", data);
          dispatch(addEventToPipeline({ pipelineId: data.pipeline_id, event: data.event }));
        } catch (err) {
          console.error("Error parsing SSE event:", err);
        }
      };

      eventSource.onerror = (err) => {
        console.error("SSE connection error:", err);
        
        // Only attempt reconnection if we haven't exceeded max attempts
        if (reconnectAttempts < maxReconnectAttempts) {
          reconnectAttempts++;
          const delay = reconnectDelay * Math.pow(2, reconnectAttempts - 1); // Exponential backoff
          
          console.log(`Attempting to reconnect in ${delay}ms (attempt ${reconnectAttempts}/${maxReconnectAttempts})`);
          
          reconnectTimeout = setTimeout(() => {
            connect();
          }, delay);
        } else {
          console.error("Max reconnection attempts reached. SSE connection closed.");
          if (eventSource) {
            eventSource.close();
          }
        }
      };
    };

    // Initial connection
    connect();

    return () => {
      if (reconnectTimeout) {
        clearTimeout(reconnectTimeout);
      }
      if (eventSource) {
        eventSource.close();
      }
    };
  }, [dispatch]);

  //Hydrate store with initial pipelines
  useEffect(() => {
    if (error) {
      showSnackbar(error, "error", 5000);
      return;
    }
    if (initialPipelines.length > 0) {
      dispatch(setPipelines(initialPipelines));
    }
  }, [error, initialPipelines]);

  const togglePipeline = (pipelineId: string) => {
    setExpandedPipeline(expandedPipeline === pipelineId ? null : pipelineId);
  };

  const handleShowDetails = (event: PipelineEvent, pipelineId: number) => {
    const user_id = userData?.user_id || 0;
    const url = `/test_cases/${user_id}/${pipelineId}/${event.id}`;
    console.log('Navigating to URL:', url);
    router.push(url);
  };
  

  const getLatestEventStatus = (pipeline: Pipeline): string => {
    if (pipeline.events && pipeline.events.length > 0) {
      // Sort events by timestamp descending to get the latest
      const sortedEvents = [...pipeline.events].sort((a, b) => 
        new Date(b.timestamp).getTime() - new Date(a.timestamp).getTime()
      );
      return sortedEvents[0].status;
    }
    return pipeline.status; // fallback to pipeline status if no events
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case "success":
        return "#3D8D7A";
      case "failed":
        return "#ef4444";
      case "running":
        return "#1B39CE";
      case "pending":
        return "#f59e0b";
      default:
        return "#8F8A8A";
    }
  };

  const getStatusIcon = (status: string) => {
    switch (status) {
      case "running":
        return <Play size={16} />;
      case "success":
        return <div className={styles.successIcon}> <Check size={16} /> </div>;
      case "failed":
        return <div className={styles.failedIcon}> <X size={16} /> </div>;
      default:
        return <Clock size={16} />;
    }
  };

  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleString();
  };

  const formatDuration = (duration: number) => {
    // Convert nanoseconds to total seconds
    let totalSeconds = Math.floor(duration / 1_000_000_000);

    const hours = Math.floor(totalSeconds / 3600);
    totalSeconds %= 3600;

    const minutes = Math.floor(totalSeconds / 60);
    const seconds = totalSeconds % 60;

    const parts = [];
    if (hours > 0) parts.push(`${hours}h`);
    if (minutes > 0) parts.push(`${minutes}m`);
    if (seconds > 0 || parts.length === 0) parts.push(`${seconds}s`);

    return parts.join(" ");
  };

  return (
    <div className={styles.pipelinesList}>
      {pipelines.map((pipeline: Pipeline, index: number) => (
        <motion.div
          key={pipeline.id?.toString() || `pipeline-${index}`}
          className={styles.pipelineCard}
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.3 }}
        >
          <div
            className={styles.pipelineHeader}
            onClick={() => togglePipeline(pipeline.id?.toString() || '')}
          >
            <div className={styles.pipelineInfo}>
              <div className={styles.pipelineName}>
                <h3>{pipeline.name}</h3>
                <div
                  className={styles.statusBadge}
                  style={{ backgroundColor: getStatusColor(getLatestEventStatus(pipeline)) }}
                >
                  {getStatusIcon(getLatestEventStatus(pipeline))}
                  <span>{getLatestEventStatus(pipeline)}</span>
                </div>
              </div>
              <p className={styles.pipelineDescription}>
                {pipeline.description}
                {pipeline.trigger_type === "commit" && (
                  <span className={styles.commitInfo}> • Commit to {pipeline.branch_name} branch</span>
                )}
              </p>
              <div className={styles.pipelineMeta}>
                <div className={styles.metaItem}>
                  <GitBranch size={14} />
                  <span>{pipeline.branch_name}</span>
                </div>
                <div className={styles.metaItem}>
                  <GitCommit size={14} />
                  <span>{pipeline.repository_path}</span>
                </div>
                <div className={styles.metaItem}>
                  <Calendar size={14} />
                  <span>Last run: {formatDate(pipeline.last_run)}</span>
                </div>
                <div className={styles.labels}>
                  {pipeline.labels.map((label) => (
                    <span key={label} className={styles.label}>
                      {label}
                    </span>
                  ))}
                </div>
              </div>
            </div>
            <div className={styles.pipelineActions}>
              <button className={styles.actionButton}>
                <Settings size={16} />
              </button>
              <button className={styles.actionButton}>
                <Trash2 size={16} />
              </button>
              {expandedPipeline === pipeline.id?.toString() ? (
                <ChevronUp size={20} className={styles.expandIcon} />
              ) : (
                <ChevronDown size={20} className={styles.expandIcon} />
              )}
            </div>
          </div>

          <AnimatePresence>
            {expandedPipeline === pipeline.id?.toString() && (
              <motion.div
                className={styles.pipelineDetails}
                initial={{ height: 0, opacity: 0 }}
                animate={{ height: "auto", opacity: 1 }}
                exit={{ height: 0, opacity: 0 }}
                transition={{ duration: 0.3 }}
              >
                <div className={styles.detailsContent}>
                  <div className={styles.eventsSection}>
                    <h4 className={styles.sectionTitle}>Recent Events</h4>
                    <div className={styles.eventsList}>
                      {pipeline.events.length > 0 ? (
                        pipeline.events.map((event: PipelineEvent) => (
                          <div key={event.id} className={styles.eventItem}>
                            <div className={styles.eventHeader}>
                              <div className={styles.eventType}>
                                <span className={styles.eventTypeBadge}>{event.type}</span>
                                <div
                                  className={styles.eventStatus}
                                  style={{ backgroundColor: getStatusColor(event.status) }}
                                >
                                  {getStatusIcon(event.status)}
                                </div>
                              </div>
                              <div className={styles.eventMeta}>
                                <span className={styles.eventDuration}>{formatDuration(event.duration)}</span>
                                <span className={styles.eventTime}>{formatDate(event.timestamp)}</span>
                              </div>
                            </div>
                            <div className={styles.eventContent}>
                              <p className={styles.eventDetails}>{event.details}</p>
                              {event.status === "success"  && (
                                <button 
                                  className={styles.showDetailsButton}
                                  onClick={() => handleShowDetails(event, pipeline.id!)}
                                >
                                  <Eye size={14} />
                                  Show Details
                                </button>
                              )}
                            </div>
                          </div>
                        ))
                      ) : (
                        <div className={styles.noEventsMessage}>
                          <p>No events exist for this pipeline yet.</p>
                        </div>
                      )}
                    </div>
                  </div>
                </div>
              </motion.div>
            )}
          </AnimatePresence>
        </motion.div>
      ))}
    </div>
  );
}
