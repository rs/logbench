all:
	go test -timeout 5h -count 10 -bench . -benchmem | tee results
	go run cmd/benchtable/main.go < results > README.md
