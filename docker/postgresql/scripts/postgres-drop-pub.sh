#!/usr/bin/env bash
set -Eeo pipefail


PRIMARY_HOST="${1:moqui-database1}"
PRIMARY_PORT="${2:-5432}"

runQuery() {
  query=$1
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" -AXqtc "$query"
}

publicationToDelete=$(runQuery "select count(*) as publication_count from pg_publication_tables WHERE pubname='moqui_pub';")
echo
echo "Publication count: "$publicationToDelete
echo

if [ $publicationToDelete == 0 ]
then
  echo "Publication not found: exit"
  echo
  exit 0
fi

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    DROP PUBLICATION IF EXISTS moqui_pub;
    SELECT pg_drop_replication_slot(slot_name) FROM pg_replication_slots where slot_name='moqui_sub';
EOSQL

count=$(runQuery "select count(*) as publication_count from pg_publication_tables WHERE pubname='moqui_pub';")

if [ $count == 0 ]
then
  echo "Deleted publications: "$publicationToDelete
  echo
  exit 0
fi

