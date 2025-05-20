#!/bin/bash

pg_connection_url="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${HOST}:${PORT}"

runSQL() {
  local script_result
  script_result="$(psql ${pg_connection_url}/${PG_ADMIN_DATABASE_NAME} -tc "$1")"
  echo "${script_result}"
}

isDatabaseExists() {
  local exists_db_query
  exists_db_query="SELECT CAST(EXISTS(SELECT datname from pg_database WHERE datname='${POSTGRES_DB}') AS TEXT);"
  db_exists=$(runSQL "${exists_db_query}")
  if [[ -n "${db_exists}" ]] && ${db_exists}; then
    printf "INFO: database %s is exist\n" "${POSTGRES_DB}"
    return 1
  fi
  printf "WARN: Database %s is not exist\n" "${POSTGRES_DB}"
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
  psql ${pg_connection_url}/${POSTGRES_DB} -f create_schema.sql
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