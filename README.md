## kubectl-pod-viewer

A kubectl plugin helps you during your daily work with kubernetes pods.

### Context
Pod's lifecycle might be affected if there is a lack of memory or misconfiguration, which lead to an unhealthy state of pod's. In that type of situation, we frequently repeat commands such as `get,describe` in debugging,
or even to check logs.

### What is this about?

`pod-viewer`, a simple CLI tool that provide a full view of kubernetes pod with information such as name, namespace, nodeName, status, conditions, events, logs, etc. For example :

```


```

### Getting started

To get started visit the [usage docs.](doc/usage.md)