#!/bin/bash

pg_scripts_path="app/db/postgres"

logInfo() {
  printf '%b' "\e[32mINFO [$(date '+%Y-%m-%d %H:%M:%S')] $*\e[0m\n"
}

logWarn() {
  printf '%b' "\e[33mWARN [$(date '+%Y-%m-%d %H:%M:%S')] $*\e[0m\n"
}

logInfo "Starting script"

logInfo "Downloading the yq lib"
curl -L -o yq https://github.com/mikefarah/yq/releases/latest/download/yq_linux_amd64
chmod +x yq

logInfo "Exporting envs"
for env in $(./yq app/resources/env.yaml -os); do
  eval "export ${env}"
done

logInfo "Removing the yq lib"
rm yq

logInfo "Run the init db script"
chmod +x ${pg_scripts_path}/initdb.sh
source ${pg_scripts_path}/initdb.sh