# SAI Kubernetes Homework
## Overview
In the main.go there is a simple HTTP server listening on port `8000`, which exposes `/stress` endpoint. It handles GET requests with `time` parameter - the length of stress function in seconds.

## Build image
Let's build our awesome container based on the golang app 
``` bash
docker build --tag <registry>/sai-stress:v0.0.1 . --push
```

## Deploy
Create Namespace, Deployment for our server, Service which will act as a load balancer and Horizontal Pod Autoscaler which will handle scaling of our application. Before creating these resources, make sure that you provide the correct image registry and tag for the server (`server-deploy.yaml`, line 25)

``` bash
kubectl apply -f server-deploy.yaml
```

Deploy our GUI - nginx with a customized index page because I'm lazy (Deployment, Service and ConfigMap).
``` bash
kubectl apply -f gui-deploy.yaml
```

Expose service via port-forward in the background:
``` bash
kubectl port-forward --namespace sai svc/gui 8080:80 &
```

As this is a very simple frontend that is executed in the user browser, the redirect URL needs to be accessible to the user, thus we will also forward ports of the server in the background.

```bash
kubectl port-forward --namespace sai svc/server 8000:80 &
```

Access our GUI via browser and start spamming the button. Watch for the changing number of container instances.