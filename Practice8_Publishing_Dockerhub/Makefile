.PHONY: build-scratch-image
build-scratch-image:
	docker build --tag myfirst-scratch-image .

# Firstly build the image, and log in to dockerhub before running this target
# https://docs.docker.com/engine/reference/commandline/login/
.PHONY: push-dockerhub
push-dockerhub:
	docker tag myfirst-scratch-image hichamshih/myfirst-scratch-image # defaults to latest
	docker push hichamshih/myfirst-scratch-image

# tagging with a version number

	docker tag myfirst-scratch-image hichamshih/myfirst-scratch-image:v1
	docker push hichamshih/myfirst-scratch-image:v1

# Firstly build the image, and log in to github packages before running this target
# https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry#authenticating-to-the-container-registry

.PHONY: push-github-packages
push-github-packages:
	docker tag myfirst-scratch-image ghcr.io/shihisham/myfirst-scratch-image # defaults to latest
	docker push ghcr.io/shihisham/myfirst-scratch-image

# tagging with a version number
	docker tag myfirst-scratch-image ghcr.io/shihisham/myfirst-scratch-image:v1 # defaults to latest
	docker push ghcr.io/shihisham/myfirst-scratch-image:v1