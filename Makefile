export-variables:
	export $(grep -v  '^#' .env | xargs)
run: export-variables
	go run cmd main.go

