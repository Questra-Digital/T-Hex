"use client";

import { Event } from "@/types/pipeline";
import { CheckCircle, XCircle, Clock, Activity, Calendar, Hash, ArrowRight, Code } from "lucide-react";
import styles from "./SeleniumEvents.module.scss";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { fetchSeleniumEventsAsync, selectEventsCache } from "@/store/testCasesSlice";
import { useEffect } from "react";
import BackButton from "@/components/BackButton";
import { useRouter } from "next/navigation";

interface SeleniumEventsProps {
  events: Event[];
  sessionId: string;
  userId: string;
  pipelineId: string;
  pipelineEventId: string;
}

export default function SeleniumEvents({ events, sessionId, userId, pipelineId, pipelineEventId }: SeleniumEventsProps) {
    const router = useRouter();
    const dispatch = useAppDispatch();
    
    // Create cache key
    const cacheKey = `${pipelineId}-${pipelineEventId}-${userId}-${sessionId}`;
    const cachedData = useAppSelector(state => selectEventsCache(state, cacheKey));
    
    // Use cached data if available, otherwise use props
    const eventsData = cachedData?.data || events;
    const loading = cachedData?.loading || false;
    const error = cachedData?.error || null;
    const formatTime = (timestamp: number) => {
        return new Date(timestamp * 1000).toLocaleString();
    };

    const getStatusIcon = (status: number) => {
        return status >= 200 && status < 300 ? (
            <div className={styles.successIcon}>
                <CheckCircle size={12} />
            </div>
        ) : (
            <div className={styles.failedIcon}>
                <XCircle size={12} />
            </div>
        );
    };

    const getStatusColor = (status: number) => {
        if (status >= 200 && status < 300) return "#3D8D7A";
        if (status >= 400 && status < 500) return "#f59e0b";
        return "#ef4444";
    };

    const getStatusText = (status: number) => {
        if (status >= 200 && status < 300) return "Success";
        if (status >= 400 && status < 500) return "Client Error";
        if (status >= 500) return "Server Error";
        return "Unknown";
    };

    const handleBackToSessions = () => {
        router.push(`/test_cases/${userId}/${pipelineId}/${pipelineEventId}`);
    };

    if (loading) {
        return (
            <div className={styles.seleniumEventsContainer}>
                <BackButton onClick={handleBackToSessions} label="Back to Sessions" />
                <div className={styles.loadingMessage}>
                    <Clock size={48} />
                    <h3>Loading Events...</h3>
                    <p>Please wait while we fetch the selenium events.</p>
                </div>
            </div>
        );
    }

    if (error) {
        return (
            <div className={styles.seleniumEventsContainer}>
                <BackButton onClick={handleBackToSessions} label="Back to Sessions" />
                <div className={styles.errorMessage}>
                    <XCircle size={48} />
                    <h3>Error Loading Events</h3>
                    <p>{error}</p>
                </div>
            </div>
        );
    }

    return (
        <div className={styles.seleniumEventsContainer}>
            <BackButton onClick={handleBackToSessions} label="Back to Sessions" />
            <div className={styles.sessionInfo}>
                <div className={styles.sessionIdDisplay}>
                    <Hash size={16} />
                    <span>Session ID: {sessionId}</span>
                </div>
                <div className={styles.eventsCount}>
                    <Activity size={16} />
                    <span>{eventsData.length} Events</span>
                </div>
            </div>

            {eventsData.length > 0 ? (
                <div className={styles.eventsTable}>
                    <div className={styles.tableHeader}>
                        <div className={styles.headerCell}>Event ID</div>
                        <div className={styles.headerCell}>Method</div>
                        <div className={styles.headerCell}>Path</div>
                        <div className={styles.headerCell}>Status</div>
                        <div className={styles.headerCell}>Request Body</div>
                        <div className={styles.headerCell}>Response</div>
                        <div className={styles.headerCell}>Timestamp</div>
                    </div>
                    
                    <div className={styles.tableBody}>
                        {eventsData.map((event) => (
                            <div key={event.id} className={styles.tableRow}>
                                <div className={styles.cell}>
                                    <div className={styles.eventIdCell}>
                                        <Hash size={14} />
                                        <span className={styles.eventId}>{event.id}</span>
                                    </div>
                                </div>
                                
                                <div className={styles.cell}>
                                    <div className={styles.methodCell}>
                                        <Code size={14} />
                                        <span className={styles.method}>{event.method}</span>
                                    </div>
                                </div>
                                
                                <div className={styles.cell}>
                                    <div className={styles.pathCell}>
                                        <span className={styles.pathText}>
                                            {event.path}
                                        </span>
                                    </div>
                                </div>
                                
                                <div className={styles.cell}>
                                    <div 
                                        className={styles.statusBadge}
                                        style={{ backgroundColor: getStatusColor(event.status) }}
                                    >
                                        {getStatusIcon(event.status)}
                                        <span>{event.status}</span>
                                    </div>
                                </div>
                                
                                <div className={styles.cell}>
                                    <div className={styles.bodyCell}>
                                        <span className={styles.bodyText}>
                                            {event.req_body || "No body"}
                                        </span>
                                    </div>
                                </div>
                                
                                <div className={styles.cell}>
                                    <div className={styles.responseCell}>
                                        <span className={styles.responseText}>
                                            {event.res || "No response"}
                                        </span>
                                    </div>
                                </div>
                                
                                <div className={styles.cell}>
                                    <div className={styles.timestampCell}>
                                        <Calendar size={14} />
                                        <span>{formatTime(event.time)}</span>
                                    </div>
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            ) : (
                <div className={styles.noEventsMessage}>
                    <Clock size={48} />
                    <h3>No Events Found</h3>
                    <p>No events are available for this session.</p>
                </div>
            )}
        </div>
    );
}
