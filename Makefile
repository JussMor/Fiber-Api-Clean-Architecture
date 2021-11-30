network:
	docker network create blog-network

postgres:
	docker run --name postgres12 --network blog-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=9135lp -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root blog_db

dropdb:
	docker exec -it postgres12 dropdb blog_db

server:
	go run server.go

.PHONY: network postgres createdb dropdb  server 