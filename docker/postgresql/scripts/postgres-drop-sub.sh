#!/usr/bin/env bash
set -Eeo pipefail

PRIMARY_HOST="${1:moqui-database1}"
PRIMARY_PORT="${2:-5432}"

runQuery() {
  query=$1
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" -AXqtc "$query"
}

subscriptionToDelete=$(runQuery "select count(*) from pg_catalog.pg_stat_subscription WHERE subname='moqui_sub';")
echo
echo "Subscription count: "$subscriptionToDelete
echo

if [ $subscriptionToDelete == 0 ]
then
  echo "Subscription not found: exit"
  echo
  exit 0
fi

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    ALTER SUBSCRIPTION moqui_sub DISABLE;
    ALTER SUBSCRIPTION moqui_sub SET (slot_name=NONE);
    DROP SUBSCRIPTION moqui_sub;
EOSQL

count=$(runQuery "select count(*) from pg_catalog.pg_stat_subscription WHERE subname='moqui_sub';")

if [ $count == 0 ]
then
  echo "Deleted subscriptions: "$subscriptionToDelete
  echo
  exit 0
fi