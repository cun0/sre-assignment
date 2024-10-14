# SRE Assignment: Monitoring & Observability Setup

The goal of this assignment is to design a monitoring strategy for a Node.js API service running on AWS. Your task is to define and implement key metrics, dashboards, and alerts to ensure observability and reliability.

## Instructions

Please follow the steps below to complete the assignment.
### 1. Metrics to Implement

Design monitoring around the following key areas:

* Application Metrics:
    * **Request Rate:** Number of requests per second (RPS) handled by the API.
    * **Error Rate:** Percentage of failed requests (e.g., HTTP 500 errors).
    * **Latency (p99):** Track the 99th percentile latency to measure user experience during peak loads.
* Infrastructure Metrics:
    * **CPU Utilization:** Monitor the percentage of CPU usage on ECS/EC2 instances.
    * **Memory Utilization:** Monitor available memory to prevent out-of-memory (OOM) issues.

### 2. Tool Selection

We recommend using prometheus and grafana for monitoring and alerting. However, you can use any other monitoring tool you are comfortable with.

### 3. Alert Configuration

Set up the following alerts:

1. **Error Rate Alert:**
    * Trigger if the error rate exceeds 5% for 5 consecutive minutes.
2. **Latency Alert:**
    * Trigger if p99 latency is greater than 500ms for 10 consecutive minutes.
3. **CPU or Memory Alert (Optional):**
    * Trigger if CPU utilization >80% or memory utilization drops below 20%.

### 4. Deliverables
1. Dashboard Screenshots:
    * Provide screenshots of dashboards with the metrics implemented.
2. Alert Configuration:
    * Include configuration files or descriptions of the alert rules (e.g., CloudWatch alarms or Prometheus alert rules).
3. Short Explanation Document:
    * A brief summary explaining:
        * Why you selected the specific metrics.
        * How your alert thresholds were defined to avoid alert fatigue.
        * Any trade-offs or assumptions made.

### 5. Submission Instructions

* Fork this repository (or create a new one).
* Add your code, configurations, and documents to the repository.
* Provide a link to the repository or submit your deliverables as a ZIP file.

