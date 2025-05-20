#!/bin/bash

removeTablesAndRelationships() {
  local pg_database_connect
  pg_database_connect="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${HOST}:${PORT}/${POSTGRES_DB}"
  psql ${pg_database_connect} -f delete_schema.sql
}

removeTablesAndRelationships
