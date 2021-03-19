### How to use `pod-viewer`

### Installation

`pod-viewer` is designed to be used as `kubectl` plugin. 

Hence, you need to make sure you have [krew installed](https://github.com/kubernetes-sigs/krew/#installation) and then you can install `pod-viewer` as follows:
```
$ kubectl krew install pod-viewer
```

If the installation fails, check if `krew` is available on your local system and also, make sure you're using the most recent index (run `kubectl krew update`) to ensure this.

### Usage

To get a full view of a given pod, as follows: 

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

Get a full view of a pod running in a different namespace 

```shell
$ kubectl pod-viewer backend -n demo
Name:       backend
Namespace:  demo
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
  Normal  Scheduled  17s   default-scheduler  Successfully assigned demo/backend to kind-control-plane
  Normal  Pulling    16s   kubelet            Pulling image "nginx"
  Normal  Pulled     15s   kubelet            Successfully pulled image "nginx" in 953.2875ms
  Normal  Created    15s   kubelet            Created container backend
  Normal  Started    15s   kubelet            Started container backend
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

A Complete view of pod frontend but only display the most recent 5 lines of output logs

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
10-listen-on-ipv6-by-default.sh: info: Getting the checksum of /etc/nginx/conf.d/default.conf
10-listen-on-ipv6-by-default.sh: info: Enabled listen on IPv6 in /etc/nginx/conf.d/default.conf
/docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
/docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
/docker-entrypoint.sh: Configuration complete; ready for start up

```

A Full view of pod store but only print logs of the front container

```shell
$ kubectl pod-viewer store -c nginx
Name:       store
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
  Normal  Scheduled  21s   default-scheduler  Successfully assigned default/store to kind-control-plane
  Normal  Pulling    20s   kubelet            Pulling image "nginx"
  Normal  Pulled     20s   kubelet            Successfully pulled image "nginx" in 923.6711ms
  Normal  Created    20s   kubelet            Created container front
  Normal  Started    19s   kubelet            Started container front
  Normal  Pulling    19s   kubelet            Pulling image "busybox"
  Normal  Pulled     17s   kubelet            Successfully pulled image "busybox" in 2.7719323s
  Normal  Created    17s   kubelet            Created container busybox
  Normal  Started    16s   kubelet            Started container busybox
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
