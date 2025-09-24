#!/bin/bash

logInfo() {
  printf '%b' "\e[32mINFO [$(date '+%Y-%m-%d %H:%M:%S')] $*\e[0m\n"
}

logWarn() {
  printf '%b' "\e[33mWARN [$(date '+%Y-%m-%d %H:%M:%S')] $*\e[0m\n"
}

exportEnvs() {
  logInfo "Exporting envs"
  envs=$(awk 'NF {print $1 $2}' app/resources/prod_env.yaml | sed 's/:/=/')
  for env in ${envs}; do
    eval "export ${env}"
  done
}

logInfo "Starting the build script"
exportEnvs
logInfo "Run the init db script"
chmod +x "${POSTGRES_SCRIPTS_PATH}/${POSTGRES_INIT_SCRIPT}"
# shellcheck disable=SC1090
source "${POSTGRES_SCRIPTS_PATH}/${POSTGRES_INIT_SCRIPT}"