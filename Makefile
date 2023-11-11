up:
	@echo "Starting app..."
	@docker-compose up --build -d

down:
	@echo "Stopping app..."
	@docker-compose down

watch: up
	@echo "Watching for file changes..."
	@docker-compose watch

logs: 
	@docker-compose logs -f --no-log-prefix user-service