# Prometheus GitLab CI Exporter

A lightweight [Prometheus](https://prometheus.io/) exporter that collects and exposes GitLab CI/CD pipeline metrics for specific projects and groups.
This allows you to monitor the health, duration, and success rates of pipelines in GitLab using Prometheus and Grafana.

---

## 🚀 Features

- Collects the **total number of pipelines** per status (success, failed, running, etc.)
- Exposes the **latest pipeline duration** (in seconds)
- Provides a **probe success metric** to track scraping health
- Supports GitLab **group/project query parameters** for flexible scraping
- Simple and minimal — ideal for per-project monitoring setups

---

## 📦 Installation

### Docker

```shell
docker run ghcr.io/aleksanderwww/prometheus-gitlabci-exporter:0.1.0
```

### From GitHub releases

```shell
wget https://github.com/AleksanderWWW/prometheus-gitlabci-exporter/releases/download/0.1.0/prometheus_gitlabci_exporter-0.1.0.tar.gz

tar -xzf prometheus_gitlabci_exporter-0.1.0.tar.gz

chmod +x prometheus_gitlabci_exporter

./prometheus_gitlabci_exporter
```

### From source

#### 1. Clone the repository
```bash
git clone https://github.com/AleksanderWWW/prometheus-gitlabci-exporter.git
cd prometheus-gitlabci-exporter
```

#### 2. Build the binary
```bash
go build -o prometheus_gitlabci_exporter
```

#### 3. Run it
You can provide your GitLab API token either via flag or environment variable.

```shell
export GITLAB_API_TOKEN="glpat-yourtoken"
./prometheus_gitlabci_exporter
```

or

```bash
./prometheus_gitlabci_exporter --api.token="glpat-yourtoken"
```

---

## ⚙️ Usage

### Command-line flags

| Flag | Description | Default |
|------|--------------|----------|
| `--api.token` | GitLab API token used to authenticate requests | *(required)* |
| `--web.port` | Address and port to listen on | `:9115` |

### Example

```bash
./gitlabci-exporter \
  --api.token="glpat-yourtoken" \
  --web.port=":9115"
```

---

## 🔍 Metrics Endpoint

The exporter exposes a `/probe` HTTP endpoint which Prometheus can scrape dynamically.

### Example Request

```bash
curl "http://localhost:9115/probe?group=my-group&project=my-project"
```

### Example Response (metrics)

```
# HELP gitlab_pipeline_total Total number of pipelines for the target project.
# TYPE gitlab_pipeline_total counter
gitlab_pipeline_total{group="my-group",project="my-project",status="success"} 12
gitlab_pipeline_total{group="my-group",project="my-project",status="failed"} 3
...

# HELP gitlab_pipeline_last_duration_seconds Duration of the latest pipeline in seconds.
# TYPE gitlab_pipeline_last_duration_seconds gauge
gitlab_pipeline_last_duration_seconds 154.23

# HELP gitlab_probe_success Whether the probe was successful (1 - success, 0 - failure).
# TYPE gitlab_probe_success gauge
gitlab_probe_success 1
```

---

## 🧠 Architecture Overview

```
+-------------------+          +-------------------+
| Prometheus Server | <------> | Exporter (/probe) |
+-------------------+          +-------------------+
                                      |
                                      v
                              GitLab API (Client)
```

The exporter:
1. Accepts HTTP requests from Prometheus with `group` and `project` query parameters.
2. Queries the GitLab API to list pipelines and get the latest one’s duration.
3. Exposes aggregated metrics over `/probe`.

---

## 📊 Example Prometheus Scrape Config

Add this job to your `prometheus.yml`:

```yaml
scrape_configs:
  - job_name: 'gitlabci_exporter'
    metrics_path: /probe
    static_configs:
      - targets:
          - 'localhost:9115'
    params:
      group: ['my-group']
      project: ['my-project']
```

---

## 📘 Exposed Metrics Summary

| Metric | Type | Description |
|--------|------|-------------|
| `gitlab_pipeline_total` | Counter | Number of pipelines by status |
| `gitlab_pipeline_last_duration_seconds` | Gauge | Duration of the latest pipeline |
| `gitlab_probe_success` | Gauge | Indicates if the exporter successfully queried GitLab |

---

## 🛠️ Configuration Environment Variables

| Variable | Description |
|-----------|-------------|
| `GITLAB_API_TOKEN` | GitLab API token (used if `--api.token` flag not provided) |

---

## 📜 License

MIT License © 2025 [AleksanderWWW](https://github.com/AleksanderWWW)
