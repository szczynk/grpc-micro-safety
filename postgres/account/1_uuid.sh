#!/usr/bin/env bash
set -e

echo "enabling uuid-ossp on database $POSTGRES_DB"
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname="$POSTGRES_DB"<<-EOSQL
  CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
EOSQL
echo "finished with exit code $?"