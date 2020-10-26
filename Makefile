.PHONY: neo4j
neo4j:
	docker-compose -f docker/docker-compose.yml \
	    up --detach


.PHONY: stop
stop:
	docker-compose -f docker/docker-compose.yml \
	    down
