#!/bin/sh

# Load variables from .env file
set -a  # Automatically export all variables
. ./.env
set +a

# Now you can use the variables
echo "Database URL is postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable"

# Example with migrate command
migrate -path cmd/apiserver/app/migrations/ -database "postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" down $1