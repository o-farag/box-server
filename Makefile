#Builds images, pushes images, builds executables

DOCKER_USERNAME ?= ofarag
APPLICATION_NAME ?= box-server
 
build-image:
	docker build --tag ${DOCKER_USERNAME}/${APPLICATION_NAME} . --network=host

push:
	docker push ${DOCKER_USERNAME}/${APPLICATION_NAME}

build:
	cd ./src && go build -o ../bin/box-server

run-image:
	docker run -p 8080:8080 ofarag/box-server:latest