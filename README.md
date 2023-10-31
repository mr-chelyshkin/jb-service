# jb-service

## Requirements
 - minikube
 - [k9s](http://k9scli.io) (optional)
 - [TaskFile](https://taskfile.dev/installation/)

## Commands

```shell
task -l # List of available commands with description
```

### Run environment
```shell
task minikube/up # run minikube and install all
```
![all-pod](https://github.com/mr-chelyshkin/jb-service/blob/main/media/all-pods.png)

### When environment up
Run service:
```shell
task minikube/service/app-lookup

# main page is swagger with routes and routes description.
``` 
![service-main](https://github.com/mr-chelyshkin/jb-service/blob/main/media/service-main.png)

Run grafana:
```shell
task minikube/infra/grafana-lookup

# run grafana in 127.0.0.1:3000, print credentials
# grafana has predefined Dashboard -> General -> kube-state
# grafana has predefined Dashboard -> General -> JB-Service
``` 
![grafana](https://github.com/mr-chelyshkin/jb-service/blob/main/media/grafana.png)
![grafana](https://github.com/mr-chelyshkin/jb-service/blob/main/media/grafana-creds.png)

Run prometheus:
```shell
minikube/infra/prometheus-lookup
``` 
![prometheus](https://github.com/mr-chelyshkin/jb-service/blob/main/media/prometheus.png)

Run Kibana:
```shell
task minikube/infra/kibana-lookup

# run kibana in 127.0.0.1:5601
# in kibana UI should create index pattern for data "service-logs*"
# after, go to "discover" for watch logs
```
![kibana](https://github.com/mr-chelyshkin/jb-service/blob/main/media/kibana.png)

## Check rateLimit (5rps)
I use [http_bench](https://github.com/linkxzhou/http_bench).
```shell
http_bench http://127.0.0.1:{port} -c 5 -d 30s
```
![rate-limit](https://github.com/mr-chelyshkin/jb-service/blob/main/media/rate-limit.png)

## JB-Service
This is a simple GoLang service that allows you to manipulate 
(on a basic level) its state. For instance, there are handles 
that toggle K8S probes (ok/not ok). The API is described in 
Swagger, and after the service is launched, the page opens 
automatically.

## FAQ
JB-service image `chelyshkin/jb-service` was built on `arm64` and pushed to dockerHub.  
Perhaps it can be problem with `amd64`.

For resolve issue:
```bash
task app/docker/build

# build JB-service image
# this Dockerfile can be build image as amd64/arm64
```

In `./k8s/service/values.yaml`:
```yaml
image:
  repository: {{ local_image }}
  pullPolicy: Never
  tag: {{ image_tag }}
```
