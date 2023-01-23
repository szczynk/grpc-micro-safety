#!/usr/bin/env bash
set -e

echo "enabling pg_cron on database $POSTGRES_DB"
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname="$POSTGRES_DB"<<-EOSQL
  CREATE EXTENSION IF NOT EXISTS "pg_cron";
EOSQL
echo "finished with exit code $?"