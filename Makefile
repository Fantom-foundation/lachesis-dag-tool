.PHONY: neo4j
neo4j:
	docker-compose -f docker/docker-compose.yml \
	    up --detach neo4j

.PHONY: neo4j-sql
neo4j-sql:
	docker-compose -f ./docker/docker-compose.yml \
	    exec neo4j cypher-shell


.PHONY: stop
stop:
	docker-compose -f docker/docker-compose.yml \
	    down
