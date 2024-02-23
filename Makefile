
build:
	docker compose -f build/docker-compose.yml build $(c)

run:
	docker compose -f build/docker-compose.yml up

restart:
	docker compose -f build/docker-compose.yml build $(c)
	docker compose -f build/docker-compose.yml up
	
down:
	docker compose -f build/docker-compose.yml down $(c)