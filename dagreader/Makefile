.PHONY: all
all: dagreader

.PHONY: dagreader
dagreader:
	GIT_COMMIT=`git rev-list -1 HEAD 2>/dev/null || echo ""` && \
	GIT_DATE=`git log -1 --date=short --pretty=format:%ct 2>/dev/null || echo ""` && \
	go build \
	    -ldflags "-s -w -X main.gitCommit=$${GIT_COMMIT} -X main.gitDate=$${GIT_DATE}" \
	    -o build/dagreader \
	    .

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
