pg_connect_url="postgres://${PG_USERNAME}:${PG_PASSWORD}@${HOST}:${PORT}"

runSQL() {
  local script_result
  script_result="$(psql ${pg_connect_url} -tc "$1")"
  echo "${script_result}"
}

checkDatabaseExists() {
  local exists_db_query
  exists_db_query="SELECT CAST(EXISTS(SELECT datname from pg_database WHERE datname='${PG_DATABASE_NAME}') AS TEXT);"
  runSQL "${exists_db_query}"
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
  local pg_database_connect
  pg_database_connect=${pg_connect_url}/${PG_DATABASE_NAME}
  psql ${pg_database_connect} -f create_schema.sql
}

main() {
  db_exists=$(checkDatabaseExists)
  if ${db_exists}; then
    echo "INFO: database ${PG_DATABASE_NAME} is exist"
  else
    echo "WARN: Database ${PG_DATABASE_NAME} is not exist"
    createDatabase
    db_exists=$(checkDatabaseExists)
    if ! ${db_exists}; then
      echo "INFO: Returning from the init script"
      exit 1
    fi
    grantPrivileges
  fi
  createTablesAndRelationships
}

main