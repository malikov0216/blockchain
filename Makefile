build:
	docker-compose build blockchain

run:
	docker-compose up blockchain

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up
	
local: 
	go run  cmd/main.go -db=localhost