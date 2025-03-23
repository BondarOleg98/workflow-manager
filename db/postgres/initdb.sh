#!/bin/bash

pg_connection_url="postgres://${PG_USERNAME}:${PG_PASSWORD}@${HOST}:${PORT}"

runSQL() {
  local script_result
  script_result="$(psql ${pg_connection_url}/${PG_ADMIN_DATABASE_NAME} -tc "$1")"
  echo "${script_result}"
}

isDatabaseExists() {
  local exists_db_query
  exists_db_query="SELECT CAST(EXISTS(SELECT datname from pg_database WHERE datname='${PG_DATABASE_NAME}') AS TEXT);"
  db_exists=$(runSQL "${exists_db_query}")
  if [[ -n "${db_exists}" ]] && ${db_exists}; then
    printf "INFO: database %s is exist\n" "${PG_DATABASE_NAME}"
    return 1
  fi
  printf "WARN: Database %s is not exist\n" "${PG_DATABASE_NAME}"
  return 0
}

createDatabase() {
  local create_db_query
  create_db_query="CREATE DATABASE ${PG_DATABASE_NAME}"
  runSQL "${create_db_query}"
}

grantPrivileges() {
  local grant_privileges_query
  grant_privileges_query="GRANT ALL PRIVILEGES ON DATABASE ${PG_DATABASE_NAME} to ${PG_USERNAME}"
  runSQL "${grant_privileges_query}"
}

createTablesAndRelationships() {
  psql ${pg_connection_url}/${PG_DATABASE_NAME} -f create_schema.sql
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