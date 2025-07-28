#!/bin/bash

exportEnvs() {
  logInfo "Exporting envs"
  envs=$(awk 'NF {print $1 $2}' ../../resources/env.yaml | sed 's/:/=/')
  for env in ${envs}; do
    eval "export ${env}"
  done
}

removeTablesAndRelationships() {
  local pg_database_connect
  pg_database_connect="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${HOST}:${PORT}/${POSTGRES_DB}"
  psql "${pg_database_connect}" -f delete_schema.sql
}

exportEnvs
removeTablesAndRelationships
