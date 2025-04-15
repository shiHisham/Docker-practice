#!/bin/sh
# Load secret from file
export POSTGRES_PASSWORD=$(cat /run/secrets/POSTGRES_PASSWORD)
# Auth method override
export POSTGRES_INITDB_ARGS="--auth=md5"
# Start the original Postgres entrypoint
exec docker-entrypoint.sh postgres