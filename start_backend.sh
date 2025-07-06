#!/bin/bash

# Set environment variables for DB connection
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=trading

# Ensure Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go to run this script."
    exit 1
fi

# Ensure PostgreSQL is running
if ! nc -z $DB_HOST $DB_PORT; then
    echo "PostgreSQL is not running on $DB_HOST:$DB_PORT. Please start PostgreSQL before running this script."
    exit 1
else
    echo "PostgreSQL is running."
fi

# Ensure Go modules are tidy and downloaded
echo "Tidying and downloading Go modules..."
go mod tidy && go mod download

# Build the Go app (binary name matches Dockerfile)
echo "Building Go backend..."
go build -o trading-system

# Run the backend
echo "Starting backend..."
./trading-system