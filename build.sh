#!/bin/bash

logInfo() {
  echo -e "\e[32mINFO [$(date '+%Y-%m-%d %H:%M:%S')] $*\e[0m"
}

logWarn() {
  echo -e "\e[33mWARN [$(date '+%Y-%m-%d %H:%M:%S')] $*\e[0m"
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
chmod +x app/db/postgres/initdb.sh
source app/db/postgres/initdb.sh