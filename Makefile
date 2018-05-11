all:
	go test -timeout 5h -count 5 -bench . -benchmem | tee results
	go run cmd/benchtable/main.go < results > README.md
