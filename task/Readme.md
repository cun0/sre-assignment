I developed an API server to send test requests and observe error conditions. I also set up a Node Exporter container to monitor node performance.

In the alert_rules.yml, I defined alert rules for key metrics, such as error rates and latency, with specific thresholds for timely issue response. The prometheus.yml was configured for scraping API and node metrics, and I set up a data source in Grafana.

To streamline deployment, I created a Docker Compose file for Grafana and Prometheus, mounting volumes for data persistence. This allowed me to start the infrastructure and effectively monitor the API server's performance.


Why I Selected the Specific Metrics

Request Rate: Monitors API load to identify sudden traffic changes, which may indicate issues.
Error Rate: Tracks failed requests to assess user satisfaction.
Latency (p99): Measures 99th percentile latency to evaluate user experience during peak loads.
CPU Utilization: Prevents performance degradation and ensures service reliability.
Memory Utilization: Helps avoid Out Of Memory (OOM) issues and maintain service stability.

How My Alert Thresholds Were Defined to Avoid Alert Fatigue

Error Rate Alert: Triggered if the error rate exceeds 5% for 5 minutes, filtering out short-lived spikes.
Latency Alert: Triggers if p99 latency exceeds 500ms for 10 minutes, focusing on ongoing performance problems.
CPU/Memory Alerts: Set at 80% CPU or 20% memory utilization, maintaining clear operational limits without overwhelming alerts.

Any Trade-offs or Assumptions Made

Traffic Patterns: Assumes predictable patterns, requiring threshold adjustments during spikes.
Resource Availability: Relies on adequate resource provisioning.
Simplicity vs. Detail: Focuses on high-impact metrics, potentially missing finer details.
Historical Context: Initial thresholds set without extensive data; adjustments expected based on performance patterns.