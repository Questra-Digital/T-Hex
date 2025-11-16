"use client";

import { SessionSelenium } from "@/types/pipeline";
import { CheckCircle, XCircle, Clock, Activity, Calendar, Hash, ArrowRight } from "lucide-react";
import styles from "./SeleniumSessions.module.scss";
import { useRouter } from "next/navigation";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { fetchSeleniumSessionsAsync, selectSessionsCache } from "@/store/testCasesSlice";
import { useEffect } from "react";
import BackButton from "@/components/BackButton";

interface SeleniumSessionsProps {
  seleniumSessions: SessionSelenium[];
  userId: string;
  pipelineId: string;
  pipelineEventId: string;
}

export default function SeleniumSessions({ seleniumSessions, userId, pipelineId, pipelineEventId }: SeleniumSessionsProps) {
    const router = useRouter();
    const dispatch = useAppDispatch();
    
    // Create cache key
    const cacheKey = `${pipelineId}-${pipelineEventId}-${userId}`;
    const cachedData = useAppSelector(state => selectSessionsCache(state, cacheKey));
    
    // Use cached data if available, otherwise use props
    const sessions = cachedData?.data || seleniumSessions;
    const loading = cachedData?.loading || false;
    const error = cachedData?.error || null;
    const formatTime = (timestamp: number) => {
        return new Date(timestamp * 1000).toLocaleString();
    };

    const getStatusIcon = (status: boolean) => {
        return status ? (
            <div className={styles.successIcon}>
                <CheckCircle size={12} />
            </div>
        ) : (
            <div className={styles.failedIcon}>
                <XCircle size={12} />
            </div>
        );
    };

    const getStatusColor = (status: boolean) => {
        return status ? "#3D8D7A" : "#ef4444";
    };

    const getStatusText = (status: boolean) => {
        return status ? "Success" : "Failed";
    };

    const getExecutionStatus = (valid: boolean) => {
        return valid ? "Running" : "Completed";
    };

    const getExecutionColor = (valid: boolean) => {
        return valid ? "#1B39CE" : "#3D8D7A";
    };

    const getExecutionIcon = (valid: boolean) => {
        return valid ? (
            <div className={styles.runningIcon}>
                <Activity size={12} />
            </div>
        ) : (
            <div className={styles.completedIcon}>
                <CheckCircle size={12} />
            </div>
        );
    };

    const handleSessionClick = (sessionId: string) => {
        router.push(`/test_cases/${userId}/${pipelineId}/${pipelineEventId}/${sessionId}`);
    };

    const handleBackToPipelines = () => {
        router.push('/pipelines');
    };

    if (loading) {
        return (
            <div className={styles.seleniumSessionsContainer}>
                <BackButton onClick={handleBackToPipelines} label="Back to Pipelines" />
                <div className={styles.loadingMessage}>
                    <Clock size={48} />
                    <h3>Loading Sessions...</h3>
                    <p>Please wait while we fetch the selenium sessions.</p>
                </div>
            </div>
        );
    }

    if (error) {
        return (
            <div className={styles.seleniumSessionsContainer}>
                <BackButton onClick={handleBackToPipelines} label="Back to Pipelines" />
                <div className={styles.errorMessage}>
                    <XCircle size={48} />
                    <h3>Error Loading Sessions</h3>
                    <p>{error}</p>
                </div>
            </div>
        );
    }

    return (
        <div className={styles.seleniumSessionsContainer}>
            <BackButton onClick={handleBackToPipelines} label="Back to Pipelines" />
            {sessions.length > 0 ? (
                <div className={styles.sessionsTable}>
                    <div className={styles.tableHeader}>
                        <div className={styles.headerCell}>Session ID</div>
                        <div className={styles.headerCell}>Test ID</div>
                        <div className={styles.headerCell}>Execution Status</div>
                        <div className={styles.headerCell}>Test Result</div>
                        <div className={styles.headerCell}>Message</div>
                        <div className={styles.headerCell}>Timestamp</div>
                        <div className={styles.headerCell}>Actions</div>
                    </div>
                    
                    <div className={styles.tableBody}>
                        {sessions.map((session) => (
                            <div 
                                key={session.session_id} 
                                className={styles.tableRow}
                                onClick={() => handleSessionClick(session.session_id)}
                            >
                                <div className={styles.cell}>
                                    <div className={styles.sessionIdCell}>
                                        <Hash size={14} />
                                        <span className={styles.sessionId}>{session.session_id}</span>
                                    </div>
                                </div>
                                
                                <div className={styles.cell}>
                                    <div className={styles.testIdCell}>
                                        <Activity size={14} />
                                        <span>{session.test_id}</span>
                                    </div>
                                </div>
                                
                                <div className={styles.cell}>
                                    <div 
                                        className={styles.executionBadge}
                                        style={{ backgroundColor: getExecutionColor(session.valid) }}
                                    >
                                        {getExecutionIcon(session.valid)}
                                        <span>{getExecutionStatus(session.valid)}</span>
                                    </div>
                                </div>
                                
                                <div className={styles.cell}>
                                    <div 
                                        className={styles.statusBadge}
                                        style={{ backgroundColor: getStatusColor(session.status) }}
                                    >
                                        {getStatusIcon(session.status)}
                                        <span>{getStatusText(session.status)}</span>
                                    </div>
                                </div>
                                
                                <div className={styles.cell}>
                                    <div className={styles.messageCell}>
                                        <span className={styles.messageText}>
                                            {session.message || "No message"}
                                        </span>
                                    </div>
                                </div>
                                
                                <div className={styles.cell}>
                                    <div className={styles.timestampCell}>
                                        <Calendar size={14} />
                                        <span>{formatTime(session.time)}</span>
                                    </div>
                                </div>
                                
                                <div className={styles.cell}>
                                    <div className={styles.actionCell}>
                                        <ArrowRight size={16} />
                                    </div>
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            ) : (
                <div className={styles.noSessionsMessage}>
                    <Clock size={48} />
                    <h3>No Selenium Sessions Found</h3>
                    <p>No selenium sessions are available for this test case.</p>
                </div>
            )}
        </div>
    );
}