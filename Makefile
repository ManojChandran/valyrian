# Define variables
APP_NAME = oouve
VERSION = 1.0
DOCKER_IMAGE = $(APP_NAME):$(VERSION)
DOCKER_CONTAINER = $(APP_NAME)-container

# Default target
all: build

# Build the Go application inside a Docker container
build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

# Run the application in a container
run:
	@echo "Running $(DOCKER_CONTAINER)..."
	docker run -d --name $(DOCKER_CONTAINER) -p 8080:8080 $(DOCKER_IMAGE)

# Stop the running container
stop:
	@echo "Stopping $(DOCKER_CONTAINER)..."
	docker stop $(DOCKER_CONTAINER) || true
	docker rm $(DOCKER_CONTAINER) || true

# View container logs
logs:
	@echo "Fetching logs from $(DOCKER_CONTAINER)..."
	docker logs -f $(DOCKER_CONTAINER)

# Test the API with curl
test:
	@echo "Testing HTTPS API endpoint..."
	curl -k https://localhost:8080

# Clean up Docker images and containers
clean:
	@echo "Cleaning up..."
	docker stop $(DOCKER_CONTAINER) || true
	docker rm $(DOCKER_CONTAINER) || true
	docker rmi $(DOCKER_IMAGE) || true