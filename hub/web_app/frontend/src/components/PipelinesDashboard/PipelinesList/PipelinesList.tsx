"use client";

import { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { ChevronDown, ChevronUp, Play, Settings, Trash2, Calendar, GitBranch, Clock, Check, X, GitCommit } from "lucide-react";
import styles from "./PipelinesList.module.scss";
import { useAppSelector } from "@/store/hooks";
import { selectPipelines } from "@/store/pipelinesSlice";
import { Pipeline } from "@/types/pipeline";

export default function PipelinesList() {
  const pipelines = useAppSelector(selectPipelines);
  const [expandedPipeline, setExpandedPipeline] = useState<string | null>(null);

  const togglePipeline = (pipelineId: string) => {
    setExpandedPipeline(expandedPipeline === pipelineId ? null : pipelineId);
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

  return (
    <div className={styles.pipelinesList}>
      {pipelines.map((pipeline: Pipeline) => (
        <motion.div
          key={pipeline.id}
          className={styles.pipelineCard}
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.3 }}
        >
          <div
            className={styles.pipelineHeader}
            onClick={() => togglePipeline(pipeline.id)}
          >
            <div className={styles.pipelineInfo}>
              <div className={styles.pipelineName}>
                <h3>{pipeline.name}</h3>
                <div
                  className={styles.statusBadge}
                  style={{ backgroundColor: getStatusColor(pipeline.status) }}
                >
                  {getStatusIcon(pipeline.status)}
                  <span>{pipeline.status}</span>
                </div>
              </div>
              <p className={styles.pipelineDescription}>
                {pipeline.description}
                {pipeline.triggerType === "commit" && (
                  <span className={styles.commitInfo}> • Commit to {pipeline.branchName} branch</span>
                )}
              </p>
              <div className={styles.pipelineMeta}>
                <div className={styles.metaItem}>
                  <GitBranch size={14} />
                  <span>{pipeline.branchName}</span>
                </div>
                <div className={styles.metaItem}>
                  <GitCommit size={14} />
                  <span>{pipeline.repositoryPath}</span>
                </div>
                <div className={styles.metaItem}>
                  <Calendar size={14} />
                  <span>Last run: {formatDate(pipeline.lastRun)}</span>
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
              {expandedPipeline === pipeline.id ? (
                <ChevronUp size={20} className={styles.expandIcon} />
              ) : (
                <ChevronDown size={20} className={styles.expandIcon} />
              )}
            </div>
          </div>

          <AnimatePresence>
            {expandedPipeline === pipeline.id && (
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
                      {pipeline.events.map((event) => (
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
                              <span className={styles.eventDuration}>{event.duration}</span>
                              <span className={styles.eventTime}>{formatDate(event.timestamp)}</span>
                            </div>
                          </div>
                          <p className={styles.eventDetails}>{event.details}</p>
                        </div>
                      ))}
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
