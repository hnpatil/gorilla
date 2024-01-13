.PHONY: deploy-all
deploy-all:
	kubectl apply -f namespace.yaml
	kubectl apply -f config.yaml
	kubectl apply -f gorilla.yaml

.PHONY: delete-all
delete-all:
	kubectl delete -f gorilla.yaml
	kubectl delete -f config.yaml
	kubectl delete -f namespace.yaml