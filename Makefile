ROOT := .
TMP_DIR := ./tmp
SV_BIN := $(TMP_DIR)/server
CSM_BIN := $(TMP_DIR)/consumer
SD_BIN := $(TMP_DIR)/seeder
SC_BIN := $(TMP_DIR)/scheduler
SV_DIR := ./cmd/server
CSM_DIR := ./cmd/consumer
SD_DIR := ./cmd/seeder
SC_DIR := ./cmd/scheduler
DOCKERFILE_DIR := .
ENVFILE_DIR := .env.local
IMAGE_NAME := instay-be
CONTAINER_SERVER := instay_server
CONTAINER_CONSUMER := instay_consumer
CONTAINER_SCHEDULER := instay_scheduler

.PHONY: build-sv run-sv build-csm run-csm build-sd run-sd build-sc run-sc clean github docker-br docker-rm

# Require Ubuntu
build-sv:
	@echo "Building..."
	@mkdir -p $(TMP_DIR)
	go build -o $(SV_BIN) $(SV_DIR)

run-sv: build-sv
	@echo "Running..."
	@$(SV_BIN)

build-sc:
	@echo "Building..."
	@mkdir -p $(TMP_DIR)
	go build -o $(SC_BIN) $(SC_DIR)

run-sc: build-sc
	@echo "Running..."
	@$(SC_BIN)

build-csm:
	@echo "Building..."
	@mkdir -p $(TMP_DIR)
	go build -o $(CSM_BIN) $(CSM_DIR)

run-csm: build-csm
	@echo "Running..."
	@$(CSM_BIN)

build-sd:
	@echo "Building..."
	@mkdir -p $(TMP_DIR)
	go build -o $(SD_BIN) $(SD_DIR)

run-sd: build-sd
	@echo "Running..."
	@$(SD_BIN)

clean:
	@echo "Cleaning..."
	@rm -rf $(TMP_DIR)

# Require Windows
github:
	@if "$(CM)"=="" ( \
		echo Usage: make github CM="commit message" && exit 1 \
	)
	git add .
	git commit -m "$(CM)"
	git push
	git push clone

# Require Docker
docker-br:
	docker build -t $(IMAGE_NAME) $(DOCKERFILE_DIR)
	docker run --env-file $(ENVFILE_DIR) -d -p 8080:8080 --name $(CONTAINER_SERVER) $(IMAGE_NAME) ./server
	docker run --env-file $(ENVFILE_DIR) --rm $(IMAGE_NAME) ./seeder
	docker run --env-file $(ENVFILE_DIR) -d --name $(CONTAINER_CONSUMER) $(IMAGE_NAME) ./consumer
	docker run --env-file $(ENVFILE_DIR) -d --name $(CONTAINER_SCHEDULER) $(IMAGE_NAME) ./scheduler

docker-rm:
	docker stop $(CONTAINER_SERVER)  $(CONTAINER_CONSUMER) $(CONTAINER_SCHEDULER)
	docker rm $(CONTAINER_SERVER)  $(CONTAINER_CONSUMER) $(CONTAINER_SCHEDULER)
	docker rmi $(IMAGE_NAME)