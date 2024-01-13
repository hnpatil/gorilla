
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

.PHONY: port-forward
port-forward:
	kubectl port-forward service/gorilla 8080:8080 -n gorilla

.PHONY: lint
lint:
	golangci-lint run

.PHONY: vendors
vendors:
	go mod download
	go mod tidy