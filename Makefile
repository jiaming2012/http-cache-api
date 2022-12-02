docker-image := jiamingnj/http-proxy-api
docker-build:
	docker build -t $(docker-image):latest .

deploy:
	docker login
	make docker-build
	docker push $(docker-image):latest
	kubectl apply -f deployments/deploy.yml
	kubectl apply -f deployments/service.yml