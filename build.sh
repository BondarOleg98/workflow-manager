#!/bin/bash

pg_scripts_path="app/db/postgres"

logInfo() {
  printf '%b' "\e[32mINFO [$(date '+%Y-%m-%d %H:%M:%S')] $*\e[0m\n"
}

logWarn() {
  printf '%b' "\e[33mWARN [$(date '+%Y-%m-%d %H:%M:%S')] $*\e[0m\n"
}

exportEnvs() {
  logInfo "Exporting envs"
  envs=$(awk 'NF {print $1 $2}' app/resources/env.yaml | sed 's/:/=/')
  for env in ${envs}; do
    eval "export ${env}"
  done
}

logInfo "Starting the build script"
exportEnvs
logInfo "Run the init db script"
chmod +x ${pg_scripts_path}/initdb.sh
source ${pg_scripts_path}/initdb.sh