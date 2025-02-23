db_exists="$(psql ${PG_URL} -tc "SELECT EXISTS(SELECT datname from pg_database WHERE datname='${PG_DATABASE_NAME}');" | sed -e 's/^[[:space:]]*//')"

if [[ ${db_exists} == 't' ]]; then
  echo "${PG_DATABASE_NAME} is exist"
else
  echo "${PG_DATABASE_NAME} is not exist"
fi
