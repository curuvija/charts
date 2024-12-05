#!/bin/bash
# to be able to apply this multiple times using presync helmfule hook you have to create it first and then apply it
kubectl -n monitoring create secret generic --from-file query-exporter/config.yaml query-exporter-config-secret --dry-run=client -o yaml | kubectl apply -f -
