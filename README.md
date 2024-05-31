# Kod-Mono

[![Build and Test](https://github.com/go-kod/kod-mono/actions/workflows/go.yml/badge.svg)](https://github.com/go-kod/kod-mono/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/go-kod/kod-mono/graph/badge.svg?token=BvQttVqscO)](https://codecov.io/gh/go-kod/kod-mono)

## Domain Driven Design

![domain driven design](./assets/ddd.excalidraw.png)

## CallGraph (Auto Generated)

![call graph](./assets/callgraph.png)

## Run

```bash
export OTEL_EXPORTER_OTLP_ENDPOINT="http://localhost:14318"
export OTEL_EXPORTER_OTLP_HEADERS="uptrace-dsn=http://project2_secret_token@localhost:14318?grpc=14317"
export OTEL_EXPORTER_OTLP_COMPRESSION=gzip
export OTEL_EXPORTER_OTLP_METRICS_DEFAULT_HISTOGRAM_AGGREGATION=BASE2_EXPONENTIAL_BUCKET_HISTOGRAM
export OTEL_EXPORTER_OTLP_METRICS_TEMPORALITY_PREFERENCE=DELTA

KOD_CONFIG=./config/server/dev.toml go run ./cmd/server
```

## Uptrace

### Overview

![overview](./assets/uptrace-overview.png)

### Traces & Logs

![traces](./assets/uptrace-traces-and-logs.png)

### Metrics

![metrics](./assets/uptrace-metrics.png)
