
## Info

* ingestion logs example https://github.com/opensearch-project/data-prepper/tree/main/examples/log-ingestion
* http-source plugin https://github.com/opensearch-project/data-prepper/tree/main/data-prepper-plugins/http-source
* all data prepper plugins https://github.com/opensearch-project/data-prepper/tree/main/data-prepper-plugins
* examples configuration https://github.com/opensearch-project/data-prepper/tree/main/examples
* trace analytics https://github.com/opensearch-project/data-prepper/blob/main/docs/trace_analytics.md

## Tracing setup

Trace analytics doc for OpenSearch you can find here https://github.com/opensearch-project/data-prepper/blob/main/docs/trace_analytics.md.

You can test traces ingestion with [otelgen](https://github.com/krzko/otelgen) by creating some traces and pushing them
to `data-prepper-server` service on port `21890`:

```bash
kubectl run otelgen --image=ghcr.io/krzko/otelgen:latest --restart=Never -n opensearch -- --otel-exporter-otlp-endpoint data-prepper-server:21890 --insecure --log-level debug traces multi
```

The process is the same for otel collector. Here is an example:

```bash
kubectl run otelgen --image=ghcr.io/krzko/otelgen:latest --restart=Never -n tracing -- --otel-exporter-otlp-endpoint tracing-opentelemetry-collector:4317 --insecure --log-level debug traces multi
```

TODO:

```bash
kubectl -n tracing run -it --rm --restart=Never otelgen --image=ghcr.io/krzko/otelgen:latest -- --otel-exporter-otlp-endpoint 10.0.17.70:4317 --insecure --log-level debug --duration 10 --rate 1 traces multi
```

## Metrics

To expose metrics add:

```yaml
config:
  ssl: false
  metricRegistries:
    - "EmbeddedMetricsFormat"
    - "Prometheus"
```

And when you deploy pod you'll see it at:

```bash
http://localhost:4900/metrics/sys
```
