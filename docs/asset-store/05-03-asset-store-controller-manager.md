---
title: Asset Store Controller Manager sub-chart
type: Configuration
---

To configure the Asset Store Controller Manager sub-chart, override the default values of its `values.yaml` file. This document describes parameters that you can configure.

>**TIP:** To learn more about how to use overrides in Kyma, see the following documents: 
>* [Helm overrides for Kyma installation](/root/kyma/#configuration-helm-overrides-for-kyma-installation)
>* [Sub-charts overrides](/root/kyma/#configuration-helm-overrides-for-kyma-installation-sub-chart-overrides)

## Configurable parameters

This table lists the configurable parameters, their descriptions, and default values:

| Parameter | Description | Default value |
|-----------|-------------|---------------|
| **resources.limits.cpu** | Defines limits for CPU resources. | `100m` |
| **resources.limits.memory** | Defines limits for memory resources. | `128Mi` |
| **resources.requests.cpu** | Defines requests for CPU resources. | `100m` |
| **resources.requests.memory** | Defines requests for memory resources. | `64Mi` |
| **maxClusterAssetConcurrentReconciles** | Defines the maximum number of cluster asset concurrent Reconciles which will run. | `3` |
| **maxAssetConcurrentReconciles** | Defines the maximum number of asset concurrent Reconciles which will run. | `3` |
| **storeUploadWorkersCount** | Defines the number of workers used in parallel to upload files to the storage bucket. | `10` |
| **validationWebhookWorkersCount** | Defines the number of workers used in parallel to validate files. | `10` |
| **mutationWebhookWorkersCount** | Defines the number of workers used in parallel to mutate files. | `10` |
