#!/usr/bin/env bash
set -Eeo pipefail

# crea pubblicazione, per elencare le tabelle pubblicate: select * from pg_publication_tables;

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE PUBLICATION moqui_pub FOR ALL TABLES;
EOSQL
