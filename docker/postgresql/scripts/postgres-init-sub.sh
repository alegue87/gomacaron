#!/usr/bin/env bash
set -Eeo pipefail

name=$(basename "$0")

if [ "$#" -ne 3 ]; then
    echo
    echo "Usage: $name primary_host primary_port copy_data[true|false]"
    echo
    echo "Example: $name moqui-database1 5432 true"
    echo "Crea sottoscrizione verso moqui-database1 copiando i dati del pubblicatore"
    echo
    echo "Example: $name moqui-database2 5432 false"
    echo "Crea sottoscrizione verso moqui-database2 senza copiare i dati del pubblicatore"
    echo

    exit 0
fi

PRIMARY_HOST="${1:-moqui-database1}"
PRIMARY_PORT="${2:-5432}"
COPY_DATA="${3:-true}"

# crea sottoscrizione, per vederla in db: select * from pg_catalog.pg_stat_subscription;

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE SUBSCRIPTION moqui_sub
    CONNECTION 'host=$PRIMARY_HOST port=$PRIMARY_PORT user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB application_name=moqui_sub'
    PUBLICATION moqui_pub
    WITH ( copy_data = $COPY_DATA, binary = true );
EOSQL

