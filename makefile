run_dev_build: 
	docker build . -t rinha-api:latest
	docker compose -f docker-compose.dev.yml up -d

run_dev: 
	docker compose -f docker-compose.dev.yml up -d

stop_dev:
	docker compose -f docker-compose.dev.yml down

up:
	docker compose up -d

down:
	docker compose down

reload:
	docker compose down
	docker build . -t rinha-api:latest
	docker compose up -d

db_con:
	docker exec -it rinha-go-db-1 psql -U postgres -d rinhadb
	
test:
	go test ./... 
