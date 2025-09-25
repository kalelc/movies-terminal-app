export:
	set -a; source .env; set +a;
run: export
	go run cmd/main.go


