#!/usr/bin/env bash
set -Eeo pipefail

name=$(basename "$0")

if [ "$#" -ne 3 ]; then
    echo "Usage: $name primary_host primary_port secondary_port"
    echo
    echo "Example: $name moqui-database1 5432 5432"
    echo "Crea schema su host secondario (localhost)"
    echo

    exit 0
fi

PRIMARY_HOST="${1:-moqui-database1}"
PRIMARY_PORT="${2:-5432}"
SECONDARY_PORT="${3:-5432}"

pg_dump -Fc -h $PRIMARY_HOST -p $PRIMARY_PORT -U $POSTGRES_USER -d $POSTGRES_DB -s | \
pg_restore -h localhost -p $SECONDARY_PORT -U $POSTGRES_USER -d $POSTGRES_DB -s

