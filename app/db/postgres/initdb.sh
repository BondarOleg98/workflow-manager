#!/bin/bash

pg_connection_url="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}"

runSQL() {
  local script_result
  script_result="$(psql "${pg_connection_url}/${POSTGRES_ADMIN_DB}" -tc "$1")"
  echo "${script_result}"
}

isDatabaseExists() {
  local exists_db_query
  exists_db_query="SELECT CAST(EXISTS(SELECT datname from pg_database WHERE datname='${POSTGRES_DB}') AS TEXT);"
  db_exists=$(runSQL "${exists_db_query}")
  if [[ -n "${db_exists}" ]] && ${db_exists}; then
    logInfo "The database ${POSTGRES_DB} is exist"
    return 1
  fi
  logWarn "The database ${POSTGRES_DB} is not exist"
  return 0
}

createDatabase() {
  local create_db_query
  create_db_query="CREATE DATABASE ${POSTGRES_DB}"
  runSQL "${create_db_query}"
}

grantPrivileges() {
  local grant_privileges_query
  grant_privileges_query="GRANT ALL PRIVILEGES ON DATABASE ${POSTGRES_DB} to ${POSTGRES_USER}"
  runSQL "${grant_privileges_query}"
}

createTablesAndRelationships() {
  psql "${pg_connection_url}/${POSTGRES_DB}" -f app/db/postgres/create_schema.sql
}

main() {
  if isDatabaseExists; then
    createDatabase
    if isDatabaseExists; then
      exit 1
    fi
    grantPrivileges
  fi
  createTablesAndRelationships
}

main