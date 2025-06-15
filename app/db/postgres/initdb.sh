#!/bin/bash

pg_connection_url="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}"

runSQL() {
  local script_result
  script_result="$(psql "${pg_connection_url}/$1" "$2" "$3")"
  echo "${script_result}"
}

isDatabaseExists() {
  local exists_db_query
  exists_db_query="SELECT CAST(EXISTS(SELECT datname from pg_database WHERE datname='${POSTGRES_DB}') AS TEXT);"
  db_exists=$(runSQL "${POSTGRES_ADMIN_DB}" "-tc" "${exists_db_query}")
  if [[ -n "${db_exists}" ]] && ${db_exists}; then
    logInfo "The database ${POSTGRES_DB} is exist"
    return 0
  fi
  logWarn "The database ${POSTGRES_DB} is not exist"
  return 1
}

areTablesExist() {
  local exist_tables_query
  local table_exists
  for table_name in "$@"; do
    exist_tables_query="
    SELECT 1 FROM information_schema.tables
    WHERE table_type='BASE TABLE'
    AND table_schema='public'
    AND table_catalog='workflow_manager'
    AND table_name='${table_name}';"
    table_exists=$(runSQL "${POSTGRES_DB}" "-tc" "${exist_tables_query}")
    if [[ ! "${table_exists}" ]]; then
      logWarn "The table ${table_name} is not exist"
      return 1
    fi
  done
  logInfo "Tables are exist"
  return 0
}

checkDataSizeInTable() {
  for table_name in "$@"; do
    table_size_query="SELECT COUNT (*) FROM ${table_name};"
    table_size=$(runSQL "${POSTGRES_DB}" "-tc" "${table_size_query}")
    logInfo "The table ${table_name} has a size: ${table_size}"
  done
}

createDatabase() {
  local create_db_query
  create_db_query="CREATE DATABASE ${POSTGRES_DB}"
  runSQL "${POSTGRES_ADMIN_DB}" "-tc" "${create_db_query}"
}

grantPrivileges() {
  local grant_privileges_query
  logInfo "Grant privileges"
  grant_privileges_query="GRANT ALL PRIVILEGES ON DATABASE ${POSTGRES_DB} to ${POSTGRES_USER}"
  runSQL "${POSTGRES_ADMIN_DB}" "-tc" "${grant_privileges_query}"
}

createTablesAndRelationships() {
  runSQL "${POSTGRES_DB}" "-f" "app/db/postgres/create_schema.sql"
}

fillDatabase() {
  logInfo "Filling DB"
  runSQL "${POSTGRES_DB}" "-f" "app/db/postgres/fill_db.sql"
}

main() {
  declare -a tables=([0]=workflows [1]=tasks [2]=actions)
  if ! isDatabaseExists; then
    createDatabase
    grantPrivileges
  fi
  if ! areTablesExist "${tables[@]}"; then
    createTablesAndRelationships
  fi

  fillDatabase
  checkDataSizeInTable "${tables[@]}"
}

main