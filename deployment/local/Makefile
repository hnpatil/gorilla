.PHONY: deploy-all
deploy-all:
	kubectl apply -f namespace.yaml
	kubectl apply -f database.yaml
	kubectl apply -f config.yaml
	kubectl apply -f gorilla.yaml
	kubectl apply -f worker.yaml

.PHONY: delete-all
delete-all:
	kubectl delete -f worker.yaml
	kubectl delete -f gorilla.yaml
	kubectl delete -f config.yaml
	kubectl delete -f database.yaml
	kubectl delete -f namespace.yaml