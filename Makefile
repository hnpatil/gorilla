
.PHONY: run
run : 
	go run main.go

.PHONY: generate
generate : 
	go generate ./entity

.PHONY: build-app
build-app:
	docker build --tag gorilla .

.PHONY: start-app
start-app:
	docker compose up -d --remove-orphans

.PHONY: stop-app
stop-app:
	docker compose down

.PHONY: deploy-all
deploy-all:
	kubectl apply -f kube/namespace.yaml
	kubectl apply -f kube/database.yaml
	kubectl apply -f kube/config.yaml
	kubectl apply -f kube/gorilla.yaml
	kubectl apply -f kube/worker.yaml

.PHONY: delete-all
delete-all:
	kubectl delete -f kube/worker.yaml
	kubectl delete -f kube/gorilla.yaml
	kubectl delete -f kube/config.yaml
	kubectl delete -f kube/database.yaml
	kubectl delete -f kube/namespace.yaml

.PHONY: port-forward
port-forward:
	kubectl port-forward service/gorilla 8080:8080 -n gorilla

.PHONY: lint
lint:
	golangci-lint run