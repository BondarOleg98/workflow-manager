#!/bin/bash

logInfo() {
  printf '%b' "\e[32mINFO [$(date '+%Y-%m-%d %H:%M:%S')] $*\e[0m\n"
}

exportEnvs() {
  logInfo "Exporting envs"
  envs=$(awk 'NF {print $1 $2}' ../../resources/env.yaml | sed 's/:/=/')
  for env in ${envs}; do
    eval "export ${env}"
  done
}

removeTablesAndRelationships() {
  logInfo "Removing tables and relationships"
  local pg_database_connect
  pg_database_connect="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${HOST}:${PORT}/${POSTGRES_DB}"
  psql "${pg_database_connect}" -f delete_schema.sql
}

logInfo "Starting the removing script"
exportEnvs
removeTablesAndRelationships