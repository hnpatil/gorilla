.PHONY: deploy-local
deploy-local:
	kubectl apply -f local/namespace.yaml
	kubectl apply -f local/database.yaml
	kubectl apply -f local/config.yaml
	kubectl apply -f local/gorilla.yaml
	kubectl apply -f local/worker.yaml

.PHONY: delete-local
delete-local:
	kubectl delete -f local/worker.yaml
	kubectl delete -f local/gorilla.yaml
	kubectl delete -f local/config.yaml
	kubectl delete -f local/database.yaml
	kubectl delete -f local/namespace.yaml