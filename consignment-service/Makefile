build:
	docker build -t eu.gcr.io/fidelis-microservice/consignment-service:latest .
	gcloud docker -- push eu.gcr.io/fidelis-microservice/consignment-service:latest
run:
	docker run -d --net="host" \
		-p 50052 \
		-e MICRO_SERVER_ADDRESS=:50052 \
		-e DISABLE_AUTH=true \
		-e MICRO_REGISTRY=mdns consignment-service
deploy:
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yml
	kubectl replace -f ./deployments/deployment.yml
