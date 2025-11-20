"use client";

import { TrendingUp, Activity, CheckCircle, Clock } from "lucide-react";
import styles from "./AnalyticsPanel.module.scss";

interface AnalyticsData {
  successRate: number;
  totalRuns: number;
  avgDuration: number;
  weeklyRuns: { day: string; runs: number }[];
  pipelineStats: { name: string; success: number; failed: number }[];
}

interface AnalyticsPanelProps {
  analytics: AnalyticsData;
}

export default function AnalyticsPanel({ analytics }: AnalyticsPanelProps) {
  // Analytics helper functions
  const getMaxRuns = () => {
    return Math.max(...analytics.weeklyRuns.map(day => day.runs));
  };

  const getBarHeight = (runs: number) => {
    const maxRuns = getMaxRuns();
    return (runs / maxRuns) * 100;
  };

  return (
    <div className={styles.analyticsPanel}>
      <div className={styles.analyticsHeader}>
        <h3 className={styles.analyticsTitle}>
          <TrendingUp size={20} />
          Analytics
        </h3>
      </div>

      {/* Key Metrics */}
      <div className={styles.metricsGrid}>
        <div className={styles.metricCard}>
          <div className={styles.metricIcon}>
            <CheckCircle size={16} />
          </div>
          <div className={styles.metricContent}>
            <div className={styles.metricValue}>{analytics.successRate}%</div>
            <div className={styles.metricLabel}>Success Rate</div>
          </div>
        </div>

        <div className={styles.metricCard}>
          <div className={styles.metricIcon}>
            <Activity size={16} />
          </div>
          <div className={styles.metricContent}>
            <div className={styles.metricValue}>{analytics.totalRuns.toLocaleString()}</div>
            <div className={styles.metricLabel}>Total Runs</div>
          </div>
        </div>

        <div className={styles.metricCard}>
          <div className={styles.metricIcon}>
            <Clock size={16} />
          </div>
          <div className={styles.metricContent}>
            <div className={styles.metricValue}>{analytics.avgDuration}m</div>
            <div className={styles.metricLabel}>Avg Duration</div>
          </div>
        </div>
      </div>

      {/* Weekly Runs Chart */}
      <div className={styles.chartSection}>
        <h4 className={styles.chartTitle}>Weekly Pipeline Runs</h4>
        <div className={styles.barChart}>
          {analytics.weeklyRuns.map((day, index) => (
            <div key={day.day} className={styles.barContainer}>
              <div 
                className={styles.bar}
                style={{ height: `${getBarHeight(day.runs)}%` }}
              />
              <div className={styles.barLabel}>{day.day}</div>
              <div className={styles.barValue}>{day.runs}</div>
            </div>
          ))}
        </div>
      </div>

      {/* Pipeline Success Rates */}
      <div className={styles.chartSection}>
        <h4 className={styles.chartTitle}>Pipeline Success Rates</h4>
        <div className={styles.pipelineStats}>
          {analytics.pipelineStats.map((pipeline, index) => {
            const total = pipeline.success + pipeline.failed;
            const successRate = (pipeline.success / total) * 100;
            
            return (
              <div key={pipeline.name} className={styles.pipelineStat}>
                <div className={styles.pipelineStatHeader}>
                  <span className={styles.pipelineStatName}>{pipeline.name}</span>
                  <span className={styles.pipelineStatRate}>{successRate.toFixed(1)}%</span>
                </div>
                <div className={styles.pipelineStatBar}>
                  <div 
                    className={styles.pipelineStatBarFill}
                    style={{ width: `${successRate}%` }}
                  />
                </div>
                <div className={styles.pipelineStatDetails}>
                  <span className={styles.successCount}>{pipeline.success} success</span>
                  <span className={styles.failedCount}>{pipeline.failed} failed</span>
                </div>
              </div>
            );
          })}
        </div>
      </div>
    </div>
  );
}
