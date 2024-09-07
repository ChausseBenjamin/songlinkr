DOCKER_IMAGE=songlinkr
CONTAINER_NAME=songlinkr-container

compile:
	docker build -t $(DOCKER_IMAGE) .

run:
	docker run --rm \
		--name $(CONTAINER_NAME) \
		-v $(PWD)/secrets:/secrets \
		-e SECRETS_PATH="/secrets" \
		$(DOCKER_IMAGE)
