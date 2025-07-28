#!/bin/bash

runSQL() {
  local script_result
  local pg_connection_url
  pg_connection_url="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}"
  script_result="$(psql "${pg_connection_url}/$1" "$2" "$3")"
  echo "${script_result}"
}

logInfo() {
  printf '%b' "\e[32mINFO [$(date '+%Y-%m-%d %H:%M:%S')] $*\e[0m\n"
}

logWarn() {
  printf '%b' "\e[33mWARN [$(date '+%Y-%m-%d %H:%M:%S')] $*\e[0m\n"
}

logInfo "Starting migrations script"

logInfo "Downloading the yq lib"
curl -L -o yq https://github.com/mikefarah/yq/releases/latest/download/yq_linux_amd64
chmod +x yq

logInfo "Exporting envs"
for env in $(./yq ../../resources/env.yaml -os); do
  eval "export ${env}"
done

logInfo "Removing the yq lib"
rm yq

logInfo "Run the db migrations script"
# shellcheck disable=SC2154
runSQL "${POSTGRES_DB}" "-f" "${pg_scripts_path}/${POSTGRES_MIGRATIONS_SCRIPT}"