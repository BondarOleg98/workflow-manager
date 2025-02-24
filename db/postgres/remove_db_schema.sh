removeTablesAndRelationships() {
  local pg_database_connect
  pg_database_connect="postgres://${PG_USERNAME}:${PG_PASSWORD}@${HOST}:${PORT}/${PG_DATABASE_NAME}"
  psql ${pg_database_connect} -f delete_schema.sql
}

removeTablesAndRelationships
