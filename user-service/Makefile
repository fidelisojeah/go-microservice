build:
	protoc -I. --go_out=plugins=micro:. \
		proto/auth/auth.proto
	docker build -t eu.gcr.io/fidelis-microservice/user-service:latest .
	gcloud docker -- push eu.gcr.io/fidelis-microservice/user-service:latest

run:
	docker run --net="host" \
		-p 50051 \
		-e DB_HOST=localhost \
		-e DB_PASS=password \
		-e DB_USER=postgres \
		-e JSONSecret=JSONSECRETKEY12**20910 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns user-service
deploy:
	sed "s/{{ UPDATED_AT }}/$(shell date)/g" ./deployments/deployment.tmpl > ./deployments/deployment.yml
	kubectl replace -f ./deployments/deployment.yml
