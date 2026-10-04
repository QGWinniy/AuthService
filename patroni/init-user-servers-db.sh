#!/bin/sh
set -eu

CONNSTR="$1"
DB_NAME="${USER_SERVERS_DB_NAME_PATRONI1:-user_servers}"
INIT_SQL="/docker-entrypoint-initdb.d/001_init.sql"

case "$DB_NAME" in
  *[!a-zA-Z0-9_]* | "")
    echo "Invalid database name: $DB_NAME" >&2
    exit 1
    ;;
esac

if ! psql "$CONNSTR" -v ON_ERROR_STOP=1 -tAc "SELECT 1 FROM pg_database WHERE datname = '$DB_NAME'" | grep -q 1; then
  psql "$CONNSTR" -v ON_ERROR_STOP=1 -c "CREATE DATABASE \"$DB_NAME\""
fi

psql "$CONNSTR" -v ON_ERROR_STOP=1 <<SQL
\\connect "$DB_NAME"
\\i $INIT_SQL
SQL
