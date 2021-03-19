## kubectl-pod-viewer

A kubectl plugin helps you during your daily work with kubernetes pods.

### Context
Pod's lifecycle might be affected if there is a lack of memory or misconfiguration, which lead to an unhealthy state of pod's. In that type of situation, we frequently repeat commands such as `get,describe`,
or even to check logs with `kubectl logs` to debug a given pod.

### What is this about?

`pod-viewer` is a simple CLI tool that provide a full view of kubernetes pod with information such as conditions, events, logs, etc. Also helps you during your debugging journey. For example :

```shell
$ kubectl pod-viewer frontend
Name:       frontend
Namespace:  default
Node:       kind-control-plane/x.x.x.x
Status:     Running
Conditions:
  Type              Status
  Initialized       True 
  Ready             True 
  ContainersReady   True 
  PodScheduled      True 
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  17s   default-scheduler  Successfully assigned default/frontend to kind-control-plane
  Normal  Pulling    16s   kubelet            Pulling image "nginx"
  Normal  Pulled     15s   kubelet            Successfully pulled image "nginx" in 953.2875ms
  Normal  Created    15s   kubelet            Created container frontend
  Normal  Started    15s   kubelet            Started container frontend
Logs:  
/docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
/docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
/docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
10-listen-on-ipv6-by-default.sh: info: Getting the checksum of /etc/nginx/conf.d/default.conf
10-listen-on-ipv6-by-default.sh: info: Enabled listen on IPv6 in /etc/nginx/conf.d/default.conf
/docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
/docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
/docker-entrypoint.sh: Configuration complete; ready for start up

```

### Getting started

To get started visit the [usage docs.](doc/usage.md)