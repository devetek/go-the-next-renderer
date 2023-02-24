run:
	@docker-compose down --remove-orphans
	@docker-compose up

ls:
	@docker-compose ps

down:
	@docker-compose down --remove-orphans