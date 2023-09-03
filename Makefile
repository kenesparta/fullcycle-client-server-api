# Run this first
init:
	go work init
	go work use ./server
	go work use ./client

run-server:
	go run ./server

run-client:
	go run ./client
