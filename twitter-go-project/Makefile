BUILD_ARCH=linux

BROKER_APP:=brokerApp
SESSION_APP:=sessionApp

up_build: build down
	@echo "Starting Docker images..."
	docker-compose up --build -d

up:
	@echo "Starting Docker images..."
	docker-compose up -d

down:
	@echo "Stopping Docker images..."
	docker-compose down

build: build_broker build_session
	@echo "Built all Docker images."

build_broker:
	@echo "Building broker service..."
	cd ../broker-service && ./build_go.sh -o -a ${BUILD_ARCH} ./api ${BROKER_APP}

build_session:
	@echo "Building session service..."
	cd ../session-service && ./build_go.sh -o -a ${BUILD_ARCH} ./api ${SESSION_APP}
