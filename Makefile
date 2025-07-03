MIGRATION_DIR := "migrations"

.PHONY: migrate-generate
migrate-generate:
	goose -s -dir $(MIGRATION_DIR) create $(name) sql
