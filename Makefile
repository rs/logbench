all:
	go test -timeout 5h -count 5 -bench . -benchmem | tee results
	(echo "---\nlayout: default\n---"; go run cmd/benchjson/main.go < results) > docs/index.html
