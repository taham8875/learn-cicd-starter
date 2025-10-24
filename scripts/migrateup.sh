#!/bin/bash

if [ -f .env ]; then
    source .env
fi

# Build a DSN that includes an auth token if it's not already present
DSN="$DATABASE_URL"
case "$DSN" in
    *authToken=*)
        : # already has token
        ;;
    *)
        if [ -n "$TURSO_AUTH_TOKEN" ]; then
            if [[ "$DSN" == *"?"* ]]; then
                DSN="${DSN}&authToken=${TURSO_AUTH_TOKEN}"
            else
                DSN="${DSN}?authToken=${TURSO_AUTH_TOKEN}"
            fi
        fi
        ;;
esac

cd sql/schema
goose turso "$DSN" up
