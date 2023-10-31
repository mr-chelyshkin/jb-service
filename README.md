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
```
![all-pod](https://github.com/mr-chelyshkin/jb-service/blob/main/media/all-pods.png)


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
# grafana has predefined Dashboard "kube-state"
``` 
![grafana](https://github.com/mr-chelyshkin/jb-service/blob/main/media/grafana.png)

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
