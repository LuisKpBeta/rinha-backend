run_dev_build: 
	docker build . -t rinha-api:latest
	docker compose -f docker-compose.dev.yml up -d

run_dev: 
	docker compose -f docker-compose.dev.yml up -d

stop_dev:
	docker compose -f docker-compose.dev.yml down