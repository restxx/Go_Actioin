#!/bin/bash

# Build the project
echo "Building the project..."
go mod tidy
go build -o main main.go
