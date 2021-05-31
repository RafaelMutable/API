.PHONY: compile-server
compile-server: #Build the server
	@echo "Building the server."
	@go get -u github.com/gorilla/mux
	@go build Server/main.go

.PHONY: run-server
run-server: #Run the server
	@echo "launching the server."
	@go run Server/main.go

.PHONY: server
server: #Build and run server
	@echo "Building and launching the server."
	@go get -u github.com/gorilla/mux
	@go build Server/main.go
	@go run Server/main.go