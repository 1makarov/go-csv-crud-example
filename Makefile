migrate-status:
	goose -dir .\schema postgres "user=root dbname=store password=test sslmode=disable" status

migrate-up:
	goose -dir .\schema postgres "user=root dbname=store password=test sslmode=disable" up

migrate-down:
	goose -dir .\schema postgres "user=root dbname=store password=test sslmode=disable" up

migrate-version:
	goose -dir .\schema postgres "user=root dbname=store password=test sslmode=disable" version