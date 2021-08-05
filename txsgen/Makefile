.PHONY: all
all: txsgen

.PHONY: txsgen
txsgen:
	GIT_COMMIT=`git rev-list -1 HEAD 2>/dev/null || echo ""` && \
	GIT_DATE=`git log -1 --date=short --pretty=format:%ct 2>/dev/null || echo ""` && \
	go build \
	    -ldflags "-s -w -X main.gitCommit=$${GIT_COMMIT} -X main.gitDate=$${GIT_DATE}" \
	    -o build/txsgen \
	    .

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -fr ./build/*
