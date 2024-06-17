## TODOs

* [x] Started http source on port 2021... - check how to push logs and verify they come at opensearch (check https://github.com/opensearch-project/data-prepper/tree/main/examples/log-ingestion)
* [x] add port 2021 to service
* [x] configure logging and tracing pipeline
* [ ] enable liveness and readiness probes
* [ ] rewrite NOTES.txt and put it back
* [x] expose prometheus metrics https://opensearch.org/docs/latest/data-prepper/managing-data-prepper/monitoring/ (http://localhost:4900/metrics/sys)
* [x] add data-prepper configuration -> https://opensearch.org/docs/latest/data-prepper/managing-data-prepper/configuring-data-prepper/
* [ ] move configuration into secret instead configmap since there are sensitive data
* [ ] add log4j2 configuration https://opensearch.org/docs/latest/data-prepper/getting-started/#additional-configurations
* [ ] check if you need to configure peer forwarder for node trace distribution https://opensearch.org/docs/latest/data-prepper/managing-data-prepper/peer-forwarder/
* [ ] configure data prepper authentication https://opensearch.org/docs/latest/data-prepper/managing-data-prepper/core-apis/#authentication